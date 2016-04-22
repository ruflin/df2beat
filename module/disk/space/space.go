package space

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/topbeat/system"
)

func init() {
	helper.Registry.AddMetricSeter("disk", "space", New)
}

// New creates new instance of MetricSeter
func New() helper.MetricSeter {
	return &MetricSeter{}
}

type MetricSeter struct{}

func (m *MetricSeter) Setup(ms *helper.MetricSet) error {
	return nil
}

func (m *MetricSeter) Fetch(ms *helper.MetricSet, host string) (event common.MapStr, err error) {

	fss, err := system.GetFileSystemList()
	if err != nil {
		logp.Warn("Getting filesystem list: %v", err)
		return nil, err
	}

	event = common.MapStr{}

	for _, fs := range fss {
		fsStat, err := system.GetFileSystemStat(fs)
		if err != nil {
			logp.Debug("topbeat", "Skip filesystem %d: %v", fsStat, err)
			continue
		}

		fsEvent := common.MapStr{
			fsStat.DevName: common.MapStr{
				"device_name": fsStat.DevName,
				"mount_point": fsStat.Mount,
				"total":       fsStat.Total,
				"used":        fsStat.Used,
				"free":        fsStat.Free,
				"avail":       fsStat.Avail,
				"files":       fsStat.Files,
				"free_files":  fsStat.FreeFiles,
			},
		}

		event.Update(fsEvent)
	}

	return event, nil
}
