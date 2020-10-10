package goabc

import (
	"fmt"
	"log"
	"reflect"
)

var abcQueue *queues

func init() {
	abcQueue = newQueues(0, nil)
}

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

func (qs *queues) add(f func()) {
	if len(qs.seqq) == 0 {
		qs.seqq = append(qs.seqq, newQueue(f, nil))
		return
	}
	qs.seqq = append(qs.seqq, newQueue(f, qs.seqq[len(qs.seqq)-1].nextNotify))
}

func (qs *queues) remove(f Func) {
	for i, fv := range qs.seqq {

		rfv := reflect.ValueOf(fv.f)
		rf := reflect.ValueOf(f)

		fmt.Println("11 = ", fmt.Sprintf("%v", rfv.String()))
		fmt.Println("22 = ", fmt.Sprintf("%v", rf.Pointer()))

		if reflect.DeepEqual(fv.f, f) {
			qs.seqq = append(qs.seqq[:i], qs.seqq[i+1:]...)
			break
		}
	}
}

func (qs *queues) flush() {
	qs.seqq = make([]*queue, 0)
}

func (qs *queues) setHooker(hook Hooker) {
	qs.hook = hook
}

func (qs *queues) start() {
	for _, fv := range qs.seqq {
		go fv.exec(false, qs.hook)
	}
}

func (qs *queues) run() {
	for _, fv := range qs.seqq {
		fv.exec(true, qs.hook)
	}
}

func (qs *queues) random() {
	for _, fv := range qs.seqq {
		go fv.exec(true, qs.hook)
	}
}
