package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	ID    int
	Total float64
}

func CreateInstanceOfType(t reflect.Type) interface{} {
	return reflect.New(t).Elem().Interface()
}

func exampleInstance() {
	orderType := reflect.TypeOf(Order{})
	order := CreateInstanceOfType(orderType).(Order)

	fmt.Printf("Созданный экземпляр: %+v\n", order)
}
