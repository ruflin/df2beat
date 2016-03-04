package disk

import (
	"github.com/elastic/beats/metricbeat/helper"
)

func init() {
	helper.Registry.AddModuler("disk", Moduler{})
}

type Moduler struct {}

func (r Moduler) Setup() error {
	return nil
}
