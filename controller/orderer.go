package controller

import (
	"lecture/WBABEProject-10/model"
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

func (p *Controller) GetOrderState(c *gin.Context) {
	phone := c.Query("phone")
	address := c.Query("address")

	result := p.md.GetOrderState(phone, address)

	c.JSON(http.StatusOK, gin.H{
		"msg":    "OK",
		"result": result,
	})
}

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

	if err != nil {
		panic(err)
	}

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

	if err != nil {
		panic(err)
	}

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
