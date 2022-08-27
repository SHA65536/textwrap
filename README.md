# textwrap
[![Go Reference](https://pkg.go.dev/badge/github.com/sha65536/textwrap.svg)](https://pkg.go.dev/github.com/sha65536/textwrap)

`textwrap` is a package made to simplify formatting text in a size limited space.

## Usage
Import the package into your project
```go
import "github.com/sha65536/textwrap"
```
Wrap a text using Wrap
```go
result, err := textwrap.Wrap("Hello There. General Kenobi.", 12)
```
Will result in:
```
Hello There.
General
Kenobi.
```
