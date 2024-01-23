package main

import (
	"fmt"
	"math/rand"
	"root/red_black_tree"
)

func main() {
	tree := red_black_tree.NewRBTree[int, int](5000)

	for i := 0; i < 5000; i++ {
		tree.Insert(rand.Intn(5000), 5)
	}

	fmt.Println(tree.InOrderTraversal())

	// liteStore := quick_store.New[int](20)

	// num := rand.Intn(20)

	// for index := 0; index < num; index++ {
	// 	liteStore.Append(index)
	// }

	// fmt.Println(num)
	// liteStore.Delete()
	// fmt.Println(liteStore.Len())
	// fmt.Println(liteStore.GetStore())

	// buff := smart_buffer.New(10)

	// buff.Buff([]byte("ja"))
	// fmt.Println(string(buff.GetStore()))

	// buff.Buff([]byte("jankoe"))
	// fmt.Println(string(buff.GetStore()))

	// buff.Buff([]byte("janko"))
	// fmt.Println(string(buff.GetStore()))

	// mem := mem_table.NewMemTable(4000)

	// mem.InsertData("15", []byte(`{"id": 123,"ime": "Branko","prezime": "Doe","dob": 30,"email": "john.doe@example.com","adresa": {"ulica": "123 Main Street", "grad": "Cityville", "država": "State", "poštanski_broj": "12345"},"telefoni": [{"tip": "mobilni","broj": "555-1234"},{"tip": "fiksni","broj": "555-5678"}]}\n`))
	// mem.InsertData("18", []byte(`{"id": 123,"ime": "Janko","prezime": "Doe","dob": 30,"email": "john.doe@example.com","adresa": {"ulica": "123 Main Street", "grad": "Cityville", "država": "State", "poštanski_broj": "12345"},"telefoni": [{"tip": "mobilni","broj": "555-1234"},{"tip": "fiksni","broj": "555-5678"}]}\n`))

	// mem.Flush()

	// fmt.Println(string(mem.GetSorted().GetStore()))
	key := 0
	for i := 0; i < 100; i++ {
		key = (key + 1) % 3
		fmt.Println(key)
	}
}
