package main

type Zero struct {
	pointer  int
	capacity int
}

func New(capacity int) *Zero {
	return &Zero{
		capacity: capacity,
	}
}

func (z *Zero) Insert() {
	z.pointer = (z.pointer + 1) % z.capacity //jako spora operacija, bolje sa if
}

func (z *Zero) InsertIf() {
	if z.pointer+1 > z.capacity {
		z.pointer = 0
		return //obavezno return za bolje performance
	}

	z.pointer++
}
