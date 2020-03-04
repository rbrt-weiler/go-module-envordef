package envordef

import (
	"os"
	"strconv"
)

// UintVal returns the value of environment variable envName as an uint or the supplied default value defaultVal if the environment variable does not exist.
func UintVal(envName string, defaultVal uint) uint {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if uintVal, parseErr := strconv.ParseUint(envVal, 10, 0); parseErr == nil {
			retVal = uint(uintVal)
		}
	}
	return retVal
}

// Uint8Val returns the value of environment variable envName as an uint8 or the supplied default value defaultVal if the environment variable does not exist.
func Uint8Val(envName string, defaultVal uint8) uint8 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if uintVal, parseErr := strconv.ParseUint(envVal, 10, 8); parseErr == nil {
			retVal = uint8(uintVal)
		}
	}
	return retVal
}

// Uint16Val returns the value of environment variable envName as an uint16 or the supplied default value defaultVal if the environment variable does not exist.
func Uint16Val(envName string, defaultVal uint16) uint16 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if uintVal, parseErr := strconv.ParseUint(envVal, 10, 16); parseErr == nil {
			retVal = uint16(uintVal)
		}
	}
	return retVal
}

// Uint32Val returns the value of environment variable envName as an uint32 or the supplied default value defaultVal if the environment variable does not exist.
func Uint32Val(envName string, defaultVal uint32) uint32 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if uintVal, parseErr := strconv.ParseUint(envVal, 10, 32); parseErr == nil {
			retVal = uint32(uintVal)
		}
	}
	return retVal
}

// Uint64Val returns the value of environment variable envName as an uint64 or the supplied default value defaultVal if the environment variable does not exist.
func Uint64Val(envName string, defaultVal uint64) uint64 {
	retVal := defaultVal
	if envVal, envExists := os.LookupEnv(envName); envExists {
		if uintVal, parseErr := strconv.ParseUint(envVal, 10, 64); parseErr == nil {
			retVal = uint64(uintVal)
		}
	}
	return retVal
}
