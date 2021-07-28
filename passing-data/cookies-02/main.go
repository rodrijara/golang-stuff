package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", set)
	http.HandleFunc("/hmt", hmt)
	http.ListenAndServe(":8080", nil)
}

var counter int

func set(res http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("cnt")

	if err != nil {
		log.Println(err.Error())
		http.SetCookie(res, &http.Cookie{
			Name:  "cnt",
			Value: fmt.Sprintf("n times: %d", counter),
		})
	} else {
		counter++
		http.SetCookie(res, &http.Cookie{
			Name:  "cnt",
			Value: fmt.Sprintf("n times: %d", counter),
		})
	}

}

func hmt(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "text/html")
	fmt.Fprintf(res, "<h1>This is the %d-th time you access %v%v<h1>", counter, req.Host, req.URL)
}
