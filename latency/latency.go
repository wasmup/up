package latency

import "runtime"

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
