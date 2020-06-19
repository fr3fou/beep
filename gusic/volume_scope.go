package gusic

type VolumeContext struct {
	Volume float64
}

func VolumeScope(octave int, f func(*OctaveContext) []Note) []Note {
	return f(&OctaveContext{Octave: octave})
}

func (v *VolumeContext) A(octave int, duration NoteDuration) Note {
	return A(octave, duration, v.Volume)
}
