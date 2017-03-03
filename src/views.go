package main


import (
	"fmt"
	"net/http"
	"strconv"
)

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Counting: %d!\n", i)
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
    	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func perimeter(w http.ResponseWriter, r *http.Request) {
	go counter()
	length, _ := strconv.Atoi(r.URL.Query().Get("length"))
	width, _ := strconv.Atoi(r.URL.Query().Get("width"))
	rect := Rectangle{length, width}
	res := rect.perimeter()
	fmt.Fprintf(w, "Here is the perimeter: %d\n", res)
	fmt.Printf("Here is the perimeter: %d\n", res)
}

func area(w http.ResponseWriter, r *http.Request) {
	length, _ := strconv.Atoi(r.URL.Query().Get("length"))
	width, _ := strconv.Atoi(r.URL.Query().Get("width"))
	rect := Rectangle{length, width}
	res := rect.area()
	fmt.Fprintf(w, "Here is the area: %d\n", res)
	fmt.Printf("Here is the area: %d\n", res)
}

func main() {
    	http.HandleFunc("/", handler)
	http.HandleFunc("/perimeter", func(w http.ResponseWriter, r *http.Request) {
		perimeter(w, r)
		area(w, r)
	})

    	http.ListenAndServe(":8080", nil)
}
