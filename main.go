package main

import "fmt"

// import (
// 	"fmt"
// 	"shovab/day4"
// )

func main() {

	// day4.StudyArray()
	// fmt.Println("\nThis shows slice:")
	// day4.StudySlice()
	// fmt.Println("\n ")
	// day4.StudyMap()
	Displaystucture()
	allBooks := Fetchallbook()
	fmt.Println(allBooks)
	var highestpricebook Book
	maxprice := 0
	for _, b := range allBooks {

		if maxprice < b.Price {
			maxprice = b.Price
			highestpricebook = b
		}
	}
	fmt.Println("The highest price book is ", highestpricebook)
}
