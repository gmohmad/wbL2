What will this program output? Explain the output of this program. Explain how 'defer' works and their execution order.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
It will print 
2
1
Because the functions in defer statement are called right before returning from the function they were called in, even
if it panics. Deferred functions execute in LIFO(last in, first out) order. The arguments passed to deferred function are 
evaluated when the function is deferred, not when it executes.
```
