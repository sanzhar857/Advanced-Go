package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "Sanzhar",
		Age:  23,
	}

	t := reflect.TypeOf(p)
	fmt.Println("Type of struct:", t.Name())

	v := reflect.ValueOf(p)
	fmt.Println("Value: ", v)

	numFields := v.NumField()
	fmt.Println("Number of fields:", numFields)

	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		fmt.Printf("Field %d: %v\n", i, field)
	}

	exampleJson()
	exampleValidate()
	exampleInstance()
}
