package viperPractice

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadMysqlConfig(fileName string) {
	//v := viper.New()
	//v.SetConfigType("json")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/Users/yinkaiyi/personal/program/GoGoGo/viperPractice")
	viper.SetConfigName("mysql")
	//v.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.Get("a.b.c"))
}
