package setting

import (
	"os"
	"strconv"
)

type Config struct {
	Port       string
	Address    string
	DBHost     string
	DBName     string
	DBUsername string
	DBPassword string
	DBTimeZone string
	JWTExpire  int
	JWTSecret  int
	JWTIssuer  string
	LogPath    string
	LogExpire  string
	LogRotate  string
}

func getEnvStr(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func getEnvInt(key string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultValue
	}

	return value
}

var CONFIG = Config{
	Port:    getEnvStr("PORT", "8080"),
	Address: getEnvStr("ADDRESS", "0.0.0.0"),
}
