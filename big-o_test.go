package bigo

import (
	"regexp"
	"testing"
)

func TestGetWorstToBestVerTwoWithTheBest(t *testing.T) {
	str := " best Big-O "
	want := regexp.MustCompile(`\b` + str + `\b`)
	msg, err := GetWorstToBestVerTwo("the best")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`GetWorstToBestVerTwo("the best") = %q, %v, want match for %#q`, msg, err, want)
	}
}
func TestGetWorstToBestVerTwoWithBest(t *testing.T) {
	str := " best Big-O "
	want := regexp.MustCompile(`\b` + str + `\b`)
	msg, err := GetWorstToBestVerTwo("best")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`GetWorstToBestVerTwo("best") = %q, %v, want match for %#q`, msg, err, want)
	}
}
func TestGetWorstToBestVerTwoWithSoso(t *testing.T) {
	str := " okay Big-O "
	want := regexp.MustCompile(`\b` + str + `\b`)
	msg, err := GetWorstToBestVerTwo("so so")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`GetWorstToBestVerTwo("so so") = %q, %v, want match for %#q`, msg, err, want)
	}
}
func TestGetWorstToBestVerTwoWithOk(t *testing.T) {
	str := " okay Big-O "
	want := regexp.MustCompile(`\b` + str + `\b`)
	msg, err := GetWorstToBestVerTwo("OK")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`GetWorstToBestVerTwo("OK") = %q, %v, want match for %#q`, msg, err, want)
	}
}
func TestGetWorstToBestVerTwoWithOkay(t *testing.T) {
	str := " okay Big-O "
	want := regexp.MustCompile(`\b` + str + `\b`)
	msg, err := GetWorstToBestVerTwo("okay")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`GetWorstToBestVerTwo("okay") = %q, %v, want match for %#q`, msg, err, want)
	}
}
