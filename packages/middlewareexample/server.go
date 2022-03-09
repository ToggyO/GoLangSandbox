package middlewareexample

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Run() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	indexHandler := http.HandlerFunc(index)
	WelcomeHandler := http.HandlerFunc(welcome)

	mux.Handle("/index", IndexLogger(indexHandler))
	mux.Handle("/welcome", CheckRequestMethod(IndexLogger(WelcomeHandler)))

	log.Printf("Serving http on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func IndexLogger(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("%s as %s", r.Method, r.URL.Path)
		next.ServeHTTP(rw, r)
		log.Printf("Finished %s in %v", r.URL.Path, time.Since(startTime))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello welcome to the index page")
}

func welcome(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome to the welcome page")
}

func CheckRequestMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Printf("Only GET requests are accepted on the path!")
			next.ServeHTTP(rw, r)
			fmt.Fprintf(rw, "Method not allowed!")
		} else {
			log.Printf("Method allowed")
			next.ServeHTTP(rw, r)
			fmt.Fprintf(rw, "Your request is valid!")
			log.Printf("Finished checking......")
		}
	})
}
