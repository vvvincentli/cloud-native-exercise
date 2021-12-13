package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	NameSpace = "http-server"
)

var (
	functinLatency = CreateExecutionTimeMetric(NameSpace, "time spent.")
)

type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	start time.Time
	last  time.Time
}

func NewTimer() *ExecutionTimer {
	return NewExecutionTimer(functinLatency)
}

func Register() error {
	err := prometheus.Register(functinLatency)
	return err
}

func NewExecutionTimer(h *prometheus.HistogramVec) *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo: h,
		start: now,
		last:  now,
	}
}

func (t *ExecutionTimer) ObserveTotal() {
	(*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}

func CreateExecutionTimeMetric(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "execution_latency_seconds",
			Help:      help,
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		}, []string{"step"},
	)
}
