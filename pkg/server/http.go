package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewServer(route *gin.Engine, port string) http.Server {

	return http.Server{
		Addr:         port,
		Handler:      route,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
