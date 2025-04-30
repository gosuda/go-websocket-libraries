package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"
)

// BenchmarkResult holds the performance metrics for a library (copied from benchmark/config.go)
type BenchmarkResult struct {
	LibraryName          string
	ConnectionsPerSecond float64
	ThroughputMBps       float64
	LatencyMs            float64
	Timestamp            time.Time
	Version              string
}

const (
	readmePath        = "README.md"
	historyPath       = "benchmark/benchmark_history.json"
	tableStartMarker  = "<!-- BENCHMARK_TABLE_START -->"
	tableEndMarker    = "<!-- BENCHMARK_TABLE_END -->"
	readmeTemplateStr = `{{.TableStartMarker}}
**Last Updated:** {{ .Timestamp }}

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
{{- range .Results }}
| [{{ .LibraryName }}](https://github.com/{{ .LibraryName }}) | {{ .Version }} | {{ printf "%.2f" .ThroughputMBps }} | {{ printf "%.2f" .LatencyMs }} |
{{- end }}
{{.TableEndMarker}}`
)

var readmeTemplate = template.Must(template.New("readme").Parse(readmeTemplateStr))

func main() {
	log.Println("Updating README benchmark table...")

	// 1. Read history
	historyData, err := os.ReadFile(historyPath)
	if err != nil {
		log.Fatalf("Failed to read benchmark history file %s: %v", historyPath, err)
	}

	var history []BenchmarkResult
	if err := json.Unmarshal(historyData, &history); err != nil {
		log.Fatalf("Failed to parse benchmark history file %s: %v", historyPath, err)
	}

	if len(history) == 0 {
		log.Println("No benchmark history found. Skipping README update.")
		return
	}

	// 2. Find latest result for each library
	latestResults := make(map[string]BenchmarkResult)
	for _, result := range history {
		if existing, ok := latestResults[result.LibraryName]; !ok || result.Timestamp.After(existing.Timestamp) {
			latestResults[result.LibraryName] = result
		}
	}

	// Convert map to slice and sort by name for consistent table order
	var sortedResults []BenchmarkResult
	for _, result := range latestResults {
		sortedResults = append(sortedResults, result)
	}
	sort.Slice(sortedResults, func(i, j int) bool {
		return sortedResults[i].LibraryName < sortedResults[j].LibraryName
	})

	// 3. Read README content
	readmeContent, err := os.ReadFile(readmePath)
	if err != nil {
		log.Fatalf("Failed to read README file %s: %v", readmePath, err)
	}

	// 4. Generate new table content
	var tableBuf bytes.Buffer
	templateData := struct {
		Results          []BenchmarkResult
		Timestamp        string
		TableStartMarker string
		TableEndMarker   string
	}{
		Results:          sortedResults,
		Timestamp:        time.Now().UTC().Format(time.RFC1123),
		TableStartMarker: tableStartMarker,
		TableEndMarker:   tableEndMarker,
	}
	if err := readmeTemplate.Execute(&tableBuf, templateData); err != nil {
		log.Fatalf("Failed to execute README template: %v", err)
	}
	newTableContent := tableBuf.String()

	// 5. Replace old table with new table in README content
	startIdx := strings.Index(string(readmeContent), tableStartMarker)
	endIdx := strings.Index(string(readmeContent), tableEndMarker)

	if startIdx == -1 || endIdx == -1 || endIdx <= startIdx {
		log.Fatalf("Could not find table markers '%s' and '%s' in %s", tableStartMarker, tableEndMarker, readmePath)
	}
	endIdx += len(tableEndMarker) // Include the end marker in the slice to replace

	newReadmeContent := string(readmeContent[:startIdx]) + newTableContent + string(readmeContent[endIdx:])

	// 6. Write updated README back
	err = os.WriteFile(readmePath, []byte(newReadmeContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write updated README file %s: %v", readmePath, err)
	}

	log.Printf("Successfully updated benchmark table in %s", readmePath)
}
