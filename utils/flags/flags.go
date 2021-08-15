package flags

import "strings"

// Flags is a csv type storage type for marking
type Flags string

// New news a *Flags
func New(flags ...string) *Flags {
	f := Flags(strings.Join(flags, ","))
	return &f
}
