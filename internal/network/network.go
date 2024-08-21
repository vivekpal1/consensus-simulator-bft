package network

import (
	"math/rand"
	"sync"
	"time"

	"github.com/vivekpal1/consensus-simulator-bft/internal/node"
	"github.com/vivekpal1/consensus-simulator-bft/internal/types"
)

type Network struct {
	Nodes []*types.Node
}

func NewNetwork(nodeCount, byzantineCount int) *Network {
	network := &Network{
		Nodes: make([]*types.Node, nodeCount),
	}

	for i := 0; i < nodeCount; i++ {
		isByzantine := i < byzantineCount
		network.Nodes[i] = node.NewNode(i, isByzantine)
	}

	for _, n := range network.Nodes {
		for _, peer := range network.Nodes {
			if n != peer {
				n.Peers = append(n.Peers, peer)
			}
		}
	}

	return network
}

func (n *Network) SimulateLatency() {
	for _, node := range n.Nodes {
		go func(node *types.Node) {
			for msg := range node.Inbox {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				node.Inbox <- msg
			}
		}(node)
	}
}

func (n *Network) Run(rounds int) types.SimulationResults {
	consensusChan := make(chan time.Time, rounds)
	var wg sync.WaitGroup

	for _, node := range n.Nodes {
		wg.Add(1)
		go func(node *types.Node) {
			defer wg.Done()
			node.Run(consensusChan)
		}(node)
	}

	n.SimulateLatency()

	startTime := time.Now()
	for i := 0; i < rounds; i++ {
		leader := n.Nodes[i%len(n.Nodes)]
		if !leader.IsCrashed {
			proposal := rand.Int()
			node.Broadcast(leader, types.Message{Type: types.Propose, From: leader.ID, Proposal: proposal})
		}
		time.Sleep(1 * time.Second)
	}

	close(consensusChan)
	wg.Wait()

	successfulRounds := 0
	var totalTime time.Duration
	for t := range consensusChan {
		successfulRounds++
		totalTime += t.Sub(startTime)
		startTime = t
	}

	return types.SimulationResults{
		SuccessfulRounds:     successfulRounds,
		AverageConsensusTime: totalTime / time.Duration(successfulRounds),
	}
}
