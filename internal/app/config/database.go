package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDatabaseConnection() (db *sql.DB, err error) {

	InitLoadConfiguration()
	host := viper.GetString("postgres.host")
	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	dbname := viper.GetString("postgres.dbname")
	connMaxIdleTime := viper.GetDuration("database.connMaxIdleTime")
	connMaxLifeTime := viper.GetDuration("database.connMaxLifeTime")
	maxIdleConn := viper.GetInt("database.maxIdleConn")
	maxOpenConn := viper.GetInt("database.maxOpenConn")

	connString := fmt.Sprintf("user=%s password=%s host=%s dbname=%s", user, password, host, dbname)
	db, err = sql.Open("postgres", connString)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}

	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	db.SetConnMaxLifetime(connMaxLifeTime * time.Second)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	return
}
