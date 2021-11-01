package main

import (
	"fmt"
	"net/http"
	"task1/config"
	"task1/internal/handlers"
	"task1/internal/server"

	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
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
	listenErr := make(chan error, 1)

	go func() {
		rout := router.New()
		var testUser handlers.TestUser
		rout.GET("/test", func(c *fasthttp.RequestCtx) {
			fmt.Println("OK")
		})
		rout.PUT("/test/user/block", cors(testUser.TestUserBlock))
		// rout.DELETE("/test/user/unBlock/:workerCount")

		listenErr <- fasthttp.ListenAndServe(cfg.TestPort, rout.Handler)
		fmt.Println("Test server listen on:", cfg.TestPort)
	}()
	server.SetupServer(*cfg).Run(cfg.Port)
}

func cors(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type,Accept")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "OPTIONS,POST,GET, PUT, DELETE")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		if string(ctx.Method()) == http.MethodOptions {
			ctx.Response.SetStatusCode(200)
			return
		}
		next(ctx)
	}
}
