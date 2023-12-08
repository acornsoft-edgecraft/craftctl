package db

import (
	"fmt"
	"time"

	"github.com/edgecraft/edge-benchmarks/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Type         string `yaml:"type"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"database_name"`
	SchemaName   string `yaml:"schema_name"`
	UserName     string `yaml:"username"`
	Password     string `yaml:"password"`
}

func New(conf *Config) (*gorm.DB, error) {
	db, err := connection(conf)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connection(conf *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Host, conf.UserName, conf.Password, conf.DatabaseName, conf.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Warnf("db connection failed")
		return nil, err
	}

	logger.Info("db connection completed.")
	return db, nil
}

type TblClusterBenchmarks struct {
	BenchmarksUid   string `gorm:"primaryKey"`
	CisVersion      string
	DetectedVersion string
	Results         string
	Totals          string
	State           int
	Reason          string
	Updater         string
	UpdatedAt       time.Time
}
