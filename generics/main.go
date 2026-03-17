package main

import "fmt"

type Number interface {
	~int64 | float64
}

type Numbers[T Number] []T

type CustomInt int64

func (ci CustomInt) IsPositive() bool {
	return ci > 0
}

func main() {

	// showSum()
	// showContains()
	// showAny()
	//unionInterfaceAndType()
	typeApproximation()

}

func typeApproximation() {
	customInts := []CustomInt{1, 2, 3, 5, 6}
	castedInts := make([]int64, len(customInts))

	for idx, val := range customInts {
		castedInts[idx] = int64(val)
	}

	fmt.Println(sumUnionInterface(customInts))
	fmt.Println(sumUnionInterface(castedInts))
}

// -----------------------------------------------------------------------------------------

func unionInterfaceAndType() {
	var ints Numbers[int64]
	ints = append(ints, []int64{1, 2, 3, 4, 5}...)

	floats := Numbers[float64]{1.0, 2, 5, 3, 5}

	fmt.Println(sumUnionInterface(ints))
	fmt.Println(sumUnionInterface(floats))
}

func sumUnionInterface[V Number](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}

	return sum
}

// -----------------------------------------------------------------------------------------

func showAny() {
	show(1, 2, 3)
	show("test1", "test2", "test3")
	show([]int64{1, 2, 3}, []int64{4, 5, 6})
	show(map[string]int64{
		"first":  1,
		"second": 2,
	})
	show(interface{}(1), interface{}("string"), any(struct{ name string }{name: "Sanzhar"}))
}

func show[T any](entitles ...T) {
	fmt.Println(entitles)
}

// -----------------------------------------------------------------------------------------

func showContains() {
	type Person struct {
		name     string
		age      int
		jobTitle string
	}

	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println(contains(ints, 4))

	strings := []string{"Sancho", "Dancho", "Erow", "Tima"}
	fmt.Println(contains(strings, "Tima"))
	fmt.Println(contains(strings, "Beka"))

	people := []Person{
		{
			name:     "Sanzhar",
			age:      23,
			jobTitle: "Backend Developer",
		},
		{
			name:     "Daniyar",
			age:      23,
			jobTitle: "Crypto Expert",
		},
	}

	fmt.Println(contains(people, Person{name: "Sanzhar", age: 23, jobTitle: "Backend Developer"}))
	fmt.Println(contains(people, Person{name: "Sanzhar", age: 22, jobTitle: "Backend Developer"}))
}

func contains[T comparable](elements []T, searchEl T) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}
	return false
}

// -----------------------------------------------------------------------------------------
func showSum() {
	floats := []float64{1.0, 2.0, 3.0}
	ints := []int64{1, 2, 3}
	// wrongFloats := []interface{}{"string", struct{}{}, true}

	fmt.Println(sum(floats))
	fmt.Println(sum[int64](ints))
	// fmt.Println(sum(wrongFloats))
}

func sum[V int64 | float64](numbers []V) V {
	var sum V
	for _, val := range numbers {
		sum += val
	}
	return sum
}
