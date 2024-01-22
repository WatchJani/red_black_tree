package smart_buffer

import "testing"

const INSERTING_DATA string = `{"id": 123,"ime": "John","prezime": "Doe","dob": 30,"email": "john.doe@example.com","adresa": {"ulica": "123 Main Street", "grad": "Cityville", "država": "State", "poštanski_broj": "12345"},"telefoni": [{"tip": "mobilni","broj": "555-1234"},{"tip": "fiksni","broj": "555-5678"}]}`

func BenchmarkInsertToSBuffer(b *testing.B) {
	b.StopTimer()
	buf := New(4096)
	data := []byte(INSERTING_DATA)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf.Buff(data)
	}
}
