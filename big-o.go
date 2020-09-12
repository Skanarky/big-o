package bigo

// This package is for exercising purposes only

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

func (m My) Worst() {
	fmt.Println("The worst Big-O is O(n!)")
}

func (m My) Worse() {
	fmt.Println("A worse Big-O is O(2^n)")
}

func (m My) Bad() {
	fmt.Println("A bad Big-O is O(n^2)")
}

func (m My) Soso() {
	fmt.Println("An okay Big-O is O(n log(n))")
}

func (m My) Good() {
	fmt.Println("A good Big-O is O(n)")
}

func (m My) Better() {
	fmt.Println("A better Big-O is O(log(n))")
}

func (m My) Best() {
	fmt.Println("The best Big-O is O(1)")
}

func setMethodMap() {
	methodMap = make(map[string]string)
	methodMap["worst"] = "Worst"
	methodMap["the worst"] = "Worst"
	methodMap["worse"] = "Worse"
	methodMap["so so"] = "Soso"
	methodMap["okay"] = "Soso"
	methodMap["ok"] = "Soso"
	methodMap["good"] = "Good"
	methodMap["better"] = "Better"
	methodMap["best"] = "Best"
	methodMap["the best"] = "Best"
}
