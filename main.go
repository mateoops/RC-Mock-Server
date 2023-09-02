package main

import (
	"RC-Mock-Server/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	ymlConfig := helpers.ReadConfigurationFromYaml()

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	for _, endpoint := range ymlConfig.Endpoints {
		endpointCopy := endpoint

		router.GET("/"+endpointCopy.Endpoint, func(c *gin.Context) {
			var rcStatus int = endpointCopy.Rc

			if rcStatus == 0 {
				headers := c.Request.Header
				rcStatus, _ = strconv.Atoi(headers.Get("RC"))
			}
			c.AbortWithStatus(rcStatus)
		})
	}

	router.Run(":" + strconv.Itoa(ymlConfig.Server.Port))
}
