package day4

import (
	"fmt"
	"maps"
)

func StudyMap() {
	person := map[string]string{
		"name":    "Ram",
		"address": "ktm",
	}
	fmt.Println("person=", person, "length=", len(person))
	name := person["name"]
	fmt.Println("name of person=", name)

	for key, value := range person {
		fmt.Println("Key=", key, "value=", value)
	}

	//Make in map

	student := make(map[string]any, 0)
	fmt.Println("length of student=", len(student))
	student["name"] = "bob"
	student["roll"] = 40
	student["status"] = true
	fmt.Println("student=", student)
	//deleting a key-vlaue
	delete(student, "status")
	fmt.Println("After delete=", student)

	for key, value := range student {
		fmt.Println("key=", key, "value=", value)
	}

	if v, ok := student["status"]; ok {
		fmt.Println("status exist", v)
	} else {
		fmt.Println("status does not exist")
	}

	// Appending map using Idiomatic way

	mymap1 := map[string]interface{}{
		"address": "ktm",
	}

	mymap := map[string]interface{}{
		"name": "ram",
		"age":  12,
	}
	for key, value := range mymap {
		mymap1[key] = value
	}
	fmt.Println(mymap1)

	//In simple
	println("\n Using simple method")
	maps.Copy(mymap1, mymap)
	fmt.Println(mymap1)
}
