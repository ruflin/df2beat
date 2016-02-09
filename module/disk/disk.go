package disk

import (

	"github.com/elastic/beats/metricbeat/helper"
)

func init() {
	Module.Register()
}

var Module = helper.NewModule("disk", Disk{})

var Config = &ApacheModuleConfig{}

type ApacheModuleConfig struct {
	Metrics map[string]interface{}
	Hosts   []string
}

type Disk struct {
	Name   string
	Config ApacheModuleConfig
}

func (r Disk) Setup() error {

	// Loads module config
	// This is module specific config object
	//Module.LoadConfig(&Config)
	return nil
}
