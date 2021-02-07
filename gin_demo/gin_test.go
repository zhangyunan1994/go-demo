package gin_demo

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGin(t *testing.T)  {
	r := gin.New()
	r.GET("/actuator/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":8010")
	if err != nil {
		panic(err)
	}
}
