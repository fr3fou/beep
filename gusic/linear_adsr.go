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
func (a *LinearADSR) Attack(notes []float64) {
	step := a.attackMultiplier / float64(len(notes))

	for i := range notes {
		notes[i] *= a.currentMultiplier
		a.currentMultiplier += step
	}
}

// Decay is the decay stage
func (a *LinearADSR) Decay(notes []float64) {
	step := a.decayMultiplier / float64(len(notes))

	for i := range notes {
		notes[i] *= a.currentMultiplier
		a.currentMultiplier -= step
	}
}

// Sustain is the sustain stage
func (a *LinearADSR) Sustain(notes []float64) {
	for i, x := range notes {
		notes[i] = x
	}
}

// Release is the release stage
func (a *LinearADSR) Release(notes []float64) {
	step := a.releaseMultiplier / float64(len(notes))

	for i := range notes {
		notes[i] *= a.currentMultiplier
		a.currentMultiplier -= step
	}
}

// GetRatios returns the ratios
func (a *LinearADSR) GetRatios() ADSRRatios {
	return a.ratios
}
