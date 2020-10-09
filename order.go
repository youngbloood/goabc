package goabc

// Add . add funcs into queue
func Add(fs ...Func) {
	for _, fv := range fs {
		abcQueue.Add(fv)
	}
}

// Remove . remove funcs from queue
func Remove(fs ...Func) {
	for _, fv := range fs {
		abcQueue.Remove(fv)
	}
}

// Start . start the funcs with concurrent
func Start() {
	for _, queue := range *abcQueue {
		go queue.f()
	}
}

// Run . run the funcs with non-concurrent
func Run() {
	for _, queue := range *abcQueue {
		queue.f()
	}
}
