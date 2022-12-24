package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Menu struct {
	Name           string    `bson:"name"`
	CanBeOrder     bool      `bson:"can_be_order"`
	Quantity       int       `bson:"quantity"`
	Price          int       `bson:"price"`
	Origin         string    `bson:"origin"`
	TodayRecommend bool      `bson:"today_recommend"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at"`
	IsDeleted      bool      `bson:"is_deleted"`
}

func (m *Model) CreateMenu(newMenu Menu) {
	newMenu.IsDeleted = false
	result, err := m.colMenu.InsertOne(context.TODO(), newMenu)

	if err != nil {
		panic(err)
	}

	fmt.Println("Document inserted with ID: %s\n", result.InsertedID)
}

func (m *Model) UpdateMenu(name string, menu Menu) {
	filter := bson.D{{"name", name}}
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
