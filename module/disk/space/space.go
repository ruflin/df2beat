package space

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/helper"
)

func init() {
	helper.Registry.AddMetricSeter("disk", "space", &MetricSeter{})
}

type MetricSeter struct {}

func (m *MetricSeter) Fetch(ms *helper.MetricSet) (events []common.MapStr, err error) {

	events = eventMapping()

	return events, nil
}
