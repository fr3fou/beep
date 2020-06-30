package main

import (
	"log"
	"math"
	"os"

	"github.com/fr3fou/gusic/gusic"
)

// Memory C# F# G#
func Memory() gusic.Melody {
	sampleRate := 48000.0
	bpm := 120
	noteLength := 4 // 4/4, therefore 4
	volume := 0.05

	m := gusic.NewMelody(
		sampleRate,
		bpm,
		noteLength,
		math.Sin,
		gusic.NewLinearADSR(gusic.NewRatios(0.25, 0.25, 0.25, 0.25), 1.35, 0.35, 1.0),
	)

	m.AddNotes(
		gusic.D(5, m.Quaver(), volume),
		gusic.B(5, m.Quaver(), volume),
		gusic.A(5, m.Quaver(), volume),
		gusic.E(5, m.Quaver(), volume),
		gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
		gusic.GS(5, m.Semiquaver()*3, volume),
		gusic.A(5, m.Quaver(), volume),
		//
		gusic.Pause(m.Quaver()),
		gusic.E(5, m.Quaver(), volume),
		gusic.A(5, m.Quaver(), volume),
		gusic.E(5, m.Quaver(), volume),
		gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
		gusic.GS(5, m.Semiquaver()*3, volume),
		gusic.A(5, m.Quaver(), volume),
		//
		gusic.D(5, m.Quaver(), volume),
		gusic.B(5, m.Quaver(), volume),
		gusic.A(5, m.Quaver(), volume),
		gusic.E(5, m.Quaver(), volume),
		gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
		gusic.GS(5, m.Semiquaver()*3, volume),
		gusic.A(5, m.Quaver(), volume),
		//
		gusic.Pause(m.Quaver()),
		gusic.E(5, m.Quaver(), volume),
		gusic.A(5, m.Quaver(), volume),
		gusic.CS(6, m.Quaver(), volume),
		gusic.B(5, m.Quaver()+m.Semiquaver(), volume),
		gusic.A(5, m.Semiquaver()*3, volume),
		gusic.B(5, m.Quaver(), volume),
		// repeat
		gusic.D(5, m.Quaver(), volume),
		gusic.B(5, m.Quaver(), volume),
		gusic.A(5, m.Quaver(), volume),
		gusic.E(5, m.Quaver(), volume),
		gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
		gusic.GS(5, m.Semiquaver()*3, volume),
		gusic.A(5, m.Quaver(), volume),
		//
		gusic.Pause(m.Quaver()),
		gusic.E(5, m.Quaver(), volume),
		gusic.A(5, m.Quaver(), volume),
		gusic.E(5, m.Quaver(), volume),
		gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
		gusic.GS(5, m.Semiquaver()*3, volume),
		gusic.A(5, m.Quaver(), volume),
		//
		gusic.D(5, m.Quaver(), volume),
		gusic.B(5, m.Quaver(), volume),
		gusic.A(5, m.Quaver(), volume),
		gusic.E(5, m.Quaver(), volume),
		gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
		gusic.GS(5, m.Semiquaver()*3, volume),
		gusic.A(5, m.Quaver(), volume),
		//
		gusic.Pause(m.Quaver()),
		gusic.E(5, m.Quaver(), volume),
		gusic.A(5, m.Quaver(), volume),
		gusic.CS(6, m.Quaver(), volume),
		gusic.B(5, m.Quaver()+m.Semiquaver(), volume),
		gusic.A(5, m.Semiquaver()*3, volume),
		gusic.B(5, m.Quaver(), volume),
	)
	return *m
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("gusic: File name argument required\nUsage: %s <filename>", os.Args[0])
	}
	FileName := os.Args[1]
	file, err := os.Create(FileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	m := Memory()
	m.Write(file)

	// pcm := m.PCM()       // returns the PCM array
	// err := m.Write(file) // write PCM to an io.Writer

	// err := wav.Write(m, file) // write wav to an io.writer
}
