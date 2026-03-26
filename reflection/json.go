package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func MarshallStructToJson(s interface{}) ([]byte, error) {

	// // Получаем значение из интерфейса
	v := reflect.ValueOf(s)

	// Проверяем, что переданный параметр - структура
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct but got %s", v.Kind())
	}

	// // Используем стандартную библиотеку для сериализации в JSON
	return json.Marshal(s)
}

func exampleJson() {
	user := User{
		Name:  "Sanzhar",
		Email: "szhubatkhanov@gmail.com",
		Age:   23,
	}

	jsonData, err := MarshallStructToJson(user)
	if err != nil {
		fmt.Println("Serialize error:", err)
		return
	}

	fmt.Println("Json:", string(jsonData))
}
