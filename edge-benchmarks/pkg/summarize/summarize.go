package summarize

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/edgecraft/edge-benchmarks/pkg/db"
	"github.com/edgecraft/edge-benchmarks/pkg/logger"
	"github.com/edgecraft/edge-benchmarks/pkg/model"
	"gorm.io/gorm"
)

type Config struct {
	BenchmarksId     string
	ResultsDirectory string
	Reason           string
}

type Summarize struct {
	DB     *gorm.DB
	Config *Config
}

func NewSummarize(db *gorm.DB, conf *Config) *Summarize {
	return &Summarize{
		DB:     db,
		Config: conf,
	}
}

const (
	ResultsSDir = "results"
	ErrorsDir   = "errors"
	Updater     = "edge-summarize"
)

func (s *Summarize) Execute() error {
	if s.Config.Reason != "" {
		err := s.updateDBFailure(s.Config.Reason)
		if err != nil {
			return err
		}
		logger.Info("saving the failure data to the database")
		return nil
	}

	results, err := s.Summarize()
	if err != nil {
		e := s.updateDBFailure(err.Error())
		if e != nil {
			return e
		}
		logger.Info("saving the failure data to the database")
		return err
	}

	logger.Info("saving the success data to the database")
	err = s.updateDBSuccess(results)
	if err != nil {
		logger.Warnf("error saving the success data to the database: %v ", err)
		return err
	}

	logger.Info("Execute done.")
	return nil
}

func (s *Summarize) Summarize() (*model.Results, error) {
	// check sonobuoy plugins results
	pluginsDir := filepath.Join(s.Config.ResultsDirectory, "plugins")

	// search plugsins's sub directory
	dirs, err := filepath.Glob(filepath.Join(pluginsDir, "*", "*"))
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		info, err := os.Stat(dir)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			continue
		}
		if strings.EqualFold(info.Name(), ResultsSDir) {
			continue
		}
		// if errors directory exists, occurred kube-bench error
		if strings.EqualFold(info.Name(), ErrorsDir) {
			logger.Errorf("error executing kube-bench. directory: %s", dir)
			return nil, errors.New("error executing kube-bench")
		}
	}

	// summarize results and totals
	files, err := filepath.Glob(filepath.Join(pluginsDir, "*", ResultsSDir, "*", "*.json"))
	if err != nil {
		return nil, err
	}

	var sumNodes []model.Node
	var totals []model.Total
	results := &model.Results{}

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			continue
		}

		name, _, _ := strings.Cut(info.Name(), ".json")
		contents, err := os.ReadFile(file)
		if err != nil {
			logger.Errorf("error reading file, results file name is %s: %v", file, err)
			return nil, err
		}

		output := &model.KbOutput{}
		if err := json.Unmarshal(contents, output); err != nil {
			logger.Errorf("error unmarshalling, results file name is %s: %v ", file, err)
			return nil, err
		}

		for _, control := range output.Controls {
			if results.Version != "" && results.DetectedVersion != "" {
				break
			}
			results.Version = control.Version
			results.DetectedVersion = control.DetectedVersion
		}

		node := model.Node{}
		node.NodeName = name
		node.Controls = output.Controls
		node.Totals = output.Totals

		sumNodes = append(sumNodes, node)
		totals = append(totals, output.Totals)
	}
	results.Nodes = sumNodes
	results.Totals = sumTotals(totals)

	logger.Info("Summarize done.")
	return results, nil
}

func (s *Summarize) updateDBSuccess(results *model.Results) error {
	jsonNodes, err := json.Marshal(results.Nodes)
	if err != nil {
		return err
	}
	jsonTotals, err := json.Marshal(results.Totals)
	if err != nil {
		return err
	}

	benchmarks := &db.TblClusterBenchmarks{
		BenchmarksUid:   s.Config.BenchmarksId,
		CisVersion:      results.Version,
		DetectedVersion: results.DetectedVersion,
		Results:         string(jsonNodes),
		Totals:          string(jsonTotals),
		State:           2,
		Updater:         Updater,
	}
	err = s.updateDB(s.DB.Save(benchmarks))
	if err != nil {
		return err
	}
	return nil
}

func (s *Summarize) updateDBFailure(reason string) error {
	benchmarks := &db.TblClusterBenchmarks{
		BenchmarksUid: s.Config.BenchmarksId,
		State:         3,
		Reason:        reason,
		Updater:       Updater,
	}
	err := s.updateDB(s.DB.Model(benchmarks).Updates(benchmarks))
	if err != nil {
		logger.Warnf("error saving the failure data to the database: %v ", err)
		return err
	}

	return nil
}

func (s *Summarize) updateDB(tx *gorm.DB) error {
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error updating database. RowsAffected is 0.")
	}

	return nil
}

func sumTotals(totals []model.Total) model.Total {
	sum := model.Total{
		Pass: 0, Fail: 0, Warn: 0, Info: 0,
	}
	for _, total := range totals {
		sum.Pass += total.Pass
		sum.Fail += total.Fail
		sum.Warn += total.Warn
		sum.Info += total.Info
	}

	return sum
}
