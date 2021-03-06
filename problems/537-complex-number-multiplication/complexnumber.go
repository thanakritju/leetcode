package complexnumbermultiplication

import (
	"fmt"
	"regexp"
	"strconv"
)

type complex struct {
	real int
	img  int
}

func (c complex) Multiply(c2 complex) complex {
	return complex{real: (c.real*c2.real - c.img*c2.img), img: (c.real*c2.img + c.img*c2.real)}
}

func (c complex) String() string {
	return fmt.Sprintf("%d+%di", c.real, c.img)
}

func complexNumberMultiply(num1 string, num2 string) string {
	com1 := createComplex(num1)
	com2 := createComplex(num2)
	return com1.Multiply(com2).String()
}

func createComplex(str string) complex {
	r, _ := regexp.Compile(`(\-?)([0-9]+)\+(\-?)([0-9]+)i`)
	matches := r.FindStringSubmatch(str)
	real, _ := strconv.Atoi(matches[2])
	img, _ := strconv.Atoi(matches[4])
	if matches[1] == "-" {
		real = real * -1
	}
	if matches[3] == "-" {
		img = img * -1
	}
	return complex{real: real, img: img}
}
