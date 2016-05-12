package space

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/topbeat/system"
	"github.com/elastic/beats/metricbeat/mb"
)

func init() {
	if err := mb.Registry.AddMetricSet("disk", "space", New); err != nil {
		panic(err)
	}
}

type MetricSet struct{
	mb.BaseMetricSet
}

// New creates and returns a new instance of MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	return &MetricSet{
		BaseMetricSet: base,
	}, nil
}

func (m *MetricSet) Fetch() ([]common.MapStr, error) {

	fss, err := system.GetFileSystemList()
	if err != nil {
		logp.Warn("Getting filesystem list: %v", err)
		return nil, err
	}

	filesSystems := make([]common.MapStr, 0, len(fss))
	for _, fs := range fss {
		fsStat, err := system.GetFileSystemStat(fs)
		if err != nil {
			logp.Debug("space", "error getting filesystem stats for '%s': %v", fs.DirName, err)
			continue
		}
		system.AddFileSystemUsedPercentage(fsStat)
		filesSystems = append(filesSystems, system.GetFilesystemEvent(fsStat))
	}

	return filesSystems, nil
}
