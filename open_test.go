package goabc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/youngbloood/goabc"
)

func initFunc() []goabc.Func {
	return []goabc.Func{
		func() { fmt.Println("11") },
		func() { fmt.Println("22") },
		func() { fmt.Println("33") },
		func() { fmt.Println("44") },
		func() { fmt.Println("55") },
		func() { fmt.Println("66") },
		func() { fmt.Println("77") },
		func() { fmt.Println("88") },
		func() { fmt.Println("99") },
	}
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
	funs := initFunc()
	goabc.Add(funs...)
	goabc.Remove(funs[0])
	goabc.Run()
	time.Sleep(2 * time.Second)
}
