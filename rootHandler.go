package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("params:", r.URL.Query()) //extract from query param

	queryParam := r.URL.Query()

	bt, err := json.Marshal(queryParam)
	if err != nil {
		fmt.Println("json marshal error", err.Error())
	}

	val := r.FormValue("name")
	fmt.Println("value from body:", val)

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bt)
}

func add(a, b int) int {
	return a + b
}
