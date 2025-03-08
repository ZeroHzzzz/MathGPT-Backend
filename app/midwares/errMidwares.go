package midwares

import (
	"fmt"
	"mathgpt/app/apiException"

	"github.com/gin-gonic/gin"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println(c.Errors)
		if length := len(c.Errors); length > 0 {
			e := c.Errors[length-1]
			err := e.Err
			if err != nil {
				var Err *apiException.Error
				if e, ok := err.(*apiException.Error); ok {
					Err = e
				} else {
					Err = apiException.OtherError(err.Error())
				}
				// TODO 建立日志系统

				c.JSON(Err.StatusCode, Err)
				return
			}
		}

	}
}

// HandleNotFound
//
//	404处理
func HandleNotFound(c *gin.Context) {
	err := apiException.NotFound
	c.JSON(err.StatusCode, err)
}
