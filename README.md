# go-edit

A simple package that provides two functions that will compute the minimum edit distance, and
weighted edit distance between two strings.

## Usage

```go
import (
    "fmt"
    edit "go-edit"
)

func main() {
    str1 := "Hello"
    str2 := "World"

    dist := edit.MinEditDistance(str1, str2)

    fmt.Printf("Edit distance between %s and %s: %d", str1, str2, dist)
}
```

## Contribution

If you see a typo or perhaps something that could be improved upon, feel free to fork the project
modify and submit a pull request.
