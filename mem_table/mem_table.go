package mem_table

import (
	"log"
	"math/rand"
	"os"
	r "root/red_black_tree"
	s "root/smart_buffer"
)

// 8_388_608
const (
	MB_8    int = 8 * 1024 * 1024
	segment int = 16 * 1024
) // for testing

type Saver struct {
	start int
	end   int
}

func NewSaver(start, end int) Saver {
	return Saver{
		start: start,
		end:   end,
	}
}

type MemTable struct {
	tree r.RBTree[string, Saver]

	sorted   s.SBuff //GreenBuffer
	unsorted s.SBuff //[]GreenBuffer

	stream chan []byte
}

// capacity need be calculated, from heigh level func
func NewMemTable(capacity int) *MemTable {
	return &MemTable{
		tree:     r.NewRBTree[string, Saver](capacity),
		sorted:   s.New(MB_8),
		unsorted: s.New(MB_8),
		stream:   make(chan []byte),
	}
}

// 500ns per node for 40_000 nodes
func (m *MemTable) InsertData(id string, data []byte) {
	saver := NewSaver(m.unsorted.Len(), m.unsorted.Len()+len(data))

	if err := m.unsorted.Buff(data); err != nil {
		m.SortBuf()

		m.SegmentFile()

		m.sorted.Reset()
		m.unsorted.Reset()
		m.tree.Reset()
	} //data in the unsorted buffer

	m.tree.Insert(id, saver) //insert to Red Black Tree
}

func (m *MemTable) SortBuf() {
	for _, value := range m.tree.InOrderTraversal() {
		m.sorted.Buff(m.unsorted.GetStore()[value.start:value.end])
	}

	m.tree.Reset()
}

func (m MemTable) GetTreeInOrderTraversal() []Saver {
	return m.tree.InOrderTraversal()
}

func (m *MemTable) SegmentFile() {
	go func() {
		for {
			data := <-m.stream

			for index := 0; index < len(data); index += segment {
				go Cutter(data[index : index+segment])
			}
		}
	}()

	m.stream <- m.sorted.GetStore()
}

func Cutter(data []byte) {
	if err := os.WriteFile(NameGenerator("./../test/", ".bin", 10), data, 0766); err != nil {
		log.Fatalln(err)
	}
}

func NameGenerator(path, extension string, size int) string {
	name := make([]byte, size)

	for index := range name {
		name[index] = byte(rand.Intn('z'-'a') + 'a')
	}

	return path + string(name) + extension
}
