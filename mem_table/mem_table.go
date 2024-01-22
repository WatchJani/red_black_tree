package mem_table

import (
	"root/red_black_tree"
	"root/smart_buffer"
)

type Saver struct {
	start int
	end   int
}

type MemTable struct {
	red_black_tree.RBTree[string, Saver]
	smart_buffer.SBuff
}
