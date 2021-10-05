package config

import (
	"log"

	"github.com/spf13/viper"
)

//## map file config.yaml
type data struct {
	//## Server
	ENV_MODE     string
	SERVER_PORT  string
	SECRET_KEY   string
	TOKEN_EXPIRE int
	//## Data base
	DATABASE_HOST          string
	DATABASE_PORT          int
	DATABASE_USER          string
	DATABASE_PASSWORD      string
	DATABASE_NAME          string
	DATABASE_AUTO_GEN_GORM bool
}

//## Server
type server struct {
	Env_Mode     string
	Port         string
	Secret_Key   string
	Token_Expire int
}

// //## Database
type database struct {
	Server                 string
	Port                   int
	User                   string
	Pass                   string
	DatabaseName           string
	DATABASE_AUTO_GEN_GORM bool
}

var config appConfig

type appConfig struct {
	Data data
}

func InitialConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("")
	v.AutomaticEnv()
	v.AddConfigPath(".")
	v.AddConfigPath("./")
	if err := v.ReadInConfig(); err != nil {
		log.Panicln(err)
		panic(err)
	}
	v.Unmarshal(&config)
}

func Server() server {
	conf := server{
		Env_Mode:     config.Data.ENV_MODE,
		Port:         config.Data.SERVER_PORT,
		Secret_Key:   config.Data.SECRET_KEY,
		Token_Expire: config.Data.TOKEN_EXPIRE,
	}
	return conf
}

func Database() database {
	conf := database{
		Server:                 config.Data.DATABASE_HOST,
		Port:                   config.Data.DATABASE_PORT,
		User:                   config.Data.DATABASE_USER,
		Pass:                   config.Data.DATABASE_PASSWORD,
		DatabaseName:           config.Data.DATABASE_NAME,
		DATABASE_AUTO_GEN_GORM: config.Data.DATABASE_AUTO_GEN_GORM,
	}
	return conf
}
