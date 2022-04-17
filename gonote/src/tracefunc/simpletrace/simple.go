package main

import "github.com/fox-wei/go_base_fox/gonote/src/tracefunc/mytrace"

func foo() {
	defer mytrace.TraceByName("foo")()
	bar()
}

func bar() {
	defer mytrace.TraceByName("bar")
}

func main() {
	defer mytrace.TraceByName("main")()
	foo()
}
