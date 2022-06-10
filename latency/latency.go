package latency

import "runtime"

// LoadCPU makes all CPU cores bussy for latency testing
func LoadCPU() {
	for i := runtime.NumCPU(); i > 0; i-- {
		go useOneCPUcore()
	}
}

func useOneCPUcore() {
	for {
		runtime.Gosched()
	}
}
