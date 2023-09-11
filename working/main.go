package main

import {

	"net/http"
	"log"
	"io/ioutil"

}

func main(){

	http.HandleFunc(*/*, func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, _ := ioutil.ReadAll(r.Body)

		log.Printf("Data %s\n", d)

	}

	http.HandleFunc(*/learning, func(w http.ResponseWriter, r *http.Request) {
		log.Println("Learning Go")

	})


	http.ListenAndServe(":8080", nil)
}