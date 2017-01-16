package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() error {
	filename := page.Title + ".txt"
	return ioutil.WriteFile(filename, page.Body, 0600)
}

func load(filename string) (*Page, error) {
	body, err := ioutil.ReadFile(filename + ".txt")
	if err != nil {
		return nil, err
	}
	return &Page{Title: filename, Body: body}, nil
}

func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(req.URL.Path[1:]))
}

func main() {
	p1 := Page{Title: "hello", Body: []byte("first body content")}
	p1.save()
	p2, _ := load("hello")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
