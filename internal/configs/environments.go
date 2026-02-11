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
	JWT_SECRET           Environment
	JWT_EXPIRES_IN_HOURS Environment
	FILES_PATH           Environment
	GOOGLE_CLIENT_ID     Environment
	GOOGLE_CLIENT_SECRET Environment
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
		JWT_SECRET:           Environment(os.Getenv("JWT_SECRET")).orPanic(),
		JWT_EXPIRES_IN_HOURS: Environment(os.Getenv("JWT_EXPIRES_IN_HOURS")).orPanic(),
		FILES_PATH:           Environment(os.Getenv("FILES_PATH")).orDefault("static"),
		GOOGLE_CLIENT_ID:     Environment(os.Getenv("GOOGLE_CLIENT_ID")).orPanic(),
		GOOGLE_CLIENT_SECRET: Environment(os.Getenv("GOOGLE_CLIENT_SECRET")).orPanic(),
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
