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

	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /index")
		http.Redirect(w, r, "data", 307)
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /data")
		//w.WriteHeader(204)
		w.Write([]byte(`
			<body>
				<h5 style="text-align: center;">You got some DATA</h5>
				<a href="info">INFO</a>
			</body>
		`))
	})

	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /info")
		//w.WriteHeader(204)
		w.Write([]byte(`
			<body>
				<h5 style="text-align: center;">You got some INFO</h5>
				<a href="data">DATA</a>
			</body>
		`))
	})

	fmt.Println("LISTEN on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))

	fmt.Println("EXITING...")
}
