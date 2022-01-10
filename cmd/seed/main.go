package main

import (
	"github.com/t1732/vsearcher/internal/config"
	lgg "github.com/t1732/vsearcher/internal/config/logger"
	"github.com/t1732/vsearcher/internal/domain/model"
	"gorm.io/gorm/logger"
)

func main() {
	dbConn, err := config.NewDB(lgg.MysqlConfig{
		LogLevel: logger.Info,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	dbConn.Create(&model.Vtuber{Name: "キズナアイ"})
	dbConn.Create(&model.Vtuber{Name: "兎鞠まり"})
	dbConn.Create(
		&model.Group{
			Name: "Re:AcT",
			Vtubers: []model.Vtuber{
				{Name: "かしこまり"},
			},
		},
	)
	dbConn.Create(
		&model.Group{
			Name: "おめがシスターズ",
			Vtubers: []model.Vtuber{
				{Name: "おめがレイ"},
				{Name: "おめがリオ"},
			},
		},
	)
	dbConn.Create(
		&model.Group{
			Name: "にじさんじ",
			Vtubers: []model.Vtuber{
				{Name: "月ノ美兎"},
				{Name: "本間ひまわり"},
			},
		},
	)
	dbConn.Create(
		&model.Group{
			Name: "ホロライブ",
			Vtubers: []model.Vtuber{
				{Name: "ときのそら"},
				{Name: "兎田ぺこら"},
			},
		},
	)
}
