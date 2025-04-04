package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from go API"))
	})
	router.HandleFunc("/students", handleGetAllStudents).Methods("GET")
	fmt.Println("Starting server on port 8060...")
	http.ListenAndServe(":8060", router)
}
func handleGetAllStudents(w http.ResponseWriter, r *http.Request) {
	studentSlice := []string{"John", "Doe", "Jane"}
	w.Header().Set("Content-Type", "application/json")
	studentByte, err := json.Marshal(studentSlice)
	if err != nil {
		w.Write([]byte("Error in getting student data"))
		return
	}
	w.Write(studentByte)
}

// import (
// 	customerror "anish/golang/customError"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"math"
// 	"strings"
// )

// func main() {
// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			fmt.Println("recovered from panic:", r)
// 		}
// 	}()

// 	resSqrt, err := ReturnSquareRoot(56)
// 	if err != nil {
// 		log.Println("error occured:", err.Error())
// 		panic(err.Error())
// 	}
// 	log.Println("result of square root:", resSqrt)

// 	name, error := CheckString("an")
// 	if error != nil {
// 		fmt.Println(error)
// 		return
// 	} else {
// 		fmt.Println(name)
// 	}

// 	e := customerror.ValidateFields("kd", "age")
// 	if e != nil {
// 		panic(e.Error())
// 	}

// }

// func ReturnSquareRoot(x float64) (float64, error) {
// 	if x < 5 {
// 		return 0, errors.New("Negative value not allowed.")
// 	}
// 	return math.Sqrt(x), nil
// }

// func CheckString(get string) (string, error) {
// 	if len(get) < 3 {
// 		return "", errors.New("Name is less than 3")
// 	}
// 	return strings.ToUpper(get), nil
// }
