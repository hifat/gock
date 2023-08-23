package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Env EnvConfig
	DB  DBConfig
}

type EnvConfig struct {
	AppHost string `envconfig:"APP_HOST"`
	AppName string `envconfig:"APP_NAME"`
	AppPort string `envconfig:"APP_PORT"`
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	Name     string `envconfig:"DB_NAME"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
}

func (c *AppConfig) Init() {
	envconfig.MustProcess("", &c.Env)
	envconfig.MustProcess("", &c.DB)
}

func LoadAppConfig() *AppConfig {
	env, ok := os.LookupEnv("ENV")
	if ok && env != "" {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		err := godotenv.Load(fmt.Sprintf("%v/.env.%v", basePath, env))
		if err != nil {
			err = godotenv.Load()
			if err != nil {
				panic(err)
			}
		}
	} else {
		godotenv.Load()
	}

	appCfg := AppConfig{}
	appCfg.Init()
	return &appCfg
}
