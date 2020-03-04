package envordef

import (
	"os"
	"strconv"
)

// IntVal returns the value of environment variable envName as an int or the supplied default value defaultVal if the environment variable does not exist.
func IntVal(envName string, defaultVal int) int {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if intVal, parseErr := strconv.ParseInt(envVal, 10, 0); parseErr == nil {
			retVal = int(intVal)
		}
	}
	return retVal
}

// Int8Val returns the value of environment variable envName as an int8 or the supplied default value defaultVal if the environment variable does not exist.
func Int8Val(envName string, defaultVal int8) int8 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if intVal, parseErr := strconv.ParseInt(envVal, 10, 8); parseErr == nil {
			retVal = int8(intVal)
		}
	}
	return retVal
}

// Int16Val returns the value of environment variable envName as an int16 or the supplied default value defaultVal if the environment variable does not exist.
func Int16Val(envName string, defaultVal int16) int16 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if intVal, parseErr := strconv.ParseInt(envVal, 10, 16); parseErr == nil {
			retVal = int16(intVal)
		}
	}
	return retVal
}

// Int32Val returns the value of environment variable envName as an int32 or the supplied default value defaultVal if the environment variable does not exist.
func Int32Val(envName string, defaultVal int32) int32 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if intVal, parseErr := strconv.ParseInt(envVal, 10, 32); parseErr == nil {
			retVal = int32(intVal)
		}
	}
	return retVal
}

// Int64Val returns the value of environment variable envName as an int64 or the supplied default value defaultVal if the environment variable does not exist.
func Int64Val(envName string, defaultVal int64) int64 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if intVal, parseErr := strconv.ParseInt(envVal, 10, 64); parseErr == nil {
			retVal = int64(intVal)
		}
	}
	return retVal
}
