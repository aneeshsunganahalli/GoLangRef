package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("/.contactInfo.json")
	fmt.Printf("\n%v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("/purchase.json")
	fmt.Printf("\n%v", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T{
	var data, _ = os.ReadFile(filePath)
	var loaded =[]T{}
	json.Unmarshal(data, &loaded)
	return loaded
}