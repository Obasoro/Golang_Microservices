package main

import "net/http"

func main(){

	http.HandleFunc(*/*, func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")

	})

	http.HandleFunc(*/*, func(w http.ResponseWriter, r *http.Request) {
		log.Println("Learning Go")

	})


	http.ListenAndServe(":8080", nil)
}