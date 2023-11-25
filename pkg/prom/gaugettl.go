package prom

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type GaugeVecTtl struct {
	gaugeVec *prometheus.GaugeVec
	ttl      int // TTL in seconds
	gauges   []GaugeWithLastUsed
}

type GaugeWithLastUsed struct {
	gaugeLabels []string
	lastUsed    time.Time
}

func NewGaugeVecTtl(opts prometheus.GaugeOpts, labelNames []string, ttl int) *GaugeVecTtl {
	var gaugeVecTtl = GaugeVecTtl{
		gaugeVec: prometheus.NewGaugeVec(opts, labelNames),
		ttl:      ttl,
		gauges:   []GaugeWithLastUsed{},
	}
	go gaugeVecTtl.garbageCollection()
	return &gaugeVecTtl
}

func (v *GaugeVecTtl) WithLabelValues(lvs ...string) prometheus.Gauge {
	var gauge = v.gaugeVec.WithLabelValues(lvs...)

	for i := 0; i < len(v.gauges); i++ {
		if reflect.DeepEqual(v.gauges[i].gaugeLabels, lvs) {
			v.gauges[i].lastUsed = time.Now()
			return gauge
		}
	}

	v.gauges = append(v.gauges, GaugeWithLastUsed{gaugeLabels: lvs, lastUsed: time.Now()})
	return gauge
}

func (v *GaugeVecTtl) garbageCollection() {
	t := time.NewTicker(10 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("Garbage collection started...")
			var t = time.Now().Add(-1 * time.Duration(v.ttl) * time.Second)
			var mainteinedGauges = []GaugeWithLastUsed{}
			for i := 0; i < len(v.gauges); i++ {
				if v.gauges[i].lastUsed.Before(t) {
					fmt.Println("Removing " + strings.Join(v.gauges[i].gaugeLabels, " "))
					v.gaugeVec.DeleteLabelValues(v.gauges[i].gaugeLabels...)
				} else {
					mainteinedGauges = append(mainteinedGauges, v.gauges[i])
				}
			}
			v.gauges = mainteinedGauges
			fmt.Println("Garbage collection finished...")
		}
	}
}
