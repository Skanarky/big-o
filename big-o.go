package bigo

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func GetWorstToBest(str string) (string, error) {
	capStr := strings.Title(strings.ToLower(str))
	if capStr != "Worst" &&
		capStr != "Worse" &&
		capStr != "Soso" &&
		capStr != "Good" &&
		capStr != "Better" &&
		capStr != "Best" {
		return "Error!", errors.New("No such method")
	}
	m := My{}
	meth := reflect.ValueOf(m).MethodByName(capStr)
	meth.Call(nil)
	return "Done!", nil
}

type My struct{}

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
