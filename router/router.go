package router

import (
	ctl "lecture/WBABEProject-10/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl}
	return r, nil
}

func (p *Router) Idx() *gin.Engine {
	r := gin.New()

	orderer := r.Group("orderer/v01")
	{
		orderer.GET("/menus")
		orderer.GET("/menus/reviews")
		orderer.POST("/menus/reviews")
		orderer.POST("/order")
		orderer.PUT("/order")
		orderer.GET("/order/state")
	}

	receipient := r.Group("receipient/v01")
	{
		receipient.POST("/menus")
		receipient.PUT("/menus")
		receipient.DELETE("/menus")
		receipient.GET("/order")
	}
	return r
}
