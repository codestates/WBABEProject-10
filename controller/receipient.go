package controller

import (
	"lecture/WBABEProject-10/model"
	"lecture/WBABEProject-10/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewMenu godoc
// @Summary call NewMenu, return ok by json.
// @Description 메뉴 추가 가능
// @name NewMenu
// @Accept  json
// @Produce  json
// @Param Name body string true "Name"
// @Param CanBeOrder body bool true "CanBeOrder"
// @Param Quantity body int true "Quantity"
// @Param Price body int true "Price"
// @Param Origin body string true "Origin"
// @Param TodayRecommend body bool true "TodayRecommend"
// @Router /receipient/v01/menus [post]
// @Success 200 {object} Controller
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

// UpdateMenu godoc
// @Summary call UpdateMenu, return ok by json.
// @Description 메뉴 업데이트 가능
// @name UpdateMenu
// @Accept  json
// @Produce  json
// @Param name path string true "name"
// @Param CanBeOrder body bool true "CanBeOrder"
// @Param Price body int true "Price"
// @Param Origin body string true "Origin"
// @Param TodayRecommend body bool true "TodayRecommend"
// @Router /receipient/v01/menus/{name} [patch]
// @Success 200 {object} Controller
func (p *Controller) UpdateMenu(c *gin.Context) {
	name := c.Param("name")
	var menu model.Menu
	if err := c.ShouldBind(&menu); err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}
	p.md.UpdateMenu(name, menu)

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

// DeleteMenu godoc
// @Summary call DeleteMenu, return ok by json.
// @Description 메뉴 삭제 기능
// @name DeleteMenu
// @Accept  json
// @Produce  json
// @Param name path string true "name"
// @Router /receipient/v01/menus/{name} [delete]
// @Success 200 {object} Controller
func (p *Controller) DeleteMenu(c *gin.Context) {
	name := c.Param("name")

	p.md.DeleteMenu(name)

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

// GetOrders godoc
// @Summary call GetOrders, return ok by json.
// @Description 주문 조회 기능
// @name GetOrders
// @Accept  json
// @Produce  json
// @Router /receipient/v01/order [get]
// @Success 200 {object} Controller
func (p *Controller) GetOrders(c *gin.Context) {
	result := p.md.GetOrders()

	c.JSON(http.StatusOK, gin.H{
		"msg":    "OK",
		"result": result,
	})
}


// UpdateOrderState godoc
// @Summary call UpdateOrderState, return ok by json.
// @Description 주문 상태 업데이트 기능
// @name UpdateOrderState
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Router /receipient/v01/order/{id}/state [patch]
// @Success 200 {object} Controller
func (p *Controller) UpdateOrderState(c *gin.Context) {
	orderId := c.Param("id")

	id, err := primitive.ObjectIDFromHex(orderId)

	util.PanicHandler(err)

	result := p.md.UpdateOrderState(id)

	c.JSON(http.StatusOK, gin.H{
		"msg":    "OK",
		"result": result,
	})
}
