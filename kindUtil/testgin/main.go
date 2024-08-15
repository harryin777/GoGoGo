/**
*   @Author: yky
*   @File: simpl_http
*   @Version: 1.0
*   @Date: 2021-07-14 21:08
 */
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"key": "hello gin",
		})
	})
	r.Run(":8080")
}
