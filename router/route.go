package router

import (
	"flag"

	"github.com/EmeraldLS/url-shortener/controller"
	"github.com/EmeraldLS/url-shortener/service"
	"github.com/gin-gonic/gin"
)

func Run() {
	port := flag.String("p", "8080", "server listening port")
	flag.Parse()

	s := gin.Default()
	s.POST("/short", controller.ShortenURL)
	s.GET("/:id", controller.RouteToURL)
	s.DELETE("/expired", controller.DeleteExpiredURLS)

	go service.AutoExpiredDeleteURL()
	// s.Delims("/")
	s.Run("0.0.0.0:" + *port)
}
