package goabc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/youngbloood/goabc"
)

func initFunc() []goabc.Func {
	list := make([]goabc.Func, 0)
	for i := 0; i < 100; i++ {
		a := i
		list = append(list, func() {
			fmt.Println(a)
			panic(a)
		})
	}
	return list
}
func TestStart(t *testing.T) {
	funs := initFunc()
	goabc.Add(funs...)
	goabc.Start()
	time.Sleep(2 * time.Second)
}

func TestRun(t *testing.T) {
	funs := initFunc()
	goabc.Add(funs...)
	goabc.Run()
	time.Sleep(2 * time.Second)
}

func TestRandom(t *testing.T) {
	funs := initFunc()
	goabc.Add(funs...)
	goabc.Random()
	time.Sleep(2 * time.Second)
}

func TestFlushStart(t *testing.T) {
	funs := initFunc()
	goabc.Add(funs...)
	goabc.Flush()
	goabc.Start()
	time.Sleep(2 * time.Second)
}

func TestFlushRun(t *testing.T) {
	funs := initFunc()
	goabc.Add(funs...)
	goabc.Flush()
	goabc.Run()
	time.Sleep(2 * time.Second)
}

func TestFlushRandom(t *testing.T) {
	funs := initFunc()
	goabc.Add(funs...)
	goabc.Flush()
	goabc.Random()
	time.Sleep(2 * time.Second)
}

func TestRemove(t *testing.T) {
	funs := initFunc()[:2]
	t.Log(len(funs))
	goabc.Add(funs...)
	goabc.Remove(funs[0])
	goabc.Run()
	time.Sleep(2 * time.Second)
}
