package confighandling

import (
	"os"
	"strconv"
	"time"
)

func defValString(checkForEnvName string, ifNotExistsVal string) string {
	envVal := os.Getenv(checkForEnvName)
	if len(envVal) != 0 {
		return envVal
	}
	return ifNotExistsVal
}
func defValBool(checkForEnvName string, ifNotExistsVal bool) bool {
	envVal := os.Getenv(checkForEnvName)
	if len(envVal) != 0 {
		return envVal == "true"
	}
	return ifNotExistsVal
}

func defValTimeDuration(checkForEnvName string, ifNotExistsVal time.Duration, modifier time.Duration) time.Duration {
	envVal := os.Getenv(checkForEnvName)
	if len(envVal) != 0 {
		d, err := strconv.Atoi(envVal)
		if err != nil {
			return ifNotExistsVal
		}
		return time.Duration(d) * modifier
	}
	return ifNotExistsVal
}
