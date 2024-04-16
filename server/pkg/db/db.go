package db

import (
	"fmt"

	"github.com/danielsobrado/ainovelprompter/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() (*gorm.DB, error) {
	dbConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.GetString("db.host"),
		config.GetString("db.port"),
		config.GetString("db.user"),
		config.GetString("db.name"),
		config.GetString("db.password"),
	)

	db, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
