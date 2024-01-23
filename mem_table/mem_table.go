package mem_table

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"path"
	r "root/red_black_tree"
	s "root/smart_buffer"
)

// 8_388_608
const MB_8 int = 1_388_608

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
	tree   r.RBTree[string, Saver]
	sorted GreenBuffer //GreenBuffer

	unsorted []GreenBuffer //[]GreenBuffer
	pointer  int
	// cap int

	stream chan []byte
}

// Just tell to us if is this buffer ready for use
type GreenBuffer struct {
	s.SBuff
	isFree bool
}

func NewGreenBuffer(capacity int) GreenBuffer {
	return GreenBuffer{
		SBuff:  s.New(capacity),
		isFree: true,
	}
}

// move on next free buffer
func (m *MemTable) Next() {
	m.unsorted[m.pointer].ReverseUsage() //make it lock for use
	m.pointer = (m.pointer + 1) % 2
}

func (g *GreenBuffer) ReverseUsage() {
	g.isFree = !g.isFree
}

// capacity need be calculated, from heigh level func
func NewMemTable(capacity int) *MemTable {
	unsortedBufferCapacity := make([]GreenBuffer, 2)
	for index := range unsortedBufferCapacity {
		unsortedBufferCapacity[index] = NewGreenBuffer(MB_8)
	}

	return &MemTable{
		tree:   r.NewRBTree[string, Saver](capacity),
		sorted: NewGreenBuffer(MB_8),
		// unsorted: s.New(MB_8),
		unsorted: unsortedBufferCapacity,
		stream:   make(chan []byte),
	}
}

// close stream -> use for testing
func (s *MemTable) Close() {
	close(s.stream)
}

// 500ns per node for 40_000 nodes
func (m *MemTable) InsertData(id string, data []byte) {
	m.tree.Insert(id, NewSaver(m.unsorted[m.pointer].Len(), m.unsorted[m.pointer].Len()+len(data))) //insert to Red Black Tree
	m.unsorted[m.pointer].Buff(data)                                                                //data in the buffer
}

// save data in data engine
func (s *MemTable) Save(id string, data []byte) {
	//if buffer is full then, write on disk
	if s.unsorted[s.pointer].Check(len(data)) {
		//sand our buffer to another process to make sorted file and write on disk
		s.stream <- s.unsorted[s.pointer].GetStore() //witch one i send on writing

		//move on another unsorted buffer
		s.Next()
	}

	//if buffer is not full, then write data in buffer
	s.InsertData(id, data)
}

// single writer on disk
func (m *MemTable) Listen() {
	go func() {
		for range m.stream {
			if err := m.Flush(); err != nil {
				log.Println(err)
			}
		}
	}()
}

// sort buffer and reset unsorted buffer
func (m *MemTable) Flush() error {
	//from unsorted make sorted buffer
	for _, data := range m.tree.InOrderTraversal() {
		m.sorted.Buff(m.unsorted[m.pointer].GetStore()[data.start:data.end])
	}

	//make usable again
	m.unsorted[m.pointer].Reset()

	//write on disk and return error
	return Write(m.sorted.GetStore())
}

// save data(sorted buffer) on disk with generated file name
func Write(data []byte) error {
	name, err := GenerateRandomString(10)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join("../test", name), data, 0644)
}

// generate random string for file store name
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

// func (m MemTable) GetUnsorted() s.SBuff {
// 	return m.unsorted
// }

// func (m MemTable) GetSorted() s.SBuff {
// 	return m.sorted
// }
