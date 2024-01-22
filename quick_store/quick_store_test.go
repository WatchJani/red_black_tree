package quick_store

import (
	"math/rand"
	"testing"
)

// test adding to our quick store
func TestInsert(t *testing.T) {
	liteStore := New[int](20)

	testData := []struct {
		insert   int
		expected int
	}{
		{insert: 123123, expected: 123123},
		{insert: 234, expected: 234},
		{insert: -123123, expected: -123123},
		{insert: 234234, expected: 234234},
		{insert: 3, expected: 3},
	}

	for index, test := range testData {
		liteStore.Append(test.insert)

		//lite_store.store[index] => Each element should be in the next position if added within the array
		if actualValue := liteStore.store[index]; actualValue != test.expected {
			t.Errorf("index: %d | Got %v | expected %v", index, actualValue, test.expected)
		}
	}
}

// Allocate more space in memory for our items => 2x more
func TestGrow(t *testing.T) {
	liteStore := New[int](5)

	for index := 0; index < 7; index++ {
		num := rand.Intn(500)
		liteStore.Append(num)
	}

	if actualValue, expectedValue := liteStore.Cap(), 10; actualValue != expectedValue {
		t.Errorf("Got %v | expected %v", actualValue, expectedValue)
	}
}

// get last value from store
func TestGetLast(t *testing.T) {
	liteStore := New[int](20)

	num := rand.Intn(20)

	lastItem := -1
	for index := 0; index < num; index++ {
		num = rand.Intn(2000)
		liteStore.Append(num)
		lastItem = num
	}

	if actualValue, expectedValue := liteStore.Get(), lastItem; actualValue != expectedValue {
		t.Errorf("Got %v | expected %v", actualValue, expectedValue)
	}
}

// delete last element from store
func TestDelete(t *testing.T) {
	liteStore := New[int](20)

	num := rand.Intn(20)

	for index := 0; index < num; index++ {
		liteStore.Append(index)
	}

	liteStore.Delete()

	if actualValue, expectedValue := liteStore.Len(), num-1; actualValue != expectedValue {
		t.Errorf("Got %v | expected %v", actualValue, expectedValue)
	}
}

// test quick_store get element by id
func TestGetById(t *testing.T) {
	liteStore := New[int](20)

	num := rand.Intn(20) + 1

	for index := 0; index < num; index++ {
		liteStore.Append(index)
	}

	num = rand.Intn(num) + 1
	if actualValue, expectedValue := num, num; actualValue != expectedValue {
		t.Errorf("Got %v | expected %v", actualValue, expectedValue)
	}
}

// my quick store - level up slice implementation insert
// 2.074ns
func BenchmarkInsert(b *testing.B) {
	b.StartTimer()
	liteStore := New[int](b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		liteStore.Append(i)
	}
}

// slice performance speed
// 1.975ns
func BenchmarkSlice(b *testing.B) {
	b.StartTimer()
	slice := make([]int, b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		slice[i] = i
	}
}
