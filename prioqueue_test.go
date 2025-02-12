package prioqueue

import (
	"cmp"
	"reflect"
	"testing"
)

func Test_PrioQueue_New(t *testing.T) {
	q := NewPrioQueue(cmp.Less, hashInt)
	assert(t, q.Len() == 0, "expected length 0")
}

func Test_PrioQueue_Push_and_pop(t *testing.T) {
	q := NewPrioQueue(cmp.Less, hashInt)
	q.Push(21)
	q.Push(12)
	q.Push(7)

	assert(t, q.Len() == 3, "expected length 3")

	v, found := q.Pop()
	assert(t, v == 7, "expected 7 to be popped first")
	assert(t, found == true, "expected found to be true")
	assert(t, q.Len() == 2, "expected length of 2 after first pop")

	v, found = q.Pop()
	assert(t, v == 12, "expected 12 to be popped second")
	assert(t, found == true, "expected found to be true")
	assert(t, q.Len() == 1, "expected length of 1 after second pop")

	v, found = q.Pop()
	assert(t, v == 21, "expected 21 to be popped last")
	assert(t, found == true, "expected found to be true")
	assert(t, q.Len() == 0, "expected length of 0 after last pop")

	v, found = q.Pop()
	assert(t, v == 0, "expected zero value to be returned")
	assert(t, found == false, "expected found to be false")
}

func Test_PrioQueue_Update(t *testing.T) {
	q := NewPrioQueue(cmp.Less, hashInt)
	q.Push(21)
	q.Push(12)
	q.Push(7)

	ok := q.Update(7, func(i int) int {
		return i * 2
	})

	assert(t, ok == true, "expected update to succeed")
	var (
		expected = []int{12, 14, 21}
		actual   = []int{}
	)

	for {
		v, found := q.Pop()
		if found == false {
			break
		}

		actual = append(actual, v)
	}

	assert(t, reflect.DeepEqual(actual, expected), "expected %+v", expected)
}

func hashInt(k int) int {
	return k
}

func assert(t *testing.T, assertion bool, format string, args ...any) {
	if !assertion {
		t.Fatalf(format, args...)
	}
}
