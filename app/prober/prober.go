package prober

import (
	"fmt"
	"runtime"
)

func Probe() string {
	return fmt.Sprintf("AARCH: %s, OS: %s\n", runtime.GOARCH, runtime.GOOS)
}
