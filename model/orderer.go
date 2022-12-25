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
	phone    string   `validate:"required"`
	address  string   `validate:"required"`
	menuName []string `validate:"required"`
}

type CreateReviewBody struct {
	score       int
	isRecommend bool
	review      string
}

type UpdateOrderBody struct {
	menuName []string `validate:"required"`
}

type AddOrderBody struct {
	menuName []string `validate:"required"`
}

type GetOrderStateBody struct {
	address string `validate:"required"`
	phone   string `validate:"required"`
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

	id := orderer.id.Hex()

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

	if order.state != 0 && order.state != 1 {
		return errors.New("주문을 추가할 수 없습니다.")
	}

	order.menuLists = append(order.menuLists, addOrderBody.menuName...)

	updateFilter := bson.D{{Key: "_id", Value: orderId}}
	update := bson.M{
		"$set": bson.M{
			"menu_lists": order.menuLists,
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

	if order.state != 0 {
		return errors.New("주문을 변경할 수 없습니다.")
	}

	updateFilter := bson.D{{"_id", orderId}}
	update := bson.D{{"$set", bson.D{{"menu_lists", updateOrderBody.menuName}}}}

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

	review.review = createReviewBody.review
	review.score = createReviewBody.score
	review.isRecommend = createReviewBody.isRecommend

	ordererId, _ := primitive.ObjectIDFromHex(order.ordererId)
	review.orderer = ordererId.String()
	review.menuLists = order.menuLists

	result, err := m.colReview.InsertOne(context.TODO(), review)

	util.PanicHandler(err)

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}

func (m *Model) CreateOrder(createOrderBody CreateOrderBody) error {
	var orderer Orderer
	orderer.address = createOrderBody.address
	orderer.phone = createOrderBody.phone

	var menu Menu
	var arrMenuId []string

	for _, value := range createOrderBody.menuName {
		filter := bson.D{{Key: "name", Value: value}}
		m.colMenu.FindOne(context.TODO(), filter).Decode(&menu)
		arrMenuId = append(arrMenuId, menu.id.String())

		if menu.quantity > 0 {
			updateFilter := bson.D{{Key: "name", Value: menu.name}}
			update := bson.D{{Key: "$set", Value: bson.D{{"quantity", menu.quantity - 1}}}}
			_, err := m.colMenu.UpdateOne(context.TODO(), updateFilter, update)

			util.PanicHandler(err)
		}

		if menu.quantity <= 0 || menu.canBeOrder == false {
			return errors.New("주문할 수 없는 메뉴")
		}
	}

	ordererResult, err := m.colOrderer.InsertOne(context.TODO(), orderer)

	util.PanicHandler(err)

	count, _ := m.colOrder.CountDocuments(context.TODO(), bson.D{{}})

	var order Order
	order.state = 0
	order.numbering = int(count) + 1
	order.ordererId = ordererResult.InsertedID.(primitive.ObjectID).Hex()

	order.menuLists = arrMenuId
	result, err := m.colOrder.InsertOne(context.TODO(), order)

	util.PanicHandler(err)

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}
