package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client     *mongo.Client
	colMenu    *mongo.Collection
	colOrder   *mongo.Collection
	colOrderer *mongo.Collection
	colReview  *mongo.Collection
}

type Orderer struct {
	Id        primitive.ObjectID `bson:"_id"`
	Phone     string             `bson:"phone"`
	Address   string             `bson:"address"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type Review struct {
	MenuLists   []string  `bson:"menu"`
	Orderer     string    `bson:"orderer"`
	Score       int       `bson:"score"`
	IsRecommend bool      `bson:"is_recommend"`
	Review      string    `bson:"review"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
	isDeleted   bool      `bson:"is_deleted"`
}

type Order struct {
	Id        primitive.ObjectID `bson:"_id"`
	MenuLists []string           `bson:"menu_lists"`
	OrdererId string             `bson:"orderer_id"`
	State     int                `bson:"state"`
	Numbering int                `bson:"numbering"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	isDeleted bool               `bson:"is_deleted"`
}

func NewModel() (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := "mongodb://127.0.0.1:27017"

	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err = r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("online-ordering-system")
		r.colMenu = db.Collection("tMenu")
		r.colOrderer = db.Collection("tOrderer")
		r.colOrder = db.Collection("tOrder")
		r.colReview = db.Collection("tReview")
	}
	return r, nil
}
