package configuration

import "github.com/spf13/viper"

type Configuration struct {
	Redis struct {
		Host     string
		Port     uint16
		Password string
	}
}

type configuration struct {
	redisHost     string `mapstructure:"REDIS_HOST"`
	redisPort     uint16 `mapstructure:"REDIS_PORT"`
	redisPassword string `mapstructure:"REDIS_PASSWORD"`
}

func NewConfiguration() Configuration {
	//loadedConfiguration, _ := loadConfiguration()
	return Configuration{
		Redis: struct {
			Host     string
			Port     uint16
			Password string
		}{Host: "localhost", Port: 6379, Password: ""},
	}
}

func loadConfiguration() (configuration configuration, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("configuration")
	viper.SetConfigType(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&configuration)
	return
}
