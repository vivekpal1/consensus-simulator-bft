# BFT Consensus Simulator

This project implements a simplified Byzantine Fault Tolerant (BFT) consensus algorithm simulator in Go. The simulator models a network of nodes trying to reach consensus on a single value, while tolerating both crash faults and Byzantine faults.

## Requirements

- Go 1.16 or higher

## Building the Project

To build the project, run the following command in the project root directory:

```
go build ./cmd/simulator
```

This will create an executable named `simulator` in the project root.

## Running the Simulator

To run the simulator, use the following command:

```
./simulator -nodes=10 -byzantine=3 -rounds=5
```

You can adjust the following parameters:

- `-nodes`: Number of nodes in the network (default: 10)
- `-byzantine`: Number of Byzantine nodes (default: 3)
- `-rounds`: Number of consensus rounds to simulate (default: 5)
