package gusic

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestLinearADSR(t *testing.T) {
	m := NewMelody(
		48000,
		120,
		4,
		math.Sin,
		NewLinearADSR(NewRatios(0.25, 0.25, 0.25, 0.25), 1.35, 0.35, 1.0),
	)
	notes := []Note{}

	n := 32
	for i := 0; i < n; i++ {
		notes = append(notes, A(4, time.Second, 0.015))
	}

	m.AddNotes(notes...)
	m.ApplyADSR()
	printVolume(m.Notes)
}

func printVolume(notes []Note) {
	for i, note := range notes {
		fmt.Printf("%d,%f\n", i, note.Volume)
	}
}
