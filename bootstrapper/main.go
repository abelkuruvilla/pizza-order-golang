package bootstrapper

import (
	"log"
	utils "pizza-delivery/apputil"

	"github.com/spf13/viper"
)

func StartUp() {
	utils.InitLog()
}

type configuration struct {
	Server, NotifyFile string
}

var AppConfig configuration

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Config file not found:", err)
	}

	AppConfig = configuration{}
	AppConfig.Server = viper.GetString("development.Server")
	AppConfig.NotifyFile = viper.GetString("development.NotifyFile")
}
