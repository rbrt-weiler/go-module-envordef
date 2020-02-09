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

// BoolVal returns the value of environment variable `name` as a bool or the supplied default value `defaultVal` if the environment variable does not exist.
func BoolVal(name string, defaultVal bool) bool {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(name)
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

// IntVal returns the value of environment variable `name` as an int or the supplied default value `defaultVal` if the environment variable does not exist.
func IntVal(name string, defaultVal int) int {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(name)
	if envExists {
		intVal, _ := strconv.Atoi(envVal)
		retVal = int(intVal)
	}
	return retVal
}

// UintVal returns the value of environment variable `name` as an uint or the supplied default value `defaultVal` if the environment variable does not exist.
func UintVal(name string, defaultVal uint) uint {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(name)
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

// StringVal returns the value of environment variable `name` as a string or the supplied default value `defaultVal` if the environment variable does not exist.
func StringVal(name string, defaultVal string) string {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(name)
	if envExists {
		retVal = envVal
	}
	return retVal
}
