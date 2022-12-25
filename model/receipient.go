package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	CanBeOrder     bool               `bson:"can_be_order"`
	Quantity       int                `bson:"quantity"`
	Price          int                `bson:"price"`
	Origin         string             `bson:"origin"`
	TodayRecommend bool               `bson:"today_recommend"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	IsDeleted      bool               `bson:"is_deleted"`
}

type GetOrdersResponse struct {
	Id        primitive.ObjectID
	MenuLists []Menu
	Orderer   Orderer
	State     int
	Numbering int
	CreatedAt time.Time
}

func (m *Model) CreateMenu(newMenu Menu) {
	newMenu.IsDeleted = false
	filter := bson.D{{Key: "name", Value: newMenu.Name}}
	count, err := m.colMenu.CountDocuments(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	if count > 0 {
		panic("이미 존재하는 이름입니다.")
	}

	result, err := m.colMenu.InsertOne(context.TODO(), newMenu)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}

func (m *Model) UpdateMenu(name string, menu Menu) {
	filter := bson.D{{Key: "name", Value: name}}
	update := bson.M{
		"$set": bson.M{
			"can_be_order":    menu.CanBeOrder,
			"price":           menu.Price,
			"origin":          menu.Origin,
			"today_recommend": menu.TodayRecommend,
		},
	}

	result, err := m.colMenu.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
}

func (m *Model) DeleteMenu(name string) {
	filter := bson.D{{Key: "name", Value: name}}
	update := bson.D{{Key: "$set", Value: bson.D{{"is_deleted", true}}}}
	result, err := m.colMenu.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents Deleted: %v\n", result.ModifiedCount)
}

func (m *Model) GetOrders() []Order {
	filter := bson.D{}
	cursor, err := m.colOrder.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

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
	if order.State == 3 {
		return "배달 완료된 상태입니다."
	} else if order.State == 0 {
		state = 1
	} else if order.State == 1 {
		state = 2
	} else if order.State == 2 {
		state = 3
	}

	updateFilter := bson.D{{Key: "_id", Value: orderId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "state", Value: state}}}}
	_, err := m.colOrder.UpdateOne(context.TODO(), updateFilter, update)

	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf("상태가 %x으로 변경되었습니다.", state)
	return s
}
