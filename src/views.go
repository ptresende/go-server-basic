package main


import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
    	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func perimeter(w http.ResponseWriter, r *http.Request, rect Rectangle) {
	res := rect.perimeter()
	fmt.Fprintf(w, "Here is the perimeter: %d", res)
}


func main() {
	rect := Rectangle{1, 2}

    	http.HandleFunc("/", handler)
	http.HandleFunc("/perimeter", func(w http.ResponseWriter, r *http.Request) {
		perimeter(w, r, rect)
	})

    	http.ListenAndServe(":8081", nil)
}
