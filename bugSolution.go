```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        m := make(map[string]int)
        mu := sync.Mutex{}

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock()
                        m["key"]++
                        mu.Unlock()
                }(i)
        }

        wg.Wait()
        fmt.Println(m["key"]) // Output will be 1000
}
```