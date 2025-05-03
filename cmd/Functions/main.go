package main
import (
	"errors"
	"fmt"
)

func main() {
	var myString string = "Hello World"
	printMe(myString)

	num := 12
	denom := 3
	var result, remainder, err = intDivision(num, denom)
	if err != nil {
		fmt.Printf("%s", err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result: %v\n", result)
	}
	fmt.Printf("Result: %v, Remainder: %v\n", result, remainder)

	switch remainder {
	case 0:
		fmt.Println("Remainder is zero")
	case 1,2:
		fmt.Println("Remainder is 1 or 2")
	default:
		fmt.Println("Remainder is something else")
	}
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int,int,error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}