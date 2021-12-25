package main

import (
	"log"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/k1/go-blog/global"
	"github.com/k1/go-blog/internal/models"
	"github.com/k1/go-blog/internal/routers"
	"github.com/k1/go-blog/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	initConfig()
	initLogger()
	initDatabase()
}

// @title 博客系统
// @version 1.0
// @description my first go web proj
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    time.Duration(global.Config.ReadTimeout),
		WriteTimeout:   time.Duration(global.Config.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func initConfig() {
	_, err := toml.DecodeFile(global.ConfigPath, &global.Config)
	if err != nil {
		panic(err)
	}
}

func initLogger() {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.Config.App.LogPath + "/" + global.Config.App.LogFile,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
}

func initDatabase() {
	var err error
	global.DBEngine, err = models.NewDBEngine(&global.Config.Mysql)
	if err != nil {
		panic(err)
	}
}
