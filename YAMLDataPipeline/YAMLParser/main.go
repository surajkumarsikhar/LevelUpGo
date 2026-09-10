package main

import (
	"fmt"
	"strings"
)

func parseYAML(input string) []map[string]string {
	// implement
	if input == "" {
		return nil
	}
	result := []map[string]string{}
	docs := strings.Split(input, "---")
	for _, doc := range docs {
		mp := make(map[string]string)
		doc = strings.TrimSpace(doc)
		lines := strings.Split(doc, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if len(line) < 1 {
				continue
			}
			if line[0] == '#' {
				continue
			}
			if idx := strings.Index(line, " #"); idx != -1 {
				line = line[:idx]
			}
			if strings.HasPrefix(line, "- ") {
				line = line[2:]
			}
			parts := strings.SplitN(line, ":", 2)
			if len(parts) < 2 {
				continue
			}
			if parts[0] == "" || parts[1] == "" {
				continue
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			mp[key] = value
		}
		if len(mp) < 1 {
			continue
		}
		result = append(result, mp)
	}
	if len(result) < 1 {
		result = nil
	}
	return result
}

func getColumn(rows []map[string]string, column string) []string {
	// implement
	if rows == nil || len(rows) < 1 {
		return nil
	}
	result := []string{}

	for _, m := range rows {
		value, ok := m[column]
		if ok {
			result = append(result, value)
		}
	}
	if len(result) < 1 {
		result = nil
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
  name: auth-service
  labels:
    team: security
spec:
  replicas: 2
  template:
    spec:
      containers:
        - image: myregistry/auth:v1.3`

	rows := parseYAML(yaml)
	fmt.Printf("Parsed %d deployments\n\n", len(rows))

	for i, row := range rows {
		fmt.Printf("Deployment %d:\n", i+1)
		for key, val := range row {
			fmt.Printf("  %s: %s\n", key, val)
		}
		fmt.Println()
	}

	names := getColumn(rows, "name")
	fmt.Printf("Service names: %v\n", names)

	teams := getColumn(rows, "team")
	fmt.Printf("Teams: %v\n", teams)

	_ = strings.Split
}
