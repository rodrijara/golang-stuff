package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", set)
	http.HandleFunc("/hmt", hmt)
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}
	counter, _ := strconv.Atoi(cookie.Value)
	counter++
	cookie.Value = strconv.Itoa(counter)
	http.SetCookie(res, cookie)
}

func hmt(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("counter")
	res.Header().Set("content-type", "text/html")
	fmt.Fprintf(res, "<h1>This is the %v-th time you access %v%v<h1>", cookie.Value, req.Host, req.URL)
}
