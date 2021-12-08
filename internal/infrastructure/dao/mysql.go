package dao

import (
	"bytes"
	"database/sql"
	"html/template"
	"os"

	"github.com/t1732/vsercher/internal/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbConfig struct {
	Username string
	Password string
	Host string
	Port string
	DBname string
}

var db *gorm.DB

func NewDB() (*sql.DB, error) {
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
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db.DB()
}

func DB() *gorm.DB {
	return db
}

func Migrate() {
	db.AutoMigrate(&model.Vtuber{})
}

func Seed() {
	if db.Migrator().HasTable(&model.Vtuber{}) {
		db.Migrator().DropTable(&model.Vtuber{})
	}
	db.Migrator().CreateTable(&model.Vtuber{})
	db.Create(&model.Vtuber{Name: "キズナアイ"})
	db.Create(&model.Vtuber{Name: "かしこまり"})
	db.Create(&model.Vtuber{Name: "おめがシスターズ"})
	db.Create(&model.Vtuber{Name: "兎田ぺこら"})
	db.Create(&model.Vtuber{Name: "本間ひまわり"})
	db.Create(&model.Vtuber{Name: "兎鞠まり"})
}
