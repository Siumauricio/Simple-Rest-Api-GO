package main

import (
	"fmt"
	"net/http"
)

type person struct {
    first string
    last  string
    age   int
}


//func main() {
//
//	//routes
//	http.HandleFunc("/", homeHandler)
//	http.HandleFunc("/contact", contactHandler)
//	//start the server
//	http.ListenAndServe(":3000",nil)
//
//}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	fmt.Println("Hello, World!")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    p1 := person{
        first: "James",
        last:  "Bond",
        age:   32,
    }
	w.Write([]byte ("My name is " + p1.last+" and I am %v years old."))
}