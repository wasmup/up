package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/wasmup/up"
	"github.com/wasmup/up/latency"
)

func main() {

	var ok int32
	time.AfterFunc(1000*time.Millisecond, func() {
		atomic.StoreInt32(&ok, 1) // async store
	})

	latency.LoadCPU()
	const n = 100_000_000
	deltaTime := make([]time.Duration, 0, n)
	t0 := time.Now()

	var i uint64
	for atomic.LoadInt32(&ok) == 0 { // async load: out of user control

		if len(deltaTime) < n {
			deltaTime, t0 = append(deltaTime, time.Since(t0)), time.Now()
		}

		i++ // e.g. CPU intensive task

		// Toggle this line then `go run .`:
		// runtime.Gosched()
		/*
			i = 3_632_642
			min latency = 44ns
			max latency = 2.964073ms
			ave = 70ns
			width  = 29.64µs
			3_632_623 x 44ns, 7 x 29.684µs, 2 x 59.324µs, 2 x 118.604µs, 1 x 148.244µs, 1 x 207.524µs, 1 x 237.164µs, 1 x 266.804µs, 1 x 533.564µs, 1 x 622.484µs, 1 x 711.404µs, 1 x 2.934404ms,
		*/
		runtime.Gosched() // fmt.Println(3_632_642 / 170_150) // 21x
		/*
			i = 170_150
			min latency = 398ns
			max latency = 2.501495ms
			ave = 3.909µs
			width  = 25.01µs
			169_947 x 398ns, 153 x 25.408µs, 24 x 50.418µs, 10 x 75.428µs, 3 x 100.438µs, 3 x 175.468µs, 2 x 225.488µs, 1 x 250.498µs, 1 x 400.558µs, 1 x 475.588µs, 1 x 825.728µs, 1 x 975.788µs, 1 x 1.150858ms, 1 x 1.250898ms, 1 x 2.476388ms,
		*/
	}

	fmt.Println("i =", up.Grouping(strconv.FormatUint(i, 10))) // 1_959_383_942

	ave, width, frequency, y := latency.Histogram(deltaTime, 100)
	fmt.Println("min latency =", y[0])
	fmt.Println("max latency =", y[len(y)-1]) // last
	fmt.Println("ave =", ave)
	fmt.Println("width  =", width)
	// fmt.Println("frequency =", frequency)
	// fmt.Println("deltaTime =", y)
	for i, f := range frequency {
		if f > 0 {
			fmt.Printf("%s x %v, ", up.Grouping(strconv.Itoa(f)), y[i])
		}
	}
	fmt.Println()
	fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.Version())
}
