package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// --- gobwas/ws Benchmark ---

func gobwasEchoHandler(conn net.Conn) {
	defer conn.Close()
	// No explicit upgrade needed here if the server handles the handshake

	// Use Reader and Writer for message handling
	reader := wsutil.NewReader(conn, ws.StateServerSide)
	writer := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText) // Assuming text messages

	for {
		header, err := reader.NextFrame()
		if err != nil {
			if err != io.EOF && !errors.Is(err, net.ErrClosed) && !strings.Contains(err.Error(), "connection reset by peer") {
				log.Printf("gobwas read header error: %v", err)
			}
			break
		}

		if header.OpCode == ws.OpClose {
			// Handle close frame
			// log.Printf("gobwas client sent close frame")
			break
		}

		// Read the payload
		payload := make([]byte, header.Length)
		_, err = io.ReadFull(reader, payload)
		if err != nil {
			log.Printf("gobwas read payload error: %v", err)
			break
		}

		// Echo back using the writer
		writer.Reset(conn, ws.StateServerSide, header.OpCode) // Reset writer for the correct opcode
		_, err = writer.Write(payload)
		if err != nil {
			log.Printf("gobwas write error: %v", err)
			break
		}
		if err = writer.Flush(); err != nil {
			log.Printf("gobwas flush error: %v", err)
			break
		}
	}
}

func startGobwasEchoServer(addr string, wg *sync.WaitGroup, serverErr chan<- error) *http.Server {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		// Send error instead of Fatalf to allow benchmark function to handle it
		serverErr <- fmt.Errorf("failed to start gobwas listener on %s: %w", addr, err)
		return nil // Return nil if listener fails
	}

	upgrader := ws.HTTPUpgrader{} // Use standard HTTP upgrader

	http.HandleFunc("/echo_gobwas", func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := upgrader.Upgrade(r, w)
		if err != nil {
			log.Printf("gobwas upgrade error: %v", err)
			return
		}
		// Hand over the raw net.Conn to the handler
		go gobwasEchoHandler(conn)
	})

	// Need an HTTP server to handle the upgrade request
	httpServer := &http.Server{Addr: addr}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Starting gobwas HTTP server for upgrade on %s", addr)
		// Use the listener with the HTTP server
		if err := httpServer.Serve(listener); err != http.ErrServerClosed {
			log.Printf("gobwas HTTP server error: %v", err)
			serverErr <- err
		} else {
			log.Printf("gobwas HTTP server on %s stopped gracefully.", addr)
		}
	}()

	// Give server a moment to start
	time.Sleep(100 * time.Millisecond)
	// Return the httpServer so it can be shut down
	return httpServer
}

// runGobwasBenchmark starts a gobwas server and runs the benchmark client.
func runGobwasBenchmark(lib Library) (BenchmarkResult, error) {
	addr := "localhost:8083" // Use a different port
	var wg sync.WaitGroup
	serverErr := make(chan error, 1)

	httpServer := startGobwasEchoServer(addr, &wg, serverErr)
	// Check if server startup failed
	if httpServer == nil {
		// Error should already be in serverErr channel from startGobwasEchoServer
		err := <-serverErr
		return BenchmarkResult{}, fmt.Errorf("failed to start server: %w", err)
	}

	defer func() {
		log.Printf("Shutting down gobwas server on %s...", addr)
		// Use context for graceful shutdown
		ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()
		if err := httpServer.Shutdown(ctxShutdown); err != nil {
			log.Printf("Error shutting down gobwas HTTP server: %v", err)
		}
		wg.Wait() // Wait for server goroutine to finish
		log.Printf("gobwas server on %s shut down.", addr)
	}()

	// --- Run Benchmark Clients ---
	targetURL := "ws://" + addr + "/echo_gobwas" // Match the handler path
	results, clientErr := runBenchmarkScenario(targetURL)
	if clientErr != nil {
		log.Printf("WARN: Client benchmark scenario for %s finished with errors: %v", lib.Name, clientErr)
	}
	// --- End Benchmark Clients ---

	// Check if server encountered an error
	select {
	case err := <-serverErr:
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
