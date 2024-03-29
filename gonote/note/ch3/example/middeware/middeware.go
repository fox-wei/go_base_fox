package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func validateAuth(s string) error {
	if s != "123456" {
		return fmt.Errorf("%s", "bad auth token")
	}

	return nil
}

func greetings(w http.ResponseWriter, r *http.Request) {
	s := time.Now()
	t := fmt.Sprint(s)
	fmt.Fprintf(w, "Hello, ying!!! "+t)
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)
		h.ServeHTTP(w, r)
	})
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := validateAuth(r.Header.Get("auth"))
		if err != nil {
			http.Error(w, "bad auth param", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.ListenAndServe(":9090", logHandler(authHandler(http.HandlerFunc(greetings))))
}
