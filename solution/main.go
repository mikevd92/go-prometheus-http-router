package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"solution/service"
	"time"
)

func Metrics(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}


func main() {
	rand.Seed(time.Now().Unix())
	registry := service.GetRegistry()
	router := httprouter.New()
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	router.GET("/monitoring/metrics", Metrics(handler))
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", HTTPActivityLogger(router)))
}