package functions

func isSlice(str, subStr string) []bool {
	length := len(str)
	outstr := make([]bool, length)
	for i := 0; i < len(str); i++ {
		if i+len(subStr) > length {
			outstr[i] = false
		} else {
			if str[i:i+len(subStr)] == subStr {
				for j := i; j < i+len(subStr); j++ {
					outstr[j] = true
				}
				i += len(subStr) - 1
			} else {
				outstr[i] = false
			}
		}
	}
	return outstr
}