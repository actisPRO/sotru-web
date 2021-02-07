package main

import (
	"./utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/logger"
	"os"
)

var (
	config utils.Config
	db     *sql.DB
	log    *logger.Logger
)

func main() {
	var err error

	// setting up logger
	lf, err := os.OpenFile("full.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
		return
	}
	log = logger.Init("Logger", true, true, lf)

	log.Info("Starting up Sea of Thieves RU webserver")

	config, err = utils.ReadConfig("config.json")
	if err != nil {
		log.Fatal("Unable to load configuration file. Error: " + err.Error())
	}

	// database connection
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		config.DBUser, config.DBPassword, config.DBHost, config.DBName))
	if err != nil {
		log.Fatal("Unable to connect to the database. Error: " + err.Error())
	}

	log.Info("Database connection established")
}
