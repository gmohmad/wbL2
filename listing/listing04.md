What will this program output? Explain the output of this program

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
It will print numbers from 0 to 9 and then deadlock, becuase we didn't close the channel, and when loop over a channel
with 'range', it loops until the channel is closed.

```
