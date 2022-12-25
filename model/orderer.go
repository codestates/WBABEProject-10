package model

import (
	"context"
	"errors"
	"fmt"
	"lecture/WBABEProject-10/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderBody struct {
	Phone    string   `validate:"required"`
	Address  string   `validate:"required"`
	MenuName []string `validate:"required"`
}

type CreateReviewBody struct {
	Score       int
	IsRecommend bool
	Review      string
}

type UpdateOrderBody struct {
	MenuName []string `validate:"required"`
}

type AddOrderBody struct {
	MenuName []string `validate:"required"`
}

type GetOrderStateBody struct {
	Address string `validate:"required"`
	Phone   string `validate:"required"`
}

func (m *Model) GetMenus() []Menu {
	filter := bson.D{}
	cursor, err := m.colMenu.Find(context.TODO(), filter)

	util.PanicHandler(err)

	var menu []Menu

	if err = cursor.All(context.TODO(), &menu); err != nil {
		panic(err)
	}

	return menu
}

func (m *Model) GetReviews() []Review {
	filter := bson.D{}
	cursor, err := m.colReview.Find(context.TODO(), filter)

	util.PanicHandler(err)

	var review []Review

	if err = cursor.All(context.TODO(), &review); err != nil {
		panic(err)
	}

	return review
}

func (m *Model) GetOrderState(phone string, address string) []Order {
	var orderer Orderer
	filter := bson.D{{Key: "phone", Value: phone}, {Key: "address", Value: address}}
	m.colOrderer.FindOne(context.TODO(), filter).Decode(&orderer)

	id := orderer.Id.Hex()

	findFilter := bson.D{{Key: "orderer_id", Value: bson.D{{"$eq", id}}}}
	cursor, err := m.colOrder.Find(context.TODO(), findFilter)

	util.PanicHandler(err)

	var order []Order

	if err = cursor.All(context.TODO(), &order); err != nil {
		panic(err)
	}

	return order
}

func (m *Model) AddOrder(orderId primitive.ObjectID, addOrderBody AddOrderBody) error {
	var order Order

	filter := bson.D{{Key: "_id", Value: orderId}}
	m.colOrder.FindOne(context.TODO(), filter).Decode(&order)

	if order.State != 0 && order.State != 1 {
		return errors.New("주문을 추가할 수 없습니다.")
	}

	order.MenuLists = append(order.MenuLists, addOrderBody.MenuName...)

	updateFilter := bson.D{{Key: "_id", Value: orderId}}
	update := bson.M{
		"$set": bson.M{
			"menu_lists": order.MenuLists,
			"state":      0,
		},
	}

	result, err := m.colOrder.UpdateOne(context.TODO(), updateFilter, update)

	util.PanicHandler(err)

	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)

	return nil
}

func (m *Model) UpdateOrder(orderId primitive.ObjectID, updateOrderBody UpdateOrderBody) error {
	var order Order

	filter := bson.D{{Key: "_id", Value: orderId}}
	m.colOrder.FindOne(context.TODO(), filter).Decode(&order)

	if order.State != 0 {
		return errors.New("주문을 변경할 수 없습니다.")
	}

	updateFilter := bson.D{{"_id", orderId}}
	update := bson.D{{"$set", bson.D{{"menu_lists", updateOrderBody.MenuName}}}}

	result, err := m.colOrder.UpdateOne(context.TODO(), updateFilter, update)

	util.PanicHandler(err)

	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)

	return nil
}

func (m *Model) CreateReview(orderId primitive.ObjectID, createReviewBody CreateReviewBody) {
	var review Review
	order := &Order{}

	filter := bson.D{{Key: "_id", Value: orderId}}
	m.colOrder.FindOne(context.TODO(), filter).Decode(order)

	review.Review = createReviewBody.Review
	review.Score = createReviewBody.Score
	review.IsRecommend = createReviewBody.IsRecommend

	ordererId, _ := primitive.ObjectIDFromHex(order.OrdererId)
	review.Orderer = ordererId.String()
	review.MenuLists = order.MenuLists

	result, err := m.colReview.InsertOne(context.TODO(), review)

	util.PanicHandler(err)

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}

func (m *Model) CreateOrder(createOrderBody CreateOrderBody) error {
	var orderer Orderer
	orderer.Address = createOrderBody.Address
	orderer.Phone = createOrderBody.Phone

	var menu Menu
	var arrMenuId []string

	for _, value := range createOrderBody.MenuName {
		filter := bson.D{{Key: "name", Value: value}}
		m.colMenu.FindOne(context.TODO(), filter).Decode(&menu)
		arrMenuId = append(arrMenuId, menu.Id.String())

		if menu.Quantity > 0 {
			updateFilter := bson.D{{Key: "name", Value: menu.Name}}
			update := bson.D{{Key: "$set", Value: bson.D{{"quantity", menu.Quantity - 1}}}}
			_, err := m.colMenu.UpdateOne(context.TODO(), updateFilter, update)

			util.PanicHandler(err)
		}

		if menu.Quantity <= 0 || menu.CanBeOrder == false {
			return errors.New("주문할 수 없는 메뉴")
		}
	}

	ordererResult, err := m.colOrderer.InsertOne(context.TODO(), orderer)

	util.PanicHandler(err)

	count, _ := m.colOrder.CountDocuments(context.TODO(), bson.D{{}})

	var order Order
	order.State = 0
	order.Numbering = int(count) + 1
	order.OrdererId = ordererResult.InsertedID.(primitive.ObjectID).Hex()

	order.MenuLists = arrMenuId
	result, err := m.colOrder.InsertOne(context.TODO(), order)

	util.PanicHandler(err)

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}
