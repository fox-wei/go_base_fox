# Go并发(Concurrency)

## 并发：Go并发实现方案

操作系统的基本调度和执行单元是进程(process)。

**并行(parallelism)，指的在同一时刻，有两个或以上的任务（进程)的代码在处理器上执行。**

多进程应用设计，将内部划分为多个模块，每个模块用一个进程承载执行，每个模块都是一个单独的执行流。

多进程应用将功能职责做了划分，可读性与维护性更好，这种将程序分成多个可独立执行的部分的结构化程序方法，就是并发设计。例如进程内部划分为多个线程。

**Go并发方案：goroutine**

Go没有使用操作系统线程作为承载分解后的代码，而由goroutine实现。**Go运行时(runtime)负责调度、轻量的用户级线程。**

Goroutine与传统操作系统的优势：

- 资源占用小，每个goroutine	初始栈大小为2KB；
- runtine调用而是操作系统，goroutine上下文切换在用户层完成，开销更小；
- 语言层面实现而不出标准库，关键字`go` + 函数(方法)创建goroutine，退出就被回收；
- channel作为goroutine间的通信原语。

**多数情况下，不用考虑goroutine的退出；一般goroutine的执行函数返回，意味着goroutine退出。**如果main goroutine退出，整个程序退出。

**Goroutine间的通信**

传统并发模型是基于对内存的共享，而Go采用CSP(communicationing sequential Process，通信顺序模型)并发模型。输入输出是基本的编程原语，数据处理逻辑(CSP中的P)只需调用输入原语获取数据，顺序地处理数据，并将结果数据通过输出原语输出。

**一个符合CSP模型的并发程序应该是一组通过输入输出原语连接起来的P集合**。

