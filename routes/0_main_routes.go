package routes

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/hhstu/gin-template/config"
	"github.com/hhstu/gin-template/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
)

func init() {
	gin.SetMode(config.AppConfig.Webserver.Mode)
}

func Routes() *gin.Engine {
	r := gin.Default()
	api := r.Group("api")

	// 基础监控
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/metrics", prometheusHandler())
	r.GET("/debug/pprof/*pprof", gin.WrapH(http.DefaultServeMux))
	api.Use(ginprom.PromMiddleware(nil))
	api.Use(HandlerRecover)

	// v1 api
	v1 := api.Group("/v1")
	example := Example{}
	v1.GET("/examples", example.List)

	return r
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func HandlerRecover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Errorf("panic: internal error %+v", r)
			debug.PrintStack()
			// returnError(500, "服务器内部异常", c)
			c.Abort()
		}
	}()
	c.Next()
}
