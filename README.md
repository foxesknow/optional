[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)

# Optional
The Optional library provides functional-like optional support to Go.

### Usage

The default value for an optional value is `none`
```go
var data optional.Value[int]
```

You use the `Some` and `None` functions to initialize an optional value:
```go
func StringToInt(s string) optional.Value[int] {
    if i, err := strconv.Atoi(s); err == nil {
        return optional.Some(i)
    }

    return optional.None[int]()
```