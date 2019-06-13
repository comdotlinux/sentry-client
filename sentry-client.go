package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

const environmentTokenKey = "OAUTH_TOKEN"
const apiBaseURL = ""

func init() {
	if token, present := os.LookupEnv(environmentTokenKey); !present || token == "" {
		log.Fatalf("Please set a environment variable named %s created on sentry", environmentTokenKey)
	}
}

func main() {

}
