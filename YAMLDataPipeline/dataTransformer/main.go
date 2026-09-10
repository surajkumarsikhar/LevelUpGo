package main

import (
	"fmt"
	"strings"
)

func addColumn(rows []map[string]string, name string, fn func(map[string]string) string) []map[string]string {
	// implement
	for _, m := range rows {
		value := fn(m)
		m[name] = value
	}
	return rows
}

func renameColumn(rows []map[string]string, oldName, newName string) []map[string]string {
	// implement
	for _, m := range rows {
		if value, ok := m[oldName]; ok {
			delete(m, oldName)
			m[newName] = value
		}
	}
	return rows
}

func filterRows(rows []map[string]string, fn func(map[string]string) bool) []map[string]string {
	// implement
	var result []map[string]string
	for _, m := range rows {
		if fn(m) {
			result = append(result, m)
		}
	}
	return result
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
        - image: myregistry/auth:v1.3  # latest stable`

	rows := parseYAML(yaml)
	fmt.Printf("Parsed %d deployments\n\n", len(rows))

	rows = addColumn(rows, "label", func(row map[string]string) string {
		return row["team"] + "/" + row["name"]
	})
	fmt.Println("Added label column:")
	for _, row := range rows {
		fmt.Printf("  %s\n", row["label"])
	}

	rows = renameColumn(rows, "image", "container_image")
	fmt.Printf("\nRenamed image -> container_image: %v\n", getColumn(rows, "container_image"))

	production := filterRows(rows, func(row map[string]string) bool {
		return row["team"] == "platform"
	})
	fmt.Printf("\nFiltered to platform team: %d services\n", len(production))
	for _, row := range production {
		fmt.Printf("  %s (replicas: %s)\n", row["label"], row["replicas"])
	}
}
