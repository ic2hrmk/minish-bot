package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetInt64Var(envVar string) (int64, error) {
	return strconv.ParseInt(os.Getenv(envVar), 10, 64)
}

func GetStringVar(envVar string) (string, error) {
	envVarValue := os.Getenv(envVar)
	if envVarValue == "" {
		return envVarValue, fmt.Errorf("env. var [%s] is empty", envVar)
	}

	return envVarValue, nil
}
