package log

type Instance struct {
	Entries []Entry
}
type Index int
type Entry struct {
	Key   string
	Value []byte
}

func (i *Instance) Write(key string, value interface{}) {}
