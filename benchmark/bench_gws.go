package main

import (
	"errors"
	"fmt" // Ensure fmt is imported
	"log"
	"net" // Added for net.Listen and net.Listener
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/lxzan/gws"
)

// --- gws Benchmark ---

type GwsEventHandler struct {
	gws.BuiltinEventHandler
}

func (h *GwsEventHandler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(payload) // Standard pong response
}

func (h *GwsEventHandler) OnMessage(socket *gws.Conn, message *gws.Message) {
	// Echo the message back
	defer message.Close() // Important to release buffer
	err := socket.WriteMessage(message.Opcode, message.Bytes())
	if err != nil {
		log.Printf("gws write error: %v", err)
	}
}

func (h *GwsEventHandler) OnClose(socket *gws.Conn, err error) {
	// gws often wraps errors, check common close errors
	if err != nil && !errors.Is(err, gws.ErrConnClosed) && !strings.Contains(err.Error(), "connection reset by peer") && !strings.Contains(err.Error(), "broken pipe") {
		log.Printf("gws client closed with error: %v", err)
	} else {
		// log.Printf("gws client disconnected: %s", socket.RemoteAddr().String()) // Optional: log normal disconnects
	}
}

// startGwsEchoServer creates a listener and starts the server with RunListener.
// It returns the server instance and the listener.
func startGwsEchoServer(addr string, wg *sync.WaitGroup, serverErr chan<- error) (*gws.Server, net.Listener) {
	handler := new(GwsEventHandler)
	server := gws.NewServer(handler, &gws.ServerOption{
		ReadBufferSize:    8192, // Increased buffer size
		Recovery:          gws.Recovery,
		PermessageDeflate: gws.PermessageDeflate{Enabled: false},
	})
	server.OnError = func(conn net.Conn, err error) {
		if err != err {
			runtime.Goexit()
		}
	}

	// Create listener manually
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		// Cannot start server if listener fails
		serverErr <- fmt.Errorf("failed to listen on %s: %w", addr, err)
		close(serverErr)   // Close channel as we can't proceed
		return server, nil // Return nil listener - Corrected
	}
	log.Printf("gws server listening on %s", listener.Addr().String())

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Starting gws server with RunListener on %s", listener.Addr().String())
		// Use RunListener
		if err := server.RunListener(listener); err != nil {
			// Check if the error is expected during shutdown (e.g., listener closed)
			// net.ErrClosed is often returned when listener.Close() is called.
			if !errors.Is(err, net.ErrClosed) && !errors.Is(err, http.ErrServerClosed) {
				log.Printf("gws server RunListener error: %v", err)
				// Only send unexpected errors to the channel
				select {
				case serverErr <- err:
				default: // Avoid blocking if channel is full or receiver is gone
				}
			} else {
				// Log expected closure without marking as error
				log.Printf("gws server RunListener stopped (expected: %v)", err)
			}
		} else {
			log.Printf("gws server RunListener stopped gracefully (returned nil).")
		}
	}()
	time.Sleep(100 * time.Millisecond) // Keep sleep for server start
	return server, listener            // Return server and listener - Corrected
}

// runGwsBenchmark starts a gws server and runs the benchmark client.
func runGwsBenchmark(lib Library) (BenchmarkResult, error) {
	addr := "localhost:8082" // Use a different port
	var wg sync.WaitGroup
	serverErr := make(chan error, 1)

	// Start the server and get the listener
	_, listener := startGwsEchoServer(addr, &wg, serverErr) // Capture listener, ignore server instance for now
	if listener == nil {
		// Error occurred during listener creation, read from channel
		err := <-serverErr
		return BenchmarkResult{}, fmt.Errorf("failed to start gws server: %w", err)
	}

	// Defer closing the listener and waiting for the server goroutine
	defer func() {
		log.Printf("Closing gws listener for %s...", addr)
		if err := listener.Close(); err != nil {
			log.Printf("Error closing gws listener: %v", err)
		}
		log.Printf("Waiting for gws server goroutine to finish for %s...", addr)
		// wg.Wait() // TODO: graceful shutdown gws server
		log.Printf("gws server goroutine finished for %s.", addr)
	}()

	// --- Run Benchmark Clients ---
	targetURL := "ws://" + addr
	results, clientErr := runBenchmarkScenario(targetURL)
	if clientErr != nil {
		log.Printf("WARN: Client benchmark scenario for %s finished with errors: %v", lib.Name, clientErr)
		// Decide if client errors should halt the process or just be logged
	}
	// --- End Benchmark Clients ---

	// --- Listener is closed in defer ---

	// Check if server encountered an *unexpected* error before listener was closed
	select {
	case err := <-serverErr:
		// This channel should only receive unexpected errors now
		log.Printf("gws server RunListener returned unexpected error: %v", err)
		// Decide if this error should fail the benchmark result
		// return BenchmarkResult{}, fmt.Errorf("server error during benchmark run: %w", err)
	default:
		// No unexpected error reported by server.RunListener
	}

	// --- Calculate Metrics ---
	// (Keep existing metric calculation logic)
	var avgLatencyMs float64
	if results.totalMessages > 0 {
		avgLatencyMs = float64(results.totalLatency.Milliseconds()) / float64(results.totalMessages)
	}
	var throughputMBps float64
	// Use configured benchmark duration for calculations, not actual run duration
	if results.totalDuration.Seconds() > 0 {
		totalBytes := float64(results.totalMessages * messageSizeBytes * 2) // Factor of 2 for echo (send + receive)
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
	return result, nil // Return nil error if clientErr is just a warning
}
