```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        m := make(map[string]int)

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m["key"]++
                }(i)
        }

        wg.Wait()
        fmt.Println(m["key"]) // Output might not be 1000
}
```