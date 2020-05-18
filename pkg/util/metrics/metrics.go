package metrics

import (
	"github.com/cuijxin/mysql-operator/pkg/cluster/innodb"
	"github.com/prometheus/client_golang/prometheus"
)

var podName string
var clusterName string

// RegisterPodName will set the name of the current pod.
func RegisterPodName(name string) {
	podName = name
}

// RegisterClusterName will set the name of the current cluster.
func RegisterClusterName(name string) {
	clusterName = name
}

// RegisterOperatorMetric will register a single operator metric.
func RegisterOperatorMetric(metric prometheus.Collector) {
	assertPodName()
	prometheus.MustRegister(metric)
}

// RegisterAgentMetric will register a single agent metric.
func RegisterAgentMetric(metric prometheus.Collector) {
	assertPodName()
	assertClusterName()
	prometheus.MustRegister(metric)
}

// NewOperatorEventCounter will build a new prometheus.CounterVec.
func NewOperatorEventCounter(name string, help string) *prometheus.CounterVec {
	return newCounter("mysql_operator", "cluster", name, help, []string{"podName"})
}

// NewOperatorEventGauge will build a new prometheus.GaugeVec.
func NewOperatorEventGauge(name string, help string) *prometheus.GaugeVec {
	return newGauge("mysql_operator", "cluster", name, help, []string{"podName"})
}

// NewAgentEventCounter will build a new prometheus.CounterVec.
func NewAgentEventCounter(name string, help string) *prometheus.CounterVec {
	return newCounter("mysql", "innodb", name, help, []string{"podName", "clusterName"})
}

// NewAgentStatusCounter will build a new prometheus.CounterVec.
func NewAgentStatusCounter(name string, help string) *prometheus.CounterVec {
	return newCounter("mysql", "innodb", name, help, []string{"podName", "clusterName", "instanceStatus"})
}

// IncEventCounter will increment a counter and set appropriate labels.
func IncEventCounter(counter *prometheus.CounterVec) {
	counter.With(eventLabels()).Inc()
}

// IncEventGauge will increment a gauge and set appropriate labels.
func IncEventGauge(gauge *prometheus.GaugeVec) {
	gauge.With(eventLabels()).Inc()
}

// DecEventGauge will decrement a gauge and set appropriate labels.
func DecEventGauge(gauge *prometheus.GaugeVec) {
	gauge.With(eventLabels()).Dec()
}

// IncStatusCounter will increment a counter and set appropriate labels.
func IncStatusCounter(counter *prometheus.CounterVec, status innodb.InstanceStatus) {
	counter.With(statusLabels(status)).Inc()
}

func newCounter(namespace string, subsystem string, name string, help string, labels []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		labels,
	)
}

func newGauge(namespace string, subsystem string, name string, help string, labels []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		labels,
	)
}

func assertPodName() {
	if podName == "" {
		panic("Metrics package requires podName. Unable to register metrics")
	}
}

func assertClusterName() {
	if clusterName == "" {
		panic("Metrics package requires clusterName. Unable to register metrics")
	}
}

func eventLabels() prometheus.Labels {
	labels := prometheus.Labels{
		"podName": podName,
	}
	if clusterName != "" {
		labels["clusterName"] = clusterName
	}
	return labels
}

func statusLabels(status innodb.InstanceStatus) prometheus.Labels {
	return prometheus.Labels{
		"podName":        podName,
		"clusterName":    clusterName,
		"instanceStatus": string(status),
	}
}
