package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("STARTING...")

	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /data")
		//w.WriteHeader(204)
		w.Write([]byte("<body><h5 style=\"text-align: center;\">You got some DATA</h5></body>"))
	})

	fmt.Println("LISTEN on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))

	fmt.Println("EXITING...")
}
