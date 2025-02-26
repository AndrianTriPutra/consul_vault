package main

import (
	"context"
	"fmt"
	"log"
	"managenv/app"
	"managenv/pkg/env"
	"managenv/pkg/logger"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

// ENV=development HOST=localhost:8500 FILE=api go run .

func init() {
	log.Println("==================================")
	log.Println("app    :manageEnv")
	log.Println("version:v1")
	log.Println("release:25.02.25")
	log.Println("==================================")
}

func main() {
	server := os.Getenv("ENV")
	if len(server) == 0 {
		server = "development"
	}
	host := os.Getenv("HOST")
	if len(host) == 0 {
		host = "localhost:8500"
	}

	file := os.Getenv("FILE")
	if len(file) == 0 {
		file = "api"
	}

	managenv := env.Setting{
		Host:  "http://" + host,
		Token: "",
		Path:  fmt.Sprintf("/v1/kv/smartcity/%s/%s", server, file),
	}

	environment, err := env.ReadEnv_Consul(managenv)
	if err != nil {
		logger.Level("fatal", "ReadEnv_Consul", err.Error())
	}

	apps := app.Setting{
		Conf: managenv,
		Env:  environment,
	}
	runtime.GOMAXPROCS(int(environment.Internal.Core))
	application := app.NewApp(apps)
	log.Println(" ================== [RUN] ================== ")

	ctx := context.Background()
	wg := new(sync.WaitGroup)
	stoped := make(chan os.Signal, 1)
	signal.Notify(stoped,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	wg.Add(2)

	go func() {
		defer wg.Done()
		runtime.Gosched()
		err := application.ReadEnv(ctx)
		if err != nil {
			logger.Level("fatal", "main", "exit ReadEnv:"+err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		runtime.Gosched()
		err := application.Schedulle(ctx)
		if err != nil {
			logger.Level("fatal", "main", "exit Schedulle:"+err.Error())
		}
	}()

	message := ""
	s := <-stoped
	switch s {
	case syscall.SIGHUP:
		message = "[hungup]"
	case syscall.SIGINT:
		message = "[interupt]"
	case syscall.SIGTERM:
		message = "[force stop]"
	case syscall.SIGQUIT:
		message = "[stop and core dump]"
	default:
		message = "[unknown signal]"
	}
	logger.Level("info", "Run", message)
}
