package main

import (
	"math/rand"
	"testing"
)

const TEST_SAMPLE = 40_000

func TestInsert(t *testing.T) {
	tree := NewRBTree(TEST_SAMPLE)

	for i := 0; i < TEST_SAMPLE; i++ {
		num := rand.Intn(TEST_SAMPLE) //random key
		tree.Insert(num)              //insert key
	}

	
}
