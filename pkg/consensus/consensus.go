package consensus

type State interface {
	Run(c *Consensus)
}
type Consensus struct {
	RunCandidate chan struct{}
	RunFollower  chan struct{}
	RunLeader    chan struct{}
}
