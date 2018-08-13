package main

import (
	"strings"
	"testing"
)

func TestAsk(t *testing.T) {
	adder := NewAdder()
	defer adder.Close()
	q := []int{1, 2, 3}
	a := adder.Ask(q)
	s := strings.Join(a, ", ")
	if s != "Sum is 6" {
		t.FailNow()
	}
}

func TestOracle(t *testing.T) {
	oracle := NewOracle(NewAdder(), NewMultiplier())
	defer oracle.Close()
	q := []int{1, 2}
	a := oracle.Ask(q)
	s := strings.Join(a, ", ")
	if s != "Sum is 3, Product is 2" {
		t.FailNow()
	}
}
