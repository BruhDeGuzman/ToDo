package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"todo/config"
)

var (
	DB  *sql.DB
	err error
)

func InitDB() {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.EnvConfig.DBUsername,
		config.EnvConfig.DBPassword,
		config.EnvConfig.DBHost,
		config.EnvConfig.DBPort,
		config.EnvConfig.DBName)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}
