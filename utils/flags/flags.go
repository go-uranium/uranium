package flags

import "strings"

type Flags string

func New(flags ...string) *Flags {
	f := Flags(strings.Join(flags, ","))
	return &f
}
