package main

import (
	"fmt"
	"strings"
)

func main() {
	// Uses underlying byte array for strings by default
	var myString = "Aresume"
	var first = myString[0]
	fmt.Printf("%v, %T\n", first, first)
	for i, v := range myString{
		fmt.Println(i,v)
	}

	// Can use rune array also, rune is an alais for int32
	var String = []rune("Awesome")
	var index = String[2]
	fmt.Printf("%v, %T\n", index, index)
	for i,v := range String {
		fmt.Println(i,v)
	}

	// String Building, keep in mind Strings are immutable (can't change after created)
	var strSlice = []string{"s","o","u","p"}
	var catStr = ""
	fmt.Println(strSlice)
	for i:= range strSlice{
		catStr += strSlice[i] // Obviosuly very inefficient as new string is formed everytime
	}
	fmt.Printf("%v", catStr)


	// Importing "strings"
	var str = []string {"q", "u", "e", "s", "t", "i", "o","n"}
	var strBuilder strings.Builder
	for i := range str {
		strBuilder.WriteString(str[i])	// String Builder 
	}
	var cat = strBuilder.String() // String created from underlying array which the values were appended to
	fmt.Printf("\n%v", cat)
}