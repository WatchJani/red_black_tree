package quick_store

import "fmt"

type QuickStore[T any] struct {
	store    []T
	pointer  int
	capacity int
}

// capacity is the space for our store (space can dynamic be expanded 2x)
func New[T any](capacity int) QuickStore[T] {
	return QuickStore[T]{
		store:    make([]T, capacity),
		capacity: capacity,
		pointer:  0,
	}
}

// Add new element to store
func (q *QuickStore[T]) Append(item T) {
	if q.pointer == q.capacity {
		q.Grow()
	}

	q.store[q.pointer] = item
	q.pointer++
}

// Allocate more space in memory for our items => 2x more
func (q *QuickStore[T]) Grow() {
	q.store = append(q.store, make([]T, q.capacity)...)
	q.capacity *= 2
}

// Return to us capacity of store
func (q QuickStore[T]) Cap() int {
	return q.capacity
}

// Delete last element in the store
func (q *QuickStore[T]) Delete() {
	if q.pointer-1 < 0 {
		panic("[QuickStore] Store is already empty")
	}

	q.pointer--
}

// return to us last value from our store
func (q QuickStore[T]) Get() T {
	return q.store[q.pointer-1]
}

// return to us number of element in our store
func (q QuickStore[T]) Len() int {
	return q.pointer
}

// return all element in our store
func (q QuickStore[T]) GetStore() []T {
	return q.store[:q.pointer]
}

// get slice
func (q *QuickStore[T]) GetStoreAll() []T {
	return q.store
}

// get element from store on specific place by index
func (q QuickStore[T]) GetById(id int) T {
	if id < 0 || id > q.capacity {
		panic(fmt.Sprintf("[QuickStore] This index [%d] is not exist", id))
	}

	return q.store[id]
}

// reset our store, we haw 0 elements in store
func (q *QuickStore[T]) Reset() {
	q.pointer = 0
}

// return pointer which are on the last element
func (q QuickStore[T]) GetPointer() int {
	return q.pointer - 1
}

// custom pointer size
func (q *QuickStore[T]) SetPointer(size int) {
	q.pointer += size
}

// return capacity initialized memory for quick store
func (q QuickStore[T]) GetCapacity() int {
	return q.capacity
}

// check is buffer full
func (q QuickStore[T]) Check(data int) bool {
	return q.pointer+data > q.capacity
}
