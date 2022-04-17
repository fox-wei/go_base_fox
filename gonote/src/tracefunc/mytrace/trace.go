package mytrace

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

//*simple trace by input the function name
func TraceByName(name string) func() {
	fmt.Println("enter:", name)
	return func() {
		fmt.Println("exit:", name)
	}
}

//&trace the function and auto get the name
func TraceByAuto() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("Not found caller!")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	fmt.Println("enter:", name)
	return func() {
		fmt.Println("exit:", name)
	}
}

//?tace the func in different goroutine
var goroutineSpace = []byte("goroutine ") //!小细节，空格

//?get the goroutine id
func ourGoroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// fmt.Println("information:", string(b))
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}

	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}

	return n
}

//?Trace function  by different gorouine
func TraceByGoroutine() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gid := ourGoroutineId()

	fmt.Printf("g[%05d]:enter:[%s]\n", gid, name)
	return func() {
		fmt.Printf("g[%05d]:exit:[%s]\n", gid, name)
	}
}

//!Trace the function by different gorotine and the trace's chain is good
func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "	"
	}

	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}

var mu sync.Mutex
var m = make(map[uint64]int)

//!Trace the call chain and output is good
func TraceFunc() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gid := ourGoroutineId()

	mu.Lock()            //^map不支持并发运行
	indents := m[gid]    //*获取当前gid的缩进层次
	m[gid] = indents + 1 //*缩进层次加1后存入map
	mu.Unlock()

	printTrace(gid, name, "->", indents)
	return func() {
		mu.Lock()
		indents := m[gid]
		m[gid] = indents - 1
		mu.Unlock()
		printTrace(gid, name, "<-", indents)
	}
}
