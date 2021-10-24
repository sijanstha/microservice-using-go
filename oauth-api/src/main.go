package main

import (
	"encoding/json"
	"fmt"
	"github.com/sijanstha/common-utils/src/logger"
	"github.com/sijanstha/common-utils/src/utils/errors"
	"github.com/sijanstha/oauth-api/src/app"
	"io/ioutil"
	"os"
	"strings"
)

func init() {
	logger.Info("Checking for environment variables")
	checkForEnvironmentVariables()
	logger.Info("Environment variables found to start application")
}

func main() {
	app.StartApplication()
}

func checkForEnvironmentVariables() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic("Missing configuration file")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string][]string
	json.Unmarshal(byteValue, &result)

	for _, value := range result["environmentVariables"] {
		if os.Getenv(value) == "" {
			logger.Error("Could not start service", errors.NewError(fmt.Sprintf("Missing envionment variable: %s", strings.ToUpper(value))))
			panic(fmt.Sprintf("Missing envionment variable: %s", strings.ToUpper(value)))
		}
	}
}
