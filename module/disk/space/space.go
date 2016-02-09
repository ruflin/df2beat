package space

import (
	//"fmt"

	"github.com/elastic/beats/libbeat/common"
	//"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"

	"github.com/ruflin/df2beat/module/disk"
)

func init() {
	MetricSet.Register()
}

var MetricSet = helper.NewMetricSet("space", MetricSeter{}, disk.Module)

type MetricSeter struct {
}

func (m MetricSeter) Setup() error {
	return nil
}

func (m MetricSeter) Fetch() (events []common.MapStr, err error) {

	events = eventMapping()

	return events, nil
}

func (m MetricSeter) Cleanup() error {
	return nil
}
