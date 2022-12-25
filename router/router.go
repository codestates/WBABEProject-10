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
		// orderer.GET("/menus")
		orderer.GET("/menus/reviews", p.ct.GetReviews)
		orderer.GET("/order/state", p.ct.GetOrderState)
		orderer.POST("/menus/reviews/:orderId", p.ct.CreateReview)
		orderer.POST("/order", p.ct.CreateOrder)
		orderer.PUT("/order/:id", p.ct.UpdateOrder)
		orderer.PATCH("/order/:id", p.ct.AddOrder)
	}

	receipient := r.Group("receipient/v01")
	{
		receipient.POST("/menus", p.ct.NewMenu)
		receipient.PATCH("/menus/:name", p.ct.UpdateMenu)
		receipient.DELETE("/menus/:name", p.ct.DeleteMenu)
		receipient.GET("/order", p.ct.GetOrders)
		receipient.PATCH("/order/:id/state", p.ct.UpdateOrderState)
	}
	return r
}
