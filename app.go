package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"gopkg.in/boj/redistore.v1"
	"net/http"
	"os"
	"sotru-web/controllers"
	"sotru-web/models"
	"sotru-web/utils"
)

var (
	config utils.Config
	db     *sql.DB
	store  *redistore.RediStore
)

func main() {
	var err error

	// setting up logger
	lf, err := os.OpenFile("full.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
		return
	}
	logger.Init("Logger", true, true, lf)

	logger.Info("Starting up Sea of Thieves RU webserver")

	config, err = utils.ReadConfig("config.json")
	if err != nil {
		logger.Fatal("Unable to load configuration file. Error: " + err.Error())
	}
	controllers.UseConfig(config)

	// database connection
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		config.DBUser, config.DBPassword, config.DBHost, config.DBName))
	if err != nil {
		logger.Fatal("Unable to connect to the database. Error: " + err.Error())
	}
	defer db.Close()
	models.UseDB(db)

	logger.Info("MySQL connection established")

	// session store
	store, err = redistore.NewRediStore(10, "tcp", ":6379", "", []byte(config.SessionSecret))
	if err != nil {
		logger.Fatal("Unable to connect to Redis session store. Error: " + err.Error())
	}
	defer store.Close()
	store.SetMaxAge(14 * 24 * 3600)
	controllers.UseStore(store)

	logger.Info("Redis connection established")

	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.LoginController)
	http.Handle("/", r)

	logger.Info("HTTP-server listening at 9900")
	err = http.ListenAndServe(":9900", nil)
	if err != nil {
		logger.Fatal("Unable to run HTTP server: " + err.Error())
	}
}
