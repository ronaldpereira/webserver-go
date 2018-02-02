package main

import	(
		"fmt"
		"time"
		"net/http"
		"math/rand"
		)

func randomic (writer http.ResponseWriter, request *http.Request) {
	rand.Seed(time.Now().UnixNano()) // Random seed based on actual system time
	number := rand.Intn(100)
	fmt.Fprintf(writer, "%d\n", number)
}

func main ()  {
	http.HandleFunc("/randomize", randomic)
	http.ListenAndServe(":8080", nil)
}