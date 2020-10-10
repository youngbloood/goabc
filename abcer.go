package goabc

type Abcer interface {
	Add(Func)
	Remove(Func)
	Flush()
	Start()
	Run()
	Random()
	SetHooker(hook Hooker)
}
