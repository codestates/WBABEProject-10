package router

import (
	ctl "lecture/WBABEProject-10/controller"

	"github.com/gin-gonic/gin"

	"lecture/WBABEProject-10/docs"

	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl}
	return r, nil
}

func (p *Router) Idx() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORS())

	r.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "localhost"

	orderer := r.Group("orderer/v01")
	{
		orderer.GET("/menus", p.ct.GetMenus)
		orderer.GET("/reviews", p.ct.GetReviews)
		orderer.GET("/order/state", p.ct.GetOrderState)
		orderer.POST("/reviews/:orderId", p.ct.CreateReview)
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
	/* [코드리뷰]
	 * Group을 사용하여 API 성격에 따라 request를 관리하는 코드는 매우 좋은 코드입니다.
     * 일반적으로 현업에서도 이와 같은 코드를 자주 사용합니다. 훌륭합니다.
	 *
	 * 또한 API의 endpoint에 version을 넣어주셔서, 이후에 api의 수정이 발생할 경우,
	 * v01 방식의 클라이언트와, v02 방식의 클라이언트를 모두 받아줄 수 있는 확장성 있는 좋은 코드입니다.
	 */

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
