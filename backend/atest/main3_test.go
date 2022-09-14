package main

import (
	"fmt"
	"reflect"
	"testing"
)

type ErrorImpl struct {
	Name string
}

func (e *ErrorImpl) Error() string {
	// e.Name = "ksdjf" // 这里在 Test3 的写法应该panic
	return "error here111 *ErrorImpl"
}

var ei *ErrorImpl
var e error

func ErrorImplFun() error {
	return ei
}

func Test3(t *testing.T) {
	f := ErrorImplFun()
	fmt.Println(f == nil)
	fmt.Println(f.Error())
	fmt.Println(reflect.TypeOf(f).String())
}

func ErrorImplFun2() error {
	if ei == nil {
		return nil
	}
	return ei
}

func Test32(t *testing.T) {
	f := ErrorImplFun2()
	fmt.Println(f == nil)
	fmt.Println(reflect.TypeOf(f).String())
}

// //输出:

// false
