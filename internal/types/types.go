package types

import "time"

type MessageType int

const (
	Propose MessageType = iota
	Vote
	Commit
)

type Message struct {
	Type     MessageType
	From     int
	Proposal int
}

type SimulationResults struct {
	SuccessfulRounds     int
	AverageConsensusTime time.Duration
}

type Node struct {
	ID           int
	IsByzantine  bool
	IsCrashed    bool
	Peers        []*Node
	Inbox        chan Message
	CurrentView  int
	CurrentPhase MessageType
	Proposal     int
	Votes        map[int]int
}
