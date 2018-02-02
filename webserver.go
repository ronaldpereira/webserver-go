package main

import "net/http"

func main ()  {
	http.HandleFunc("/randomize", randomic)
	http.ListenAndServe(":8080", nil)
}