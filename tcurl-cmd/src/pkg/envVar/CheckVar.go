package envVar

import (
	"os"
	"strconv"
)

type DefaultEnv struct {
	Key   string
	Value string
}

func SetDefaultEnv(envs []DefaultEnv) {
	for _, v := range envs {
		GetSetString(v.Key, v.Value)
	}
}

func GetSetString(key string, value string) {
	if os.Getenv(key) == "" {
		_ = os.Setenv(key, value) // 默认认为不会出错
	}
}

func GetEnvInt(key string) int {
	data, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return -999
	}
	return data
}

func GetEnvString(key string) string {
	return os.Getenv(key)
}

func GetEnvBool(key string) bool {
	if os.Getenv(key) == "true" || os.Getenv(key) == "True" || os.Getenv(key) == "TRUE" {
		return true
	} else {
		return false
	}
}
