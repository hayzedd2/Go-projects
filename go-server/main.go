package main

import (
	"fmt"
	"log"
	"net/http"
)

func formServer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform error %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful!")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w,"Name recived is %v\n", name)
	fmt.Fprintf(w,"Email received is %v\n", email)
}
func helloServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not allowed", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello!")

}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formServer)
	http.HandleFunc("/hello", helloServer)

	fmt.Printf("Starting server at 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
