package main

import (

	"log"
	"net/http"
	"github.com/nicholasjackson/building-microservices-youtube/product-api/handlers"
	"os/signal"
	"os"
	"context"
	"time"

)

// func main(){

// 	http.HandleFunc(*/*, func(rw http.ResponseWriter, r *http.Request) {
// 		log.Println("Hello World")
// 		d, _ := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(rw, "Ooops", http.StatusBadRequest)
// 			// rw.WriteHeder(http.StatusBadRequest)
// 			// rw.Write([]bytes("Ooops"))
// 			return
// 		})

// 		// log.Printf("Data %s\n", d)


// 		fmt.Fprintf("rw, Hello %s", )

// 	}

// 	http.HandleFunc(*/learning, func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Learning Go")

// 	})


// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello()
	gh := handlers.Goodbye()


	sm := http.NewServerMux()
	sm.Handler("/", hh)
	sm.Handler("/Goodbye", gh)

	s := &http.Server{
		Addr: ":9090"
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.kill)

	sig := <- sigChan
	l.Println("recieved terminate, graceful shutdown", sig)

	tc := context.WithTimeout(context.Background(), 30*time.second)

	s.shutdown



	http.ListenAndServe("9090", nil, sm)
}