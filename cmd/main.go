package main

import (
	"fmt"

	"github.com/tmw/go-prioqueue"
)

type Color struct {
	name string
	prio int
}

func compareColor(a, b Color) bool {
	return a.prio < b.prio
}

func hashColor(a Color) string {
	return a.name
}

func main() {
	q := prioqueue.NewPrioQueue(compareColor, hashColor)

	q.Push(Color{name: "red", prio: 15})
	q.Push(Color{name: "blue", prio: 21})
	q.Push(Color{name: "orange", prio: 2})

	ok := q.Update("orange", func(p Color) Color {
		p.prio += 200
		return p
	})

	if !ok {
		panic("nok.")
	}

	v, _ := q.Pop()
	fmt.Printf("first = %s\n", v.name)

	v, _ = q.Pop()
	fmt.Printf("second = %s\n", v.name)

	v, _ = q.Pop()
	fmt.Printf("third = %s\n", v.name)
}
