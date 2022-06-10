package latency

import (
	"runtime"
	"time"
)

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

// Histogram organizes data points into `n` user-specified ranges and returns:
// ave, frequency, and ranges for latency measurement.
//
// First make slice of `time.Duration`:
// const n = 100_000_000
// dt := make([]time.Duration, 0, n)
// t0 := time.Now()
//
// Then inside your main loop:
// if len(dt) < n {
// 	dt, t0 = append(dt, time.Since(t0)), time.Now()
// }
//
// After the loop, show the result:
//
// ave, frequency, y := latency.Histogram(dt, 10)
// fmt.Println("min =", y[0])
// fmt.Println("max =", y[1])
// fmt.Println("ave =", ave)
// // fmt.Println("frequency =", frequency)
// // fmt.Println("deltaTime =", y)
// for i, f := range frequency {
// 	if f > 0 {
// 		fmt.Printf("%s x %v, ", up.Grouping(strconv.Itoa(f)), y[i])
// 	}
// }
// fmt.Println()
// fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.Version())
func Histogram(dt []time.Duration, n int) (ave, width time.Duration, frequency []int, y []time.Duration) {
	frequency = make([]int, n)
	y = make([]time.Duration, n+1)
	min := dt[0]
	max := min
	ave = min
	for _, v := range dt[1:] {
		ave += v
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	ave /= time.Duration(len(dt))
	width = (max - min) / time.Duration(n)
	v := min
	for i := range y {
		y[i] = v
		v += width
	}
	y[len(y)-1] = max // compansate division error
	for _, v := range dt {
		i := int((v - min) / width)
		if i >= n {
			i--
		}
		frequency[i]++ // frequency count
	}
	return
}
