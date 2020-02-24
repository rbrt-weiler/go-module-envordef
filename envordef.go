// Package envordef returns a value from environment or, if the environment variable does not exist, a defined default value.
// It is intended to be used in conjunction with the flags module in order to create environment aware applications.
package envordef

import (
	"os"
	"strconv"
)

const (
	moduleName    string = "envordef"
	moduleVersion string = "0.1.0"
	moduleURL     string = "https://gitlab.com/rbrt-weiler/go-module-envordef"
)

/*
########   #######   #######  ##       ########    ###    ##    ##  ######
##     ## ##     ## ##     ## ##       ##         ## ##   ###   ## ##    ##
##     ## ##     ## ##     ## ##       ##        ##   ##  ####  ## ##
########  ##     ## ##     ## ##       ######   ##     ## ## ## ##  ######
##     ## ##     ## ##     ## ##       ##       ######### ##  ####       ##
##     ## ##     ## ##     ## ##       ##       ##     ## ##   ### ##    ##
########   #######   #######  ######## ######## ##     ## ##    ##  ######
*/

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

/*
######## ##        #######     ###    ########  ######
##       ##       ##     ##   ## ##      ##    ##    ##
##       ##       ##     ##  ##   ##     ##    ##
######   ##       ##     ## ##     ##    ##     ######
##       ##       ##     ## #########    ##          ##
##       ##       ##     ## ##     ##    ##    ##    ##
##       ########  #######  ##     ##    ##     ######
*/

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

/*
#### ##    ## ########  ######
 ##  ###   ##    ##    ##    ##
 ##  ####  ##    ##    ##
 ##  ## ## ##    ##     ######
 ##  ##  ####    ##          ##
 ##  ##   ###    ##    ##    ##
#### ##    ##    ##     ######
*/

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

/*
##     ## #### ##    ## ########  ######
##     ##  ##  ###   ##    ##    ##    ##
##     ##  ##  ####  ##    ##    ##
##     ##  ##  ## ## ##    ##     ######
##     ##  ##  ##  ####    ##          ##
##     ##  ##  ##   ###    ##    ##    ##
 #######  #### ##    ##    ##     ######
*/

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

/*
 ######  ######## ########  #### ##    ##  ######    ######
##    ##    ##    ##     ##  ##  ###   ## ##    ##  ##    ##
##          ##    ##     ##  ##  ####  ## ##        ##
 ######     ##    ########   ##  ## ## ## ##   ####  ######
      ##    ##    ##   ##    ##  ##  #### ##    ##        ##
##    ##    ##    ##    ##   ##  ##   ### ##    ##  ##    ##
 ######     ##    ##     ## #### ##    ##  ######    ######
*/

// StringVal returns the value of environment variable envName as a string or the supplied default value defaultVal if the environment variable does not exist.
func StringVal(envName string, defaultVal string) string {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(envName)
	if envExists {
		retVal = envVal
	}
	return retVal
}
