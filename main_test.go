package main

import (
	"fmt"
	"testing"
)

func TestEnvelope(t *testing.T) {
	n := 32

	notes := []Sample{}

	for i := 0; i < n; i++ {
		notes = append(notes, 0.5)
	}

	printVolume(Envelope(notes, 0.25, 0.25, 0.25))
}

func printVolume(notes []Sample) {
	for i, note := range notes {
		fmt.Printf("%d,%f\n", i, note)
	}
}
