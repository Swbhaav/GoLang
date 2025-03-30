package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name  string
	Roll  int
	Marks float32
	Age   time.Time
}

func Displaystucture() {
	person := Person{
		Name:  "Ram",
		Roll:  20,
		Marks: 50.10,
		Age:   time.Date(2012, 2, 25, 6, 25, 22, 125, time.UTC),
	}
	fmt.Println(person)
	person.calculateAge()

}

func (p Person) calculateAge() {
	nowtime := time.Now()
	age := nowtime.Year() - p.Age.Year()
	fmt.Println("age=", age)

}
