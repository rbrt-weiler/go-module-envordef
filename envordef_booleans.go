package envordef

import (
	"os"
	"strconv"
)

// BoolVal returns the value of environment variable envName as a bool or the supplied default value defaultVal if the environment variable does not exist.
func BoolVal(envName string, defaultVal bool) bool {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if boolVal, parseErr := strconv.ParseBool(envVal); parseErr == nil {
			retVal = boolVal
		}
	}
	return retVal
}
