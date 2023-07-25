package base

import "fmt"

type MyStuff struct {
	Field string
}

func (m *MyStuff) String() string {
	return fmt.Sprintf("v2: %s\n", m.Field)
}
