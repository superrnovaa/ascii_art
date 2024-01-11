package functions

import (
	"fmt"
	"os"
	"strings"
)

func Result() error {
	var index = map[int]int{ //checking for no. of flags in the arguments passed
		5: 2,
		4: 2,
		3: 2,
		2: 1,
		1: 0,
	}
	if len(os.Args) < 2 || len(os.Args) > 6 { //checking the no. of the arguments passed
		PrintError()
		return nil
	}
	colors := []string{"red", "blue", "green", "white", "black", "yellow", "pink", "grey", "purple", "brown", "orange", "cyan"}
	colorAvailable := false
	color, str, slice, res := "", "", "", "" //res = resultant string
	output := ""
	isrgb := false             //rgb flag check
	template := "standard.txt" //Keeping standard banner as the default template
	arg := os.Args[1:]
	length := len(arg)
	isoutput, iscolor := false, false
	oindex, cindex := 0, 0
	var err int

	for i := 0; i < index[length]; i++ {
		if strings.Contains(arg[i], "--output=") {
			output = strings.TrimPrefix(arg[i], "--output=")
			oindex, isoutput = i, true
		} else if strings.Contains(arg[i], "--color=") {
			colorarg := arg[i]
			if len(arg[i]) == 15 && colorarg[8] == '#' {
				color, err = hex(colorarg[9:])
				if err == 0 {
					fmt.Print("error: Incorrect usage for hex color")
					return nil
				}
				isrgb = true
			} else if strings.Contains(arg[i], "rgb(") {
				color = colorarg[12 : len(colorarg)-1]
				isrgb = true
			} else {
				color = strings.TrimPrefix(arg[i], "--color=")
				for j, _ := range colors {
					if colors[j] == color {
						colorAvailable = true
						break
					}
				}
				if !colorAvailable {
					fmt.Println("This color is not available, These are the available colors:\nred, blue, green, white, black, yellow, pink, grey, purple, brown, orange, cyan")
				}
			}
			cindex, iscolor = i, true
		} else if length != 2 && length != 3 {
			// print error
			PrintError()
			return nil
		}
	}
	if length == 5 {
		slice = arg[2]
		str = arg[3]
		template = Template(arg[4])
	} else if length == 4 {
		if oindex < cindex {
			slice = arg[2]
			str = arg[3]
		} else if oindex > cindex {
			str = arg[2]
			template = Template(arg[3])
		}
	} else if length == 3 {
		if iscolor {
			slice = arg[1]
			str = arg[2]
		} else if isoutput {
			str = arg[1]
			template = Template(arg[2])
		} else {
			fmt.Println("Syntax error: no flag detected")
			PrintError()
			return nil
		}
	} else if length == 2 {
		if iscolor || isoutput {
			str = arg[1]
		} else {
			str = arg[0]
			template = Template(arg[length-1])
		}
	} else if length == 1 {
		str = arg[0]
		if len(str) > 8 && (str[0:9] == "--output=" || str[0:8] == "--color=") {
			fmt.Println("Please enter text")
			PrintError()
			return nil
		}
	}
	args := strings.Split(str, "\\n")
	isempty := true
	for _, word := range args {
		if word != "" {
			isempty = false
		}
		if isempty {
			size := len(args) - 1
			for h := 0; h < size; h++ {
				fmt.Println()
				res += "\n"
			}
			//// exit program
		}
	}
	if iscolor {
		Printcolor(args, color, slice, template, isrgb)
	}
	if isoutput {
		res = Printoutput(args, template)
		Print(res, output)
	}
	if !isoutput && !iscolor {
		res = Printoutput(args, template)
		fmt.Print(res)
	}
	return nil
}
