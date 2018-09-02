package cmd

import (
	"github.com/KM/sysusage/beater"

	cmd "github.com/elastic/beats/libbeat/cmd"
)

// Name of this beat
var Name = "sysusage"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
