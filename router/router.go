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
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORS())

	orderer := r.Group("orderer/v01")
	{
		orderer.GET("/menus", p.ct.GetMenus)
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

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
