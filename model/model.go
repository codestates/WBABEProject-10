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

type Menu struct {
	id             primitive.ObjectID `bson:"_id,omitempty"`
	name           string             `bson:"name"`
	canBeOrder     bool               `bson:"can_be_order"`
	quantity       int                `bson:"quantity"`
	price          int                `bson:"price"`
	origin         string             `bson:"origin"`
	todayRecommend bool               `bson:"today_recommend"`
	createdAt      time.Time          `bson:"created_at"`
	updatedAt      time.Time          `bson:"updated_at"`
	isDeleted      bool               `bson:"is_deleted"`
}

type Orderer struct {
	id        primitive.ObjectID `bson:"_id"`
	phone     string             `bson:"phone"`
	address   string             `bson:"address"`
	createdAt time.Time          `bson:"created_at"`
	updatedAt time.Time          `bson:"updated_at"`
}

type Review struct {
	menuLists   []string  `bson:"menu"`
	orderer     string    `bson:"orderer"`
	score       int       `bson:"score"`
	isRecommend bool      `bson:"is_recommend"`
	review      string    `bson:"review"`
	createdAt   time.Time `bson:"created_at"`
	updatedAt   time.Time `bson:"updated_at"`
	isDeleted   bool      `bson:"is_deleted"`
}

type Order struct {
	id        primitive.ObjectID `bson:"_id"`
	menuLists []string           `bson:"menu_lists"`
	ordererId string             `bson:"orderer_id"`
	state     int                `bson:"state"`
	numbering int                `bson:"numbering"`
	createdAt time.Time          `bson:"created_at"`
	updatedAt time.Time          `bson:"updated_at"`
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
