package gusic

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"time"
)

// NewMelody is a constructor for a melody
func NewMelody(sampleRate float64, bpm int, noteLength int, generator Generator, adsr ADSR) *Melody {
	m := &Melody{
		SampleRate: sampleRate,
		BPM:        bpm,
		NoteLength: noteLength,
		Envelope:   adsr,
		Generator:  generator,
		Runs:       []Run{{}},
	}

	m.calculateDurations()

	return m
}

func (m *Melody) calculateDurations() {
	m.breve = time.Minute / NoteDuration(m.BPM) * NoteDuration(m.NoteLength) * 2
	m.semibreve = m.breve / 2
	m.minim = m.semibreve / 2
	m.crotchet = m.semibreve / 4
	m.quaver = m.semibreve / 8
	m.semiquaver = m.semibreve / 16
	m.demisemiquaver = m.semibreve / 32
}

// AddNote adds a note
func (m *Melody) AddNote(note Note) {
	m.Runs[0].AddNote(note)
}

// AddNotes adds notes
func (m *Melody) AddNotes(notes ...Note) {
	m.Runs[0].AddNotes(notes...)
}

// NewRun creates and accepts notes for a new run
func (m *Melody) NewRun(notes ...Note) {
	m.Runs = append(m.Runs, Run{Notes: notes})
}

// ApplyADSR applies all the stages of an ADSR to an array of notes
func (m *Melody) applyADSR(noteSamples []float64) {
	length := len(noteSamples)

	ratios := m.Envelope.GetRatios()

	attackEnd := int(ratios.AttackRatio * float64(length))
	decayEnd := int(ratios.DecayRatio*float64(length)) + attackEnd
	sustainEnd := int(ratios.SustainRatio*float64(length)) + decayEnd

	m.Envelope.Attack(noteSamples[:attackEnd])
	m.Envelope.Decay(noteSamples[attackEnd:decayEnd])
	m.Envelope.Sustain(noteSamples[decayEnd:sustainEnd])
	m.Envelope.Release(noteSamples[sustainEnd:])
}

func (m *Melody) compute() []float64 {
	runSamples := [][]float64{}

	longestSample := 0.0
	for i, run := range m.Runs {
		runSamples = append(runSamples, []float64{})
		for _, note := range run.Notes {
			samples := []float64{}

			for j := 1.0; j <= math.Ceil(m.SampleRate*note.Duration.Seconds()); j++ {
				val := m.Generator(j*note.Frequency*2*math.Pi/m.SampleRate) * note.Volume
				samples = append(samples, val)
			}

			m.applyADSR(samples)
			runSamples[i] = append(runSamples[i], samples...)
		}
		longestSample = math.Max(longestSample, float64(len(runSamples[i])))
	}

	// Combine raw wave data with basic addition
	finalSamples := make([]float64, int64(longestSample))

	for i := int64(0); i < int64(longestSample); i++ {
		for _, sample := range runSamples {
			if i < int64(len(sample)) {
				finalSamples[i] += sample[i]
			}
		}
	}

	return finalSamples
}

// PCM generates the PCM
func (m *Melody) PCM() ([]byte, error) {
	buf := &bytes.Buffer{}

	if err := binary.Write(buf, binary.LittleEndian, m.compute()); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *Melody) Write(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, m.compute())
}
