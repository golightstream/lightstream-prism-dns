package dns64

import (
	"github.com/coredns/coredns/plugin"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// RequestsTranslatedCount is the number of DNS requests translated by dns64.
	RequestsTranslatedCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "dns",
		Name:      "requests_dns64_translated_total",
		Help:      "Counter of DNS requests translated by dns64.",
	}, []string{"server"})
)
