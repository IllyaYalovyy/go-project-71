package reverseproxy

import (
	"log"
	"net/http"
	"time"
)

func withLogging(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handling request '%v'", r.RequestURI)
		f(w, r)
	}
}

func withElapsedTime(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startedTime := time.Now()
		f(w, r)
		elapsedTime := time.Since(startedTime)
		log.Printf("Request hadling finished. Elapsed time: %v", elapsedTime)
	}
}
