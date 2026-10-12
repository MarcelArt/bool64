// Package bool64 provides a fast, lightweight 64-bit boolean flag set implementation
// using Go's built-in uint64 type. It allows managing and querying up to 64 independent
// boolean flags using efficient bitwise operations.
package bool64

// Bool64 represents a set of up to 64 boolean flags stored as a single uint64 value.
type Bool64 uint64

// Satisfies returns true if all of the specified flags are set in b.
func (b Bool64) Satisfies(flags Bool64) bool {
	return b&flags == flags
}

// Has returns true if at least one of the specified flags is set in b.
func (b Bool64) Has(flags Bool64) bool {
	return b&flags != 0
}

// Add sets the specified flags in b.
func (b *Bool64) Add(flags Bool64) {
	*b |= flags
}

// Remove clears the specified flags in b.
func (b *Bool64) Remove(flags Bool64) {
	*b &^= flags
}

// Toggle flips the state of the specified flags in b.
func (b *Bool64) Toggle(flags Bool64) {
	*b ^= flags
}
