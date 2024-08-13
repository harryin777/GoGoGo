package strategy

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 另一种策略模式 https://www.cnblogs.com/amunote/p/15553288.html

// 以下主要实现了为所有接口都提供一个校验参数的逻辑，具体这个逻辑可以定制化

// 校验策略接口
type ValidationStrategy interface {
	Validate(c *gin.Context) error
}

// 具体的校验策略1
type ValidateName struct{}

func (v *ValidateName) Validate(c *gin.Context) error {
	name := c.Query("name")
	if name == "" {
		return fmt.Errorf("name parameter is required")
	}
	return nil
}

// 具体的校验策略2
type ValidateAge struct{}

func (v *ValidateAge) Validate(c *gin.Context) error {
	age := c.Query("age")
	if age == "" {
		return fmt.Errorf("age parameter is required")
	}
	return nil
}

// 使用策略进行参数校验的中间件
func UseValidation(strategy ValidationStrategy) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := strategy.Validate(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func T() {
	r := gin.Default()

	// 为每个接口应用不同的校验策略
	r.GET("/hello", UseValidation(&ValidateName{}), func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	r.GET("/age", UseValidation(&ValidateAge{}), func(c *gin.Context) {
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Your age is %s", age),
		})
	})

	r.Run(":8080")
}
