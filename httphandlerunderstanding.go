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

func (d database) bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s: %s", "bar", d["bar"])
}

func (d database) foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s: %s", "foo", d["foo"])
}

func (d database) baz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s: %s", "baz", d["baz"])
}

func main() {
	var db = database{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", db.foo)
	mux.HandleFunc("/bar", db.bar)
	mux.Handle("/baz", http.HandlerFunc(db.baz))

	http.ListenAndServe(":8080", mux)
}
