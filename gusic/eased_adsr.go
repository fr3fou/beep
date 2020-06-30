package gusic

// NewEasedADSR is a constructor for an ADSR
func NewEasedADSR(easer Easer, ratios ADSRRatios, attackMultiplier, decayMultiplier float64) ADSR {
	return &EasedADSR{
		easing:            easer,
		ratios:            ratios,
		attackMultiplier:  attackMultiplier,
		decayMultiplier:   decayMultiplier,
		sustainMultiplier: attackMultiplier - decayMultiplier,
	}
}

// Attack is the attack stage
func (a *EasedADSR) Attack(samples []float64) {
	var t float64
	for i := range samples {
		t = float64(i) / float64(len(samples))
		samples[i] *= a.attackMultiplier * a.easing(t)
	}
}

// Decay is the decay stage
func (a *EasedADSR) Decay(samples []float64) {
	var t float64
	for i := range samples {
		t = float64(i) / float64(len(samples))
		samples[i] *= a.attackMultiplier - a.decayMultiplier*a.easing(t)
	}
}

// Sustain is the sustain stage
func (a *EasedADSR) Sustain(samples []float64) {
	for i, x := range samples {
		samples[i] = x * a.sustainMultiplier
	}
}

// Release is the release stage
func (a *EasedADSR) Release(samples []float64) {
	var t float64

	for i := range samples {
		t = float64(i) / float64(len(samples))
		samples[i] *= a.sustainMultiplier - a.sustainMultiplier*a.easing(t)
	}
}

// GetRatios returns the ratios
func (a *EasedADSR) GetRatios() ADSRRatios {
	return a.ratios
}
