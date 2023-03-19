package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		var err error

		dsn := getPostgresDSN()
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if nil != err {
			fmt.Println("Failed to create DB Connection ", err.Error())
		}

	}

	return db
}

func getPostgresDSN() string {
	return fmt.Sprintf(GetValue(DATABASE_CONNECTION_STRING), GetValue(DATABASE_HOST), GetValue(DATABASE_USER), GetValue(DATABASE_PASS), GetValue(DATABASE_NAME), GetValue(DATABASE_PORT), GetValue(DATABASE_SSL))
}
