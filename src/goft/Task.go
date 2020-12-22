package goft

import "sync"

func init() {
	c := GetTaskChan()
	go func() {
		for {
			te := <-c
			doTask(te)
		}
	}()
}
func doTask(t *TaskExector) {
	go func() {
		defer func() {
			t.callback()
		}()
		t.Exec()
	}()
}

var TaskChan chan *TaskExector
var once sync.Once

func GetTaskChan() chan *TaskExector {
	once.Do(func() {
		TaskChan = make(chan *TaskExector, 1)
	})
	return TaskChan
}

type TaskFunc func(param ...interface{})

type TaskExector struct {
	f        TaskFunc
	param    []interface{}
	callback func()
}

func NewTaskExector(f TaskFunc, param []interface{}, callback func()) *TaskExector {
	return &TaskExector{f: f, param: param, callback: callback}
}

func (this *TaskExector) Exec() {
	this.f(this.param...)
}

func Task(f TaskFunc, cf func(), param ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		GetTaskChan() <- NewTaskExector(f, param, cf)
	}()

}
