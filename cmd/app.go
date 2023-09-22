package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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

func init() {
	gin.SetMode(config.AppConfig.Webserver.Mode)
}

func main() {
	defer log.Logger.Sync()
	webServerPort := config.AppConfig.Webserver.Port
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%s", webServerPort),
		Handler:        routes.Routes(),
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	go func() {
		log.Logger.Infof("http server start at %d", webServerPort)
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
}
