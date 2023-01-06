package controller

import (
	"lecture/WBABEProject-10/model"
	"lecture/WBABEProject-10/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewMenu godoc
// @Summary 기존 메뉴에 새로운 메뉴를 추가합니다.
// @Description 메뉴 추가 가능
// @name NewMenu
// @Accept  json
// @Produce  json
// @Param Menu body model.Menu true "Menu"
// @Router /receipient/v01/menus [post]
// @Success 200 {object} string
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
// @Summary 기존 메뉴를 수정합니다.
// @Description 메뉴 업데이트 가능
// @name UpdateMenu
// @Accept  json
// @Produce  json
// @Param name path string true "name"
// @Param Menu body model.Menu true "Menu"
// @Router /receipient/v01/menus/{name} [patch]
// @Success 200 {object} string
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
// @Summary 메뉴를 삭제합니다.
// @Description 메뉴 삭제 기능
// @name DeleteMenu
// @Accept  json
// @Produce  json
// @Param name path string true "name"
// @Router /receipient/v01/menus/{name} [delete]
// @Success 200 {object} string
func (p *Controller) DeleteMenu(c *gin.Context) {
	name := c.Param("name")

	p.md.DeleteMenu(name)

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

// GetOrders godoc
// @Summary 전체 주문을 조회합니다.
// @Description 주문 조회 기능
// @name GetOrders
// @Accept  json
// @Produce  json
// @Router /receipient/v01/order [get]
// @Success 200 {object} string
func (p *Controller) GetOrders(c *gin.Context) {
	result := p.md.GetOrders()

	c.JSON(http.StatusOK, gin.H{
		"msg":    "OK",
		"result": result,
	})
}

// UpdateOrderState godoc
// @Summary 주문 받은 상태를 업데이트합니다.
// @Description 주문 상태 업데이트 기능
// @name UpdateOrderState
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Router /receipient/v01/order/{id}/state [patch]
// @Success 200 {object} string
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
