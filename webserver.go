package main

import	(
		"fmt"
		"time"
		"net/http"
		"math/rand"
		)

func randomic (writer http.ResponseWriter, request *http.Request) {
	rand.Seed(time.Now().UnixNano()) // System time seed
	number := rand.Intn(100)

	if number < 90 { // 90% chance (from 0 to 89) of printing a random number
		fmt.Fprintf(writer, "%d", rand.Intn(1000000))
	} else { // 10% chance (from 90 to 99) of printing a random string
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") // Possible letters to mount the random string
		strSize := rand.Intn(199) + 1 // Random size from 1 to 200
		str := make([]rune, strSize)

		for i := range str {
			str[i] = letters[rand.Intn(len(letters))] // Pick a random letter from letters
		}

		fmt.Fprintf(writer, "%s", string(str))
	}
}

func main ()  {
	http.HandleFunc("/randomize", randomic)
	http.ListenAndServe(":8080", nil)
}