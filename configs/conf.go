package configs

import (
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	RunMode  string `toml:"run_mode"`
	App      `toml:"app"`
	Server   `toml:"server"`
	Database `toml:"database"`
}

type App struct {
	PageSize  int    `toml:"page_size"`
	JWTSecret string `toml:"jwt_secret"`
}

type Server struct {
	Port         string
	ReadTimeout  int `toml:"read_timeout"`
	WriteTimeout int `toml:"write_timeout"`
}

type Database struct {
	Mysql `toml:"mysql"`
}

type Mysql struct {
	User        string
	Password    string
	Host        string
	Port        string
	DBName      string `toml:"dbname"`
	TablePrefix string `toml:"table_prefix"`
}

var (
	Conf Config

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

const confPath = "configs/app.toml"

func Init() {
	_, err := toml.DecodeFile(confPath, &Conf)
	if err != nil {
		panic(err)
	}
	ReadTimeout = time.Duration(Conf.ReadTimeout) * time.Second
	WriteTimeout = time.Duration(Conf.WriteTimeout) * time.Second
}
