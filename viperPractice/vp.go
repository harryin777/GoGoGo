package viperPractice

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadMysqlConfig(fileName string) {
	//v := viper.New()
	//v.SetConfigType("json")
	viper.AddConfigPath("/Users/yinkaiyi/Downloads")
	viper.SetConfigName("mysql")
	//v.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.Get("a"))
}
