package main

// TODO: Define Weekday type and constants
// type Weekday int
// const ( Sunday Weekday = iota ... )
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// TODO: Define permission bit flags
// const ( Read = 1 << iota ... )
const (
	Read = 1 << iota
	Write
	Execute
)

// TODO: Define byte size constants
// const ( KB = 1 << (10 * iota) after skipping 0... )
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

func main() {
	// Your constants will be tested automatically.
	// Just define the three const blocks above.
}
