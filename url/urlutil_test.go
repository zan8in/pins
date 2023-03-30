package urlutil

import (
	"fmt"
	"testing"
)

func TestDomainName(t *testing.T) {
	a, _ := DomainName("http://example.com:8080/abc.php?ada")
	fmt.Println(a)

	a1, _ := DomainName("xxx.example.com:8080")
	fmt.Println("++++", a1)
}
