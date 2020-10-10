package goabc

import (
	"fmt"
	"log"
	"reflect"
)

var abcQueue *queues

var _ Abcer = (*queues)(nil)

type Func func()
type queue struct {
	f          Func
	nextNotify chan struct{}
	preNotify  <-chan struct{}
}

func newQueue(f Func, preNotify <-chan struct{}) *queue {
	return &queue{
		f:          f,
		nextNotify: make(chan struct{}, 0),
		preNotify:  preNotify,
	}
}

func (q *queue) exec(closeNow bool, hook Hooker) {
	if closeNow {
		close(q.nextNotify)
	} else {
		defer close(q.nextNotify)
	}

	defer q.rec(hook)
	if q.preNotify != nil {
		<-q.preNotify
	}
	q.f()
}

func (q *queue) rec(hook Hooker) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			if hook != nil {
				hook.HookErr(err)
				return
			}
			log.Println(err)
		}
	}
}

type queues struct {
	seqq []*queue
	hook Hooker
}

func newQueues(length int, hook Hooker) *queues {
	return &queues{
		seqq: make([]*queue, length),
		hook: hook,
	}
}

// the same func can be added many times. [相同地址的函数可以被多次添加到队列中]
func (qs *queues) Add(f Func) {
	if len(qs.seqq) == 0 {
		qs.seqq = append(qs.seqq, newQueue(f, nil))
		return
	}
	qs.seqq = append(qs.seqq, newQueue(f, qs.seqq[len(qs.seqq)-1].nextNotify))
}

// the same func will be clear that same in slice. [队列中的相同的地址的函数都会被移除掉]
func (qs *queues) Remove(f Func) {
LOOP:
	for i, fv := range qs.seqq {
		rfv := reflect.ValueOf(fv.f)
		rf := reflect.ValueOf(f)
		if rfv.Pointer() == rf.Pointer() {
			// 处理preNotify和nextNotify
			if i == 0 && len(qs.seqq) > 1 {
				qs.seqq[i+1].preNotify = nil
			} else if i < len(qs.seqq)-1 {
				qs.seqq[i+1].preNotify = qs.seqq[i-1].nextNotify
			}

			qs.seqq = append(qs.seqq[:i], qs.seqq[i+1:]...)
			goto LOOP
		}
	}
}

func (qs *queues) Flush() {
	qs.seqq = make([]*queue, 0)
}

func (qs *queues) SetHooker(hook Hooker) {
	qs.hook = hook
}

func (qs *queues) Start() {
	for _, fv := range qs.seqq {
		go fv.exec(false, qs.hook)
	}
}

func (qs *queues) Run() {
	for _, fv := range qs.seqq {
		fv.exec(true, qs.hook)
	}
}

func (qs *queues) Random() {
	for _, fv := range qs.seqq {
		go fv.exec(true, qs.hook)
	}
}

func (qs *queues) print() {
	for i, fv := range qs.seqq {
		fmt.Printf("fv[%d] = %v\n", i, fv.f)
	}
}
