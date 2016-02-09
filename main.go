package main

import (
	"os"

	"github.com/elastic/beats/metricbeat/beater"
	_ "github.com/ruflin/df2beat/module/disk/space"

	"github.com/elastic/beats/libbeat/beat"
)

var Name = "df2beat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
