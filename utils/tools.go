package utils

import (
	"fmt"
	"os"
)

func WriteToCSV(st, filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(f, "%s", st)

	f.Close()
}
