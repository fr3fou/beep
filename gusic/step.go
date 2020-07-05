package gusic

// 12th root of 2
const twelfthrootof2 float64 = 1.059463094359

const (
	c  = -9
	cs = -8
	d  = -7
	ds = -6
	e  = -5
	f  = -4
	fs = -3
	g  = -2
	gs = -1
	a  = 0
	as = 1
	b  = 2
)

// A is A
func A(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+a, duration, volume)
}

// AS is A#
func AS(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+as, duration, volume)
}

// B is B
func B(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+b, duration, volume)
}

// C is C
func C(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+c, duration, volume)
}

// CS is C#
func CS(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+cs, duration, volume)
}

// D is D
func D(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+d, duration, volume)
}

// DS is D#
func DS(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+ds, duration, volume)
}

// E is E
func E(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+e, duration, volume)
}

// F is F
func F(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+f, duration, volume)
}

// FS is F#
func FS(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+fs, duration, volume)
}

// G is G
func G(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+g, duration, volume)
}

// GS is G#
func GS(octave int, duration NoteDuration, volume float64) SingleNote {
	return NewNote((octave-4)*12+gs, duration, volume)
}
