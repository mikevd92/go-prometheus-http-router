package service

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"solution/service/metrics"
	"strconv"
)

type service struct {
	httpRequestTotal    *prometheus.CounterVec
	httpErrorCount      *prometheus.CounterVec
	httpRequestDuration *prometheus.HistogramVec
}

var (
	instance *service
	registry *prometheus.Registry
)

func init()  {
	httpRequestDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "dca_api",
		Name:      "response_time",
		Help:      "The lantency of the HTTPBase request",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method"})
	httpErrorCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "dca_api",
		Name:      "error_count",
		Help:      "The frequency of http error status codes",
	}, []string{"handler", "method", "status"})
	httpRequestTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "dca_api",
		Name:      "http_requests_total",
		Help:      "The frequency of http requests",
	}, []string{"handler", "method"})
	instance = &service{
		httpRequestDuration: httpRequestDuration,
		httpErrorCount:      httpErrorCount,
		httpRequestTotal:    httpRequestTotal,
	}
	registry = prometheus.NewRegistry()

	//goMetrics := collectors.NewGoCollector()
	goProcessMetrics := collectors.NewProcessCollector(collectors.ProcessCollectorOpts{Namespace: "dca_api"})

	//default metrics
	registry.MustRegister(goProcessMetrics)
	//registry.MustRegister(goMetrics)

	//custom metrics
	registry.MustRegister(instance.httpRequestTotal)
	registry.MustRegister(instance.httpErrorCount)
	registry.MustRegister(instance.httpRequestDuration)
}

func PrometheusService() *service{
	return instance
}

func GetRegistry() *prometheus.Registry {
	return registry
}

func (s *service) SaveHTTPCount(h *metrics.HTTPBase) {
	s.httpRequestTotal.WithLabelValues(h.Handler, h.Method).Inc()
}

func (s *service) SaveHTTPErrorCount(h *metrics.HTTPError) {
	s.httpErrorCount.WithLabelValues(h.Handler, h.Method, strconv.Itoa(h.Status)).Inc()
}

func (s *service) SaveHTTPDuration(h *metrics.HTTPDuration) {
	s.httpRequestDuration.WithLabelValues(h.Handler, h.Method).Observe(h.Duration)
}
