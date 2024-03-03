package mem_table

import (
	"fmt"
	"math/rand"
	"testing"
)

// // 181 words 2,547 characters
var TEST_INPUT []byte = []byte(`[
    {"color": "red", "value": "#f00"},
    {"color": "green", "value": "#0f0"},
    {"color": "blue", "value": "#00f"},
    {"color": "cyan", "value": "#0ff"},
    {"color": "magenta", "value": "#f0f"},
    {"color": "yellow", "value": "#ff0"},
    {"color": "black", "value": "#000"}
]
`)

func BenchmarkInsertDataToBuf(b *testing.B) {
	b.StopTimer()
	mem := NewMemTable(40000)
	generatedID := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		generatedID[i] = fmt.Sprintf("%d", rand.Intn(500_000))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 40_000; j++ {
			mem.InsertData(generatedID[i], TEST_INPUT)
		}
	}
}

func BenchmarkStressTest(b *testing.B) {
	mem := NewMemTable(40000)
	generatedID := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		generatedID[i] = fmt.Sprintf("%d", rand.Intn(500_000))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mem.InsertData(generatedID[i], TEST_INPUT)
	}
}
