package config

import (
	"bytes"
	"html/template"
	"os"

	"github.com/t1732/vsercher/internal/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	dsn := b.String()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(conn *gorm.DB) {
	conn.AutoMigrate(&model.Vtuber{})
}

func Seed(conn *gorm.DB) {
	if conn.Migrator().HasTable(&model.Vtuber{}) {
		conn.Migrator().DropTable(&model.Vtuber{})
	}
	conn.Migrator().CreateTable(&model.Vtuber{})
	conn.Create(&model.Vtuber{Name: "キズナアイ"})
	conn.Create(&model.Vtuber{Name: "かしこまり"})
	conn.Create(&model.Vtuber{Name: "おめがシスターズ"})
	conn.Create(&model.Vtuber{Name: "兎田ぺこら"})
	conn.Create(&model.Vtuber{Name: "本間ひまわり"})
	conn.Create(&model.Vtuber{Name: "兎鞠まり"})
}
