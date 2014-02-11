// +build darwin freebsd linux netbsd openbsd

package main

import (
	"os"
	"path"
)

const (
	rcFname            = "rc"
	alwaysBlockedFname = "blocked"
	alwaysDirectFname  = "direct"
	statFname          = "stat"

	newLine = "\n"
)

func initConfigDir() {
	home, _ := os.Getwd()
	configPath.dir = path.Join(home, "conf")
}
