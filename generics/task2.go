package main

import "fmt"

// Эта задача заставит тебя вспомнить про ограничение comparable,
// так как обычный any нельзя сравнивать через ==.

// Задание: Напиши функцию Contains[T comparable](slice []T, target T) bool.

// Она должна возвращать true, если элемент target есть в слайсе, и false, если нет.

func testContains() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	strs := []string{"men", "women", "male", "female"}

	fmt.Println(Containss(ints, 4))
	fmt.Println(Containss(strs, "men"))
}

func Containss[T comparable](slice []T, target T) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}
