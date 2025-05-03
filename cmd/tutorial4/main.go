package main
import (
	"fmt"
	"time"
)

func main() {

	var intArr [3]int32
	var arr [3]int32 = [3]int32{1,2,3} // or
	array := [...]int32{1,2,3}

	fmt.Println(array)

	fmt.Println(arr)

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// Slices are like arrays, can append to them & more flexible, used more commonly

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("Length and capacity are: %v and %v\n", len(intSlice), cap(intSlice))
	
	intSlice = append(intSlice, 7)
	fmt.Printf("Length and capacity are: %v and %v\n", len(intSlice), cap(intSlice))

	intSlice2 := []int32{8,9,10}
	intSlice =  append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3,8)
	fmt.Println(intSlice3)

	// Maps
	var myMap map[string]uint8 = make(map[string]uint8) // Refers to string keys and unsigned int as values
	fmt.Println(myMap)

	var map2 = map[string]uint8{"Adam": 23, "Sarah": 24}
	fmt.Println(map2["Sarah"])

	var age, ok = map2["Jason"]

	delete(map2, "Adam") // delete from map

	if ok {
		fmt.Printf("Age is %v\n", age)
	} else {
		fmt.Printf("Invalid\n")
	}

	// for loops
	for name := range map2 { 			// Iterates through the keys of the map
		fmt.Printf("Name: %v\n", name)
	}

	for name, age := range map2 {			// Iterates through keys and values
		fmt.Printf("Name: %v and Age: %v\n", name, age)
	}

	for i,v := range arr {
		fmt.Printf("Index: %v and Value: %v\n", i, v)
	}

	// No while loops built in but can use for loop sytax to acheive

	// for i:= 0; i < 10; i++ {
	// 	fmt.Print(i)
	// }

	// Try to specify capacity as its faster
	var n int= 10000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Without Preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("With Preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
