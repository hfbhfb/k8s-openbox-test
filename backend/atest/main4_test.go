package main

import (
	"fmt"
	"testing"
)

type ErrorImpl4 struct {
	Name string
}

func (e *ErrorImpl4) Error() string {
	e.Name = "in point opt effect out call" // 这里在 Test3 的写法应该panic
	fmt.Println(e.Name)
	return "error here111 *ErrorImpl"
}

func (e ErrorImpl4) ErrorNotInPoing() string {
	e.Name = "not in poing" // 这里在 Test3 的写法应该panic
	fmt.Println(e.Name)
	return "error here111 *ErrorImpl"
}

func Test4(t *testing.T) {
	a := ErrorImpl4{}
	a.Name = "11"
	fmt.Println(a.Name)
	a.Error()
	fmt.Println(a.Name)
	a.ErrorNotInPoing()
	fmt.Println(a.Name)

}
