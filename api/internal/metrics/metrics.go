package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Compteur simple pour les requêtes HTTP
	RequestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Nombre total de requêtes HTTP par endpoint",
		},
		[]string{"method", "path"},
	)
)
