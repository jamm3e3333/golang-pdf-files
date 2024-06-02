package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jamm3e3333/start/pkg/http"
	"github.com/jamm3e3333/start/pkg/logger"
)

func main() {
	lg := logger.NewLogger()

	gin.SetMode(gin.ReleaseMode)
	ge := gin.Default()

	httpServer := http.NewServer(&http.Config{
		Port:            8080,
		ReadTimeout:     5,
		WriteTimeout:    5,
		ShutdownTimeout: 5,
		Handler:         ge,
	})
	errChan := httpServer.Start()
	lg.Info(fmt.Sprintf("starting http server on port: %d", 8080))

	err := <-errChan
	lg.Error(err)
}
