package model

type CandidateID string

type LogEntry struct {
	Term  uint
	Data  []byte
	Index uint
}
