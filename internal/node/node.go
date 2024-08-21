package node

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vivekpal1/consensus-simulator-bft/internal/types"
)

func NewNode(id int, isByzantine bool) *types.Node {
	return &types.Node{
		ID:           id,
		IsByzantine:  isByzantine,
		IsCrashed:    false,
		Inbox:        make(chan types.Message, 100),
		CurrentView:  0,
		CurrentPhase: types.Propose,
		Votes:        make(map[int]int),
	}
}

func Run(n *types.Node, consensusChan chan<- time.Time) {
	for {
		select {
		case msg := <-n.Inbox:
			if n.IsCrashed {
				continue
			}
			HandleMessage(n, msg, consensusChan)
		}
	}
}

func HandleMessage(n *types.Node, msg types.Message, consensusChan chan<- time.Time) {
	switch msg.Type {
	case types.Propose:
		if n.CurrentPhase == types.Propose {
			n.Proposal = msg.Proposal
			n.CurrentPhase = types.Vote
			Broadcast(n, types.Message{Type: types.Vote, From: n.ID, Proposal: n.Proposal})
		}
	case types.Vote:
		if n.CurrentPhase == types.Vote {
			n.Votes[msg.Proposal]++
			if n.Votes[msg.Proposal] >= 2*len(n.Peers)/3 {
				n.CurrentPhase = types.Commit
				Broadcast(n, types.Message{Type: types.Commit, From: n.ID, Proposal: msg.Proposal})
			}
		}
	case types.Commit:
		if n.CurrentPhase == types.Vote || n.CurrentPhase == types.Commit {
			n.CurrentPhase = types.Propose
			n.CurrentView++
			n.Votes = make(map[int]int)
			consensusChan <- time.Now()
			fmt.Printf("Node %d: Consensus reached for value %d in view %d\n", n.ID, msg.Proposal, n.CurrentView-1)
		}
	}
}

func Broadcast(n *types.Node, msg types.Message) {
	if n.IsByzantine {
		for _, peer := range n.Peers {
			if rand.Float32() < 0.5 {
				peer.Inbox <- msg
			} else {
				peer.Inbox <- types.Message{Type: msg.Type, From: n.ID, Proposal: rand.Int()}
			}
		}
	} else {
		for _, peer := range n.Peers {
			peer.Inbox <- msg
		}
	}
}
