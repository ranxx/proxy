package main_test

import (
	"fmt"
	"reflect"
	"testing"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectFunc(fc interface{}) {
	t := reflect.TypeOf(fc)
	fmt.Println(t.Kind(), t.NumIn())

	name := reflect.New(t.In(0)).Elem()
	if t.In(0).Kind() == reflect.String {
		// name = reflect.ValueOf("axing")
		name.SetString("axing")
	}

	reflect.ValueOf(fc).Call([]reflect.Value{name})
}

func Test01(t *testing.T) {
	reflectFunc(func(name string) {
		fmt.Println("我是", name)
	})
}
