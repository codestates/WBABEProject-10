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
// @Summary call GetMenus, return ok by json.
// @Description 메뉴 조회 기능
// @name GetMenus
// @Accept  json
// @Produce  json
// @Router /orderer/v01/menus [get]
// @Success 200 {object} Controller
func (p *Controller) GetMenus(c *gin.Context) {
	result := p.md.GetMenus()

	c.JSON(http.StatusOK, gin.H{
		"msg":    "OK",
		"result": result,
	})
}

// GetReviews godoc
// @Summary call GetReviews, return ok by json.
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
// @Summary call GetOrderState, return ok by json.
// @Description 주문 상태 조회 기능
// @name GetOrderState
// @Accept  json
// @Produce  json
// @Param phone query string true "phone"
// @Param address query string true "address"
// @Router /orderer/v01/order/state [get]
// @Success 200 {object} Controller
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
// @Summary call AddOrder, return ok by json.
// @Description 주문을 추가하는 기능
// @name AddOrder
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Router /orderer/v01/order/{id} [patch]
// @Success 200 {object} Controller
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
// @Summary call AddOrder, return ok by json.
// @Description 주문 수정 기능
// @name UpdateOrder
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param MenuName body []string true "MenuName"
// @Router /orderer/v01/order/{id} [put]
// @Success 200 {object} Controller
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
// @Summary call CreateReview, return ok by json.
// @Description 리뷰 기능
// @name CreateReview
// @Accept  json
// @Produce  json
// @Param orderId path string true "orderId"
// @Param Score body int true "Score"
// @Param IsRecommend bool string true "IsRecommend"
// @Param Review path string true "Review"
// @Router /orderer/v01/reviews/{orderId} [post]
// @Success 200 {object} Controller
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
// @Summary call CreateOrder, return ok by json.
// @Description 주문 기능
// @name CreateOrder
// @Accept  json
// @Produce  json
// @Param Phone body string true "Phone"
// @Param Address body string true "Address"
// @Param MenuName body []string true "MenuName"
// @Router /orderer/v01/order [post]
// @Success 200 {object} Controller
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

/* [코드리뷰]
	 * 들어오는 api request에 대해서 validation 코드를 넣어보시는 건 어떨까요?
	 * server에서는 항상 client를 의심하는 방어적 코딩 스타일을 수행해야 합니다.
	 * 
	 * 에러 상황에 대한 간단한 메세지 들이 많이 발생하고 있네요.
	 * 해당 값을, 하나의 map 함수의 key, value로 관리하는 것을 추천드립니다.
	 * 대신 mapping되는 string 타입의 key 값이 상황이 잘 설명되는 
	 * naming convention이 있으면 깔끔해질 것으로 보여집니다.
*/
