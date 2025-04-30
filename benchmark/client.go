package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

// clientResult holds results from a single client goroutine
type clientResult struct {
	messagesSent int
	bytesSent    int64
	bytesRecv    int64
	latencies    []time.Duration
	err          error
}

// aggregatedResult holds combined results from all clients
type aggregatedResult struct {
	totalMessages int64
	totalLatency  time.Duration
	totalDuration time.Duration // Actual duration benchmark was intended to run
	runDuration   time.Duration // Actual duration clients were running before stopping
	clientErrors  []error
}

// runBenchmarkScenario runs the echo benchmark with concurrent clients for a specific duration.
func runBenchmarkScenario(targetURL string) (aggregatedResult, error) {
	log.Printf("Starting benchmark scenario for %s with %d clients for %v...", targetURL, numClients, benchmarkDuration)
	var wg sync.WaitGroup
	resultsChan := make(chan clientResult, numClients)
	startBarrier := make(chan struct{}) // To synchronize client start times
	stopSignal := make(chan struct{})   // To signal clients to stop

	scenarioStartTime := time.Now()

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			// Wait for start signal
			<-startBarrier
			// Pass stopSignal to the client
			res := runEchoBenchmarkClient(clientID, targetURL, stopSignal)
			resultsChan <- res
		}(i)
	}

	// Give goroutines a moment to set up before starting
	time.Sleep(50 * time.Millisecond)
	// Add 1-second delay before starting the benchmark timer
	log.Println("Waiting 1 second before starting benchmark timer...")
	time.Sleep(1 * time.Second)
	clientStartTime := time.Now() // Record when clients actually start
	close(startBarrier)           // Signal clients to start

	// Let benchmark run for the specified duration
	time.Sleep(benchmarkDuration)
	close(stopSignal) // Signal clients to stop

	// Wait for all clients to finish gracefully (with a timeout)
	waitTimeout := time.After(10 * time.Second) // Add timeout for waiting
	waitDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(waitDone)
	}()

	select {
	case <-waitDone:
		// All clients finished
	case <-waitTimeout:
		log.Printf("WARN: Timeout waiting for clients to finish for %s", targetURL)
		// Proceed with aggregation, but results might be incomplete
	}
	close(resultsChan) // Close channel after waiting

	actualRunDuration := time.Since(clientStartTime) // Duration clients were actually running
	totalScenarioDuration := time.Since(scenarioStartTime)
	log.Printf("Benchmark scenario for %s finished. Client run duration: %v, Total duration: %v.", targetURL, actualRunDuration, totalScenarioDuration)

	// Aggregate results
	aggRes := aggregatedResult{
		totalDuration: benchmarkDuration, // Use configured duration for calculations
		runDuration:   actualRunDuration, // Store actual run duration
	}
	var totalLatencySum time.Duration
	var totalMessagesCount int64

	for res := range resultsChan {
		if res.err != nil {
			aggRes.clientErrors = append(aggRes.clientErrors, res.err)
			// Optionally skip metrics from failed clients
			continue
		}
		atomic.AddInt64(&totalMessagesCount, int64(res.messagesSent))
		for _, lat := range res.latencies {
			totalLatencySum += lat
		}
	}

	aggRes.totalMessages = totalMessagesCount
	aggRes.totalLatency = totalLatencySum

	if len(aggRes.clientErrors) > 0 {
		// Return a summary error if any client failed
		return aggRes, fmt.Errorf("%d client(s) failed during benchmark", len(aggRes.clientErrors))
	}

	return aggRes, nil
}

// runEchoBenchmarkClient simulates a single client performing echo requests until stopSignal is closed.
func runEchoBenchmarkClient(clientID int, targetURL string, stopSignal <-chan struct{}) clientResult {
	dialer := websocket.Dialer{
		HandshakeTimeout: clientConnectTimeout,
	}
	conn, _, err := dialer.Dial(targetURL, nil)
	if err != nil {
		return clientResult{err: fmt.Errorf("client %d: dial error: %w", clientID, err)}
	}
	defer conn.Close()

	message := make([]byte, messageSizeBytes)
	latencies := make([]time.Duration, 0) // Removed messagesPerClient
	var bytesSent, bytesRecv int64
	var messagesSent int

	// Set read/write deadlines
	conn.SetWriteDeadline(time.Now().Add(clientWriteTimeout)) // Initial deadline

	// Loop until stopSignal is closed
	for {
		select {
		case <-stopSignal:
			// Stop signal received, exit loop gracefully
			// log.Printf("Client %d received stop signal.", clientID) // Optional logging
			return clientResult{ // Return current results
				messagesSent: messagesSent,
				bytesSent:    bytesSent,
				bytesRecv:    bytesRecv,
				latencies:    latencies,
				err:          nil,
			}
		default:
			// Continue sending/receiving
		}

		sendTime := time.Now()
		// Reset write deadline before each write
		conn.SetWriteDeadline(time.Now().Add(clientWriteTimeout))
		err = conn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			// Check if error is due to stop signal causing connection closure
			select {
			case <-stopSignal: // If stop signal was just received, this might be expected
				return clientResult{messagesSent: messagesSent, bytesSent: bytesSent, bytesRecv: bytesRecv, latencies: latencies, err: nil} // Return success on graceful stop
			default:
				// Otherwise, it's a real write error
				return clientResult{messagesSent: messagesSent, bytesSent: bytesSent, bytesRecv: bytesRecv, latencies: latencies, err: fmt.Errorf("client %d: write error: %w", clientID, err)}
			}
		}
		bytesSent += int64(len(message))

		// Reset read deadline before each read
		conn.SetReadDeadline(time.Now().Add(clientReadTimeout))

		_, recvMsg, err := conn.ReadMessage()
		if err != nil {
			// Check if error is due to stop signal causing connection closure
			select {
			case <-stopSignal: // If stop signal was just received, this might be expected
				return clientResult{messagesSent: messagesSent, bytesSent: bytesSent, bytesRecv: bytesRecv, latencies: latencies, err: nil} // Return success on graceful stop
			default:
				// Otherwise, it's a real read error
				return clientResult{messagesSent: messagesSent, bytesSent: bytesSent, bytesRecv: bytesRecv, latencies: latencies, err: fmt.Errorf("client %d: read error: %w", clientID, err)}
			}
		}
		recvTime := time.Now()
		bytesRecv += int64(len(recvMsg))

		// Basic validation (optional)
		if len(recvMsg) != len(message) {
			log.Printf("WARN: Client %d received message size mismatch (%d != %d)", clientID, len(recvMsg), len(message))
		}

		latencies = append(latencies, recvTime.Sub(sendTime))
		messagesSent++

		// No need to reset write deadline here, it's done at the start of the loop
	}
	// This part is now unreachable due to the infinite loop and return inside select
}
