package fileutil

import (
	"strconv"
	"sync"
	"testing"
)

func TestWrite(t *testing.T) {

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			sf := SafeFile{}
			sf.Write("test.txt", strconv.Itoa(i)+"\n")
		}(i)
	}
	wg.Wait()
}
