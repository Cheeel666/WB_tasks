package main

import "task1/internal/server"

// Usage:
// GET: http://localhost:8080/api/v1/about
// GET: http://localhost:8080/api/v1/user/:userid
func main() {
	server.SetupServer().Run()
}
