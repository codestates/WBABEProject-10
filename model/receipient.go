package model

import (
	"context"
	"fmt"
	"lecture/WBABEProject-10/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetOrdersResponse struct {
	id        primitive.ObjectID
	menuLists []Menu
	orderer   Orderer
	state     int
	numbering int
	createdAt time.Time
}

func (m *Model) CreateMenu(newMenu Menu) {
	newMenu.isDeleted = false
	filter := bson.D{{Key: "name", Value: newMenu.name}}
	count, err := m.colMenu.CountDocuments(context.TODO(), filter)

	util.PanicHandler(err)

	if count > 0 {
		panic("이미 존재하는 이름입니다.")
	}

	result, err := m.colMenu.InsertOne(context.TODO(), newMenu)

	util.PanicHandler(err)

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}

func (m *Model) UpdateMenu(name string, menu Menu) {
	filter := bson.D{{Key: "name", Value: name}}
	update := bson.M{
		"$set": bson.M{
			"can_be_order":    menu.canBeOrder,
			"price":           menu.price,
			"origin":          menu.origin,
			"today_recommend": menu.todayRecommend,
		},
	}

	result, err := m.colMenu.UpdateOne(context.TODO(), filter, update)

	util.PanicHandler(err)

	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
}

func (m *Model) DeleteMenu(name string) {
	filter := bson.D{{Key: "name", Value: name}}
	update := bson.D{{Key: "$set", Value: bson.D{{"is_deleted", true}}}}
	result, err := m.colMenu.UpdateOne(context.TODO(), filter, update)

	util.PanicHandler(err)

	fmt.Printf("Documents Deleted: %v\n", result.ModifiedCount)
}

func (m *Model) GetOrders() []Order {
	filter := bson.D{}
	cursor, err := m.colOrder.Find(context.TODO(), filter)

	util.PanicHandler(err)

	var orders []Order

	if err = cursor.All(context.TODO(), &orders); err != nil {
		panic(err)
	}

	return orders
}

func (m *Model) UpdateOrderState(orderId primitive.ObjectID) string {
	order := &Order{}

	filter := bson.D{{Key: "_id", Value: orderId}}
	m.colOrder.FindOne(context.TODO(), filter).Decode(order)

	var state int
	if order.state == 3 {
		return "배달 완료된 상태입니다."
	} else if order.state == 0 {
		state = 1
	} else if order.state == 1 {
		state = 2
	} else if order.state == 2 {
		state = 3
	}

	updateFilter := bson.D{{Key: "_id", Value: orderId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "state", Value: state}}}}
	_, err := m.colOrder.UpdateOne(context.TODO(), updateFilter, update)

	util.PanicHandler(err)

	s := fmt.Sprintf("상태가 %x으로 변경되었습니다.", state)
	return s
}
