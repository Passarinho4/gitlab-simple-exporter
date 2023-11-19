package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	PipelineCounter         *prometheus.CounterVec
	SuccessPipelinesCounter *prometheus.CounterVec
	FailedPipelinesCounter  *prometheus.CounterVec
	PipelineDurations       *prometheus.GaugeVec
	BuildDurations          *prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		PipelineCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "pipeline_counter",
				Help: "Number of executed pipelines",
			},
			[]string{"namespace", "project_name", "branch"},
		),
		SuccessPipelinesCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "success_pileline_counter",
				Help: "Number of successfully executed pipelines",
			},
			[]string{"namespace", "project_name", "branch"},
		),
		FailedPipelinesCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "failed_pipeline_counter",
				Help: "Number of unsuccessfully executed pipelines",
			},
			[]string{"namespace", "project_name", "branch"},
		),
		PipelineDurations: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "pipelines_durations",
				Help: "Duration of pipeline execution",
			},
			[]string{
				"namespace",
				"project_name",
				"branch",
				"pipeline_id",
			},
		),
		BuildDurations: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "builds_durations",
				Help: "Duration of build in pipeline execution",
			},
			[]string{
				"namespace",
				"project_name",
				"branch",
				"pipeline_id",
				"build_stage",
			},
		),
	}
	reg.MustRegister(m.PipelineCounter)
	reg.MustRegister(m.SuccessPipelinesCounter)
	reg.MustRegister(m.FailedPipelinesCounter)
	reg.MustRegister(m.PipelineDurations)
	reg.MustRegister(m.BuildDurations)
	return m
}
