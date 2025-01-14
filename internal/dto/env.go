package dto

import (
	"fmt"
	"os"

	"go.arcalot.io/log"
)

// LookupEnvVar returns the value of an enviornment variable.
// lookupEnvVar will return an error if the enviornment variable is not set or empty.
// The error includes the registry url to destinguish which registry encountered the error if found.
func LookupEnvVar(registry_url string, key string, logger log.Logger) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		err := fmt.Errorf("%s environment variable not set to push to %s", key, registry_url)
		return "", err
	} else if len(val) == 0 {
		err := fmt.Errorf("%s environment variable empty to push to %s", key, registry_url)
		return "", err
	}
	return val, nil
}
