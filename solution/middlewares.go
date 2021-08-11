package main

import (
	"net/http"
	"solution/service"
	"solution/service/metrics"
)

func HTTPActivityLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		duration := metrics.NewHTTPDuration(r.RequestURI, r.Method)
		duration.Started()
		next.ServeHTTP(w, r)
		duration.Finished()
		httpBase := metrics.NewHTTPBase(r.RequestURI, r.Method)
		prometheusService := service.PrometheusService()
		prometheusService.SaveHTTPDuration(duration)
		prometheusService.SaveHTTPCount(httpBase)
	})
}
