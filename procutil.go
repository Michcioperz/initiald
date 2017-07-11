package main

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func amIInit() bool {
	procFiles, err := ioutil.ReadDir("/proc")
	if err == nil {
		myPid := os.Getpid()
		for _, procFile := range procFiles {
			if !procFile.IsDir() {
				continue
			}
			if procPid, err := strconv.Atoi(procFile.Name()); err != nil || procPid == myPid {
				continue
			}
			commFileName := path.Join("/proc", procFile.Name(), "comm")
			commBytes, err := ioutil.ReadFile(commFileName)
			if err != nil {
				continue
			}
			comm := strings.TrimSpace(string(commBytes))
			if comm == "initiald" {
				return false
			}
		}
	}
	return true
}
