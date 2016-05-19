package stack_test

import (
	"github.com/matteo.angelino/stack"
	"testing"
)

func CreateTestSet() *stack.Stack {
	s := new(stack.Stack)
	s.Push(42)
	s.Push("String")
	s.Push(20.5)

	return s
}

func TestStackCreate(t *testing.T) {
	var s stack.Stack
	for i := 0; i < 1000; i++ {
		s.Push(i)
	}

	if s.Size() != 1000 {
		t.Fail()
	}

	if s.Top().Value().(int) != 999 {
		t.Fail()
	}
}

func TestStackSize(t *testing.T) {
	s := CreateTestSet()
	if s.Size() != 3 {
		t.Fail()
	}
}

func TestStackPop(t *testing.T) {
	s := CreateTestSet()
	e := s.Pop()

	if e.Value().(float64) != 20.5 {
		t.Fail()
	}

	if s.Size() != 2 {
		t.Fail()
	}
}

func TestEmptyStack(t *testing.T) {
	s := CreateTestSet()
	for e := s.Pop(); e != nil; e = s.Pop() {

	}

	if s.Size() != 0 {
		t.Fail()
	}
}

func TestListCompliance(t *testing.T) {
	s := CreateTestSet()
	for e := s.Front(); e != nil; e = s.Next() {

	}

	if s.Size() != 0 {
		t.Fail()
	}
}
