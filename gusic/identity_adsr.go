package gusic

// NewIdentityADSR is a constructor for a no-op ADSR
func NewIdentityADSR() ADSR {
	return &IdentityADSR{}
}

// Attack is the attack stage
func (a *IdentityADSR) Attack(samples []float64) {}

// Decay is the decay stage
func (a *IdentityADSR) Decay(samples []float64) {}

// Sustain is the sustain stage
func (a *IdentityADSR) Sustain(samples []float64) {}

// Release is the release stage
func (a *IdentityADSR) Release(samples []float64) {}

// Ratios returns the ratios
func (a *IdentityADSR) Ratios() ADSRRatios {
	return a.ratios
}
