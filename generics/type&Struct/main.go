package main

import "fmt"

// TODO: Define UndoHistory[T any] struct and implement Save, Undo, CanUndo
type UndoHistory[T any] struct {
	states []T
}

func (h *UndoHistory[T]) Save(state T) {
	h.states = append(h.states, state)
}

func (h *UndoHistory[T]) Undo() (T, bool) {
	var res T
	if l := len(h.states); l > 0 {
		res = h.states[l-1]
		h.states = h.states[:l-1]
		return res, true
	}
	return res, false
}

func (h *UndoHistory[T]) CanUndo() bool {
	return len(h.states) > 0
}

func main() {
	var h UndoHistory[string]

	fmt.Println("Can undo:", h.CanUndo())

	h.Save("draft 1")
	h.Save("draft 2")
	h.Save("draft 3")

	fmt.Println("Can undo:", h.CanUndo())

	val, ok := h.Undo()
	fmt.Printf("Undo: %s, %v\n", val, ok)

	val, ok = h.Undo()
	fmt.Printf("Undo: %s, %v\n", val, ok)
}
