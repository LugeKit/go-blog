package global

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/k1/go-blog/configs"
	"github.com/k1/go-blog/pkg/logger"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const confPath = "configs/app.toml"

var (
	Logger *logger.Logger
	Config configs.Config
)

func Init() {
	InitConfig()
	InitLogger()
}

func InitConfig() {
	_, err := toml.DecodeFile(confPath, &Config)
	if err != nil {
		panic(err)
	}
}

func InitLogger() {
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  Config.App.LogPath + "/" + Config.App.LogFile,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
}
