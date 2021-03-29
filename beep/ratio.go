package beep

// NewRatios is a constructor for ADSR ratios.
// Values should be 0.0-1.0 and should sum up to 1.0.
// They represent a percentage of what duration should each state take.
func NewRatios(attack, decay, sustain, release float64) ADSRRatios {
	if attack+decay+sustain+release != 1.0 {
		panic("beep: ratios should sum up to 1.0")
	}

	return ADSRRatios{
		attack, decay, sustain, release,
	}
}
