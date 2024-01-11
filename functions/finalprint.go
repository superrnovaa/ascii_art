package functions

import (
	"path/filepath"
	"fmt"
	"os"
)

func Print(result string, filename string) {
	if filepath.Ext(filename) != ".txt" { //To verify the args passed is text file only
		fmt.Println("Only text files required!")
	} else {
		outputText, err := os.Create(filename) //creating output file
		if err != nil {
			fmt.Println("Error creating output file!")
			os.Exit(1)
		}
		//Writing the resultant string to the text file
		outputText.WriteString(result)
		//closing the file
		defer outputText.Close()
	}
}