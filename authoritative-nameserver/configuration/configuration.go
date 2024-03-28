package configuration

type Configuration struct {
	Redis struct {
		Host     string
		Port     uint16
		Password string
	}
}

func LoadConfiguration() Configuration {
	return Configuration{
		Redis: struct {
			Host     string
			Port     uint16
			Password string
		}{Host: "localhost", Port: 6379, Password: ""},
	}
}
