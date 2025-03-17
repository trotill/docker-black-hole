package env

import (
	"flag"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Port                 string `env:"PORT" envDefault:"9080"`
	Docker               int    `env:"DOCKER" envDefault:"1"`
	ExecuteFromUser      string `env:"EXECUTE_FROM_USER" envDefault:"root"`
	ExecuteMaxTimeoutSec int    `env:"EXECUTE_MAX_TIMEOUT_SEC" envDefault:"600"`
	ScriptPath           string `env:"SCRIPT_PATH" envDefault:"scripts"`
	AllowAbsoluteMode    int    `env:"ALLOW_ABSOLUTE_MODE" envDefault:"0"`
	ShellPath            string `env:"SHELL_PATH" envDefault:"bash"`
	DisableLogs          int    `env:"DISABLE_LOGS" envDefault:"0"`
}

var cfg = Config{}
var loaded = false

func load() {
	loaded = true
	envFilePath := flag.String("env", ".env", "Config file to load")
	flag.Parse()
	err := godotenv.Load(*envFilePath)
	if err != nil {
		log.Println("Config file not found. System environment variables are used. For set config file use option env, example [docker-black-hole --env=.env.prod]")

	} else {
		log.Printf("Loading env variables from %s", *envFilePath)

	}

	if err := env.Parse(&cfg); err != nil {
		log.Printf("Error parsing environment variables: %v\n", err)
	}
}

func GetEnv() Config {
	if !loaded {
		load()
	}
	return cfg
}
