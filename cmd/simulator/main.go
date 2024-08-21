package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/vivekpal1/consensus-simulator-bft/internal/network"
)

func main() {
	nodeCount := flag.Int("nodes", 10, "Number of nodes in the network")
	byzantineCount := flag.Int("byzantine", 3, "Number of Byzantine nodes")
	rounds := flag.Int("rounds", 5, "Number of consensus rounds to simulate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	net := network.NewNetwork(*nodeCount, *byzantineCount)
	results := net.Run(*rounds)

	fmt.Printf("Simulation completed.\n")
	fmt.Printf("Total rounds: %d\n", *rounds)
	fmt.Printf("Successful consensus rounds: %d\n", results.SuccessfulRounds)
	fmt.Printf("Average time to consensus: %v\n", results.AverageConsensusTime)
}
