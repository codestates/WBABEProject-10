package controller

import (
	"lecture/WBABEProject-10/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

func (p *Controller) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "health",
	})
	return
}
