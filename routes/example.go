package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hhstu/gin-template/apis"
)

type Example struct{}

func (*Example) List(c *gin.Context) {
	var u apis.Example
	if err := c.Bind(&u); err != nil {
		return
	}
}
