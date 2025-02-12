package prioqueue

import (
	"container/heap"
)

type compareFn[T any] func(a, j T) bool
type hashFn[T any, H comparable] func(a T) H
type inner[T any, H comparable] struct {
	items     []T
	indexes   map[H]int
	compareFn compareFn[T]
	hashFn    hashFn[T, H]
}

func (in inner[T, H]) Len() int {
	return len(in.items)
}

func (in inner[T, H]) At(index int) T {
	return in.items[index]
}

func (in inner[T, H]) Less(a, j int) bool {
	itemA := in.At(a)
	itemJ := in.At(j)
	return in.compareFn(itemA, itemJ)
}

func (in inner[T, H]) Swap(i, j int) {
	itemI, itemJ := in.items[i], in.items[j]
	hashI, hashJ := in.hashFn(itemI), in.hashFn(itemJ)
	in.items[i], in.items[j] = itemJ, itemI
	in.indexes[hashI], in.indexes[hashJ] = j, i
}

func (in *inner[T, H]) Push(val any) {
	v := val.(T)
	in.items = append(in.items, v)
	hash := in.hashFn(v)
	in.indexes[hash] = len(in.items) - 1
}

func (in *inner[T, H]) Pop() any {
	end := len(in.items) - 1
	item := in.items[end]
	in.items = in.items[:end]
	hash := in.hashFn(item)
	delete(in.indexes, hash)
	return item
}

type PrioQueue[T any, H comparable] struct {
	inner inner[T, H]
}

func (pq *PrioQueue[T, H]) Push(v T) {
	heap.Push(&pq.inner, v)
}

func (pq *PrioQueue[T, H]) Pop() (T, bool) {
	var res T
	if pq.inner.Len() == 0 {
		return res, false
	}

	v := heap.Pop(&pq.inner)
	if res, ok := v.(T); ok {
		return res, true
	}

	return res, false
}

func (pq *PrioQueue[T, H]) Update(hash H, updateFn func(T) T) bool {
	idx, found := pq.inner.indexes[hash]
	if !found {
		return false
	}

	pq.inner.items[idx] = updateFn(pq.inner.items[idx])
	heap.Fix(&pq.inner, idx)

	return true
}

func (pq *PrioQueue[T, H]) Len() int {
	return pq.inner.Len()
}

func NewPrioQueue[T any, H comparable](
	compareFn compareFn[T],
	hashFn hashFn[T, H],
) PrioQueue[T, H] {
	pq := PrioQueue[T, H]{
		inner: inner[T, H]{
			compareFn: compareFn,
			hashFn:    hashFn,
			indexes:   make(map[H]int),
		},
	}

	heap.Init(&pq.inner)
	return pq
}
