package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Product struct {
	Name  string  `validate:"required"`
	Price float64 `validate:"min=0"`
}

func validateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
		return errors.New("expected struct")
	}

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		tag := fieldType.Tag.Get("validate")

		if strings.Contains(tag, "required") && fieldValue.IsZero() {
			return fmt.Errorf("Field %s required", fieldType.Name)
		}

		if strings.Contains(tag, "min=0") && fieldValue.Kind() == reflect.Float64 && fieldValue.Float() < 0 {
			return fmt.Errorf("значение поля %s должно быть неотрицательным", fieldType.Name)
		}
	}
	return nil
}

func exampleValidate() {
	product := Product{
		Name:  "Apple",
		Price: -15.0,
	}

	err := validateStruct(product)
	if err != nil {
		fmt.Println("Ошибка валидации:", err)
	} else {
		fmt.Println("Структура валидна")
	}
}


