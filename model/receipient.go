package model

import (
	"context"
	"fmt"
	"lecture/WBABEProject-10/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *Model) CreateMenu(newMenu Menu) {
	newMenu.IsDeleted = false
	filter := bson.D{{Key: "name", Value: newMenu.Name}}
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
			"can_be_order":    menu.CanBeOrder,
			"price":           menu.Price,
			"origin":          menu.Origin,
			"today_recommend": menu.TodayRecommend,
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

	util.PanicHandler(err)

	s := fmt.Sprintf("상태가 %x으로 변경되었습니다.", state)
	return s
	/* [코드리뷰]
	 * 조건문을 통해 분기를 잘 태워주셨습니다.
	 * 또한 key,value 방식으로 state도 잘 구현해주셨습니다.
	 * 그러나 해당 코드에는 return도 발생이 되고, panic도 발생할 것입니다.
	 * 해당 코드를 호출하는 다른 코드에서는 어떤 일이 일어날 지 예상할 수 없는 코드로 보여집니다.
	 * 항상 function에서 발생하는 일들이 예상가능한 획일적인 return이 이루어지게 만들어주어야 합니다.
	 *
	 * 두번째로, 현재는 대략 30 lines으로 하나의 화면에 코드가 모두 나와 괜찮은 코드이지만,
	 * 이후 코드의 복잡성을 고려한다면, 하나의 function에서 최소한의 return 이 수행되게 만들어주어야 합니다.
	 * to-be:
	 var value string

	 if ... {
		value = "배달 완료된 상태입니다."
	 } else if ...{
		value = fmt.Sprintf("상태가 %x으로 변경되었습니다.", state)
	 } else{
		value = "블라 블라"
	 }

	 return value
	 */
}
