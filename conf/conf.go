package conf

import (
	"fmt"

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

var Conf Config

const confPath = "conf/app.toml"

func init() {
	_, err := toml.DecodeFile(confPath, &Conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(Conf)
}
