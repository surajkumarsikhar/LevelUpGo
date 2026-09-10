package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sum(rows []map[string]string, column string) float64 {
	// implement
	var sum float64
	for _, row := range rows {
		val, err := strconv.ParseFloat(row[column], 64)
		if err == nil {
			sum += val
		}
	}
	return sum
}

func average(rows []map[string]string, column string) float64 {
	// implement
	count := 0
	for _, row := range rows {
		_, err := strconv.ParseFloat(row[column], 64)
		if err == nil {
			count++
		}
	}
	sum := sum(rows, column)
	if count == 0 {
		return 0.0
	}
	return sum / float64(count)
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
        - image: myregistry/api:v2.1  # production
          resources:
            cpu: 500
            memory: 512
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
        - image: myregistry/auth:v1.3
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
            memory: 2048`

	rows := parseYAML(yaml)
	fmt.Printf("Parsed %d deployments\n\n", len(rows))

	fmt.Printf("Total CPU:     %.0f millicores\n", sum(rows, "cpu"))
	fmt.Printf("Total Memory:  %.0f MB\n", sum(rows, "memory"))
	fmt.Printf("Total Replicas: %.0f\n\n", sum(rows, "replicas"))

	fmt.Printf("Avg CPU:       %.0f millicores\n", average(rows, "cpu"))
	fmt.Printf("Avg Memory:    %.0f MB\n", average(rows, "memory"))
	fmt.Printf("Avg Replicas:  %.1f\n\n", average(rows, "replicas"))

	platformRows := groupBy(rows, "team")["platform"]
	fmt.Printf("Platform team:\n")
	fmt.Printf("  CPU:      %.0f millicores\n", sum(platformRows, "cpu"))
	fmt.Printf("  Memory:   %.0f MB\n", sum(platformRows, "memory"))
	fmt.Printf("  Replicas: %.0f\n", sum(platformRows, "replicas"))

	_ = strconv.ParseFloat
}
