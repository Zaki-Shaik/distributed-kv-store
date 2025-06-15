# Distributed Key-Value Store (Go)

This is a basic distributed key-value store built in Go.  
The first iteration runs as a single-node in-memory store, designed with concurrency safety and extensibility in mind.

## Features

- ✅ In-memory key-value storage  
- ✅ Thread-safe reads and writes using `sync.RWMutex`  
- 🔜 Pluggable consensus layer (Raft via HashiCorp)  
- 🔜 Multi-node support  
- 🔜 gRPC-based communication between nodes

## Getting Started

### Prerequisites

- Go 1.20+ installed
- Git

### Running the Server (Single Node)

```bash
go run main.go
