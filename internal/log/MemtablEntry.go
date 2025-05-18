package log

import "time"

type MemtableEntry struct {
	Key       string
	Value     *string
	Timestamp time.Time
	Deleted bool
}
