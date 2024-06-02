package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jamm3e3333/start/pkg/http"
	"github.com/jamm3e3333/start/pkg/logger"
)

func main() {
	lg := logger.NewLogger()

	gin.SetMode(gin.ReleaseMode)
	ge := gin.New()

	ge.Any("/", func(c *gin.Context) {
		c.AbortWithStatusJSON(200, gin.H{
			"message": "Hello World",
		})
	})

	httpServer := http.NewServer(&http.Config{
		Port:            8080,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    30 * time.Second,
		ShutdownTimeout: 15 * time.Second,
		Handler:         ge,
	})
	errChan := httpServer.Start()
	lg.Info(fmt.Sprintf("starting http server on port: %d", 8080))

	err := <-errChan
	lg.Error(err)
}
