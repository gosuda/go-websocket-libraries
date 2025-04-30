package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	historyPath = "benchmark/benchmark_history.json"
	graphPath   = "benchmark_performance.png"
)

// BenchmarkResult matches the structure in benchmark/main.go
type BenchmarkResult struct {
	LibraryName          string    `json:"LibraryName"`
	ConnectionsPerSecond float64   `json:"ConnectionsPerSecond"`
	ThroughputMBps       float64   `json:"ThroughputMBps"`
	LatencyMs            float64   `json:"LatencyMs"`
	Timestamp            time.Time `json:"Timestamp"`
	Version              string    `json:"Version"`
}

// timeValue holds a timestamp and a corresponding performance value.
type timeValue struct {
	t time.Time
	v float64
}

// timeValues implements plotter.XYer.
type timeValues []timeValue

func (tv timeValues) Len() int { return len(tv) }
func (tv timeValues) XY(i int) (x, y float64) {
	// Convert time to Unix timestamp (seconds since epoch) for plotting
	return float64(tv[i].t.Unix()), tv[i].v
}

func main() {
	log.Println("Starting graph generation...")

	// Read history data
	historyData, err := os.ReadFile(historyPath)
	if err != nil {
		log.Fatalf("Failed to read benchmark history file %s: %v", historyPath, err)
	}

	var history []BenchmarkResult
	if err := json.Unmarshal(historyData, &history); err != nil {
		log.Fatalf("Failed to parse benchmark history file %s: %v", historyPath, err)
	}

	if len(history) == 0 {
		log.Println("No benchmark results found in history. Skipping graph generation.")
		return
	}

	// Group results by library
	resultsByLib := make(map[string]timeValues)
	for _, r := range history {
		resultsByLib[r.LibraryName] = append(resultsByLib[r.LibraryName], timeValue{t: r.Timestamp, v: r.ThroughputMBps})
	}

	// Create a new plot
	p := plot.New()

	p.Title.Text = "WebSocket Library Performance Over Time"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Throughput (MB/s)"

	// Use a time ticker for the X-axis
	p.X.Tick.Marker = plot.TimeTicks{Format: "15:04:05"} // Format for time (HH:MM:SS)
	p.X.Tick.Label.Rotation = 0.8                        // Rotate labels more to prevent overlap
	// p.X.Tick.Label.XAlign = vg.Right // Removed due to undefined error
	// p.X.Tick.Label.YAlign = vg.Center // Removed due to undefined error

	p.Add(plotter.NewGrid())

	// Add lines for each library
	i := 0
	libNames := make([]string, 0, len(resultsByLib))
	for name := range resultsByLib {
		libNames = append(libNames, name)
	}
	sort.Strings(libNames) // Sort names for consistent color assignment

	for _, name := range libNames {
		values := resultsByLib[name]
		// Sort values by time for correct line plotting
		sort.Slice(values, func(i, j int) bool {
			return values[i].t.Before(values[j].t)
		})

		line, points, err := plotter.NewLinePoints(values)
		if err != nil {
			log.Printf("WARN: Could not create line/points for %s: %v", name, err)
			continue
		}
		line.Color = plotutil.Color(i)
		points.Color = plotutil.Color(i)
		points.Shape = plotutil.Shape(i) // Use different shapes

		p.Add(line, points)
		p.Legend.Add(name, line, points)
		i++
	}

	p.Legend.Top = true
	p.Legend.Left = false // Position legend top-right

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 6*vg.Inch, graphPath); err != nil { // Increased size for better resolution/layout
		log.Fatalf("Could not save plot: %v", err)
	}

	log.Printf("Graph saved successfully to %s", graphPath)
}
