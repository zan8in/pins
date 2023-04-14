package fileutil

import (
	"fmt"
	"testing"
)

func TestCover(t *testing.T) {
	err := CoverFile("test.txt", ",]", "]")
	fmt.Println(err)
}
