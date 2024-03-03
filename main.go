package main

import (
	"fmt"
	"math/rand"
	m "root/mem_table"
)

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

func main() {
	mem := m.NewMemTable(40000)
	generatedID := make([]string, 40000)

	for i := 0; i < 40000; i++ {
		generatedID[i] = fmt.Sprintf("%d", rand.Intn(500_000))
	}

	for i := 0; i < 40000; i++ {
		mem.InsertData(generatedID[i], TEST_INPUT)
	}

	for _, value := range mem.GetTreeInOrderTraversal() {
		fmt.Println(value)
	}
}
