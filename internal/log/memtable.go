package log

type Memtable struct {
	entries MemtableEntry
	size    uint64
}

func NewMemtable() (*Memtable, error) {
	return &Memtable{
	}, nil
}

func (m *Memtable) Append(MemtableEntry) {}
