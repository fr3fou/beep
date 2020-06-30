package gusic

// NewEasedADSR is a constructor for an ADSR
func NewEasedADSR(easer Easer, ratios ADSRRatios, attackMultiplier, decayMultiplier, releaseMultiplier float64) ADSR {
	return &EasedADSR{
		easing:            easer,
		ratios:            ratios,
		attackMultiplier:  attackMultiplier,
		decayMultiplier:   decayMultiplier,
		releaseMultiplier: releaseMultiplier,
		currentMultiplier: 0.0,
	}
}

// Attack is the attack stage
func (a *EasedADSR) Attack(samples []float64) {
	for i := range samples {
		samples[i] *= a.attackMultiplier * a.easing(float64(i)/float64(len(samples)))
	}
}

// Decay is the decay stage
func (a *EasedADSR) Decay(samples []float64) {
	for i := range samples {
		samples[i] *= a.attackMultiplier - a.decayMultiplier*a.easing(float64(i)/float64(len(samples)))
	}
}

// Sustain is the sustain stage
func (a *EasedADSR) Sustain(samples []float64) {
	for i, x := range samples {
		samples[i] = x
	}
}

// Release is the release stage
func (a *EasedADSR) Release(samples []float64) {
	for i := range samples {
		samples[i] *= 1 - a.releaseMultiplier*a.easing(float64(i)/float64(len(samples)))
	}
}

// GetRatios returns the ratios
func (a *EasedADSR) GetRatios() ADSRRatios {
	return a.ratios
}
