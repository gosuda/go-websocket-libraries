package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/lesismal/nbio/nbhttp"
	nbioWebsocket "github.com/lesismal/nbio/nbhttp/websocket"
)

// --- lesismal/nbio Benchmark ---

func newNbioWebsocketUpgrader() *nbioWebsocket.Upgrader {
	u := nbioWebsocket.NewUpgrader()
	u.OnMessage(func(c *nbioWebsocket.Conn, messageType nbioWebsocket.MessageType, data []byte) {
		// Echo
		err := c.WriteMessage(messageType, data)
		if err != nil {
			log.Printf("nbio write error: %v", err)
			c.Close() // Close connection on write error
		}
	})
	u.OnClose(func(c *nbioWebsocket.Conn, err error) {
		// Log unexpected errors
		if err != nil && !errors.Is(err, net.ErrClosed) && !strings.Contains(err.Error(), "closed") { // More generic check for closed errors
			log.Printf("nbio connection closed with error: %v", err)
		}
	})
	// Add OnOpen, OnPing, OnPong handlers if needed for more complex benchmarks
	return u
}

func startNbioEchoServer(addr string, serverErr chan<- error) (*nbhttp.Engine, error) { // Removed wg, added error return
	upgrader := newNbioWebsocketUpgrader()

	mux := &http.ServeMux{}
	mux.HandleFunc("/echo_nbio", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("nbio upgrade error: %v", err)
			return
		}
		// nbio handles the connection loop internally based on the upgrader's event handlers
		_ = conn // conn variable might be unused directly here, depending on setup
		// log.Printf("nbio client connected: %s", conn.RemoteAddr().String())
	})

	engine := nbhttp.NewEngine(nbhttp.Config{
		Network: "tcp",
		Addrs:   []string{addr},
		Handler: mux,
	})

	log.Printf("Starting nbio echo server on %s", addr)
	err := engine.Start() // Call Start directly
	if err != nil {
		log.Printf("nbio engine failed to start: %v", err)
		// Send error non-blockingly in case receiver isn't ready or benchmark exits early
		select {
		case serverErr <- err:
		default:
		}
		return nil, err // Return error immediately
	}
	// engine.Start() is non-blocking for nbio, server runs in background threads managed by the engine.
	log.Printf("nbio engine started successfully on %s", addr)
	return engine, nil // Return engine and nil error
}

// runNbioBenchmark starts an nbio server and runs the benchmark client.
func runNbioBenchmark(lib Library) (BenchmarkResult, error) {
	addr := "localhost:8084"         // Use a different port
	serverErr := make(chan error, 1) // Keep channel for potential async errors if needed later, though Start() error is now direct

	engine, startErr := startNbioEchoServer(addr, serverErr) // Call updated function
	if startErr != nil {
		return BenchmarkResult{}, fmt.Errorf("failed to start nbio server: %w", startErr)
	}
	if engine == nil {
		// Should be caught by startErr, but double-check
		return BenchmarkResult{}, fmt.Errorf("failed to start nbio server for unknown reason (nil engine)")
	}

	defer func() {
		log.Printf("Stopping nbio server on %s...", addr)
		engine.Stop() // Use nbio specific stop method
		// No wg.Wait() needed anymore
		log.Printf("nbio server on %s stopped.", addr)
	}()

	// --- Run Benchmark Clients ---
	targetURL := "ws://" + addr + "/echo_nbio" // Match the handler path
	results, clientErr := runBenchmarkScenario(targetURL)
	if clientErr != nil {
		log.Printf("WARN: Client benchmark scenario for %s finished with errors: %v", lib.Name, clientErr)
	}
	// --- End Benchmark Clients ---

	// Check if server encountered an error
	select {
	case err := <-serverErr:
		// Check if the error is just the server stopping normally (might not be an error type)
		// This depends on how nbio signals normal shutdown via Start() return.
		// Assuming for now any error reported here is problematic for the benchmark.
		return BenchmarkResult{}, fmt.Errorf("server error during benchmark: %w", err)
	default:
	}

	// --- Calculate Metrics ---
	var avgLatencyMs float64
	if results.totalMessages > 0 {
		avgLatencyMs = float64(results.totalLatency.Milliseconds()) / float64(results.totalMessages)
	}
	var throughputMBps float64
	if results.totalDuration > 0 {
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
