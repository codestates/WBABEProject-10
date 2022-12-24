package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client   *mongo.Client
	colOrder *mongo.Collection
}

type Menu struct {
	Name           string    `bson:"name"`
	CanBeOrder     bool      `bson:"can_be_order"`
	Limited        int       `bson:"limited"`
	Price          int       `bson:"price"`
	Origin         string    `bson:"origin"`
	TodayRecommend bool      `bson:"today_recommend"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at"`
}

type Orderer struct {
	Name      string    `bson:"name"`
	Phone     string    `bson:"phone"`
	Address   string    `bson:"address"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type MenuReview struct {
	Menu          Menu      `bson:"menu"`
	Orderer       Orderer   `bson:"orderer"`
	Score         int       `bson:"score"`
	IsRecommended bool      `bson:"is_recommended"`
	review        string    `bson:"review"`
	CreatedAt     time.Time `bson:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"`
}

type Order struct {
	Menu      Menu      `bson:"menu"`
	Orderer   Orderer   `bson:"orderer"`
	State     string    `bson:"state"`
	Numbering int       `bson:"numbering"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type Reorder struct {
	Menu      Menu      `bson:"menu"`
	Orderer   Orderer   `bson:"orderer"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
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
		r.colOrder = db.Collection("order")
	}
	return r, nil
}
