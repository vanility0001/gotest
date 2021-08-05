package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func Assert(x interface{}, y interface{}) (bool, error) {
	if x == y && reflect.TypeOf(x).String() == reflect.TypeOf(y).String() {
		return true, nil
	} else {
		return false, errors.New("the two arguments do not match")
	}
}

func AssertType(x interface{}, y interface{}) (bool, error) {
	if reflect.TypeOf(x).String() == reflect.TypeOf(y).String() {
		return true, nil
	} else {
		return false, errors.New("the two types do not match")
	}
}

func AssertValue(x interface{}, y interface{}) (bool, error) {
	if x == y {
		return true, nil
	} else {
		return false, errors.New("the two values do not match")
	}
}

func main() {
	b, err := AssertValue("hello", 8)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(b)
}
