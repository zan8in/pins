package sliceutil

import (
	"testing"
)

func TestSafe(t *testing.T) {
	ss := SafeSlice{}
	ss.items = append(ss.items, Items{value: "a"})
	ss.items = append(ss.items, Items{value: "b"})
	ss.items = append(ss.items, Items{value: "c"})
	ss.items = append(ss.items, Items{value: "d"})
	ss.items = append(ss.items, Items{value: "e"})
	ss.items = append(ss.items, Items{value: "f"})

	s := ss.Key("d")
	t.Log(s)
}
