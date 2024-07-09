package utils

import "os"

var env envInterface = realEnv{}

func setEnv(e envInterface) {
	env = e
}

type envInterface interface {
	Getenv(key string) string
}

type realEnv struct{}

func (r realEnv) Getenv(key string) string {
	return os.Getenv(key)
}

func IsProduction() bool {
	return env.Getenv("APP_ENV") == "production"
}
