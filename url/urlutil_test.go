package urlutil

import (
	"fmt"
	"testing"
)

func TestDomainName(t *testing.T) {
	a, _ := Hostname("http://www.example.com:8080/abc.php?ada")
	fmt.Println(a)

	a2, _ := Hostname("http://121.37.66.33:16868/ROOT.zip")
	fmt.Println(a2)

	a3, _ := Domain("http://121.37.66.33:16868/ROOT.zip")
	fmt.Println(a3)

	a1, _ := Domain("xxx.example.com:8080")
	fmt.Println(a1)

	a4, _ := Domain("com:8080")
	fmt.Println(a4)
}
