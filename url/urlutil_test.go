package urlutil

import (
	"fmt"
	"testing"
)

func TestDomainName(t *testing.T) {
	a, _ := Hostname("http://example.com:8080/abc.php?ada")
	fmt.Println(a)

	a1, _ := Hostname("xxx.example.com:8080")
	fmt.Println("++++", a1)
}
