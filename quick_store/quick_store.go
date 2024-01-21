package quick_store

type QuickStore[T any] struct {
	store    []T
	pointer  int
	capacity int
}

// func New[T int | string | bool | float32 | float64](capacity int) *QuickStore[T] {

// capacity is the space for our store (space can dynamic be expanded 2x)
func New[T any](capacity int) QuickStore[T] {
	return QuickStore[T]{
		store:    make([]T, capacity),
		capacity: capacity,
		pointer:  -1,
	}
}

// Add new element to store
func (q *QuickStore[T]) Append(item T) {
	q.pointer++

	if q.pointer > q.capacity {
		q.Grow()
	}

	q.store[q.pointer] = item
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
	return q.store[q.pointer]
}

// return to us number of element in our store
func (q QuickStore[T]) Len() int {
	return q.pointer
}

// return
func (q QuickStore[T]) GetStore() []T {
	return q.store[:q.pointer]
}
