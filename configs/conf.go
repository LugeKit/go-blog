package configs

import "time"

type Config struct {
	App      `toml:"app"`
	Server   `toml:"server"`
	Database `toml:"database"`
}

type App struct {
	PageSize  int    `toml:"page_size"`
	JWTSecret string `toml:"jwt_secret"`
	LogPath   string `toml:"log_path"`
	LogFile   string `toml:"log_file"`
}

type Server struct {
	RunMode      string `toml:"run_mode"`
	Port         string
	ReadTimeout  Duration `toml:"read_timeout"`
	WriteTimeout Duration `toml:"write_timeout"`
}

type Duration time.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	nd, err := time.ParseDuration(string(text))
	*d = Duration(nd)
	return err
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
