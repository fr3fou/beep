package gusic

import (
	"fmt"
	"testing"
	"time"
)

func TestADSR(t *testing.T) {
	adsr := NewADSR(0.25, 1.35, 0.25, 0.35, 0.25, 0.25, 1.0)

	n := 32

	notes := []Note{}

	for i := 0; i < n; i++ {
		notes = append(notes, A(4, time.Second, 0.015))
	}

	adsr.Apply(notes)
	printVolume(notes)
}

func printVolume(notes []Note) {
	for i, note := range notes {
		fmt.Printf("%d,%f\n", i, note.Volume)
	}
}
