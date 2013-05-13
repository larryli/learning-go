package main

import (
	"fmt"
	"reflect"
)

type Brid struct {
	Name           string
	LifeExpectance int
}

func (b *Brid) Fly() {
	fmt.Println("我会飞~~~")
}

func main() {
	sparrow := &Brid{"Sparrow", 3}
	sparrow.Fly()
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
