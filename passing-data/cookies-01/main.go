package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// write a cookie
func set(res http.ResponseWriter, req *http.Request) {
	cookie := &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	}
	http.SetCookie(res, cookie)
	//alternatively
	// http.SetCookie(res, &http.Cookie{
	// 	Name: "my-cookie",
	// 	Value: "some value",
	// })
	fmt.Fprintln(res, "Cookie WRITTEN, check your browser")
	fmt.Fprintln(res, "Go to devtools > application > cookies")
}

// read my cookie
func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(res, "YOUR COOKIE:", c)
}
