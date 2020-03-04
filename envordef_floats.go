package envordef

import (
	"os"
	"strconv"
)

// Float32Val returns the value of environment variable envName as a float32 or the supplied default value defaultVal if the environment variable does not exist.
func Float32Val(envName string, defaultVal float32) float32 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if floatVal, parseErr := strconv.ParseFloat(envVal, 32); parseErr == nil {
			retVal = float32(floatVal)
		}
	}
	return retVal
}

// Float64Val returns the value of environment variable envName as a float64 or the supplied default value defaultVal if the environment variable does not exist.
func Float64Val(envName string, defaultVal float64) float64 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if floatVal, parseErr := strconv.ParseFloat(envVal, 64); parseErr == nil {
			retVal = float64(floatVal)
		}
	}
	return retVal
}
