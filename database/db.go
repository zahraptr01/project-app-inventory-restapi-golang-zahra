package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDB() *sql.DB {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	name := viper.GetString("DB_NAME")
	ssl := viper.GetString("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, name, ssl)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	return db
}
