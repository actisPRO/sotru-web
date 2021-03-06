package main

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"gopkg.in/boj/redistore.v1"
	"net/http"
	"os"
	"sotru-web/cache"
	"sotru-web/controllers"
	"sotru-web/models"
	"sotru-web/utils"
)

var (
	config utils.Config
	db     *sql.DB
	store  *redistore.RediStore
	bot    *discordgo.Session
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
	models.UseConfig(config)
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
		logger.Fatal("Unable to connect to Redis session storage. Error: " + err.Error())
	}
	defer store.Close()
	store.SetMaxAge(14 * 24 * 3600)
	controllers.UseStore(store)

	logger.Info("Redis connection established")

	// discord bot connection
	bot, err = discordgo.New("Bot " + config.DiscordBotToken)
	if err != nil {
		logger.Fatal("Unable to create Discord session. Error: " + err.Error())
	}
	err = bot.Open()
	if err != nil {
		logger.Fatal("Unable to open Discord connection. Error: " + err.Error())
	}
	models.UseDiscord(config.DiscordGuild, bot)

	cache.UseDiscord(config.DiscordGuild, bot)
	cache.Connection, err = redis.Dial("tcp", ":6379")
	if err != nil {
		logger.Fatal("Unable to connect to Redis (cache). Error: " + err.Error())
	}

	logger.Info("Discord connection established")

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.IndexController)
	r.HandleFunc("/login", controllers.LoginController)
	r.HandleFunc("/logout", controllers.LogoutController)
	r.HandleFunc("/refresh", controllers.RefreshController)
	r.HandleFunc("/blacklist", controllers.BlacklistController)

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)
	http.Handle("/", r)

	logger.Info("HTTP-server listening at 9900")
	err = http.ListenAndServe(":9900", nil)
	if err != nil {
		logger.Fatal("Unable to run HTTP server: " + err.Error())
	}
}
