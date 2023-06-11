package csvutil

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	fileutil "github.com/zan8in/pins/file"
)

func ReadFile(filename string) (chan []string, error) {
	var err error

	if !fileutil.FileExists(filename) {
		return nil, fmt.Errorf("%s does not exist", filename)
	}

	out := make(chan []string)
	go func() {
		defer close(out)

		file, err := os.Open(filename)
		if err != nil {
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}

			out <- record
		}

	}()

	return out, err
}
