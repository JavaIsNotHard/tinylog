package log

var (
	maximum_write_buffer = 1024
)

type Memtable struct {
	entries *skiplist
	size    uint64
}

func NewMemtable() (*Memtable, error) {
	return &Memtable{
	}, nil
}

// parameters -> MemtableEntry
// return value -> size of the inserted data and error
func (m *Memtable) Insert(MemtableEntry) (uint64, error) {return 0, nil}

// parameter -> key, type: string
// return value -> reference to the MemtableEntry or error
func (m* Memtable) Get(key string) (*MemtableEntry, error) {return nil, nil}

func (m* Memtable) Delete(key string) error {return nil}
