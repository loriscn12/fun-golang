// This is the first Golang code ever written.

package main


import (
	"fmt"
	"log"
	"net/http"
)
	

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is the first Golang project!!") // write data to response
}


func main() {
	http.HandleFunc("/", welcomePage) // setting router rule

	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

