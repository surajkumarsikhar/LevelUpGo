package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

func GetCallerInfo(skip int) string {
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fn := runtime.FuncForPC(pc)
		return fmt.Sprintf("%s@%s:%d\n", fn.Name(), file, line)
	}
	return "unknown"
}

func LogWithCaller(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Sprintf("[%s:%d] %s\n", filepath.Base(file), line, msg)
	}
}

func WrapError(err error) error {
	_, file, line, ok := runtime.Caller(1)
	var warpErr error
	if ok {
		warpErr = fmt.Errorf("[%s:%d] %s\n", file, line, err)
	}
	return warpErr
}

func main() {
	// Test GetCallerInfo
	info := GetCallerInfo(0)
	fmt.Printf("Caller info: %s\n", info)

	// Test LogWithCaller
	LogWithCaller("Application started")
	LogWithCaller("Processing request")

	// Test WrapError
	originalErr := errors.New("original error")
	wrapped := WrapError(originalErr)
	fmt.Printf("Wrapped error: %v\n", wrapped)
}
