package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func generateReport(rows []map[string]string) string {
	// implement
	seen := make(map[string]bool)
	for _, row := range rows {
		for key := range row {
			seen[key] = true
		}
	}

	columns := make([]string, 0, len(seen))

	for key := range seen {
		columns = append(columns, key)
	}

	sort.Strings(columns)

	columnStr := "(none)"
	if len(columns) > 0 {
		columnStr = strings.Join(columns, ", ")
	}
	return fmt.Sprintf("=== Pipeline Report ===\nTotal Rows: %d\nColumns: %s", len(rows), columnStr)
}

// --- Completed from Lesson 1 ---

func parseYAML(input string) []map[string]string {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	docs := strings.Split(input, "---")
	var rows []map[string]string

	for _, doc := range docs {
		doc = strings.TrimSpace(doc)
		if doc == "" {
			continue
		}

		lines := strings.Split(doc, "\n")
		row := make(map[string]string)

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			if strings.HasPrefix(line, "- ") {
				line = line[2:]
			}

			if strings.HasPrefix(line, "#") {
				continue
			}

			if idx := strings.Index(line, " #"); idx != -1 {
				line = line[:idx]
			}

			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if key == "" || value == "" {
				continue
			}

			row[key] = value
		}

		if len(row) > 0 {
			rows = append(rows, row)
		}
	}

	return rows
}

func getColumn(rows []map[string]string, column string) []string {
	if len(rows) == 0 {
		return nil
	}

	var result []string
	for _, row := range rows {
		if val, ok := row[column]; ok {
			result = append(result, val)
		}
	}

	return result
}

// --- Completed from Lesson 2 ---

func addColumn(rows []map[string]string, name string, fn func(map[string]string) string) []map[string]string {
	if len(rows) == 0 {
		return rows
	}

	for _, row := range rows {
		row[name] = fn(row)
	}

	return rows
}

func renameColumn(rows []map[string]string, oldName, newName string) []map[string]string {
	for _, row := range rows {
		if val, ok := row[oldName]; ok {
			row[newName] = val
			delete(row, oldName)
		}
	}

	return rows
}

func filterRows(rows []map[string]string, fn func(map[string]string) bool) []map[string]string {
	if len(rows) == 0 {
		return nil
	}

	var result []map[string]string
	for _, row := range rows {
		if fn(row) {
			result = append(result, row)
		}
	}

	return result
}

// --- Completed from Lesson 3 ---

func groupBy(rows []map[string]string, column string) map[string][]map[string]string {
	groups := make(map[string][]map[string]string)
	if len(rows) == 0 {
		return groups
	}
	for _, row := range rows {
		key, ok := row[column]
		if !ok {
			continue
		}
		groups[key] = append(groups[key], row)
	}
	return groups
}

func countBy(rows []map[string]string, column string) map[string]int {
	counts := make(map[string]int)
	if len(rows) == 0 {
		return counts
	}
	for _, row := range rows {
		val, ok := row[column]
		if !ok {
			continue
		}
		counts[val]++
	}
	return counts
}

// --- Completed from Lesson 4 ---

func sum(rows []map[string]string, column string) float64 {
	if len(rows) == 0 {
		return 0.0
	}
	total := 0.0
	for _, row := range rows {
		val, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			continue
		}
		total += val
	}
	return total
}

func average(rows []map[string]string, column string) float64 {
	if len(rows) == 0 {
		return 0.0
	}
	total := 0.0
	count := 0
	for _, row := range rows {
		val, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			continue
		}
		total += val
		count++
	}
	if count == 0 {
		return 0.0
	}
	return total / float64(count)
}

// --- Completed from Lesson 5 ---

type PipelineStep func([]map[string]string) []map[string]string

type Pipeline struct {
	steps []PipelineStep
}

func NewPipeline() *Pipeline {
	return &Pipeline{steps: []PipelineStep{}}
}

func (p *Pipeline) AddStep(step PipelineStep) *Pipeline {
	p.steps = append(p.steps, step)
	return p
}

func (p *Pipeline) Run(rows []map[string]string) []map[string]string {
	result := rows
	for _, step := range p.steps {
		result = step(result)
	}
	return result
}

func main() {
	yaml := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-gateway
  labels:
    team: platform
spec:
  replicas: 3
  template:
    spec:
      containers:
        - image: myregistry/api:v2.1
          resources:
            cpu: 500
            memory: 512
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-runner
  labels:
    team: qa
spec:
  replicas: 1
  template:
    spec:
      containers:
        - image: myregistry/test:v0.9
          resources:
            cpu: 100
            memory: 128
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auth-service
  labels:
    team: platform
spec:
  replicas: 2
  template:
    spec:
      containers:
        - image: myregistry/auth:v1.3  # latest stable
          resources:
            cpu: 250
            memory: 256
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: etl-worker
  labels:
    team: data
spec:
  replicas: 5
  template:
    spec:
      containers:
        - image: myregistry/etl:v3.0
          resources:
            cpu: 1000
            memory: 2048
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: metrics-collector
  labels:
    team: infra
spec:
  replicas: 4
  template:
    spec:
      containers:
        - image: myregistry/metrics:v1.8
          resources:
            cpu: 800
            memory: 1024`

	rows := parseYAML(yaml)
	fmt.Printf("Parsed %d deployments from YAML\n\n", len(rows))

	pipeline := NewPipeline()
	pipeline.
		AddStep(func(rows []map[string]string) []map[string]string {
			return filterRows(rows, func(row map[string]string) bool {
				return row["team"] != "qa"
			})
		}).
		AddStep(func(rows []map[string]string) []map[string]string {
			return addColumn(rows, "label", func(row map[string]string) string {
				return row["team"] + "/" + row["name"]
			})
		}).
		AddStep(func(rows []map[string]string) []map[string]string {
			return renameColumn(rows, "cpu", "cpu_millicores")
		})

	result := pipeline.Run(rows)

	fmt.Println(generateReport(result))
	fmt.Println()
	for _, row := range result {
		fmt.Printf("%-25s cpu=%-6s mem=%-6s replicas=%s\n",
			row["label"], row["cpu_millicores"], row["memory"], row["replicas"])
	}
	fmt.Printf("\nTotal CPU:     %.0f millicores\n", sum(result, "cpu_millicores"))
	fmt.Printf("Total Memory:  %.0f MB\n", sum(result, "memory"))
	fmt.Printf("Avg Replicas:  %.1f\n", average(result, "replicas"))
	fmt.Println()
	fmt.Println("Services per team:")
	for team, count := range countBy(result, "team") {
		fmt.Printf("  %s: %d\n", team, count)
	}

	_ = sort.Strings
	_ = strconv.ParseFloat
}
