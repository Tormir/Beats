package main

import (
	"os"

	"github.com/KM/sysusage/cmd"
//	"github.com/elastic/beats/libbeat/beat"
	"github.com/KM/sysusage/beater"
	_ "github.com/KM/sysusage/include"
)

var RootCmd = cmd.GenRootCmd("sysusage", "", beater.New)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
