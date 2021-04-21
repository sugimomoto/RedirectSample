package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		param := c.Query("param")
		query := c.DefaultQuery("query", "unknown")

		c.JSON(200, gin.H{
			"message": "query result : " + query + "/ param :" + param,
		})
	})
	r.Run()
}
