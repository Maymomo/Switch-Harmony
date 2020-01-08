package common

import (
	"log"
	"runtime"
)

func PrintFileAndLine() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("panic %s:%d\n", file, line)
}
