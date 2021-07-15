package main

import (
	"basic_golang_echo/internal/config"
	"basic_golang_echo/internal/handler"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	configs := config.Configs{}
	if err := configs.InitAllConfigs(); err != nil {
		fmt.Println(err)
		return
	}
	e := echo.New()
	e.Use(handler.Recover)

	closes, err := handler.NewRoutes(e, &configs)
	if err != nil {
		log.Panic("new routes error:", err)
		return
	}

	e.Logger.Fatal(e.Start(":8888"))

	// graceful shutdown
	//น่าจะทำงานหลังจากรับ ^C interrupt
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals

	srvCtx, srvCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer srvCancel()

	for _, close := range closes {
		if err := close(); err != nil {
			log.Errorf("failed to cleanup resources: %v", err)
		}
	}

	log.Info("shutting down http server...")
	if err := e.Server.Shutdown(srvCtx); err != nil {
		log.Panic("http server shutdown with error:", err)
	}
}
