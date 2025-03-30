package day4

import "fmt"

func StudyArray() {
	arr := [5]int{2, 3, 4, 5, 6}
	fmt.Println("arr=", arr, "length=", len(arr))
	fmt.Println("Value of index2:", arr[2])
	// array is value type, so when assigned to new variable, copy it
	arr1 := arr
	fmt.Println("arr1=", arr1)
	arr1[2] = 40
	fmt.Println("original arr=", arr)
	fmt.Println("New arr1=", arr1, "capacity=", cap(arr1))
}
