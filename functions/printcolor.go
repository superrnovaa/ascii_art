package functions

import (
	"fmt"
)

func Printcolor(args []string, color string, slice string, template string, isrgb bool) {
	for _, word := range args {
		strArray := []bool{}
		if slice != "" {
			strArray = isSlice(word, slice)
		}
		for i := 1; i <= 8; i++ {
			for n, letter := range word {
				if slice != "" {
					if strArray[n] {
						fmt.Print(ESCseq(color, isrgb), GetLine(int(letter-' ')*9+i, template))
					} else {
						fmt.Print(ESCseq("white", isrgb), GetLine(int(letter-' ')*9+i, template))
					}
				} else {
					fmt.Print(ESCseq(color,isrgb), GetLine(int(letter-' ')*9+i, template))
				}
			}
			if word != "" {
				fmt.Println()
			}
		}
		if word == "" {
			fmt.Println()
		}
	}
}