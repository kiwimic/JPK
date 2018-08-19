package utils

import (
	"fmt"
	"os"
	"strings"
)

func WriteToCSV(st, filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(f, "%s", st)

	f.Close()
}

func RemoveStringFromSliceOfString(value []string, pattern, replacement string, n int) []string {
	N := len(value)
	ret := make([]string, N)
	temp := ""
	for i := 0; i < N; i++ {
		temp = value[i]
		ret[i] = strings.Replace(temp, pattern, replacement, n)
	}
	return ret
}
