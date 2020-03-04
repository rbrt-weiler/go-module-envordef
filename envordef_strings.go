package envordef

import (
	"os"
)

// StringVal returns the value of environment variable envName as a string or the supplied default value defaultVal if the environment variable does not exist.
func StringVal(envName string, defaultVal string) string {
	retVal := defaultVal
	envVal, envExists := os.LookupEnv(envName)
	if envExists {
		retVal = envVal
	}
	return retVal
}
