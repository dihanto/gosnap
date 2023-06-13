package config

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewDb() (db *sql.DB, err error) {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.port")
	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	dbname := viper.GetString("postgres.dbname")
	connMaxIdleTime := viper.GetDuration("database.connMaxIdleTime")
	connMaxLifeTime := viper.GetDuration("database.connMaxLifeTime")
	maxIdleConn := viper.GetInt("database.maxIdleConn")
	maxOpenConn := viper.GetInt("database.maxOpenConn")

	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	db.SetConnMaxLifetime(connMaxLifeTime * time.Second)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	return
}
