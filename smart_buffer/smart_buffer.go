package smart_buffer

import (
	"errors"
	q "root/quick_store"
)

type SBuff struct {
	q.QuickStore[byte]
}

func New(capacity int) SBuff {
	return SBuff{
		QuickStore: q.New[byte](capacity),
	}
}

func (s *SBuff) Buff(data []byte) error {
	if s.Len()+len(data) > s.Cap() {
		return errors.New("buffer is full")
		// s.Reset()
	}

	copy(s.GetStoreAll()[s.Len():], data)
	s.SetPointer(len(data))

	return nil
}
