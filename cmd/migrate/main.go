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

	dbConn.Migrator().DropTable(
		&model.Vtuber{},
		&model.Group{},
		"memberships",
	)

	dbConn.AutoMigrate(
		&model.Vtuber{},
		&model.Group{},
	)
}
