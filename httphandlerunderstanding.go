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
	switch r.URL.Path {
	case "/foo":
		fmt.Fprintf(w, "%s: %s", "foo", db["foo"])
	case "/bar":
		fmt.Fprintf(w, "%s: %s", "bar", db["bar"])
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s Not Found!", r.URL.Path)
	}
}

func main() {
	db := database{"foo": 1, "bar": 2}
	http.ListenAndServe(":8080", db)
}
