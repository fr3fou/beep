package main

import (
	"log"
	"math"
	"os"

	"github.com/fr3fou/beep/beep"
)

// Memory C# F# G#
func Memory() beep.Melody {
	sampleRate := 48000.0
	bpm := 80
	noteLength := 4 // 4/4, therefore 4
	volume := 0.125

	m := beep.NewMelody(
		sampleRate,
		bpm,
		noteLength,
		math.Sin,
		beep.NewEasedADSR(func(t float64) float64 {
			return t * t
		}, beep.NewRatios(0.25, 0.25, 0.25, 0.25), 1.35, 0.35),
	)

	m.AddNotes(
		beep.D(5, m.Quaver(), volume),
		beep.B(5, m.Quaver(), volume),
		beep.A(5, m.Quaver(), volume),
		beep.E(5, m.Quaver(), volume),
		beep.GS(5, m.Quaver()+m.Semiquaver(), volume),
		beep.GS(5, m.Semiquaver()*3, volume),
		beep.A(5, m.Quaver(), volume),
		//
		beep.Rest(m.Quaver()),
		beep.E(5, m.Quaver(), volume),
		beep.A(5, m.Quaver(), volume),
		beep.E(5, m.Quaver(), volume),
		beep.GS(5, m.Quaver()+m.Semiquaver(), volume),
		beep.GS(5, m.Semiquaver()*3, volume),
		beep.A(5, m.Quaver(), volume),
		//
		beep.D(5, m.Quaver(), volume),
		beep.B(5, m.Quaver(), volume),
		beep.A(5, m.Quaver(), volume),
		beep.E(5, m.Quaver(), volume),
		beep.GS(5, m.Quaver()+m.Semiquaver(), volume),
		beep.GS(5, m.Semiquaver()*3, volume),
		beep.A(5, m.Quaver(), volume),
		//
		beep.Rest(m.Quaver()),
		beep.E(5, m.Quaver(), volume),
		beep.A(5, m.Quaver(), volume),
		beep.CS(6, m.Quaver(), volume),
		beep.B(5, m.Quaver()+m.Semiquaver(), volume),
		beep.A(5, m.Semiquaver()*3, volume),
		beep.B(5, m.Quaver(), volume),
		// repeat
		beep.D(5, m.Quaver(), volume),
		beep.B(5, m.Quaver(), volume),
		beep.A(5, m.Quaver(), volume),
		beep.E(5, m.Quaver(), volume),
		beep.GS(5, m.Quaver()+m.Semiquaver(), volume),
		beep.GS(5, m.Semiquaver()*3, volume),
		beep.A(5, m.Quaver(), volume),
		//
		beep.Rest(m.Quaver()),
		beep.E(5, m.Quaver(), volume),
		beep.A(5, m.Quaver(), volume),
		beep.E(5, m.Quaver(), volume),
		beep.GS(5, m.Quaver()+m.Semiquaver(), volume),
		beep.GS(5, m.Semiquaver()*3, volume),
		beep.A(5, m.Quaver(), volume),
		//
		beep.D(5, m.Quaver(), volume),
		beep.B(5, m.Quaver(), volume),
		beep.A(5, m.Quaver(), volume),
		beep.E(5, m.Quaver(), volume),
		beep.GS(5, m.Quaver()+m.Semiquaver(), volume),
		beep.GS(5, m.Semiquaver()*3, volume),
		beep.A(5, m.Quaver(), volume),
		//
		beep.Rest(m.Quaver()),
		beep.E(5, m.Quaver(), volume),
		beep.A(5, m.Quaver(), volume),
		beep.CS(6, m.Quaver(), volume),
		beep.B(5, m.Quaver()+m.Semiquaver(), volume),
		beep.A(5, m.Semiquaver()*3, volume),
		beep.B(5, m.Quaver(), volume),
	)

	m.AddRun(
		//
		beep.Rest(m.Quaver()),
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3), //
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3), //
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3), //
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3),
		beep.D(4, m.Crotchet(), volume*2/3), //
		beep.C(4, m.Crotchet(), volume*2/3),
		beep.C(4, m.Crotchet(), volume*2/3),
		beep.C(4, m.Crotchet(), volume*2/3),
		beep.C(4, m.Crotchet(), volume*2/3), //
		beep.C(4, m.Crotchet(), volume*2/3),
		beep.C(4, m.Crotchet(), volume*2/3),
		beep.C(4, m.Crotchet(), volume*2/3),
		beep.C(4, m.Crotchet(), volume*2/3), //
		beep.C(4, m.Crotchet(), volume*2/3),
		beep.C(4, m.Crotchet(), volume*2/3),
		beep.C(4, m.Crotchet(), volume*2/3),
		beep.C(4, m.Crotchet(), volume*2/3), //
		beep.FS(4, m.Crotchet(), volume*2/3),
		beep.FS(4, m.Crotchet(), volume*2/3),
		beep.FS(4, m.Crotchet(), volume*2/3),
		beep.FS(4, m.Quaver(), volume*2/3),
	)

	return *m
}

func BlindingLights() beep.Melody {
	sampleRate := 48000.0
	bpm := 171
	noteLength := 4 // 4/4, therefore 4
	volume := 0.125

	m := beep.NewMelody(
		sampleRate,
		bpm,
		noteLength,
		math.Sin,
		beep.NewEasedADSR(math.Sin, beep.NewRatios(0.25, 0.25, 0.25, 0.25), 1.35, 0.35),
	)

	m.AddNotes(
		beep.F(5, m.Half(), volume),
		beep.F(5, m.Dotted(m.Quarter()), volume),
		beep.DS(5, m.Eighth(), volume),
		//
		beep.F(5, m.Eighth(), volume),
		beep.G(5, m.Quarter(), volume),
		beep.C(5, m.Eighth(), volume),
		beep.C(5, m.Eighth(), volume),
		beep.DS(5, m.Dotted(m.Quarter()), volume),
		//
		beep.F(5, m.Half(), volume),
		beep.F(5, m.Dotted(m.Quarter()), volume),
		beep.DS(5, m.Eighth(), volume),
		//
		beep.F(5, m.Eighth(), volume),
		beep.G(5, m.Quarter(), volume),
		beep.C(5, m.Eighth(), volume),
		beep.C(5, m.Eighth(), volume),
		beep.DS(5, m.Dotted(m.Quarter()), volume),
		//
		beep.B(5, m.Eighth(), volume),
		beep.G(5, m.Quarter(), volume),
		beep.F(5, m.Quarter(), volume),
		beep.F(5, m.Quarter(), volume),
		beep.DS(5, m.Eighth(), volume),
		beep.F(5, m.Eighth(), volume),
		//
		beep.F(5, m.Whole(), volume),
		beep.F(5, m.Half(), volume),
		beep.C(5, m.Eighth(), volume),
		beep.B(5, m.Eighth(), volume),
		beep.C(5, m.Eighth(), volume),
		beep.B(5, m.Eighth(), volume),
		//
		//
	)

	return *m
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("beep: File name argument required\nUsage: %s <filename>", os.Args[0])
	}

	fileName := os.Args[1]
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	m := Memory()

	if err := m.Write(file); err != nil {
		panic(err)
	}
}
