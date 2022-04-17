package main

import (
	"sync"

	"github.com/fox-wei/go_base_fox/gonote/src/tracefunc/mytrace"
)

func A1() {
	defer mytrace.TraceFunc()()
	B1()
}

func B1() {
	defer mytrace.TraceFunc()()
	C1()
}

func C1() {
	defer mytrace.TraceFunc()()
	D()
}

func D() {
	defer mytrace.TraceFunc()()
}

func A2() {
	defer mytrace.TraceFunc()()
	B2()
}

func B2() {
	defer mytrace.TraceFunc()()
	C2()
}

func C2() {
	defer mytrace.TraceFunc()()
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
