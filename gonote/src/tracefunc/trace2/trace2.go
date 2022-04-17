package main

import (
	"sync"

	"github.com/fox-wei/go_base_fox/gonote/src/tracefunc/mytrace"
)

func A1() {
	defer mytrace.TraceByGoroutine()()
	B1()
}

func B1() {
	defer mytrace.TraceByGoroutine()()
	C1()
}

func C1() {
	defer mytrace.TraceByGoroutine()()
	D()
}

func D() {
	defer mytrace.TraceByGoroutine()()
}

func A2() {
	defer mytrace.TraceByGoroutine()()
	B2()
}

func B2() {
	defer mytrace.TraceByGoroutine()()
	C2()
}

func C2() {
	defer mytrace.TraceByGoroutine()()
	D()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		A2()
		wg.Done()
	}()

	A1()
	wg.Wait()
}
