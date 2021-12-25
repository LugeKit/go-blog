package global

import (
	"github.com/k1/go-blog/configs"
	"github.com/k1/go-blog/pkg/logger"
)

const ConfigPath = "configs/app.toml"

var (
	Logger *logger.Logger
	Config configs.Config
)
