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
	bpm := 80
	noteLength := 4 // 4/4, therefore 4
	volume := 0.125

	m := gusic.NewMelody(
		sampleRate,
		bpm,
		noteLength,
		math.Sin,
		gusic.NewEasedADSR(func(t float64) float64 {
			return t * t
		}, gusic.NewRatios(0.25, 0.25, 0.25, 0.25), 1.35, 0.35),
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
		gusic.Rest(m.Quaver()),
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
		gusic.Rest(m.Quaver()),
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
		gusic.Rest(m.Quaver()),
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
		gusic.Rest(m.Quaver()),
		gusic.E(5, m.Quaver(), volume),
		gusic.A(5, m.Quaver(), volume),
		gusic.CS(6, m.Quaver(), volume),
		gusic.B(5, m.Quaver()+m.Semiquaver(), volume),
		gusic.A(5, m.Semiquaver()*3, volume),
		gusic.B(5, m.Quaver(), volume),
	)

	m.AddRun(
		//
		gusic.Rest(m.Quaver()),
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3), //
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3), //
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3), //
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3),
		gusic.D(4, m.Crotchet(), volume*2/3), //
		gusic.C(4, m.Crotchet(), volume*2/3),
		gusic.C(4, m.Crotchet(), volume*2/3),
		gusic.C(4, m.Crotchet(), volume*2/3),
		gusic.C(4, m.Crotchet(), volume*2/3), //
		gusic.C(4, m.Crotchet(), volume*2/3),
		gusic.C(4, m.Crotchet(), volume*2/3),
		gusic.C(4, m.Crotchet(), volume*2/3),
		gusic.C(4, m.Crotchet(), volume*2/3), //
		gusic.C(4, m.Crotchet(), volume*2/3),
		gusic.C(4, m.Crotchet(), volume*2/3),
		gusic.C(4, m.Crotchet(), volume*2/3),
		gusic.C(4, m.Crotchet(), volume*2/3), //
		gusic.FS(4, m.Crotchet(), volume*2/3),
		gusic.FS(4, m.Crotchet(), volume*2/3),
		gusic.FS(4, m.Crotchet(), volume*2/3),
		gusic.FS(4, m.Quaver(), volume*2/3),
	)

	return *m
}

func BlindingLights() gusic.Melody {
	sampleRate := 48000.0
	bpm := 171
	noteLength := 4 // 4/4, therefore 4
	volume := 0.125

	m := gusic.NewMelody(
		sampleRate,
		bpm,
		noteLength,
		math.Sin,
		gusic.NewEasedADSR(math.Sin, gusic.NewRatios(0.25, 0.25, 0.25, 0.25), 1.35, 0.35),
	)

	m.AddNotes(
		gusic.F(5, m.Half(), volume),
		gusic.F(5, m.Dotted(m.Quarter()), volume),
		gusic.DS(5, m.Eighth(), volume),
		//
		gusic.F(5, m.Eighth(), volume),
		gusic.G(5, m.Quarter(), volume),
		gusic.C(5, m.Eighth(), volume),
		gusic.C(5, m.Eighth(), volume),
		gusic.DS(5, m.Dotted(m.Quarter()), volume),
		//
		gusic.F(5, m.Half(), volume),
		gusic.F(5, m.Dotted(m.Quarter()), volume),
		gusic.DS(5, m.Eighth(), volume),
		//
		gusic.F(5, m.Eighth(), volume),
		gusic.G(5, m.Quarter(), volume),
		gusic.C(5, m.Eighth(), volume),
		gusic.C(5, m.Eighth(), volume),
		gusic.DS(5, m.Dotted(m.Quarter()), volume),
		//
		gusic.B(5, m.Eighth(), volume),
		gusic.G(5, m.Quarter(), volume),
		gusic.F(5, m.Quarter(), volume),
		gusic.F(5, m.Quarter(), volume),
		gusic.DS(5, m.Eighth(), volume),
		gusic.F(5, m.Eighth(), volume),
		//
		gusic.F(5, m.Whole(), volume),
		gusic.F(5, m.Half(), volume),
		gusic.C(5, m.Eighth(), volume),
		gusic.B(5, m.Eighth(), volume),
		gusic.C(5, m.Eighth(), volume),
		gusic.B(5, m.Eighth(), volume),
		//
		//
	)

	return *m
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("gusic: File name argument required\nUsage: %s <filename>", os.Args[0])
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
