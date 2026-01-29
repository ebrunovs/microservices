package config

import (
    "log"
    "os"
    "strconv"
)

func GetEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func GetEnvAsInt(key string, defaultValue int) int {
    valueStr := os.Getenv(key)
    if valueStr == "" {
        return defaultValue
    }

    value, err := strconv.Atoi(valueStr)
    if err != nil {
        log.Printf("Erro ao converter %s para int, usando valor padrão: %d", key, defaultValue)
        return defaultValue
    }
    return value
}

func GetDataSourceURL() string {
    return GetEnv("DATA_SOURCE_URL", "root:password@tcp(localhost:3306)/shipping?parseTime=true")
}

func GetApplicationPort() int {
    return GetEnvAsInt("APPLICATION_PORT", 9090)
}