package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func saveResults(newResults []BenchmarkResult) {
	historyFilePath := "benchmark/benchmark_history.json"
	log.Printf("Saving results to %s...", historyFilePath)

	var history []BenchmarkResult

	// Read existing history
	data, err := os.ReadFile(historyFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("WARN: Could not read existing history file %s: %v. Starting fresh.", historyFilePath, err)
		}
		// If file doesn't exist or other read error, history remains empty slice
	} else {
		if err := json.Unmarshal(data, &history); err != nil {
			log.Printf("WARN: Could not parse existing history file %s: %v. Starting fresh.", historyFilePath, err)
			history = []BenchmarkResult{} // Reset history if parsing fails
		}
	}

	// Append new results
	history = append(history, newResults...)

	// Write updated history back to file
	updatedData, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal updated history: %v", err)
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(historyFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Failed to create directory %s: %v", dir, err)
	}

	err = os.WriteFile(historyFilePath, updatedData, 0644)
	if err != nil {
		log.Fatalf("Failed to write updated history file %s: %v", historyFilePath, err)
	}
	log.Printf("Successfully saved %d new results.", len(newResults))
}
