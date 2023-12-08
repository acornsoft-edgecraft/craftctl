package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/edgecraft/edge-benchmarks/pkg/config"
	"github.com/edgecraft/edge-benchmarks/pkg/db"
	"github.com/edgecraft/edge-benchmarks/pkg/logger"
	sum "github.com/edgecraft/edge-benchmarks/pkg/summarize"
	"github.com/spf13/cobra"
)

var summarize *sum.Summarize

var rootCmd = &cobra.Command{
	Use:   "EdgeCraft CIS-Benchmarks",
	Short: "CIS Benchmarks for Edgecraft Cluster",
	Long:  "CIS Benchmarks for Edgecraft Cluster",
	Run: func(cmd *cobra.Command, args []string) {

		// Interrupt handler
		c := make(chan os.Signal)
		go func() {
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			logger.Infof("Recieved %s signal", <-c)
			os.Exit(1)
		}()

		if err := summarize.Execute(); err != nil {
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().String("benchmarks-id", "", "benchmarks-id for edgecraft")
	rootCmd.Flags().String("results-dir", "/tmp/edge", "results directory (default is /tmp/edge)")
	rootCmd.Flags().String("reason", "", "reason for error")
}

func initConfig() {
	// create default logger
	err := logger.New()
	if err != nil {
		logger.Fatalf("Could not instantiate log %ss", err.Error())
	}

	// load config file
	conf, err := config.Load()
	if err != nil {
		logger.Fatalf("Could not load configuration: %s", err.Error())
		os.Exit(0)
	}

	DB, err := db.New(conf.DB)
	if err != nil {
		logger.Fatalf("Could not connect DB: %s", err.Error())
	}

	// get command flags
	benchmarksId, err := rootCmd.Flags().GetString("benchmarks-id")
	if err != nil {
		logger.Warnf("invalid benchmarks-id", err.Error())
	} else if benchmarksId == "" {
		logger.Fatal("benchmarks id is required")
		os.Exit(0)
	}
	logger.Infof("specified benchmarks-id is %s", benchmarksId)

	resultsDir, err := rootCmd.Flags().GetString("results-dir")
	if err != nil {
		logger.Warnf("invalid results directory", err.Error())
	}
	logger.Infof("specified results directory is %s", resultsDir)

	reason, err := rootCmd.Flags().GetString("reason")
	if err != nil {
		logger.Warnf("invalid reason", err.Error())
	}
	logger.Infof("specified reason is %s", resultsDir)

	params := &sum.Config{
		BenchmarksId:     benchmarksId,
		ResultsDirectory: resultsDir,
		Reason:           reason,
	}
	summarize = sum.NewSummarize(DB, params)

	logger.Info("Intialize done.")
}
