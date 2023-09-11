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
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			// rw.WriteHeder(http.StatusBadRequest)
			// rw.Write([]bytes("Ooops"))
			return
		})

		// log.Printf("Data %s\n", d)


		fmt.Fprintf("rw, Hello %s", )

	}

	http.HandleFunc(*/learning, func(w http.ResponseWriter, r *http.Request) {
		log.Println("Learning Go")

	})


	http.ListenAndServe(":8080", nil)
}