package mysql

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
	Host     string
	Port     string
	DBname   string
}

var Connection *gorm.DB

func NewConnection() (*sql.DB, error) {
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
	Connection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return Connection.DB()
}

func Migrate() {
	Connection.AutoMigrate(&model.Vtuber{})
}

func Seed() {
	if Connection.Migrator().HasTable(&model.Vtuber{}) {
		Connection.Migrator().DropTable(&model.Vtuber{})
	}
	Connection.Migrator().CreateTable(&model.Vtuber{})
	Connection.Create(&model.Vtuber{Name: "キズナアイ"})
	Connection.Create(&model.Vtuber{Name: "かしこまり"})
	Connection.Create(&model.Vtuber{Name: "おめがシスターズ"})
	Connection.Create(&model.Vtuber{Name: "兎田ぺこら"})
	Connection.Create(&model.Vtuber{Name: "本間ひまわり"})
	Connection.Create(&model.Vtuber{Name: "兎鞠まり"})
}
