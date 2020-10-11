package goabc

type Abcer interface {
	Add(Func)
	Remove(Func)
	Sort()
	Flush()
	Start()
	Run()
	Random()
	SetHooker(hook Hooker)
}
