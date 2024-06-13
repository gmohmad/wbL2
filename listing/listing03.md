What will this program output? Explain the output of this program. Explain internals of interfaces and how they differ 
from empty interfaces.
```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Answer:

It will print
```
<nil>
false
```
fmt.Println(err) prints <nil> because the value of the err returned by Foo() is nil, but fmt.Println(err == nil) prints
false, because an interface value is nil only when both it's type and it's value are nil. In our case, the value is nil,
but the type is '*os.PathError', that's why it prints false. 
An interface in Go consists of a type and a value. Type is the static type of the interface itself, and value contains the 
dynamic type and the dynamic value of the object assigned to the interface. Empty interfaces have no methods that objects 
must implement, so we can use them with any type.
