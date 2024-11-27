package signal

type Signaller struct {
	RunCandidate chan struct{}
	RunFollower  chan struct{}
	RunLeader    chan struct{}
}

func Init() *Signaller {
	return &Signaller{
		RunCandidate: make(chan struct{}),
		RunFollower:  make(chan struct{}),
		RunLeader:    make(chan struct{}),
	}
}

func (s *Signaller) SignalCandidate() {
	s.RunCandidate <- struct{}{}
}

func (s *Signaller) SignalFollower() {
	s.RunFollower <- struct{}{}
}

func (s *Signaller) SignalLeader() {
	s.RunLeader <- struct{}{}
}
