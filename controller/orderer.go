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
