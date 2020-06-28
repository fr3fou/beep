package gusic

// NewLinearADSR is a constructor for an ADSR
func NewLinearADSR(ratios ADSRRatios, attackMultiplier, decayMultiplier, releaseMultiplier float64) ADSR {
	return &LinearADSR{
		ratios:            ratios,
		attackMultiplier:  attackMultiplier,
		decayMultiplier:   decayMultiplier,
		releaseMultiplier: releaseMultiplier,
		currentMultiplier: 0.0,
	}
}

// Attack is the attack stage
func (a *LinearADSR) Attack(notes []Note) {
	step := a.attackMultiplier / float64(len(notes))

	for i := range notes {
		notes[i].Volume *= a.currentMultiplier
		a.currentMultiplier += step
	}
}

// Decay is the decay stage
func (a *LinearADSR) Decay(notes []Note) {
	step := a.decayMultiplier / float64(len(notes))

	for i := range notes {
		notes[i].Volume *= a.currentMultiplier
		a.currentMultiplier -= step
	}
}

// Sustain is the sustain stage
func (a *LinearADSR) Sustain(notes []Note) {
	for i, x := range notes {
		notes[i].Volume = x.Volume
	}
}

// Release is the release stage
func (a *LinearADSR) Release(notes []Note) {
	step := a.releaseMultiplier / float64(len(notes))

	for i := range notes {
		notes[i].Volume *= a.currentMultiplier
		a.currentMultiplier -= step
	}
}

// GetRatios returns the ratios
func (a *LinearADSR) GetRatios() ADSRRatios {
	return a.ratios
}
