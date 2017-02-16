// +build linux

package devoops

import (
	"fmt"
	"io/ioutil"
	"os"
)

func envForPid(pid int) []string {
	f, err := os.Open(fmt.Sprintf("/proc/%d/environ", pid))
	if err != nil { // untested, but better safe than sorry
		return []string{}
	}
	defer f.Close()

	// read file
	//	environment variables are in the Golang format (key=val)
	//	but separated by \0 :(
	//	hence, some hackish parsing needs to take place
	stream, err := ioutil.ReadAll(f)
	if err != nil { // untested, but better safe than sorry
		return []string{}
	}

	// convert it to strings
	//	...the hackish parsing I mentioned before
	lastZero := -1
	envVars := []string{}
	for i := 0; i < len(stream); i++ {
		if stream[i] == 0x0 {
			envVars = append(envVars, string(stream[lastZero+1:i]))
			lastZero = i
		}
	}

	return envVars
}
