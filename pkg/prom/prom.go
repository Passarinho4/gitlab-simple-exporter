package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	PipelineCounter    *prometheus.CounterVec
	PipelineDurations  *GaugeVecTtl
	PipelineTimestamps *GaugeVecTtl
	BuildDurations     *GaugeVecTtl
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
				Help: "Duration in seconds of the most recent pipeline",
			},
			[]string{
				"repo",
				"branch",
				"pipeline_id",
			},
			600,
		),
		PipelineTimestamps: NewGaugeVecTtl(
			prometheus.GaugeOpts{
				Name: "gitlab_ci_pipeline_job_timestamp",
				Help: "Creation date timestamp of the most recent pipeline",
			},
			[]string{
				"repo",
				"branch",
				"pipeline_id",
			},
			600,
		),
		BuildDurations: NewGaugeVecTtl(
			prometheus.GaugeOpts{
				Name: "gitlab_ci_build_duration",
				Help: "Duration in seconds of the most recent job",
			},
			[]string{
				"repo",
				"branch",
				"pipeline_id",
				"build_stage",
				"build_status",
			},
			600,
		),
	}
	reg.MustRegister(m.PipelineCounter)
	reg.MustRegister(m.PipelineDurations.gaugeVec)
	reg.MustRegister(m.PipelineTimestamps.gaugeVec)
	reg.MustRegister(m.BuildDurations.gaugeVec)
	return m
}
