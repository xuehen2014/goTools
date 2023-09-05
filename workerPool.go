package goTools

import (
	"sync"
)

type TaskWorker interface {
	InitData()
	ConsumeTask() (error, TaskResult)
}

type TaskResult interface {
	ToString() string
}

type WorkerPool struct {
	consumeCh chan TaskWorker
	ResultCh  chan TaskResult
	wg        sync.WaitGroup
}

func NewWorkerPool() *WorkerPool {
	p := &WorkerPool{
		consumeCh: make(chan TaskWorker),
		ResultCh:  make(chan TaskResult),
	}
	return p
}

func (p *WorkerPool) Consumer(maxGoroutines int) {
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			defer p.wg.Done()
			for task := range p.consumeCh {
				task.InitData()
				err, res := task.ConsumeTask()
				if err != nil {
					p.ResultCh <- res
				}
			}
		}()
	}
	p.CloseResultCh()
}

func (p *WorkerPool) Producer(w TaskWorker) {
	p.consumeCh <- w
}

func (p *WorkerPool) CloseConsumeCh() {
	close(p.consumeCh)
}

func (p *WorkerPool) CloseResultCh() {
	go func() {
		p.wg.Wait()
		close(p.ResultCh)
	}()
}

// 使用示例

func Example() {
	consumePool := NewWorkerPool()
	consumePool.Consumer(6)
	//发送数据
	go func() {
		for i := 1; i <= 1000; i++ {
			worker := NewReportWorker(i)
			consumePool.Producer(worker)
		}
		consumePool.CloseConsumeCh()
	}()
	//读取结果数据
	for res := range consumePool.ResultCh {
		if errInfoList, ok := res.(*ReportErrResult); ok {
			for _ = range errInfoList.ErrorList {
				// TODO
			}
		}
	}
}

//接口自检
var _ TaskWorker = &ReportWorker{}

type ReportWorker struct {
	taskId int
}

func NewReportWorker(id int) *ReportWorker {
	return &ReportWorker{
		taskId: id,
	}
}

func (r *ReportWorker) InitData() {
	// TODO
}

func (r *ReportWorker) ConsumeTask() (error, TaskResult) {
	taskResult := &ReportErrResult{}
	return nil, taskResult
}

var _ TaskResult = &ReportErrResult{}

type ReportErrResult struct {
	ErrorList []error
}

func (r *ReportErrResult) ToString() string {
	var errString string
	for _, rowVal := range r.ErrorList {
		errString = errString + rowVal.Error() + ";"
	}
	return errString
}

//追加错误消费
func (r *ReportErrResult) AddError(errInfo error) {
	r.ErrorList = append(r.ErrorList, errInfo)
}
