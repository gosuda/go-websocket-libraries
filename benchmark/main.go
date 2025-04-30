package main

import (
	"fmt"
	"log"
	"time"
	// Add import back
)

// List of libraries from README.md
var libraries = []Library{
	{Name: "gorilla/websocket", URL: "https://github.com/gorilla/websocket"},
	{Name: "lxzan/gws", URL: "https://github.com/lxzan/gws"},
	{Name: "gobwas/ws", URL: "https://github.com/gobwas/ws"},
	{Name: "lesismal/nbio", URL: "https://github.com/lesismal/nbio"},
	{Name: "coder/websocket", URL: "https://github.com/coder/websocket"},
}

func main() {
	log.Println("Starting WebSocket library benchmark...")

	results := []BenchmarkResult{}
	benchmarkTimestamp := time.Now().UTC() // Capture timestamp at the start

	// TODO: Implement benchmarking logic for each library
	// 1. Fetch latest version for each library
	// 2. Run benchmark tests (e.g., echo server, broadcast)
	// 3. Collect metrics (connections/sec, throughput, latency)
	// 4. Store results with timestamp and version

	for i := range libraries {
		lib := &libraries[i] // Use pointer to modify original slice element
		fmt.Printf("Fetching version for %s...\n", lib.Name)
		version, err := getLatestVersion(lib.URL)
		if err != nil {
			log.Printf("WARN: Failed to get version for %s: %v. Skipping benchmark.", lib.Name, err)
			continue // Skip if version fetch fails
		}
		lib.Version = version
		fmt.Printf("Benchmarking %s@%s...\n", lib.Name, lib.Version)

		var benchResult BenchmarkResult
		var benchErr error

		switch lib.Name {
		case "gorilla/websocket":
			benchResult, benchErr = runGorillaBenchmark(*lib)
		case "lxzan/gws":
			benchResult, benchErr = runGwsBenchmark(*lib)
		case "gobwas/ws":
			benchResult, benchErr = runGobwasBenchmark(*lib)
		case "lesismal/nbio":
			benchResult, benchErr = runNbioBenchmark(*lib)
		case "coder/websocket":
			benchResult, benchErr = runCoderWebsocketBenchmark(*lib)
		// Add cases for other libraries here
		default:
			log.Printf("WARN: No benchmark implementation for %s. Using placeholder data.", lib.Name)
			// --- Placeholder for libraries without specific implementation ---
			benchResult = BenchmarkResult{
				LibraryName:    lib.Name,
				ThroughputMBps: 1.0,                // Indicate placeholder
				LatencyMs:      1.0,                // Indicate placeholder
				Timestamp:      benchmarkTimestamp, // Use the captured timestamp
				Version:        lib.Version,
			}
			benchErr = nil
			// --- End Placeholder ---
		}

		if benchErr != nil {
			log.Printf("ERROR: Benchmark failed for %s: %v", lib.Name, benchErr)
			continue // Skip adding results if benchmark failed
		}

		// Use the captured timestamp for successful benchmark results
		benchResult.Timestamp = benchmarkTimestamp
		results = append(results, benchResult)
	}

	if len(results) > 0 {
		saveResults(results)
	} else {
		log.Println("No benchmarks were successfully run.")
	}

	log.Println("Benchmark finished.")
}
