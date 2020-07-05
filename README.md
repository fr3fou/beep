# gusic

go music.

## TODO

- [x] Refactor to idiomatic go code
- [x] Rename to a proper name (?)
- [x] Make examples
- [x] Make note durations methods on `gusic.Melody` (as they are dependant on the note length and BPM)
- [ ] Implement a simple format - ~~mp3~~ / wav
- [x] Define all notes in all octaves a seperate file
- [ ] GUI + keybinds
- [ ] Play multiple notes at once (chords)
  - `f(x) + g(x)` where `f` and `g` produce different notes, should result in them playing at the same time
  - [ ] Note -> SingleNote, add Chord (multiple notes). Both implement new Note interface
  ```go
  type Note interface {
      Compute(sampleRate float64) []float64
  }

  ```
- [x] Support for merging melodies (playing 2 melodies at the same time)
  ```go
  m := mego.NewMelody(...)
  m.AddNotes(...)
  
  // or
  m.Runs[0].AddNotes(...)

  // more runs, i.e. staves 
  m.NewRun(notes...)
  ```
- [ ] Support for concatenating melodies / have melodies with differing BPM
  ```go
  m := mego.NewMelody(...)
  m.AddNotes(...)
  n := mego.NewMelody(...)
  n.AddNotes(...)
  o := m.Concat(n) // combines both melodies and makes a longer one
  ```
- [ ] Dual Channel support (left and right ear)
- [x] ~~Fix clipping in release of linear ADSR~~  _(theoretically negligible)_
- [x] Implement logarithmic / exp ADSR


## How to play music

```console
ffplay megolovania.bin -autoexit -showmode 1 -f f64le -ar 48000
```

## Shoutouts

[viktordanov](https://github.com/viktordanov)

## References

- <https://pages.mtu.edu/~suits/NoteFreqCalcs.html>
- <https://medium.com/@MicroPyramid/understanding-audio-quality-bit-rate-sample-rate-14286953d71f>
- <https://www.youtube.com/watch?v=FYTZkE5BZ-0>
- <https://www.reddit.com/r/explainlikeimfive/comments/4d4krv/eli5_what_is_the_difference_between_sample_rate/>
- <https://www.geogebra.org/graphing>
- <https://golang.org/pkg/encoding/binary>
- <https://www.wikiwand.com/en/A440_(pitch_standard)>
- <https://www.wikiwand.com/en/Envelope_(music)>
- <https://mathbitsnotebook.com/Algebra1/FunctionGraphs/FNGContinuousDiscrete.html>
- <https://xiph.org/video/vid2.shtml>
- <https://wiki.multimedia.cx/index.php/PCM>
- <https://stackoverflow.com/questions/1073606/is-there-a-one-line-function-that-generates-a-triangle-wave>
- <https://www.desmos.com/calculator/nduy9l2pez>
- <https://dsp.stackexchange.com/questions/2555/help-with-equations-for-exponential-adsr-envelope>
- <http://web.mit.edu/6.111/www/f2004/projects/ghs_report.pdf>
- <http://www.martin-finke.de/blog/articles/audio-plugins-011-envelopes/>
- <https://www.geogebra.org/graphing/ehmusxyd>
