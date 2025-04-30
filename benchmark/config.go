package main

import "time"

// --- Benchmark Configuration ---
const (
	benchmarkDuration    = 10 * time.Second // How long each benchmark runs
	numClients           = 1000             // Number of concurrent clients
	messageSizeBytes     = 4096             // Size of messages to echo (Increased as per user request)
	clientConnectTimeout = 5 * time.Second
	clientWriteTimeout   = 5 * time.Second // Increased timeout
	clientReadTimeout    = 5 * time.Second // Increased timeout
)

// --- Data Structures ---

// Library represents a WebSocket library to benchmark
type Library struct {
	Name    string
	URL     string
	Version string // To be fetched
}

// BenchmarkResult holds the performance metrics for a library
type BenchmarkResult struct {
	LibraryName    string
	ThroughputMBps float64
	LatencyMs      float64
	Timestamp      time.Time
	Version        string
}
