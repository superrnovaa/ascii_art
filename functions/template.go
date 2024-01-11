package functions

import (
	"fmt"
)

func Template(x string) string {
	template := "standard.txt" //Keeping standard banner as the default template
	if x == "shadow" {
		template = "shadow.txt"
	} else if x == "thinkertoy" {
		template = "thinkertoy.txt"
	} else if x != "standard" {
		// print error
		fmt.Println("no such template will use standard instead")
	}
	return template
}
