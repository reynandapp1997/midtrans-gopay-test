package main

import (
	"flag"
	"fmt"
	"midtrans-gopay-test/config"
	"midtrans-gopay-test/controller"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configFile := flag.String("config", "config/dev.json", "Config environment path")
	flag.Parse()
	config.LoadConfiguration(*configFile)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowWildcard = true
	corsConfig.AllowOrigins = []string{"*"} // contain whitelist domain
	corsConfig.AllowHeaders = []string{"*", "user", "domain", "verb", "object", "Content-Type", "Accept"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")

	router.Use(cors.New(corsConfig))

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	v1.GET("/midtrans/charge", controller.RequestChart)
	v1.GET("/midtrans/status", controller.GetTransactionStatus)

	err := router.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		panic(err)
	}
}
