What will this program output? Explain the output of this program

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Answer:
It will print 'error', because in 'test' function we return value of type '*customError', and in main function we assing 
result of 'test' to err, so err != nil is true, because only the value of err is nil, but the type is *customError.
