package mem_table_test

import (
	"fmt"
	"math/rand"
	"root/mem_table"
	"testing"
)

var TEST_INPUT []byte = []byte(`{"id": 123,"ime": "Branko","prezime": "Doe","dob": 30,"email": "john.doe@example.com","adresa": {"ulica": "123 Main Street", "grad": "Cityville", "država": "State", "poštanski_broj": "12345"},"telefoni": [{"tip": "mobilni","broj": "555-1234"},{"tip": "fiksni","broj": "555-5678"}]}\n`)

func BenchmarkInsertDataToBuf(b *testing.B) {
	b.StopTimer()
	mem := mem_table.NewMemTable(40000)
	generatedID := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		generatedID[i] = fmt.Sprintf("%d", rand.Intn(500_000))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		mem.InsertData(generatedID[i], TEST_INPUT)
	}
}

func Benchmark(b *testing.B) {
	b.StopTimer()
	mem := mem_table.NewMemTable(40000)
	generatedID := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		generatedID[i] = fmt.Sprintf("%d", rand.Intn(500_000))
	}

	b.StartTimer()

	mem.Listen()

	for i := 0; i < b.N; i++ {
		mem.Save(generatedID[i], TEST_INPUT)
	}

	mem.Close()
}
