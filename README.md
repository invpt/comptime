# comptime
Some compile-time utilities for Go, built using generics. No codegen needed.

## Examples

An example for `comptime.Bool`:
```go
package somelibrary

import (
    "log"

    "github.com/invpt/comptime"
)

// ... some unrelated code ...

// A function that performs a task really quickly and prints debug messages
// whenever its generic DebugInfo argument is comptime.True.
func PerformTask[DebugInfo comptime.Bool](/* ... */) {
    // ... some code that would be slowed down by unnecessary printing ...

    // if DebugInfo is comptime.True, this if statement will be present
    // in the compiled program. If it is comptime.False, the if statement should
    // be absent from the final binary, resulting in zero performance loss.
    if comptime.GetBool[DebugInfo]() {
        log.Println("some info that's only useful when you're debugging stuff")
    }

    // ... maybe some more performance-intensive code ...
}

// ... more unrelated code ...
```

TODO: Add examples for `comptime.Ordinal` and `comptime.Union`.
