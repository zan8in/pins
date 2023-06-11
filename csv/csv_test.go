package csvutil

import (
	"fmt"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	files, err := ReadFile("D:\\代码审计\\泛微\\泛微OA e-cology9 sql注入漏洞 CNVD-2023-12632\\go\\assets_2023610.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for f := range files {
		fmt.Println(strings.Join(f, ","))
	}
}