CSP中P(进程)，是一个抽象概念，它代表任何顺序处理逻辑的封装。它获取输入数据(或从其他P获取)，并产生可以被其他P消费的输出数据。模型示意图：
![alt](https://static001.geekbang.org/resource/image/e7/8c/e7c4fcc00ece399601de800e3a7f598c.jpg?wh=1920x465)


## Goroutine调度器(Goroutine Scheduler)

Go程序对于操作系统来说只是一个用户层程序，操作系统眼中只有线程，甚至不知道gorotine的存在，gorotine的调度由Go自己实现。
Gorotine之间“公平”竞争"CPU"资源的任务，落到runtime头上。

操作系统层面，线程竞争的CPU资源是真实的物理CPU，go程序整体运行在一个或多个操作系统上，所以gorotine竞争的“CPU”资源是操作系统的线程。调度器的任务时按照一定算法将gorotine放到不同操作系统的线程中执行。

**调度器模型**

最开是的模型是G-M模型，G是goroutine运行时对应一个抽象结构；物理CPU的操作系统线程对应结构M(machine)。

随着不断演化，目前是G-M-P模型，加入中间层P。有人说过：“计算机科学领域的任何问题都可以通过增加一个间接的中间层来解决”

![alt](https://static001.geekbang.org/resource/image/43/a8/43ffdbc6b2203d9400ac98423192caa8.png?wh=1224x1142)


P是一个“逻辑processor"，每个G要想运行必须被分配到一个P中。对于G，P就是它的CPU，可以说G眼里只有P。但从调度器角度看，M才是真正的CPU，P和M绑定，才能让G运行起来。

开始不支持抢占式调度，导致某个G实现死循环的代码逻辑，G将永久占用P和M使其他G出现饿死的情况。Go1.12后实现了基于协作的“抢占式”调度。

这个抢占式调度在**函数或方法入口处**加入检测代码(runtime.morestack_noctxt)，让runtime有机会检查是否需要执行抢占式调度。对于纯算法循环计算的G，Go调度器无法依然抢占。

Go1.14加入对于非协作的抢占式调度支持，这种基于系统信号向线程发送的方式来抢占。

**深入GPM模型**

- G：代表Goroutine，存储Goroutine的执行栈信息，状态，和任务函数等；G对象可以重用；
- P；逻辑processor，P的数量决定系统最大可并行的G的数量，P拥有各种G对象队列、链表、一些缓存等。
- M：代表真正的执行资源，绑定有效P后，进入一个循环调度：P从本地队列或全局队列获取G，切换到G执行栈并执行G函数，调用goexit做清理工作并回到M。

**G被抢占式调度**

如果某个G，没有进行系统调度，没有进行I/O操作，没有阻塞在一个channel操作，则这个G将被抢占式调度。

Go程序启动时，runtime启动一个名为sysmon的M进行监控(监控线程)。

如果G被阻塞在某个channel操作或网络I/O操作上时，M可以不被阻塞，这避免了大量创建 M 导致的开销。但如果G因慢系统调用而阻塞，那么M也会一起阻塞，但在阻塞前会与P解绑，P 会尝试与其他 M 绑定继续运行其他 G。但若没有现成的 M，Go 运行时会建立新的 M，这也是系统调用可能导致系统线程数量增加的原因。

## 并发：channel

**作为一等公民的channel**

定义channel类型变量并赋值，作为函数\方法参数,作为返回值，将channel发送到其他channel，这大大简化channel的使用。

创建channel：
```go
var ch chan int //*默认值nil
//为channel赋值
ch1 := make(chan int) //无缓冲的chan
ch2 := make(chan int, 4) //带缓冲的chan
```

**发送与接收**
使用`<-`进行发送：
```go

ch1 <- 13    // 将整型字面值13发送到无缓冲channel类型变量ch1中
n := <- ch1  // 从无缓冲channel类型变量ch1中接收一个整型值存储到整型变量n中
ch2 <- 17    // 将整型字面值17发送到带缓冲channel类型变量ch2中
m := <- ch2  // 从带缓冲channel类型变量ch2中接收一个整型值存储到整型变量m中
```

channel是用于在goroutine间的通信，绝大多数的channel的读写分布在不同的goroutine上中。

对于无缓冲channel，发送与接收是同步的。也就是说只有对它进行接收操作的goroutine和进行发送的goroutine同时存在的情况下才能进行通信，否则单方面操作会让goroutine挂起。例如：
```go
package main

func main() {
	ch1 := make(chan int)
	ch1 <- 23 //fatal error: deadlock
	n := <-ch1
	println(n)
}
```
**一个无缓冲的channel,读写操作都在同一个goroutine中,导致所有goroutine进入休眠，程序处于死锁状态**。所以对无缓冲的channel，读写操作要在不同的goroutine中，修改如下：
```go
func main() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 22 //写操作在一个新的goroutine
	}

	n := <- ch1
	println(n)
}
```

带缓冲的channel的运行时有缓冲区，对其进行发送在缓冲区未满，进行接收操作类缓冲区非空的情况下是异步的（发送或接收操作不需要阻塞等待）。

- 缓冲区未满： 发送不会阻塞挂起，否则阻塞挂起
- 缓冲区非空：接收不会阻塞挂起，否则阻塞挂起

操作符`<-`还可以声明只发送类型(send-only)和只接受类型(receive-only)的channel。
```go
ch1 := make(chan<- int, 1) // 只发送channel类型
ch2 := make(<-chan int, 1) // 只接收channel类型

<-ch1       // invalid operation: <-ch1 (receive from send-only type chan<- int)
ch2 <- 13   // invalid operation: ch2 <- 13 (send to receive-only type <-chan int)
```
通常这两种类型作为函数的参数或返回值，用来限制对channel的操作。
```go
func produce(ch chan<- int) {
    for i := 0; i < 10; i++ {
        ch <- i + 1
        time.Sleep(time.Second)
    }
    close(ch)
}

func consume(ch <-chan int) {
    for n := range ch {
        println(n)
    }
}

func main() {
    ch := make(chan int, 5)
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        produce(ch)
        wg.Done()
    }()

    go func() {
        consume(ch)
        wg.Done()
    }()

    wg.Wait()
}
```
这是一个简单的消费者-生产者模型，生产者向channel进行发送(chan <-int)，消费者从channel接收操作(<- chan int)。for range从channel中接收数据，for range会阻塞在对channel的接受操作，直到有数据或channel，才能继续运行。当channel关闭循环结束。

**关闭channel**

调用`close`函数关闭channel，所有等待的从这个channel接收的数据的操作都将返回。
channel关闭后返回值的情况：
```go
n := <- ch      // 当ch被关闭后，n将被赋值为ch元素类型的零值
m, ok := <-ch   // 当ch被关闭后，m将被赋值为ch元素类型的零值, ok值为false
for v := range ch { // 当ch被关闭后，for range循环结束
    ... ...
}
```
通过“comma,ok"或for-range方式可以判断channel是否被关闭。从生产者-消费者模型可知，**发送端负责关闭channel**。如果向已关闭channel发送数据则会产生panic。

**select**

当涉及同时多个channel操作时，可以使用select：
```go
select {
case x := <-ch1:     // 从channel ch1接收数据
  ... ...

case y, ok := <-ch2: // 从channel ch2接收数据，并根据ok值判断ch2是否已经关闭
  ... ...

case ch3 <- z:       // 将z值发送到channel ch3中:
  ... ...

default:             // 当上面case中的channel通信均无法实施时，执行该默认分支
}
```
如果没有default，case中的channel操作都阻塞的时候，整个select都会阻塞，直到某个channel可以进行发送或接收。

**无缓冲channel的使用**

**第一种用法：用作信号传递**

**1:n信号通知机制：**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

//& 1:n信号通知机制
type signal struct{}

func worker(i int) {
	fmt.Printf("Worker %d: is working...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d: is done!\n", i)
}

func spawnGroup(f func(int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-groupSignal //~所有goroutine阻塞，直到channel关闭才能继续执行
			fmt.Printf("Worker %d:start to working...\n", i)
			f(i)
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()

	return c
}

func main() {
	fmt.Printf("Start a group of workers......\n\n")
	groupSignal := make(chan signal)

	c := spawnGroup(worker, 5, groupSignal) //^5个goroutine发生阻塞
	time.Sleep(time.Second * 5)

	fmt.Printf("The group of workers start to work...\n\n")
	close(groupSignal) //*发送信号，关闭channel让阻塞的goroutine开始工作，实现广播
	<-c                //?等待子goroutine退出
	fmt.Println()
	fmt.Printf("The group of workers work done!\n")
}
```
创建5个goroutine，这个5个会阻塞在<-groupSignal(groupSignal无缓冲channel，未发送数据)，直到关闭groupSignal(发送信号)才能继续运行。

**第二种用法：替代锁机制**
```go

type counter struct {
    c chan int
    i int
}

func NewCounter() *counter {
    cter := &counter{
        c: make(chan int),
    }
    go func() {
        for {
            cter.i++
            cter.c <- cter.i
        }
    }()
    return cter
}

func (cter *counter) Increase() int {
    return <-cter.c
}

func main() {
    cter := NewCounter()
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            v := cter.Increase()
            fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```
计数器交给一个独立的goroutine处理，并利用无缓冲channel的同步阻塞特性，实现对计数器的控制。这样的计数器值自增的动作变为从channel的接受动作。

**带缓冲channel的惯用法**

**第一种用法：用作消息队列**

带缓冲channel作为消息队列，性能比无缓冲队列在数据收发性能明显要好于无缓冲channel。

- 无论是1收1发还是多收多发，带缓冲 channel 的收发性能都要好于无缓冲 channel；
- 对于带缓冲 channel 而言，发送与接收的 Goroutine 数量越多，收发性能会有所下降；
- 对于带缓冲 channel 而言，选择适当容量会在一定程度上提升收发性能。

**第二种用法：用作计数信号量(counting semaphore)**

channel当前数据个数代表的是，当前同时处于活动状态的goroutine的数量，容量则代表允许同时处于活动状态的goroutine的最大数量。发送代表获取一个信号量，接收则代表释放一个信号量。

```go
var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i + 1
		}
		close(jobs)
	}()

	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			active <- struct{}{}
			log.Printf("handle job:%d\n", j)
			time.Sleep(time.Second)
			<-active
		}(j)
	}

	wg.Wait()
}
```
同一个时间最多允许3个goroutine处于活动状态。

**len(channel)的应用**

`len`对两种channel的定义：

- 无缓冲channel，`len`的返回值是0;
- 带缓冲channel，返回未读取的元素个数。

如果直接使用`len`进行判空或判满的操作，因为存在goroutine间的竞争，数据可能被其他gorotine读取，从而导致当前goroutine阻塞无法完成后续逻辑。为了不出现将“判空与读取”放在同一事务，“判满与写入”放在同一事务。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

//*对channel判空和读取同时完成
//*使用select

func trySend(ch chan<- int, i int) bool {
	select {
	case ch <- i:
		return true
	default:
		return false
	}
}

func tryRecv(ch <-chan int) (int, bool) {
	select {
	case i := <-ch:
		return i, true
	default:
		return 0, false
	}
}

//^生产者
func producer(ch chan<- int) {
	var i int = 0
	for {
		if i > 10 {
			return
		}

		time.Sleep(time.Second * 2)

		ok := trySend(ch, i)
		if ok {
			fmt.Printf("[producer] send [%d] to channel\n", i)
			i++
			continue
		}

		fmt.Printf("[producer]: try send [%d], but channel is full\n", i)
	}
}

//^消费者
func consumer(ch <-chan int) {
	for {
		i, ok := tryRecv(ch)
		if !ok {
			fmt.Println("[consumer]: try to recv from channel, but the channel is empty")
			time.Sleep(time.Second)
			continue
		}

		fmt.Printf("[consumer]: recv [%d] from channel\n", i)
		if i > 6 {
			fmt.Println("[consumer]: exit......")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)

	wg.Add(2)
	go func() {
		defer wg.Done()
		producer(ch)
	}()

	go func() {
		defer wg.Done()
		consumer(ch)
	}()

	wg.Wait()

}
```
从上面的例子中，当channel满时，trySend不会阻塞；空时tryRece不会阻塞。所以生产者和消费者不会发生阻塞。

