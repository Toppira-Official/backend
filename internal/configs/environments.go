package configs

import (
	"log"
	"os"
	"slices"

	"github.com/joho/godotenv"
)

type Environment string

type Environments struct {
	PORT                 Environment
	MODE                 Environment
	POSTGRES_USER        Environment
	POSTGRES_PASSWORD    Environment
	POSTGRES_HOST        Environment
	POSTGRES_PORT        Environment
	POSTGRES_DB          Environment
	REDIS_HOST           Environment
	REDIS_PORT           Environment
	REDIS_PASSWORD       Environment
	REDIS_DB             Environment
	JWT_SECRET           Environment
	JWT_EXPIRES_IN_HOURS Environment
	FILES_PATH           Environment
	GOOGLE_CLIENT_ID     Environment
	GOOGLE_CLIENT_SECRET Environment
	GOOGLE_REDIRECT_URL  Environment
}

func LoadEnvironmentsFromEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("cannot load environments from given env file")
	}
}

func GetEnvironments() Environments {
	return Environments{
		PORT:                 Environment(os.Getenv("PORT")).orDefault("3000"),
		MODE:                 Environment(os.Getenv("MODE")).orPanic().mustBeIn("develop", "production"),
		POSTGRES_USER:        Environment(os.Getenv("POSTGRES_USER")).orPanic(),
		POSTGRES_PASSWORD:    Environment(os.Getenv("POSTGRES_PASSWORD")).orPanic(),
		POSTGRES_HOST:        Environment(os.Getenv("POSTGRES_HOST")).orDefault("localhost"),
		POSTGRES_PORT:        Environment(os.Getenv("POSTGRES_PORT")).orDefault("5432"),
		POSTGRES_DB:          Environment(os.Getenv("POSTGRES_DB")).orPanic(),
		REDIS_HOST:           Environment(os.Getenv("REDIS_HOST")).orDefault("localhost"),
		REDIS_PORT:           Environment(os.Getenv("REDIS_PORT")).orDefault("6379"),
		REDIS_PASSWORD:       Environment(os.Getenv("REDIS_PASSWORD")).orPanic(),
		REDIS_DB:             Environment(os.Getenv("REDIS_DB")).orDefault("0"),
		JWT_SECRET:           Environment(os.Getenv("JWT_SECRET")).orPanic(),
		JWT_EXPIRES_IN_HOURS: Environment(os.Getenv("JWT_EXPIRES_IN_HOURS")).orPanic(),
		FILES_PATH:           Environment(os.Getenv("FILES_PATH")).orDefault("static"),
		GOOGLE_CLIENT_ID:     Environment(os.Getenv("GOOGLE_CLIENT_ID")).orPanic(),
		GOOGLE_CLIENT_SECRET: Environment(os.Getenv("GOOGLE_CLIENT_SECRET")).orPanic(),
		GOOGLE_REDIRECT_URL:  Environment(os.Getenv("GOOGLE_REDIRECT_URL")).orPanic(),
	}
}

func (env Environment) String() string {
	return string(env)
}

func (env Environment) orDefault(defaultValue string) Environment {
	if env == "" {
		return Environment(defaultValue)
	}

	return env
}

func (env Environment) orPanic() Environment {
	if env == "" {
		panic("env not set")
	}

	return env
}

func (env Environment) mustBeIn(allowedValues ...string) Environment {
	doesContain := slices.Contains(allowedValues, string(env))

	if doesContain {
		return env
	}

	panic("env is not allowed")
}
