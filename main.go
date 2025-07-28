package main

import (
	"fmt"
	"net/http"
)

// "strings"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// this is when user is try to access the / default path of Servers
	fmt.Fprintln(w, "Welcome to Home Page")
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	// when http request is made to this ,do the following things or send it
	fmt.Fprintln(w, " Welcome to About Us pages bro ....")
}
func ContactUs(w http.ResponseWriter, r *http.Request) {
	// when http request is made to this do the following things
	fmt.Fprintln(w, "welcome to Contact Us page ....")
}

func main() {
	// backend Journey Started now ....
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutUs)
	http.HandleFunc("/contact", ContactUs)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
