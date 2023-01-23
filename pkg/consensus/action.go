package consensus

type ActionName string

func (a ActionName) String() string {
	return string(a)
}

type Action struct {
	Name ActionName
	Data interface{}
}

const (
	TIMEOUT ActionName = "Timeout"
)
