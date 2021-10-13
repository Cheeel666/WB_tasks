package main

import (
	"task1/config"
	"task1/internal/server"

	"github.com/sirupsen/logrus"
)

const configPath = "config/config.json"

// Usage:
// GET: http://localhost:8080/api/v1/about
// GET: http://localhost:8080/api/v1/user/:userid
func main() {
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}
	server.SetupServer(*cfg).Run()
}
