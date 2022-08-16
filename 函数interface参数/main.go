package main

import (
	"fmt"
	"reflect"
)

func main() {

}

//TestSliceInterface test slice interface
func TestSliceInterface(parameters ...interface{}) {

	for _, i := range parameters {
		switch t := i.(type) {
		case string:
			fmt.Println(t, "is string")
		case int:
			fmt.Println(t, "is int")
		case []string:
			fmt.Println(t, "is string slice")
		default:
			fmt.Println(t, "is other type")
		}
	}
	//how to get element inside interface slice and interface
	insideInterface := parameters[0]
	fmt.Println("inside interface:", insideInterface)
	fmt.Println("get the int slice:", insideInterface.([]int))
	fmt.Println("get the first element in the int slice:", insideInterface.([]int)[0])

}

//TestInterface test interface
func TestInterface(parameters interface{}) {

	//for i.(type) is also ok
	v := reflect.ValueOf(parameters).Kind()
	switch v {
	case reflect.String:
		fmt.Println(parameters, "is string")
	case reflect.Int:
		fmt.Println(parameters, "is int")
	case reflect.Slice:
		fmt.Println(parameters, "is slice")
	}

}
