package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/hhstu/gin-template/log"
	"github.com/hhstu/gin-template/routes"
	"syscall"

	"github.com/hhstu/gin-template/config"
	_ "github.com/hhstu/gin-template/utils"
	_ "go.uber.org/automaxprocs"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	defer log.Logger.Sync()
	webServerPort := config.AppConfig.Webserver.Port
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%s", webServerPort),
		Handler:        routes.Routes(),
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	go func() {
		log.Logger.Infof("http server start at 0.0.0.0:%s", webServerPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Logger.Error("listen error: %s", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Logger.Infof("get OS shutdown signal, shutting down...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Error("shutdown error: %s", err)
	}
	log.Logger.Infof("shutdown success")
}
