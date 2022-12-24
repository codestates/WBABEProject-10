package model

import (
	"context"
	"fmt"
)

func (m *Model) CreateMenu(newMenu Menu) {
	result, err := m.colMenu.InsertOne(context.TODO(), newMenu)

	if err != nil {
		panic(err)
	}

	fmt.Println("Document inserted with ID: %s\n", result.InsertedID)
}
