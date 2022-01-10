package config

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/t1732/vsercher/internal/domain/model"
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

func init() {
	conn, err := NewDB()
	if err != nil {
		panic(err)
	}

	Migrate(conn)
	Seed(conn)
}

func NewDB() (*gorm.DB, error) {
	config := dbConfig{
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
	if err = t.Execute(&b, config); err != nil {
		return nil, err
	}

	logger, err := newLogger()
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

func newLogger() (logger.Interface, error) {
	file, err := os.OpenFile(RootPath+"/log/gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return logger.New(
		log.New(file, "\r\n", log.LstdFlags|log.Lshortfile),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	), nil
}

func Migrate(conn *gorm.DB) {
	if conn.Migrator().HasTable(&model.Vtuber{}) {
		conn.Migrator().DropTable(&model.Vtuber{}, &model.Group{}, &model.Membership{})
	}
	conn.AutoMigrate(&model.Vtuber{}, &model.Group{}, &model.Membership{})
}

func Seed(conn *gorm.DB) {
	conn.Create(&model.Vtuber{Name: "キズナアイ"})
	conn.Create(&model.Vtuber{Name: "兎鞠まり"})
	conn.Create(
		&model.Group{
			Name: "Re:AcT",
			Memberships: []model.Membership{
				{Vtuber: model.Vtuber{Name: "かしこまり"}},
			},
		},
	)
	conn.Create(
		&model.Group{
			Name: "おめがシスターズ",
			Memberships: []model.Membership{
				{Vtuber: model.Vtuber{Name: "おめがレイ"}},
				{Vtuber: model.Vtuber{Name: "おめがリオ"}},
			},
		},
	)
	conn.Create(
		&model.Group{
			Name: "にじさんじ",
			Memberships: []model.Membership{
				{Vtuber: model.Vtuber{Name: "月ノ美兎"}},
				{Vtuber: model.Vtuber{Name: "本間ひまわり"}},
			},
		},
	)
	conn.Create(
		&model.Group{
			Name: "ホロライブ",
			Memberships: []model.Membership{
				{Vtuber: model.Vtuber{Name: "ときのそら"}},
				{Vtuber: model.Vtuber{Name: "兎田ぺこら"}},
			},
		},
	)
}
