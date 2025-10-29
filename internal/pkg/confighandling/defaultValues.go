package confighandling

import (
	"os"
	"strconv"
	"time"
)

func defValString(checkForEnvName string, ifNotExistsVal string) string {
	if len(os.Getenv(checkForEnvName)) != 0 {
		return os.Getenv(checkForEnvName)
	} else {
		return ifNotExistsVal
	}
}
func defValBool(checkForEnvName string, ifNotExistsVal bool) bool {
	if len(os.Getenv(checkForEnvName)) != 0 {
		return os.Getenv(checkForEnvName) == "true"
	} else {
		return ifNotExistsVal
	}
}

func defValTimeDuration(checkForEnvName string, ifNotExistsVal time.Duration, modifier time.Duration) time.Duration {
	if len(os.Getenv(checkForEnvName)) != 0 {
		d, err := strconv.Atoi(os.Getenv(checkForEnvName))
		if err != nil {
			return ifNotExistsVal
		}
		return time.Duration(d) * modifier
	}
	return ifNotExistsVal
}
