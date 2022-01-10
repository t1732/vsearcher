package logger

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

type MysqlConfig struct {
	IsFile   bool
	FilePath string
	LogLevel logger.LogLevel
}

func NewLogger(c MysqlConfig) (logger.Interface, error) {
	lgg := log.New(os.Stdout, "\r\n", log.LstdFlags)
	cnf := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  c.LogLevel,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}

	if c.IsFile {
		file, err := os.OpenFile(c.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		lgg = log.New(file, "\r\n", log.LstdFlags)
		cnf.Colorful = false
	}

	return logger.New(lgg, cnf), nil
}
