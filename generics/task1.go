package main

import "fmt"

// В Go часто нужно отфильтровать слайс (например, оставить только четные числа или только длинные строки).
// Раньше пришлось бы писать две разные функции.

// Задание: Напиши дженерик-функцию Filter[T any](slice []T, test func(T) bool) []T.

func testFilter() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Filter(ints, func(n int) bool { return n%2 == 0 }))

	strs := []string{"aoba", "test", "testtest", "apple", "orange", "cat"}
	fmt.Println(Filter(strs, func(s string) bool { return len(s) > 4 }))
}

func Filter[T any](slice []T, test func(T) bool) []T {
	new := []T{}
	for _, val := range slice {
		if test(val) {
			new = append(new, val)
		}
	}
	return new
}
