package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/passarinho4/gitlab-simple-exporter/pkg/gitlab"
	"github.com/passarinho4/gitlab-simple-exporter/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var reg = prometheus.NewRegistry()
var metrics = prom.NewMetrics(reg)

func main() {

	go garbageCollection(metrics)

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

	metrics.PipelineCounter.WithLabelValues(r.Project.Namespace, r.Project.Name, r.Object_attributes.Ref).Inc()

	if r.Object_attributes.Status == "success" {
		metrics.SuccessPipelinesCounter.WithLabelValues(r.Project.Namespace, r.Project.Name, r.Object_attributes.Ref).Inc()
	} else {
		metrics.FailedPipelinesCounter.WithLabelValues(r.Project.Namespace, r.Project.Name, r.Object_attributes.Ref).Inc()
	}

	metrics.PipelineDurations.WithLabelValues(r.Project.Namespace, r.Project.Name,
		r.Object_attributes.Ref, strconv.Itoa(r.Object_attributes.Id)).Add(float64(r.Object_attributes.Duration))

	for i := 0; i < len(r.Builds); i++ {
		metrics.BuildDurations.WithLabelValues(r.Project.Namespace, r.Project.Name,
			r.Object_attributes.Ref, strconv.Itoa(r.Object_attributes.Id), r.Builds[i].Stage).Add(float64(r.Builds[i].Duration))
	}

	return
}

func garbageCollection(metrics *prom.Metrics) {
	t := time.NewTicker(3600 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("Reseting Gauge!")
			metrics.PipelineDurations.Reset()
		}
	}
}
