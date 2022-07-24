package config

type Config struct {
	Host           string
	Port           string
	DBName         string
	CollectionName string
	UserName       string
	Password       string
	MaxPageLimit   int64
}

var EnvConfig = map[string]Config{

	"local": {
		Host:     "localhost",
		Port:     "5672",
		UserName: "user",
		Password: "password",
	},
}
