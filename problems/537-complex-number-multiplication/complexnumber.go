package complexnumbermultiplication

import (
	"strconv"
	"strings"
)

type complex struct {
	real int
	img  int
}

func complexNumberMultiply(num1 string, num2 string) string {
	return ""
}

func createComplex(str string) complex {
	s := strings.Split(str, "+")
	real, err := strconv.Atoi(s[0])
}
