package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	appleResult := []int32{}
	for _, apple := range apples {
		newValue := apple + a
		if newValue >= s && newValue <= t {
			appleResult = append(appleResult, newValue)
		}
	}
	orangeResult := []int32{}
	for _, orange := range oranges {
		newValue := orange + b
		if newValue >= s && newValue <= t {
			orangeResult = append(orangeResult, newValue)
		}
	}
	fmt.Println(len(appleResult))
	fmt.Println(len(orangeResult))
}

func appleAndOrangeSample() {
	s := int32(7)
	t := int32(11)
	a := int32(5)
	b := int32(15)
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

	countApplesAndOranges(
		s,
		t,
		a,
		b,
		apples,
		oranges,
	)
}
