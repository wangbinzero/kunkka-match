package conf

import "github.com/spf13/viper"

var Gconfig Config

type Config struct {
	*viper.Viper
}

// 加载配置文件
func LoadConfig() {
	conf := Config{viper.New()}
	conf.SetConfigName("config")
	conf.AddConfigPath("./conf/")
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	Gconfig = conf
}
