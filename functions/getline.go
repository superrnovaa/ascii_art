package functions

import (
	"bufio"
	"os"
	"fmt"
)

// GetLine is a function to read font from files
func GetLine(num int, filename string) string {
	//f, e := os.Open("../fonts/" + filename)
	f, e := os.Open(filename)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	scanner := bufio.NewScanner(f)
	lineNum := 0
	line := ""
	for scanner.Scan() {
		if lineNum == num {
			line = scanner.Text()
		}
		lineNum++
	}
	return line
}