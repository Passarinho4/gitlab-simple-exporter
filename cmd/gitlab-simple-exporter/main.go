package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/passarinho4/gitlab-simple-exporter/pkg/gitlab"
	"github.com/passarinho4/gitlab-simple-exporter/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var reg = prometheus.NewRegistry()
var metrics = prom.NewMetrics(reg)

func main() {

	http.HandleFunc("/webhook", handleGitlabHook)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleGitlabHook(w http.ResponseWriter, req *http.Request) {
	var r, err = gitlab.ParseGitlabHook(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	metrics.PipelineCounter.WithLabelValues(r.Project.Web_url, r.Object_attributes.Ref, r.Object_attributes.Status).Inc()

	metrics.PipelineDurations.WithLabelValues(r.Project.Web_url,
		r.Object_attributes.Ref, strconv.Itoa(r.Object_attributes.Id)).Add(float64(r.Object_attributes.Duration))

	metrics.PipelineTimestamps.WithLabelValues(r.Project.Web_url,
		r.Object_attributes.Ref, strconv.Itoa(r.Object_attributes.Id)).Add(float64(r.Object_attributes.Created_at.Sec()))

	for i := 0; i < len(r.Builds); i++ {
		metrics.BuildDurations.WithLabelValues(r.Project.Web_url,
			r.Object_attributes.Ref, strconv.Itoa(r.Object_attributes.Id), r.Builds[i].Stage, r.Builds[i].Status).Add(float64(r.Builds[i].Duration))
	}

	return
}
