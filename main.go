package main

import (
	"fmt"
	"shovab/books"
	my_interface "shovab/interface"
)

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

	// Displaystucture()
	// allBooks := Fetchallbook()
	// fmt.Println(allBooks)
	// var highestpricebook Book
	// maxprice := 0
	// for _, b := range allBooks {

	// 	if maxprice < b.Price {
	// 		maxprice = b.Price
	// 		highestpricebook = b
	// 	}
	// }
	// fmt.Println("The highest price book is ", highestpricebook)

	rect := my_interface.Rectangle{
		Length:  34,
		Breadth: 20,
	}

	my_interface.DisplayAreaAndPerimeter(rect)
	myCircle := my_interface.Circle{
		Radius: 40,
	}
	my_interface.DisplayAreaAndPerimeter(myCircle)

	//getting books
	allBooks := books.GetAllBooks()
	fmt.Println("all books:", allBooks)

	booksname := books.GetBookNameList()
	fmt.Println("Names:", booksname)
}
