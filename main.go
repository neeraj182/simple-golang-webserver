package main

import (
	"fmt"
	"log"
	"net/http" // import the net/http package
)

// func fomrhandler created
func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err) // parseForm() parses the form data
		return
	}
	fmt.Fprintf(w, "POST request successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Name = %s\n", address)
}

// func helloHandler created
func helloHandler(w http.ResponseWriter, r *http.Request) { // w here is the response writer and r is the request
	if r.URL.Path != "/hello" { // if the path is not /hello
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" { // if the request is not a GET request
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Golang World")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // telling golang to check directory ./static for files
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // listen and serve on port 8080
		log.Fatal(err)
	}
}
