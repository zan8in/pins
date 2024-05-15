package timeutil

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	t.Log(Format(FormatFullDateTime))
	t.Log(Format(FormatShortDateTime))
	t.Log(Format(FormatShortMonthDateTime))
	t.Log(Format(Format_1))
	t.Log(GetAnyTimeAgo(180))
	fmt.Println(GetAnyTimeAgo(180))
}
