package mem_table

import (
	r "root/red_black_tree"
	s "root/smart_buffer"
)

const MB_8 int = 8_388_608

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
	tree     r.RBTree[string, Saver]
	sorted   s.SBuff
	unsorted s.SBuff
}

// capacity need be calculated
func NewMemTable(capacity int) *MemTable {
	return &MemTable{
		tree:     r.NewRBTree[string, Saver](capacity),
		sorted:   s.New(MB_8),
		unsorted: s.New(MB_8),
	}
}

func (m *MemTable) InsertData(id string, data []byte) {
	m.tree.Insert(id, NewSaver(m.unsorted.Len(), len(data))) //insert to Red Black Tree
	m.unsorted.Buff(data)                                    //data in the buffer
}
