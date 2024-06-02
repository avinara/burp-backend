package database

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/burp-backend/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dbConfig *config.DatabaseConfig) (*sql.DB, error) {
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.ConnectionString + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbConfig.Database + "?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
		return nil, err
	}

	DB = db
	return DB, nil
}
