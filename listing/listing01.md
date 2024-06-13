What will this program output? Explain the output of this program.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
It will print [77 78 79]
We just slice 'a' using a[1:4] which will take all element from 1 index to 4 index non-inclusive, and assign the result
to 'b'.
```
