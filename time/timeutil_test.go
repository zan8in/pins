package timeutil

import "testing"

func TestFormat(t *testing.T) {
	t.Log(Format(FormatFullDateTime))
	t.Log(Format(FormatShortDateTime))
	t.Log(Format(FormatShortMonthDateTime))
	t.Log(Format(Format_1))
}
