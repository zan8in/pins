package sliceutil

import (
	"testing"
)

func TestSafe(t *testing.T) {
	ss := SafeSlice{}
	ss.Append("a")
	ss.Append("b")

	s := ss.Key("b")
	t.Log(s)

	n := ss.Num("a")
	t.Log(n)
	ss.UpdateNum("a", 99)
	t.Log(ss.Num("a"))

	for _, v := range ss.List() {
		t.Logf("found %s", v.(string))
	}

	ss.Update(0, "x")

	for _, v := range ss.List() {
		t.Logf("found %s", v.(string))
	}

	t.Logf("Get 1 = %s", ss.Get(1).(string))
	t.Logf("Key x = %d", ss.Key("b"))

}
