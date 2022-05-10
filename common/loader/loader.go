package loader

import (
	"encoding/json"
	"fmt"
	"github.com/msbranco/goconfig"
	"os"
	"path/filepath"
)

const (
	ConfPathEnv = "SERVER_CONF_PATH"
)

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

var (
	configPath   string
	appConfigMap map[string]*goconfig.ConfigFile
)

func init() {
	var err error
	var workPath string
	workPath = os.Getenv(ConfPathEnv)
	if len(workPath) == 0 {
		workPath, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	configPath = filepath.Join(workPath, "config")
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

func LoadElasticSearchConfig(config interface{}, filename string) {
	var err error
	var decoder *json.Decoder

	file := OpenFile(filename)
	defer func() {
		_ = file.Close()
	}()

	decoder = json.NewDecoder(file)
	if err = decoder.Decode(config); err != nil {
		msg := fmt.Sprintf("Decode json fail for config file at %s. Error: %v", filename, err)
		panic(msg)
	}
}

func OpenFile(filename string) *os.File {
	fullPath := filepath.Join(configPath, filename)

	var file *os.File
	var err error

	if file, err = os.Open(fullPath); err != nil {
		msg := fmt.Sprintf("Can not load config at %s. Error: %v", fullPath, err)
		panic(msg)
	}

	return file
}
