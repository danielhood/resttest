package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "you have hit the root")
}

func GetName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	age, err := strconv.Atoi(vars["age"])
	if err != nil {
		age = 0
	}

	fmt.Fprintln(w, GetNameForAge(age))
}

func GetNameForAge(age int) string {
	return "somebody"
}
