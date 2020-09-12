package bigo

// This package is for exercising purposes only

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type My struct{}

// ver I
var methodMap map[string]string

// ver II
var methodMapWithActualMethods map[string]func() string

// ver I
func GetWorstToBestVerOne(str string) (string, error) {
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

// ver II
func GetWorstToBestVerTwo(str string) (string, error) {
	lwrStr := strings.ToLower(str)
	setMethodMap()
	if methodMapWithActualMethods[lwrStr] == nil {
		return "Error!", errors.New(("No such method: " + strings.Title(lwrStr)))
	}
	m := methodMapWithActualMethods[lwrStr]()
	return m, nil
}

// ver I
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

// ver II
func worst() string {
	return "The worst Big-O is O(n!)"
}
func worse() string {
	return "A worse Big-O is O(2^n)"
}
func bad() string {
	return "A bad Big-O is O(n^2)"
}
func soso() string {
	return "An okay Big-O is O(n log(n))"
}
func good() string {
	return "A good Big-O is O(n)"
}
func better() string {
	return "A better Big-O is O(log(n))"
}
func best() string {
	return "The best Big-O is O(1)"
}

func setMethodMap() {
	// ver I
	methodMap = make(map[string]string)
	methodMap["worst"] = "Worst"
	methodMap["the worst"] = "Worst"
	methodMap["worse"] = "Worse"
	methodMap["bad"] = "Bad"
	methodMap["so so"] = "Soso"
	methodMap["okay"] = "Soso"
	methodMap["ok"] = "Soso"
	methodMap["good"] = "Good"
	methodMap["better"] = "Better"
	methodMap["best"] = "Best"
	methodMap["the best"] = "Best"

	// ver II
	methodMapWithActualMethods = make(map[string]func() string)
	methodMapWithActualMethods["worst"] = worst
	methodMapWithActualMethods["the worst"] = worst
	methodMapWithActualMethods["worse"] = worse
	methodMapWithActualMethods["bad"] = bad
	methodMapWithActualMethods["so so"] = soso
	methodMapWithActualMethods["okay"] = soso
	methodMapWithActualMethods["ok"] = soso
	methodMapWithActualMethods["good"] = good
	methodMapWithActualMethods["better"] = better
	methodMapWithActualMethods["best"] = best
	methodMapWithActualMethods["the best"] = best
}
