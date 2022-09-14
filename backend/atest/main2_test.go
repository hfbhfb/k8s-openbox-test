package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Foo struct {
	One string
	Two int
}

func (f Foo) CallFuncNoArgs() {
	fmt.Printf("One is: %s, Two is: %d\n", f.One, f.Two)

}
func (f Foo) CallFuncWithArgs(name string, a int) {
	fmt.Printf("name:%s a:%d\n", name, a)
}

func (f Foo) CallFuncWithArgsAndRet(one, two int) (int, string) {
	fmt.Printf("one:%d two:%d\n", one, two)
	return one + two, "CallFuncWithArgsAndRet string"
}

func TestGG(t *testing.T) {
	var foo = &Foo{"Award", 666}
	// 方法名必须大写
	// if name not FuncName will panic  reflect: call of reflect.Value.Call on zero Value
	reflect.ValueOf(foo).MethodByName("CallFuncNoArgs").Call([]reflect.Value{})

	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf("Colin")
	in[1] = reflect.ValueOf(111)
	_ = reflect.ValueOf(foo).MethodByName("CallFuncWithArgs").Call(in)

	in[0] = reflect.ValueOf(1)
	in[1] = reflect.ValueOf(2)
	values := reflect.ValueOf(foo).MethodByName("CallFuncWithArgsAndRet").Call(in)
	//fmt.Printf("return value:%v\n", value[0].Interface().(int))
	for _, value := range values {
		fmt.Printf("value:%v\n", value)
	}
}
