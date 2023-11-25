package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	PipelineCounter   *prometheus.CounterVec
	PipelineDurations *GaugeVecTtl
	BuildDurations    *GaugeVecTtl
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		PipelineCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "gitlab_ci_pipeline_total",
				Help: "Number of executed pipelines",
			},
			[]string{"repo", "branch", "status"},
		),
		PipelineDurations: NewGaugeVecTtl(
			prometheus.GaugeOpts{
				Name: "gitlab_ci_pipeline_duration",
				Help: "Duration of pipeline execution",
			},
			[]string{
				"repo",
				"branch",
				"pipeline_id",
			},
			300,
		),
		BuildDurations: NewGaugeVecTtl(
			prometheus.GaugeOpts{
				Name: "gitlab_ci_build_duration",
				Help: "Duration of build in pipeline execution",
			},
			[]string{
				"repo",
				"branch",
				"pipeline_id",
				"build_stage",
				"build_status",
			},
			300,
		),
	}
	reg.MustRegister(m.PipelineCounter)
	reg.MustRegister(m.PipelineDurations.gaugeVec)
	reg.MustRegister(m.BuildDurations.gaugeVec)
	return m
}
