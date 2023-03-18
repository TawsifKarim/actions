package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	const port = "8080"
	r.HandleFunc("/", rootHandler)

	fmt.Println("starting server at port:", 8080)
	err := http.ListenAndServe("localhost:"+port, r) //notice the r in place of nil
	if err != nil {
		log.Fatal("Something went wrong starting the server", err)
	}

}
