package gusic

func identity(t float64) float64 {
	return t
}

// NewLinearADSR is a constructor for an ADSR
func NewLinearADSR(ratios ADSRRatios, attackMultiplier, decayMultiplier, releaseMultiplier float64) ADSR {
	return &EasedADSR{
		easing:            identity,
		ratios:            ratios,
		attackMultiplier:  attackMultiplier,
		decayMultiplier:   decayMultiplier,
		releaseMultiplier: releaseMultiplier,
	}
}
