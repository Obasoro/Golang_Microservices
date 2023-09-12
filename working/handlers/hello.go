package handlers

import {
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
}

type Hello struct {

	l * log.Logger

}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h*Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, _ := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			// rw.WriteHeder(http.StatusBadRequest)
			// rw.Write([]bytes("Ooops"))
			return
		}

		// log.Printf("Data %s\n", d)


		fmt.Fprintf("rw, Hello %s", )

}
