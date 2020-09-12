package bigo

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type My struct{}

var methodMap map[string]string

func GetWorstToBest(str string) (string, error) {
	lwrStr := strings.ToLower(str)
	setMethodMap()
	if methodMap[lwrStr] == "" {
		return "Error!", errors.New(("No such method: " + strings.Title(lwrStr)))
	}
	m := My{}
	meth := reflect.ValueOf(m).MethodByName(methodMap[lwrStr])
	meth.Call(nil)
	return "Done!", nil
}

func (m My) Worst() string {
	str := "The worst Big-O is O(n!)"
	fmt.Println(str)
	return str
}

func (m My) Worse() string {
	str := "A worse Big-O is O(2^n)"
	fmt.Println(str)
	return str
}

func (m My) Bad() string {
	str := "A bad Big-O is O(n^2)"
	fmt.Println(str)
	return str
}

func (m My) Soso() string {
	str := "An okay Big-O is O(n log(n))"
	fmt.Println(str)
	return str
}

func (m My) Good() string {
	str := "A good Big-O is O(n)"
	fmt.Println(str)
	return str
}

func (m My) Better() string {
	str := "A better Big-O is O(log(n))"
	fmt.Println(str)
	return str
}

func (m My) Best() string {
	str := "The best Big-O is O(1)"
	fmt.Println(str)
	return str
}

func setMethodMap() {
	methodMap = make(map[string]string)
	methodMap["worst"] = "Worst"
	methodMap["the worst"] = "Worst"
	methodMap["wrse"] = "Worse"
	methodMap["so so"] = "Soso"
	methodMap["okay"] = "Soso"
	methodMap["ok"] = "Soso"
	methodMap["good"] = "Good"
	methodMap["better"] = "Better"
	methodMap["best"] = "Best"
	methodMap["the best"] = "Best"
}
