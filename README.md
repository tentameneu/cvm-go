# CVM Go

CVM algorithm implemented in Go based on paper [The CVM Algorithm for Estimating Distinct Elements in Streams](https://cs.stanford.edu/~knuth/papers/cvm-note.pdf).
This library can be used to estimate number of distinct elements in a stream when total number of elements is bigger than available buffer size for storring data about stream elements.

## Installation

Install module with:

```bash
go get github.com/tentameneu/cvm-go
```

## Usage

Usage for integer elements:

```go
package main

import (
    "fmt"
    
    "github.com/tentameneu/cvm-go"
)

func main() {
    streamInt := []int{3, 4, 1, 3, 2, 8, 9, 6, 7, 5}
    cvmInt := cvm.NewCVM(10, func(x, y int) int { return x - y })
    for _, element := range streamInt {
        fmt.Println(cvmInt.Process(element))
    }
    // Output:
    // 1
    // 2
    // 3
    // 3
    // 4
    // 5
    // 6
    // 7
    // 8
    // 9
}    
```

Usage for float elements:

```go
package main

import (
    "fmt"
    
    "github.com/tentameneu/cvm-go"
)
func main() {
    streamFloat := []float64{3.3, 4.4, 1.1, 3.3, 2.2, 8.8, 9.9, 6.6, 7.7, 5.5}
    cvmFloat := cvm.NewCVM(10, func(x, y float64) int { return int(x - y) })
    for _, element := range streamFloat {
        fmt.Println(cvmFloat.Process(element))
    }
    // Output:
    // 1
    // 2
    // 3
    // 3
    // 4
    // 5
    // 6
    // 7
    // 8
    // 9
}    
```

Usage for struct elements:

```go
package main

import (
    "fmt"

    "github.com/tentameneu/cvm-go"
)

type Person struct {
    ID   int
    Name string
}

func main() {
    streamStruct := []*Person{
        {ID: 3, Name: "Bruce"},
        {ID: 4, Name: "Clark"},
        {ID: 1, Name: "John"},
        {ID: 3, Name: "Bruce"},
        {ID: 2, Name: "Pamela"},
        {ID: 8, Name: "Selina"},
        {ID: 9, Name: "Barry"},
        {ID: 6, Name: "Harley"},
        {ID: 7, Name: "Barbara"},
        {ID: 5, Name: "Joker"},
    }
    cvmStruct := cvm.NewCVM(10, func(x, y *Person) int { return x.ID - y.ID })
    for _, element := range streamStruct {
        fmt.Println(cvmStruct.Process(element))
    }
    // Output:
    // 1
    // 2
    // 3
    // 3
    // 4
    // 5
    // 6
    // 7
    // 8
    // 9
}
```

Example of usage for buffer smaller than stream of elements. In this example stream of elements has 1_000_000 total elements, of which 50_000 are distinct.
Buffer can contain only 10_000 elements. CVM algorithm is used to estimate number of distinct elements, it naturally includes randomness, so your result may differ, but it should be close to 50_000.

```go
package main

import (
    "fmt"

    "github.com/tentameneu/cvm-go"
)

func main() {
    total, distinct := 1_000_000, 50_000
    bufferSize := 10_000

    stream := make([]int, total)
    for i := 0; i < total; i++ {
        stream[i] = i % distinct
    }

    cvmSim := cvm.NewCVM(bufferSize, func(x, y int) int { return x - y })

    for _, element := range stream {
        cvmSim.Process(element)
    }

    fmt.Printf("Estimated number of distinct elements is %d\n", cvmSim.N())
    // Output:
    // Estimated number of distinct elements is 50161
}
```
