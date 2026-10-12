# bool64

[![Go Reference](https://pkg.go.dev/badge/github.com/MarcelArt/bool64.svg)](https://pkg.go.dev/github.com/MarcelArt/bool64)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`bool64` is a fast, lightweight, and zero-allocation Go library for managing a set of up to 64 boolean flags in a single `uint64` wrapper using bitwise operations. It is perfect for handling user permissions, settings/configuration flags, state sets, and any other bitmask-oriented logic efficiently.

## Installation

To install `bool64`, use `go get`:

```bash
go get github.com/MarcelArt/bool64
```

## Features

- **Zero Allocation**: Utilizes standard bitwise operations under the hood, making it extremely fast with zero memory overhead.
- **64 Independent Flags**: Up to 64 flags can be stored in a single variable.
- **Fluent & Readable API**: Replaces complex bitwise syntax (`|=`, `&^=`, `^=`, `&`) with clean method calls (`Add`, `Remove`, `Toggle`, `Has`, `Satisfies`).

## Usage

### 1. Defining Flags

Define your flags as `bool64.Bool64` constants using bit shifts (`1 << iota`):

```go
package main

import (
	"fmt"
	"github.com/MarcelArt/bool64"
)

const (
	UserRead bool64.Bool64 = 1 << iota
	UserWrite
	UserExecute
	GroupRead
	GroupWrite
	GroupExecute
	OtherRead
)
```

### 2. Manipulating Flags

Use `Add`, `Remove`, and `Toggle` to modify a flag set. These methods modify the variable in-place (requiring a pointer receiver).

```go
func main() {
	var permissions bool64.Bool64

	// Add flags
	permissions.Add(UserRead | UserWrite)
	fmt.Printf("After Add: %08b\n", permissions) // 00000011

	// Remove a flag
	permissions.Remove(UserWrite)
	fmt.Printf("After Remove: %08b\n", permissions) // 00000001

	// Toggle a flag (Off -> On)
	permissions.Toggle(UserExecute)
	fmt.Printf("After Toggle (On): %08b\n", permissions) // 00000101

	// Toggle a flag (On -> Off)
	permissions.Toggle(UserRead)
	fmt.Printf("After Toggle (Off): %08b\n", permissions) // 00000100
}
```

### 3. Querying Flags

Use `Has` and `Satisfies` to check the state of the flags:

- **`Has(flags)`**: Returns `true` if **at least one** of the specified flags is set (equivalent to `b & flags != 0`).
- **`Satisfies(flags)`**: Returns `true` if **all** of the specified flags are set (equivalent to `b & flags == flags`).

```go
func checkPermissions(p bool64.Bool64) {
	// Check if the user has AT LEAST UserRead OR GroupRead permissions
	if p.Has(UserRead | GroupRead) {
		fmt.Println("Allowed to read (either user or group read flag is present)")
	}

	// Check if the user satisfies BOTH UserRead AND UserWrite permissions
	if p.Satisfies(UserRead | UserWrite) {
		fmt.Println("Allowed to read and write (both flags are present)")
	}
}
```

## API Reference

### `type Bool64 uint64`
The underlying type representing a set of up to 64 flags.

#### `func (b Bool64) Satisfies(flags Bool64) bool`
Returns `true` if all of the specified flags are set.

#### `func (b Bool64) Has(flags Bool64) bool`
Returns `true` if at least one of the specified flags is set.

#### `func (b *Bool64) Add(flags Bool64)`
Sets the specified flags.

#### `func (b *Bool64) Remove(flags Bool64)`
Clears/unsets the specified flags.

#### `func (b *Bool64) Toggle(flags Bool64)`
Toggles the specified flags.

## Testing

To run the unit tests:

```bash
go test -v ./...
```

For coverage report:

```bash
go test -coverprofile=cover.out && go tool cover -html=cover.out
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
