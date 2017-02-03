package main

import (
	"fmt"
	"net/http"
)

type pound float32
type database map[string]pound

func (p pound) String() string {
	return fmt.Sprintf("£%.2f", p)
}

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"foo": 1, "bar": 2}
	http.ListenAndServe(":8080", db)
}
