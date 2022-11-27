# Base91
A simple Base91 encoder/decoder written in Go.

## How to use
```go
package main

import (
    "github.com/eloonstra/base91"
)

func main() {
    encoded := base91.Encode([]byte("Hello World!"))
    decoded := base91.Decode(encoded)
}
```

That's all.