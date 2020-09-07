package bigo

import (
	"fmt"
	"reflect"
)

func GetWorstToBest(str string) {
	m := My{}
	meth := reflect.ValueOf(m).MethodByName(str)
	meth.Call(nil)
}

type My struct{}

func (m My) Worst() string {
	str := "The worst Big-O is O(n!)"
	fmt.Println(str)
	return str
}
