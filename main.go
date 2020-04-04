package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/kit"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	AppName    = "app"
	AppVersion = "0.1.0"
)

var (
	configFile  string
	showVersion bool
)

func init() {
	flag.StringVar(&configFile, "f", kit.DefaultConfig, "config file")
	flag.BoolVar(&showVersion, "v", false, "show version and exit")
	flag.Parse()
}

func main() {
	if showVersion {
		fmt.Printf("Name: %s, Version: %s\n", AppName, AppVersion)
		return
	}

	config, err := kit.NewConfig(configFile)
	if err != nil {
		log.Fatalf("Load config file failed: %s", err.Error())
	}

	kit.InitLog(config)
	logrus.Infof("Config: %+v", config)

	startListener(config)
}

func startListener(config *kit.Config) {
	if config.Log.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
	}

	router := getRouter()
	srv := &http.Server{
		Addr:    "0.0.0.0:" + config.App.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Listen failed: %s", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Shutdown server failed: %s", err.Error())
	} else {
		logrus.Info("Shutdown server success")
	}
}
