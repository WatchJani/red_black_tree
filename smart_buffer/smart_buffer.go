package smart_buffer

import q "root/quick_store"

type SBuff struct {
	q.QuickStore[byte]
}

func New(capacity int) *SBuff {
	return &SBuff{
		QuickStore: q.New[byte](capacity), //need be custom because my pointer count from -1
	}
}

func (s *SBuff) Buff(data []byte) {
	if s.Len()+len(data) > s.Cap() {
		//flush to disk
		s.Reset()
	}

	copy(s.GetStoreAll()[s.Len():], data)
	s.SetPointer(len(data))
}
