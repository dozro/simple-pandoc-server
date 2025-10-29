package confighandling

import "os"

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
