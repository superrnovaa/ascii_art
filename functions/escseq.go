package functions

import (
	"strings"
)
// ESCseq is a function that return an esc seqnuence of necessary color
func ESCseq(color string, isrgb bool) string {
	if isrgb {
		y := strings.ReplaceAll(color, " ", "")
		x := strings.Split(y, ",")
		if len(x) != 3 {
			//fmt.Println("error: --color=rgb(r,g,b))
			return "\u001b[38;2;255;255;255m"
		}
return "\u001b[38;2;" + x[0] + ";" + x[1] + ";" + x[2] + "m"
	} 
switch color {
case "white":
	return "\u001b[38;2;255;255;255m"
case "black":
	return "\u001b[38;2;0;0;0m"
case "red":
	return "\u001b[38;2;255;0;0m"
case "green":
	return "\u001b[38;2;0;255;0m"
case "blue":
	return "\u001b[38;2;0;0;255m"
case "yellow":
	return "\u001b[38;2;255;255;0m"
case "pink":
	return "\u001b[38;2;255;0;255m"
case "grey":
	return "\u001b[38;2;128;128;128m"
case "purple":
	return "\u001b[38;2;160;32;255m"
case "brown":
	return "\u001b[38;2;160;128;96m"
case "orange":
	return "\u001b[38;2;255;160;16m"
case "cyan":
	return "\u001b[38;2;0;183;235m"
}
return "\u001b[38;2;255;255;255m"
}
