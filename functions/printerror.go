package functions

import (
	"fmt"
)

func PrintError() {
	fmt.Println("Possible usage cases:\ngo run . [STRING]\n1.EX: go run . \"something\"\ngo run . [STRING] [BANNER]\n2.EX: go run . something standard")
	fmt.Println("go run . [OPTION] [STRING] [BANNER]\n3.EX: go run . --output=<fileName.txt> something standard")
	fmt.Println("go run . [OPTION] [SUBSTRING] [STRING] \n4.EX: go run . --color=<color> <letters to be colored> \"something\"")
	fmt.Println("5.EX: go run . --output=<filename.txt> --color=<color> <letters to be colored> \"something\"")
	/////////////////////////////////
	fmt.Println("go run . [OPTION] [STRING]\n6.EX: go run . --color=<color> \"something\"\n7.EX: go run . --output=<filename.txt> \"something\"")
	fmt.Println("8.EX: go run . --output=<filename.txt> --color=<color> \"something\"")
	fmt.Println("go run . [OPTION] [STRING] [BANNER]\n9.EX: go run . --color=<color> --output=<filename.txt> \"string\" banner")
	fmt.Println("10.EX: go run . --color=<color> --output=<filename.txt> <letters to be colored> \"something\" banner")
}