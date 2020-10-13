# goabc
goroutine sequential execution

# Usage

Define functions
```
func func1(){fmt.Println(11)}
func func2(){fmt.Println(22)}
```

## Add
Add functions into `goabc`'s queue, the same function can add to queue many times.
```
// add func1 two times, add func2 one time
goabc.Add(goabc.Func(func1),goabc.Func(func1),goabc.Func(func2))
```

## Remove
Remove functions from `goabc`'s queue, it will remove all the functions that in the queue.
```
// remove all func1
goabc.Remove(goabc.Func(func1))
```
## Flush
flush all functions.
```
// flush the queue
goabc.Flush()
```

## Hooker
Implement the hooker to handle the error customer.
```
goabc.SetHooker(hook)
```

## Start
Start to execute the functions in goroutine ordered.
```
goabc.Start()
```

## Run
Start to execute the functions orderd, not in goroutine.
```
goabc.Run()
```

## Random
Start to execute the functions in goroutine, the execute order is random(decide by goroutine-invoke tool).
```
goabc.Random()
```

## Customer Define Abcer
You can implement the `Abcer` interface to controll the execute ordered by youself.