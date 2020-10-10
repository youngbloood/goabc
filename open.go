package goabc

// Add . add funcs into queue
func Add(fs ...Func) {
	for _, fv := range fs {
		abcQueue.add(fv)
	}
}

// Remove . remove funcs from queue
func Remove(fs ...Func) {
	for _, fv := range fs {
		abcQueue.remove(fv)
	}
}

// Flush . flush all funcs
func Flush() {
	abcQueue.flush()
}

// Start . start the funcs with concurrent
func Start() {
	abcQueue.start()
}

// Run . run the funcs with non-concurrent
func Run() {
	abcQueue.run()
}

// Random . random run the func concurrent
func Random() {
	abcQueue.random()
}

// SetHook . set a hook into queue
func SetHook(hook Hooker) {
	abcQueue.setHooker(hook)
}
