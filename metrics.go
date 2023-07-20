package main

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
)

// Supported prometheus metrics
type Metrics struct {
	httpRequestsDuration *prometheus.HistogramVec
	httpRequestsTotal    *prometheus.CounterVec
}

var metricsRegistry *prometheus.Registry
var metrics *Metrics

// Register prometheus metrics
func registerMetrics() {
	metricsRegistry = prometheus.NewRegistry()
	reg := metricsRegistry

	metrics = &Metrics{
		httpRequestsDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "http_requests_duration",
				Help: "Total HTTP request duration",
			}, []string{"method", "code", "path"}),

		httpRequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			}, []string{"method", "code", "path"}),
	}

	reg.MustRegister(metrics.httpRequestsDuration)
	reg.MustRegister(metrics.httpRequestsTotal)
}

// Request metrics handler wraps prometheus metrics evaluation for HTTP
// request hadler. It currently evaluates count and duration.
func requestMetrics(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		st := time.Now()

		// http metric labels
		labels := prometheus.Labels{
			"method": string(ctx.Method()),
			"code":   strconv.Itoa(ctx.Response.StatusCode()),
			"path":   string(ctx.Path()),
		}

		next(ctx)
		duration := time.Since(st).Seconds()

		// set metrics
		metrics.httpRequestsTotal.With(labels).Inc()
		metrics.httpRequestsDuration.With(labels).Observe(duration)
	}
}
