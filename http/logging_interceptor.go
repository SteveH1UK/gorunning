package http

import (
	"log"
	"net/http"
	"time"
)

func handlerLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(writer, request)
		log.Printf("%s - %s (%v)\n", request.Method, request.URL.Path, time.Since(startTime))
	})
}
