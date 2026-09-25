package main

import (
	"fmt"
	"os"
	"time"
)

func WriteReport(stats map[string]int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating report: %w", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "Log Analysis Report\n")
	fmt.Fprintf(file, "===================\n\n")

	total := 0
	for _, count := range stats {
		total += count
	}
	fmt.Fprintf(file, "Total Entries: %d\n\n", total)

	fmt.Fprintf(file, "Breakdown:\n")
	fmt.Fprintf(file, "  ERROR: %d\n", stats["ERROR"])
	fmt.Fprintf(file, "  WARN:  %d\n", stats["WARN"])
	fmt.Fprintf(file, "  INFO:  %d\n", stats["INFO"])

	return nil
}

func AppendLog(filename, level, message string) error {
	file, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return fmt.Errorf("opening log file: %w", err)
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_, err = fmt.Fprintf(file, "[%s] %s - %s\n", level, timestamp, message)
	if err != nil {
		return fmt.Errorf("writing log: %w", err)
	}

	return nil
}

func main() {
	stats := map[string]int{
		"ERROR": 2,
		"WARN":  3,
		"INFO":  10,
	}

	err := WriteReport(stats, "report.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Report written")

	err = AppendLog("app.log", "INFO", "Application started")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Log appended")
}
