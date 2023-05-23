package randutil

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		r, err := IntN(999999)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(r)
	}
}

func TestRandLowercase(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandLowercase(18))
		fmt.Println(RandUppercase(18))
		fmt.Println(Randcase(18))
		fmt.Print("--------------------\n")
	}
}
