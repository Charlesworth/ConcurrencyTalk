package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	var ops uint64 = 0

	for i := 0; i < 3; i++ {
		fmt.Println("goRoutine", i, "started")
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

func setProc() {
	procNo := runtime.NumCPU()
	runtime.GOMAXPROCS(procNo)
	fmt.Println("Using", procNo, "processors for maximum thread count")
}
