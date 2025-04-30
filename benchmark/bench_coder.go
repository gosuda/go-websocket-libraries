package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	coderWebsocket "github.com/coder/websocket"
)

// --- coder/websocket Benchmark ---

func coderEchoHandler(w http.ResponseWriter, r *http.Request) {
	// Use Accept directly as per coder/websocket documentation
	c, err := coderWebsocket.Accept(w, r, &coderWebsocket.AcceptOptions{
		// Add options if needed, e.g., Subprotocols, CompressionMode
		// InsecureSkipVerify: true, // Allow cross-origin if necessary for benchmark client
	})
	if err != nil {
		log.Printf("coder/websocket Accept error: %v", err)
		return
	}
	// CloseNow() is recommended by docs for simpler closing than Close() handshake
	defer c.CloseNow()

	// Use a background context for reads/writes within the handler
	ctx := context.Background() // Or r.Context() if appropriate, but docs advise against it after hijack

	for {
		// Use Read method
		msgType, msg, err := c.Read(ctx)
		if err != nil {
			// Check for normal closure
			var ce coderWebsocket.CloseError
			if errors.As(err, &ce) {
				if ce.Code == coderWebsocket.StatusNormalClosure || ce.Code == coderWebsocket.StatusGoingAway {
					// log.Printf("coder/websocket client closed normally: %v", err)
					break
				}
			}
			log.Printf("coder/websocket read error: %v", err)
			break
		}

		// Use Write method
		err = c.Write(ctx, msgType, msg)
		if err != nil {
			log.Printf("coder/websocket write error: %v", err)
			break
		}
	}
}

func startCoderWebsocketEchoServer(addr string, wg *sync.WaitGroup, serverErr chan<- error) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo_coder", coderEchoHandler) // Use a unique path
	server := &http.Server{Addr: addr, Handler: mux}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Starting coder/websocket echo server on %s", addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("coder/websocket server error: %v", err)
			serverErr <- err
		} else {
			log.Printf("coder/websocket server on %s stopped gracefully.", addr)
		}
	}()
	time.Sleep(100 * time.Millisecond) // Allow server startup
	return server
}

// runCoderWebsocketBenchmark starts a server and runs the benchmark client.
func runCoderWebsocketBenchmark(lib Library) (BenchmarkResult, error) {
	addr := "localhost:8085" // Use a different port
	var wg sync.WaitGroup
	serverErr := make(chan error, 1)

	server := startCoderWebsocketEchoServer(addr, &wg, serverErr)
	if server == nil { // Basic check if startup failed
		select {
		case err := <-serverErr:
			return BenchmarkResult{}, fmt.Errorf("failed to start server: %w", err)
		default:
			// This case might happen if start function returns nil without sending error
			return BenchmarkResult{}, fmt.Errorf("failed to start server for unknown reason")
		}
	}

	defer func() {
		log.Printf("Shutting down coder/websocket server on %s...", addr)
		// Use context for graceful shutdown
		ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()
		if err := server.Shutdown(ctxShutdown); err != nil {
			log.Printf("Error shutting down coder/websocket server: %v", err)
		}
		wg.Wait()
		log.Printf("coder/websocket server on %s shut down.", addr)
	}()

	// --- Run Benchmark Clients ---
	targetURL := "ws://" + addr + "/echo_coder" // Match the handler path
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
