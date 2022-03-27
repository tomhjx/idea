package support

import (
	"log"
	"runtime"
)

func LogPrintTraces(depth int) {
	if depth < 1 {
		depth = 1
	}
	start := 2
	pc := make([]uintptr, start+depth)
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[start:n])
	for {
		frame, more := frames.Next()
		log.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}
