package gusic

func NewADSR(attack, attackMultiplier, decay, decayMultiplier, sustain, release, releaseMultiplier float64) *ADSR {
	return &ADSR{
		attackRatio:       attack,
		attackMultiplier:  attackMultiplier,
		decayRatio:        decay,
		decayMultiplier:   decayMultiplier,
		sustainRatio:      sustain,
		releaseRatio:      release,
		releaseMultiplier: releaseMultiplier,
	}
}

func (a *ADSR) attack(notes []Note) {
	multiplier := a.attackMultiplier / float64(len(notes))

	for i := range notes {
		notes[i].Volume *= multiplier * float64(i)
	}
}

func (a *ADSR) decay(notes []Note) {
	multiplier := a.decayMultiplier / float64(len(notes))

	for i := range notes {
		notes[i].Volume *= multiplier * float64(i)
	}
}

func (a *ADSR) sustain(notes []Note) {
	for i, x := range notes {
		notes[i].Volume = x.Volume
	}
}

func (a *ADSR) release(notes []Note) {
	multiplier := a.releaseMultiplier / float64(len(notes))

	for i := range notes {
		notes[i].Volume *= multiplier * float64(i)
	}
}
func (a *ADSR) apply(notes []Note) {
	length := len(notes)

	attackEnd := int(a.attackRatio * float64(length))
	decayEnd := int(a.decayRatio*float64(length)) + attackEnd
	sustainEnd := int(a.sustainRatio*float64(length)) + decayEnd

	a.attack(notes[:attackEnd])
	a.decay(notes[attackEnd:decayEnd])
	a.sustain(notes[decayEnd:sustainEnd])
	a.release(notes[sustainEnd:])
}
