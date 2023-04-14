package fileutil

import (
	"bufio"
	"os"
	"sync"
)

type SafeFile struct {
	sync.RWMutex
	of *os.File
}

func (sf *SafeFile) Write(filename, data string) error {
	if sf.of == nil {
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		sf.of = f
	}

	sf.Lock()

	go func() {
		defer sf.Unlock()

		wbuf := bufio.NewWriterSize(sf.of, len(data))
		wbuf.WriteString(data)
		wbuf.Flush()
	}()

	return nil
}

func (sf *SafeFile) WriteAppend(filename, data string) error {
	if sf.of == nil {
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		sf.of = f
	}

	sf.Lock()

	go func() {
		defer sf.Unlock()

		wbuf := bufio.NewWriterSize(sf.of, len(data))
		wbuf.WriteString(data)
		wbuf.Flush()
	}()

	return nil
}
