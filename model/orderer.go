package model

import (
	"context"
	"errors"
	"fmt"

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

func (m *Model) CreateReview(orderId primitive.ObjectID, createReviewBody CreateReviewBody) {
	var review Review
	order := &Order{}

	filter := bson.D{{Key: "_id", Value: orderId}}
	m.colOrder.FindOne(context.TODO(), filter).Decode(order)

	review.Review = createReviewBody.Review
	review.Score = createReviewBody.Score
	review.IsRecommend = createReviewBody.IsRecommend

	ordererId, _ := primitive.ObjectIDFromHex(order.OrdererId)
	review.Orderer = ordererId
	review.MenuLists = order.MenuLists

	result, err := m.colMenuReview.InsertOne(context.TODO(), review)

	if err != nil {
		panic(err)
	}

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

			if err != nil {
				panic(err)
			}
		}

		if menu.Quantity <= 0 || menu.CanBeOrder == false {
			return errors.New("주문할 수 없는 메뉴")
		}
	}

	ordererResult, err := m.colOrderer.InsertOne(context.TODO(), orderer)

	if err != nil {
		panic(err)
	}

	count, _ := m.colOrder.CountDocuments(context.TODO(), bson.D{{}})

	var order Order
	order.State = 0
	order.Numbering = int(count) + 1
	order.OrdererId = ordererResult.InsertedID.(primitive.ObjectID).Hex()

	order.MenuLists = arrMenuId
	result, err := m.colOrder.InsertOne(context.TODO(), order)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
}
