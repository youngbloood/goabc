package goabc

func init() {
	abcQueue = newQueues(0, nil)
	iopen = &open{
		abc: abcQueue,
	}
}

var iopen *open

type open struct {
	abc Abcer
}

// Register . register the abc
func Register(abc Abcer) {
	iopen.abc = abc
}

// Add . add funcs into queue
func Add(fs ...Func) {
	for _, fv := range fs {
		iopen.abc.Add(fv)
	}
}

// Remove . remove funcs from queue
func Remove(fs ...Func) {
	for _, fv := range fs {
		iopen.abc.Remove(fv)
	}
}

// Flush . flush all funcs
func Flush() {
	iopen.abc.Flush()
}

// Start . start the funcs with concurrent
func Start() {
	iopen.abc.Start()
}

// Run . run the funcs with non-concurrent
func Run() {
	iopen.abc.Run()
}

// Random . random run the func concurrent
func Random() {
	iopen.abc.Random()
}

// SetHook . set a hook into queue
func SetHook(hook Hooker) {
	iopen.abc.SetHooker(hook)
}
