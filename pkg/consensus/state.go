package consensus

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/revolyssup/bargedb/pkg/model"
)

type PersistentState interface {
	SetCurrentTerm(uint)
	GetCurrentTerm() uint
	SetVotedFor(model.CandidateID)
	GetVotedFor() model.CandidateID
	Close()
}

type persistentState struct {
	f           *os.File
	mx          sync.RWMutex
	currTerm    uint
	CandidateID model.CandidateID
}

// Stores currTerm and votedFor in a file seperated by newline.
// While running, serves from memorys
func NewFileBasedPersistentState(filepath string) (PersistentState, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	// Read the first line to get the current term

	var term uint
	var candidateID model.CandidateID

	data, err := io.ReadAll(f)
	if err != nil {
		f.Close()
		return nil, err
	}

	if len(data) > 0 {
		parts := strings.Split(string(data), "\n")
		if len(parts) >= 1 {
			if t, err := strconv.ParseUint(parts[0], 10, 64); err == nil {
				term = uint(t)
			}
		}
		if len(parts) >= 2 {
			candidateID = model.CandidateID(parts[1])
		}
	}

	return &persistentState{f: f, currTerm: uint(term), CandidateID: candidateID}, nil
}

func (p *persistentState) Close() {
	p.f.Close()
}

func (p *persistentState) SetCurrentTerm(term uint) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.currTerm = term
	p.writeState()
}

func (p *persistentState) GetCurrentTerm() uint {
	p.mx.RLock()
	defer p.mx.RUnlock()
	return p.currTerm
}

func (p *persistentState) SetVotedFor(candidateID model.CandidateID) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.CandidateID = candidateID
	p.writeState()
}

func (p *persistentState) GetVotedFor() model.CandidateID {
	p.mx.RLock()
	defer p.mx.RUnlock()
	return p.CandidateID
}

func (p *persistentState) writeState() {
	if err := p.f.Truncate(0); err != nil {
		fmt.Println("truncate error:", err)
		return
	}
	if _, err := p.f.Seek(0, 0); err != nil {
		fmt.Println("seek error:", err)
		return
	}
	if _, err := fmt.Fprintf(p.f, "%d\n%s\n", p.currTerm, p.CandidateID); err != nil {
		fmt.Println("write error:", err)
	}
	if err := p.f.Sync(); err != nil {
		fmt.Println("sync error:", err)
	}
}
