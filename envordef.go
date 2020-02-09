// Package envordef returns a value from environment or, if the environment variable does not exist, a defined default value.
// It is intended to be used in conjunction with the flags module in order to create environment aware applications.
package envordef

import (
	"os"
	"strconv"
)

const (
	moduleName    string = "envordef"
	moduleVersion string = "0.1.0-dev"
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
	envVal, envExists := os.LookupEnv(envName)
	if envExists {
		retVal, _ = strconv.ParseBool(envVal)
	}
	return retVal
}

/*
##    ## ##     ## ##     ## ######## ########  ####  ######   ######
###   ## ##     ## ###   ### ##       ##     ##  ##  ##    ## ##    ##
####  ## ##     ## #### #### ##       ##     ##  ##  ##       ##
## ## ## ##     ## ## ### ## ######   ########   ##  ##        ######
##  #### ##     ## ##     ## ##       ##   ##    ##  ##             ##
##   ### ##     ## ##     ## ##       ##    ##   ##  ##    ## ##    ##
##    ##  #######  ##     ## ######## ##     ## ####  ######   ######
*/

// IntVal returns the value of environment variable envName as an int or the supplied default value defaultVal if the environment variable does not exist.
func IntVal(envName string, defaultVal int) int {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(envName)
	if envExists {
		intVal, _ := strconv.Atoi(envVal)
		retVal = int(intVal)
	}
	return retVal
}

// Int8Val returns the value of environment variable envName as an int or the supplied default value defaultVal if the environment variable does not exist.
func Int8Val(envName string, defaultVal int8) int8 {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(envName)
	if envExists {
		intVal, _ := strconv.Atoi(envVal)
		retVal = int8(intVal)
	}
	return retVal
}

// Int16Val returns the value of environment variable envName as an int or the supplied default value defaultVal if the environment variable does not exist.
func Int16Val(envName string, defaultVal int16) int16 {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(envName)
	if envExists {
		intVal, _ := strconv.Atoi(envVal)
		retVal = int16(intVal)
	}
	return retVal
}

// Int32Val returns the value of environment variable envName as an int or the supplied default value defaultVal if the environment variable does not exist.
func Int32Val(envName string, defaultVal int32) int32 {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(envName)
	if envExists {
		intVal, _ := strconv.Atoi(envVal)
		retVal = int32(intVal)
	}
	return retVal
}

// Int64Val returns the value of environment variable envName as an int or the supplied default value defaultVal if the environment variable does not exist.
func Int64Val(envName string, defaultVal int64) int64 {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(envName)
	if envExists {
		intVal, _ := strconv.Atoi(envVal)
		retVal = int64(intVal)
	}
	return retVal
}

// UintVal returns the value of environment variable envName as an uint or the supplied default value defaultVal if the environment variable does not exist.
func UintVal(envName string, defaultVal uint) uint {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(envName)
	if envExists {
		intVal, _ := strconv.Atoi(envVal)
		retVal = uint(intVal)
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
