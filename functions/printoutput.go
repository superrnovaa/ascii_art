package functions

func Printoutput(args []string, template string) string {
	res := ""
	for _, word := range args {
		for i := 1; i <= 8; i++ {
			for _, letter := range word {
				res += GetLine(int(letter-' ')*9+i, template)
			}
			if word != "" {
				res += "\n"
			}
		}
		if word == "" {
			res += "\n"
		}
	}
	return res
}