这种方式使用大多数场合，但是会改变channel的状态，让channel接收或发送一个数据。目前没有仅测试channel状态而不让其状态改变的方法，但是一些特殊场景可以使用`len`判断channel：
![alt](https://static001.geekbang.org/resource/image/b3/37/b31d081fcced758b8f99c938a0b75237.jpg?wh=1920x1047)

- 多发送单接受者，len(chan) > 0;
- 多接收单发送，len(chan) < cap(chan)；

**nil channel的妙用**

对`nil channel`进行读写会发生阻塞，且从已经关闭channel读取数据是不会发生阻塞，返回是chan类型的零值(chan int->0)。

```go

func main() {
    ch1, ch2 := make(chan int), make(chan int)
    go func() {
        time.Sleep(time.Second * 5)
        ch1 <- 5
        close(ch1)
    }()

    go func() {
        time.Sleep(time.Second * 7)
        ch2 <- 7
        close(ch2)
    }()

    for {
        select {
        case x, ok := <-ch1: //判断chan是否关闭，防止读取到零值
            if !ok {
                ch1 = nil
            } else {
                fmt.Println(x)
            }
        case x, ok := <-ch2:
            if !ok {
                ch2 = nil
            } else {
                fmt.Println(x)
            }
        }
        if ch1 == nil && ch2 == nil {
            break
        }
    }
    fmt.Println("program end")
}
```

判断两个chan是否关闭，显示设置为nil；这样就可发生阻塞，被设置为nil的select分支不会被执行。

**与select使用的惯用法**

1. 利用`default`分支避免阻塞；例如上面的trySend和tryRece函数的。
2. 实现超时机制，
```go

func worker() {
  select {
  case <-c:
       // ... do some stuff
  case <-time.After(30 *time.Second):
      return
  }
}
```

3. 实现心跳机制，执行周期性任务
```go

func worker() {
  heartbeat := time.NewTicker(30 * time.Second)
  defer heartbeat.Stop()
  for {
    select {
    case <-c:
      // ... do some stuff
    case <- heartbeat.C:
      //... do heartbeat stuff
    }
  }
}
```

## 并发：如何共享变量

在并发环境下，实现变量共享，除了使用channel实现，还可以使用sync包提房的一系列方法。

sync包低级同步原语的应用场景：

1. 需要高性能的临界区(critical section)同步机制场景；
```go
import "sync"

var mutex sync.Mutex //零值可用，无需赋值

//临界区上锁
mutex.Lock()
//critical section
mutex.Unlock()
```

2. 不想转移结构体对象所有权，但要保证结构体内部状态数据的同步访问的场景；使用channel的特点：在goroutine间通过channel转移对象的所有权。

**sync使用注意事项**

sync包源代码有一个提示：
```go
// Values containing the types defined in this package should not be copied.

//Mutex结构
type Mutex struct {
	 state int32 //互斥锁状态
	 sema uint32 //控制所信号量状态
}
```
意思是：**不应复制那些包含了此包中类型的值**。

Mutext是值类型，一旦发生赋值，实际上是两个不同的对象，两个没有关联，自然就不能实现同步。

```go
func main() {
	i := 0
	var wg sync.WaitGroup
	var mut sync.Mutex

	wg.Add(1)
	go func(mu sync.Mutex) { //值传递会提示警告；应使用指针
		defer wg.Done()
		defer mu.Unlock()
		mu.Lock()
		i = 22
		time.Sleep(time.second * 5)
		fmt.Printf("g1:%d\n", i)
	}(mut)

	time.Sleep(time.Second) //确保先执行g1

	mut.Lock()
	i = 24
	fmt.Printf("g2:%d\n", i)
	mut.Unlock()

	wg.Wait()
}

//可能结果
//g1:24
//g2:24
```
结果没有实现同步，因为两个Mutex不是同一个对象，赋值给匿名函数的时候两个对象没有关联。甚至，如果拷贝时机不对，比如一个mutex处于Locked状态时对它进行拷贝，就会对副本进行加锁，导致gorotine一直阻塞。

在使用 sync 包中的类型的时候，我们推荐通过**闭包**方式，或者是**传递类型实例（或包裹该类型的类型实例）的地址（指针）**的方式进行。

**互斥锁与读写锁**

互斥锁的使用原则：
1. 尽量减少锁中的操作
2. 一定要调用Unlock解锁，可以使用`defer`，将加锁与解锁放在同一位置。

读写锁与互斥锁使用方法大致相同：
```go
var rwmu sync.RWMutex

rwmu.RLock()   //加读锁
readSomething()
rwmu.RUnlock() //解读锁

rwmu.Lock()    //加写锁
changeSomething()
rwmu.Unlock()  //解写锁
```
写锁与互斥锁类型，某个goroutine持有写锁，其他goroutine无论尝试加读锁或写锁都会被阻塞；读锁较为宽松，不会阻塞其他goroutine加读锁，但会阻塞写锁。

互斥锁是临界区同步首选，读写锁适合应用在具有一定并发量且读多写少的场合。 

在大量并发读的情况下，多个Goroutine可以持有读锁，从而减少在锁竞争中等待的时间。互斥锁，即便在读请求场合，同一个时刻只有一个Goroutine持有锁。其他只能阻塞等掉被调度。



## 实践：实现一个轻量级线程池

**线程池的三个部分：**

- Pool的创建与销毁；
- Pool中Worker(goroutine)管理；
- task的提交与调度

workpool后两部分实现原理：
![alt](https://static001.geekbang.org/resource/image/d4/fd/d48ba3a204ca6e8961a4425573afa0fd.jpg?wh=1920x1047)

用户要提交给workerpool的请求抽象为Task，Task的提交与调度通过Schedule函数，提交到task channel中，已经创建的worker将从task channel中读取。

Pool 的结构：
```go
type Pool struct {
	capacity int //?workerpool 大小

	active chan struct{} //*计数器
	tasks  chan Task //*task channel

	wg   sync.WaitGroup //&用于在pool销毁时等所有worker退出
	quit chan struct{} //&通知各个worker跳出退出
}
```
- capacity：pool的容量
- active：一种带缓冲的channel，作为计数器；
- - active channel可写入，创建一个worker，
- - active channel满了，停止创建直到某个worker退出(worker == goroutine)
- task的提交与调度

workerpool对外的主要三个API:

- workerpool.New：创建pool类型实例，将pool池的worker管理机制运行起来；
- workerpool.Free：销毁一个pool池，停掉所有的worker；
- Pool.Schedule：这是Pool类型的导出方法，用户通过此方法向pool池提交待执行的任务(Task)。


**workerpool.New的创建**

```go
//!Pool类型创建
func New(capacity int) *Pool {
	//~参数错误处理
	if capacity <= 0 {
		capacity = defaultCapacity
	}

	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	//?Pool创建
	p := &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		quit:     make(chan struct{}),
		active:   make(chan struct{}, capacity),
	}

    fmt.Printf("workerpool start\n")

    go p.run() //?p初始化完成开始运行workerpool

	return p
}
```

指定接受workerpool池的容量，Pool类型实例p完成初始化后，创建一个goroutine,用户对workerpool管理，它执行Pool类型的run方法。

**run方法实现：**

```go
//!启动workerpool
func (p *Pool) run() {
	idx := 0

	for  {
		select {
		case <-p.quit:
			return
         case p.active <- struct{}{}: //当active可写入，创建worker
			idx++
			//?create new worker
            p.newWorker(idx)
		}
	}
}
```

run方法内无限循环，用来监听Pool类型的两个属性：quit和active。
当受到quit信号后这个goroutine结束运行；当active可写入时，run方法会创建一个worker goroutine。为了区分不同goroutine的输出日志，使用变量idx作为worker的编号，以参数的形式传到到newWorker方法中。

**newWorker方法实现：**
```go
//&create a worker
func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		//*异常处理
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]:recover panic[%s] and exit\n", i, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: start\n", i)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]: exit\n", i)
				<-p.active //*worker退出channel
				return
			case t := <-p.tasks:
				//*执行task
				fmt.Printf("worker[%03d]: receive a task\n", i)
				t()
			}
		}
	}()
}
```

newWorker方法创建新的goroutine作为worker，新worker核心与run方法类似，不断监听quit和tasks两个channel。task channel放置的是用户通过Schedule方法提交的请求。

Task是用户提交的一个抽象，本质上是一个函数。
```go
type Task func()
```

**Schedule方法实现：**

```go
//!Schedule方法，用户提交请求
func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPooLFree //?workerpool可能已销毁
	case p.tasks <- t:
		return nil
	}
}
```

Schdule方法的核心是将Task实例发送到workerpool的task channel。

