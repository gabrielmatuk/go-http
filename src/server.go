package main

import (
	"log"
	"net/http"
)

const PORT string = ":3333"

func main() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	log.Print("Server listening on ", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
