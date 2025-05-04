package main

import (
	"fmt"
)
// Redudant to make similar functions for diff types, so use generics
// func main() {
// 	var intSlice = []int{1, 2, 3}
// 	fmt.Println(sumIntSlice(intSlice))

// 	var floatSlice = []float32{1, 2, 3}
// 	fmt.Println(sumFloatSlice(floatSlice))
// }

// func sumIntSlice(slice []int) int{
// 	var sum int
// 	for _,v := range slice{
// 		sum += v
// 	}
// 	return sum
// }

// func sumFloatSlice(slice []float32) float32{
// 	var sum float32
// 	for _,v := range slice{
// 		sum += v
// 	}
// 	return sum
// }

// Generics work similarly int TS
func main(){
	var intSlice = []int{1,2,3}
	var floatSlice = []float32{1,2,3}
	fmt.Println(sumSlice(intSlice))
	fmt.Println(sumSlice(floatSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _,v := range slice{
		sum += v
	}
	return sum
}