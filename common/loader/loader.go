package loader

import (
	"github.com/msbranco/goconfig"
	"path/filepath"
)

const (
	RunModeEnv    = "RUN_MODE"
	RunEnv        = "RUN_ENV"
	ConfPathEnv   = "SERVER_CONF_PATH"
	ServerChannel = "SERVER_CHANNEL"
	ServerRegion  = "SERVER_REGION"
)

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

const (
	AWSProductEnv = "aws_product_env"
	AWSDevelopEnv = "aws_develop_env"
)

var (
	configPath   string
	appConfigMap map[string]*goconfig.ConfigFile
)

func init() {
	appConfigMap = make(map[string]*goconfig.ConfigFile)
}

func LoadConfigFile(fileName string) *goconfig.ConfigFile {
	if configFile, exist := appConfigMap[fileName]; exist {
		return configFile
	} else {
		var err error
		fullPath := filepath.Join(configPath, fileName)
		configFile, err = goconfig.ReadConfigFile(fullPath)
		appConfigMap[fileName] = configFile
		if err != nil {
			panic(err)
		}
		return configFile
	}
}
