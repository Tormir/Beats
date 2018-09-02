package main

import (
	"os"

	"github.com/KM/sysusage/cmd"

	_ "github.com/KM/sysusage/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
