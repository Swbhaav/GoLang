package day4

import "fmt"

func StudySlice() {
	//slice
	slc := []int{2, 3, 4, 5, 6}
	fmt.Println("slc=", slc, "length", len(slc))
	fmt.Println("Value of index2 ", slc[2])

	slc1 := slc
	fmt.Println("slc1", slc1)
	slc1[2] = 40
	fmt.Println("Original", slc)
	fmt.Println("New slc1=", slc1, "Capacity=", cap(slc1))

	//empty slice
	copiedSlc := make([]int, 5)
	copy(copiedSlc, slc)
	fmt.Println("Copied slc=", copiedSlc)
	copiedSlc[0] = 99
	fmt.Println("Original", slc)
	fmt.Println("New copied slc=", copiedSlc)

	//Make function, syntax= make(type, length, capacity)

	slcnew := make([]string, 3, 5)
	fmt.Println("slc new=", slcnew, len(slcnew), "capacity=", cap(slcnew))
	slcnew = append(slcnew, "40", "Hello", "hi")
	fmt.Println("Slc after append:", slcnew, "length", len(slcnew), "capacity", cap(slcnew))

}
