package router

import (
	"fmt"

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

	account := r.Group("acc/v01")
	{
		fmt.Println(account)
		account.GET("/health", p.ct.Health)
	}
	return r
}
