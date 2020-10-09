package goabc

import (
	"reflect"
)

type Func func()

type queue struct {
	f      Func
	notify chan struct{}
}

type queues []*queue

var abcQueue *queues

func init() {
	aq := make(queues, 0)
	abcQueue = &aq
}

func (qs *queues) Add(f func()) {
	*qs = append(*qs, &queue{
		f: f,
	})
}

func (qu *queues) Remove(f Func) {
	for i, fv := range *qu {
		if reflect.DeepEqual(fv.f, f) {
			*qu = append((*qu)[:i], (*qu)[i+1:]...)
		}
	}
}
