package main

import "github.com/fox-wei/go_base_fox/gonote/src/tracefunc/mytrace"

func foo() {
	defer mytrace.TraceByAuto()()
	bar()
}

func bar() {
	defer mytrace.TraceByAuto()()
}

func main() {
	defer mytrace.TraceByAuto()()
	foo()
}
