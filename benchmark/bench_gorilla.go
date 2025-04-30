package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  8192, // Increased buffer size
	WriteBufferSize: 8192, // Increased buffer size
	CheckOrigin: func(r *http.Request) bool { // Allow all connections for benchmark
		return true
	},
}

func gorillaEchoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Gorilla upgrade error: %v", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Gorilla read error: %v", err)
			}
			break // Exit loop on error or close
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Gorilla write error: %v", err)
			break
		}
	}
}

func startGorillaEchoServer(addr string, wg *sync.WaitGroup, serverErr chan<- error) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", gorillaEchoHandler)
	server := &http.Server{Addr: addr, Handler: mux}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Starting Gorilla echo server on %s", addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Gorilla server error: %v", err)
			serverErr <- err // Report error
		} else {
			log.Printf("Gorilla server on %s stopped gracefully.", addr)
		}
	}()
	// Give server a moment to start - replace with more robust check if needed
	time.Sleep(100 * time.Millisecond)
	return server
}

// runGorillaBenchmark starts a server and runs the benchmark client.
func runGorillaBenchmark(lib Library) (BenchmarkResult, error) {
	addr := "localhost:8081" // Use a specific port for this benchmark
	var wg sync.WaitGroup
	serverErr := make(chan error, 1)

	server := startGorillaEchoServer(addr, &wg, serverErr)
	defer func() {
		log.Printf("Shutting down Gorilla server on %s...", addr)
		// Use context for graceful shutdown if needed, simple close for now
		if err := server.Close(); err != nil {
			log.Printf("Error shutting down Gorilla server: %v", err)
		}
		wg.Wait() // Wait for server goroutine to finish
		log.Printf("Gorilla server on %s shut down.", addr)
	}()

	// --- Run Benchmark Clients ---
	targetURL := "ws://" + addr + "/echo" // Adjust path if needed
	results, clientErr := runBenchmarkScenario(targetURL)
	if clientErr != nil {
		// Log client-side errors but potentially still return server metrics if server was ok
		log.Printf("WARN: Client benchmark scenario for %s finished with errors: %v", lib.Name, clientErr)
		// Decide if this should be a fatal error for the benchmark result
		// For now, we'll continue and report potentially zero/low metrics
	}
	// --- End Benchmark Clients ---

	// Check if server encountered an error during benchmark run
	select {
	case err := <-serverErr:
		return BenchmarkResult{}, fmt.Errorf("server error during benchmark: %w", err)
	default:
		// No server error
	}

	// --- Calculate Metrics ---
	var avgLatencyMs float64
	if results.totalMessages > 0 {
		avgLatencyMs = float64(results.totalLatency.Milliseconds()) / float64(results.totalMessages)
	}

	var throughputMBps float64
	if results.totalDuration > 0 {
		// Total bytes = messages * size * 2 (send + receive)
		totalBytes := float64(results.totalMessages * messageSizeBytes * 2)
		throughputMBps = (totalBytes / (1024 * 1024)) / results.totalDuration.Seconds()
	}
	// --- End Calculate Metrics ---

	result := BenchmarkResult{
		LibraryName:    lib.Name,
		ThroughputMBps: throughputMBps,
		LatencyMs:      avgLatencyMs,
		Timestamp:      time.Now().UTC(),
		Version:        lib.Version,
	}
	return result, nil
}
