# goset

set集合，基于map实现。

## example

```go
package main

import (
	"fmt"
	"github.com/Hami-Lemon/goset"
)

func main() {
	set := goset.New[int]()
	set.Add(1)
	set.Size()       // 1
	set.IsEmpty()    // false
	set.IsNotEmpty() //true
	set.Remove(1)
	set.Contains(1) //true

	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.ForEach(func(v int) {
		fmt.Printf("%d ", v) //1 2 3
	})
	fmt.Println()
}

```