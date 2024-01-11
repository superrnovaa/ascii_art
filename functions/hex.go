package functions

import (
	"strconv"
)

func hex(x string) (string, int){
	r,i := strconv.ParseUint(string(x[0:2]), 16, 64)
	g,j := strconv.ParseUint(string(x[2:4]), 16, 64)
	b,k := strconv.ParseUint(string(x[4:6]), 16, 64)
	if i!= nil || j!=nil || k != nil{
		return "0", 0
	}
	return strconv.Itoa(int(r))+","+ strconv.Itoa(int(g))+ ","+strconv.Itoa(int(b)), 1
}
