package base

import "fmt"

type MyStuff struct {
	Field string
	Extra string
}

func (m *MyStuff) String() string {
	return fmt.Sprintf("v2: %s %s\n", m.Field, m.Extra)
}
