package main
import "fmt"
import "unicode/utf8"

func main(){
	// var num int = 12
	// var num2 int32 = 23
	// var num3 int64 = 2335
	// var num4 int16 = 134

	var intNum int = 12345
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 1234.53
	fmt.Println(floatNum)

	var result float32 = float32(floatNum) + 12.3
	fmt.Println(result)

	var myString string = "Hello World"
	fmt.Println(myString)

	var secondString string = "Hello" + "World"
	fmt.Println(secondString)

	fmt.Println(len(myString)) // length of string
	fmt.Println(utf8.RuneCountInString(myString)) // number of runes in string

	var myRune rune = 'A'
	fmt.Println(myRune) // rune is an alias for int32

	var myBool bool = true
	fmt.Println(myBool)

	v1, v2 := 1, 2
	fmt.Println(v1, v2)

	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float64 = 3.141
	fmt.Println(pi)
}