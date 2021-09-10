package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	sync.Map
	f func() error
}

func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}

	return &t
}

func (t *Task) Execute() {
	_ = t.f()
}

type Pool struct {
	//对外接收Task的入口
	EntryChannel chan *Task

	//协程池最大worker数量,限定Goroutine的个数
	workerNum int

	//协程池内部的任务就绪队列
	JobsChannel chan *Task
}

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		workerNum:   cap,
		JobsChannel:  make(chan *Task),
	}

	return &p
}

func (p *Pool) worker(workID int) {
	//worker不断的从JobsChannel内部任务队列中拿任务
	for task := range p.JobsChannel {
		//如果拿到任务,则执行task任务
		task.Execute()
		fmt.Println("worker ID ", workID, " 执行完毕任务")
	}
}

func (p *Pool) Run() {
	//1,首先根据协程池的worker数量限定,开启固定数量的Worker,
	//  每一个Worker用一个Goroutine承载
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}

	//2, 从EntryChannel协程池入口取外界传递过来的任务
	//   并且将任务送进JobsChannel中
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}

	//3, 执行完毕需要关闭JobsChannel
	close(p.JobsChannel)

	//4, 执行完毕需要关闭EntryChannel
	close(p.EntryChannel)
}

//主函数
func main() {
	//创建一个Task
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	p := NewPool(3)

	//开一个协程 不断的向 Pool 输送打印一条时间的task任务
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	//启动协程池p
	p.Run()

}
