package main

import (
	"fmt"
	"strings"
)

func groupBy(rows []map[string]string, column string) map[string][]map[string]string {
	// implement
	groups := make(map[string][]map[string]string)
	for _, row := range rows {
		val, ok := row[column]
		if ok {
			groups[val] = append(groups[val], row)
		}
	}
	return groups
}

func countBy(rows []map[string]string, column string) map[string]int {
	// implement\
	counts := make(map[string]int)

	for _, row := range rows {
		if val, ok := row[column]; ok {
			counts[val]++
		}
	}
	return counts
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

			// Strip YAML list item markers
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
        - image: myregistry/etl:v3.0`

	rows := parseYAML(yaml)
	fmt.Printf("Parsed %d deployments\n\n", len(rows))

	groups := groupBy(rows, "team")
	fmt.Println("Grouped by team:")
	for team, members := range groups {
		fmt.Printf("  %s:\n", team)
		for _, row := range members {
			fmt.Printf("    - %s\n", row["name"])
		}
	}

	fmt.Println()
	counts := countBy(rows, "team")
	fmt.Println("Count by team:")
	for team, count := range counts {
		fmt.Printf("  %s: %d\n", team, count)
	}

	_ = strings.Split
}
