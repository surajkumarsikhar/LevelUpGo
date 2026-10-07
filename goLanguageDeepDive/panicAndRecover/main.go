package main

import (
	"fmt"
	"strconv"
)

func MustParseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		msg := fmt.Sprintf("invalid integer: %s", s)
		panic(msg)
	}
	return num
}

func SafeParseInt(s string) (result int, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	result = MustParseInt(s)
	return result, nil
}

func ExecuteWithRecovery(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	f()
	return nil
}

func main() {
	fmt.Printf("MustParseInt(\"42\"): %d\n", MustParseInt("42"))

	n, err := SafeParseInt("123")
	fmt.Printf("SafeParseInt(\"123\"): %d,%v\n", n, err)

	n, err = SafeParseInt("abc")
	fmt.Printf("SafeParseInt(\"abc\"): %d, %v\n", n, err)

	safeFunc := func() {
		fmt.Print("") // Do nothing dangerous
	}

	panickyFunc := func() {
		panic("intentional panic")
	}

	err = ExecuteWithRecovery(safeFunc)
	fmt.Printf("ExecuteWithRecovery(safe): %v\n", err)

	err = ExecuteWithRecovery(panickyFunc)
	fmt.Printf("ExecuteWithRecovery(panicky): %v\n", err)
}
