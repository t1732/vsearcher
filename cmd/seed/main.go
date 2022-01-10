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
			Memberships: []model.Membership{
				{Vtuber: model.Vtuber{Name: "かしこまり"}},
			},
		},
	)
	dbConn.Create(
		&model.Group{
			Name: "おめがシスターズ",
			Memberships: []model.Membership{
				{Vtuber: model.Vtuber{Name: "おめがレイ"}},
				{Vtuber: model.Vtuber{Name: "おめがリオ"}},
			},
		},
	)
	dbConn.Create(
		&model.Group{
			Name: "にじさんじ",
			Memberships: []model.Membership{
				{Vtuber: model.Vtuber{Name: "月ノ美兎"}},
				{Vtuber: model.Vtuber{Name: "本間ひまわり"}},
			},
		},
	)
	dbConn.Create(
		&model.Group{
			Name: "ホロライブ",
			Memberships: []model.Membership{
				{Vtuber: model.Vtuber{Name: "ときのそら"}},
				{Vtuber: model.Vtuber{Name: "兎田ぺこら"}},
			},
		},
	)
}
