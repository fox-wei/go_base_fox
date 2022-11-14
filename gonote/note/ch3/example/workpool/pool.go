package workpool

import (
	"errors"
	"fmt"
	"sync"
)

//*错误信息
var (
	ErrNoIdleWorkerInPool = errors.New("no idle worker in pool") //^Workerpool中任务已满，没有空闲处理
	ErrWorkerPooLFree     = errors.New("workerPool free")        //^Workerpool 终止运行
)

type Task func() //*用户提交的请求

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

type Pool struct {
	capacity int //?workerpool 大小

	active chan struct{} //*计数器
	tasks  chan Task     //*task channel

	wg   sync.WaitGroup //&用于在pool销毁时等所有worker退出
	quit chan struct{}  //&通知各个worker跳出退出
}

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

//!启动workerpool
func (p *Pool) run() {
	idx := 0

	for {
		select {
		case <-p.quit: //*close(quit)
			return
		case p.active <- struct{}{}: //!细节失误
			idx++
			//?create new worker
			p.newWorker(idx)
		}
	}
}

//&create a worker
func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		//*异常处理，防止task pani导致整个pool受到影响
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]:recover panic[%s] and exit\n", i, err)
				<-p.active //*worker退出pool池
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

//!Schedule方法，用户提交请求
func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPooLFree //*workerpool可能已销毁
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) Free() {
	close(p.quit) //*make sure the worker and p.run eixt
	p.wg.Wait()
	fmt.Printf("workerpool freed\n")
}
