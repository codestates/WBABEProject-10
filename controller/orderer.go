package controller

import (
	"lecture/WBABEProject-10/model"
	"lecture/WBABEProject-10/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateReviewBody struct {
	Score       int
	IsRecommend bool
	Review      string
}

// GetMenus godoc
// @Summary 메뉴를 조회합니다.
// @Description 메뉴 조회 기능
// @name GetMenus
// @Accept  json
// @Produce  json
// @Router /orderer/v01/menus [get]
// @Success 200 {object} []model.Menu
func (p *Controller) GetMenus(c *gin.Context) {
	result := p.md.GetMenus()

	c.JSON(http.StatusOK, gin.H{
		"msg":    "OK",
		"result": result,
	})
}

// GetReviews godoc
// @Summary 리뷰를 조회합니다.
// @Description 리뷰 조회 기능
// @name GetReviews
// @Accept  json
// @Produce  json
// @Router /orderer/v01/reviews [get]
// @Success 200 {object} Controller
func (p *Controller) GetReviews(c *gin.Context) {
	result := p.md.GetReviews()

	c.JSON(http.StatusOK, gin.H{
		"msg":    "OK",
		"result": result,
	})
}

// GetOrderState godoc
// @Summary 주문 상태를 조회합니다.
// @Description 주문 상태 조회 기능
// @name GetOrderState
// @Accept  json
// @Produce  json
// @Param phone query string true "Phone"
// @Param address query string true "Address"
// @Router /orderer/v01/order/state [get]
// @Success 200 {object} string
func (p *Controller) GetOrderState(c *gin.Context) {
	phone := c.Query("phone")
	address := c.Query("address")

	result := p.md.GetOrderState(phone, address)

	c.JSON(http.StatusOK, gin.H{
		"msg":    "OK",
		"result": result,
	})
}

// AddOrder godoc
// @Summary 기존 주문에 새로운 주문을 추가합니다.
// @Description 주문을 추가하는 기능
// @name AddOrder
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Router /orderer/v01/order/{id} [patch]
// @Success 200 {object} string
// @Failure 400 {object} string
func (p *Controller) AddOrder(c *gin.Context) {
	id := c.Param("id")

	orderId, _ := primitive.ObjectIDFromHex(id)

	var addOrderBody model.AddOrderBody

	if err := c.ShouldBind(&addOrderBody); err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	v := validator.New()
	err := v.Struct(addOrderBody)

	util.PanicHandler(err)

	result := p.md.AddOrder(orderId, addOrderBody)

	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "메뉴를 추가할 수 없습니다.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

// UpdateOrder godoc
// @Summary 이미 전달한 주문을 수정합니다.
// @Description 주문 수정 기능
// @name UpdateOrder
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param MenuName body []string true "MenuName"
// @Router /orderer/v01/order/{id} [put]
// @Success 200 {object} string
// @Failure 400 {object} string
func (p *Controller) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	orderId, _ := primitive.ObjectIDFromHex(id)

	var updateOrderBody model.UpdateOrderBody

	if err := c.ShouldBind(&updateOrderBody); err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	v := validator.New()
	err := v.Struct(updateOrderBody)

	util.PanicHandler(err)

	result := p.md.UpdateOrder(orderId, updateOrderBody)

	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "메뉴를 변경할 수 없습니다",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

// CreateReview godoc
// @Summary 리뷰를 작성합니다.
// @Description 리뷰 기능
// @name CreateReview
// @Accept  json
// @Produce  json
// @Param orderId path string true "orderId"
// @Param CreateReviewBody body model.CreateReviewBody true "CreateReviewBody"
// @Router /orderer/v01/reviews/{orderId} [post]
// @Success 200 {object} string
func (p *Controller) CreateReview(c *gin.Context) {
	orderId := c.Param("orderId")

	id, _ := primitive.ObjectIDFromHex(orderId)

	var createReviewBody model.CreateReviewBody

	if err := c.ShouldBind(&createReviewBody); err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	p.md.CreateReview(id, createReviewBody)

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

// CreateOrder godoc
// @Summary 원하는 메뉴를 주문합니다.
// @Description 주문 기능
// @name CreateOrder
// @Accept  json
// @Produce  json
// @Param CreateOrderBody body model.CreateOrderBody true "CreateOrderBody"
// @Router /orderer/v01/order [post]
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 400 {object} string
func (p *Controller) CreateOrder(c *gin.Context) {
	var createOrderBody model.CreateOrderBody

	if err := c.ShouldBind(&createOrderBody); err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	v := validator.New()
	err := v.Struct(createOrderBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "필수 값이 존재하지 않습니다.",
		})
		return
	}

	result := p.md.CreateOrder(createOrderBody)

	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "해당 메뉴는 현재 주문할 수 없습니다.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
