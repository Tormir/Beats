package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/KM/sysusage/beater"
//	_ "github.com/KM/sysusage/include"
)

func main() {
	err := beat.Run("sysusage", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
