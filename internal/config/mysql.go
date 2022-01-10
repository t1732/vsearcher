package config

import (
	"bytes"
	"html/template"
	"os"

	"github.com/imdario/mergo"
	lgg "github.com/t1732/vsearcher/internal/config/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBname   string
}

func NewDB(c lgg.MysqlConfig) (*gorm.DB, error) {
	dbConfig := dbConfig{
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	}
	t, err := template.New("dsn").Parse("{{.Username}}:{{.Password}}@tcp({{.Host}}:{{.Port}})/{{.DBname}}?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err = t.Execute(&b, dbConfig); err != nil {
		return nil, err
	}

	loggerConfig := lgg.MysqlConfig{
		IsFile:   false,
		FilePath: App.RootPath.Join("/log/gorm.log"),
		LogLevel: logger.Error,
	}

	if err := mergo.Merge(&loggerConfig, c, mergo.WithOverride); err != nil {
		return nil, err
	}

	logger, err := lgg.NewLogger(loggerConfig)
	if err != nil {
		return nil, err
	}

	dsn := b.String()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // デフォルトのトランザクション機能を無効化
		PrepareStmt:            true, // プリペアードステートメントキャッシュ有効化
		Logger:                 logger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
