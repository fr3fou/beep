package gusic

// NewADSR is a constructor for an ADSR
func NewADSR(attack, attackMultiplier, decay, decayMultiplier, sustain, release, releaseMultiplier float64) *ADSR {
	return &ADSR{
		attackRatio:       attack,
		attackMultiplier:  attackMultiplier,
		decayRatio:        decay,
		decayMultiplier:   decayMultiplier,
		sustainRatio:      sustain,
		releaseRatio:      release,
		releaseMultiplier: releaseMultiplier,
		currentMultiplier: 0.0,
	}
}

func (a *ADSR) attack(notes []Note) {
	step := a.attackMultiplier / float64(len(notes))

	for i := range notes {
		notes[i].Volume *= a.currentMultiplier
		a.currentMultiplier += step
	}
}

func (a *ADSR) decay(notes []Note) {
	step := a.decayMultiplier / float64(len(notes))

	for i := range notes {
		notes[i].Volume *= a.currentMultiplier
		a.currentMultiplier -= step
	}
}

func (a *ADSR) sustain(notes []Note) {
	for i, x := range notes {
		notes[i].Volume = x.Volume
	}
}

func (a *ADSR) release(notes []Note) {
	step := a.releaseMultiplier / float64(len(notes))

	for i := range notes {
		notes[i].Volume *= a.currentMultiplier
		a.currentMultiplier -= step
	}
}

func (a *ADSR) Apply(notes []Note) {
	length := len(notes)

	attackEnd := int(a.attackRatio * float64(length))
	decayEnd := int(a.decayRatio*float64(length)) + attackEnd
	sustainEnd := int(a.sustainRatio*float64(length)) + decayEnd

	a.attack(notes[:attackEnd])
	a.decay(notes[attackEnd:decayEnd])
	a.sustain(notes[decayEnd:sustainEnd])
	a.release(notes[sustainEnd:])
}
