package database

import (
	"fmt"
	"portfoyum-api/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the underlying database connection
var DB *gorm.DB

// Connect initiate the database connection and migrate all the tables
func Connect() {
	db, err := gorm.Open(postgres.Open(config.Settings.Database.ConnectionUri), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err.Error())
	}

	DB = db

	fmt.Println("[DATABASE]::CONNECTED")
}

// Migrate migrates all the database tables
func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
