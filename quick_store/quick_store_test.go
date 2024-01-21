package quick_store

import (
	"testing"
)

// test adding to our quick store
func TestInsert(t *testing.T) {
	lite_store := New[int](20)

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
		lite_store.Append(test.insert)

		//lite_store.store[index] => Each element should be in the next position if added within the array
		if actualValue := lite_store.store[index]; actualValue != test.expected {
			t.Errorf("index: %d | Got %v | expected %v", index, actualValue, test.expected)
		}
	}
}


