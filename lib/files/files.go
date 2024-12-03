package files

import (
	"bufio"
	"fmt"
	"os"
)

func Readfilelines(path string) []string {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}