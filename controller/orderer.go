package controller

import (
	"lecture/WBABEProject-10/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Controller) NewMenu(c *gin.Context) {
	var menu model.Menu

	if err := c.ShouldBind(&menu); err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}
	p.md.CreateMenu(menu)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
