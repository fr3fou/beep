# gusic

go music.

## TODO

- [ ] Refactor to idiomatic go code
- [ ] "Context" for making notes of the same octave / noteduration / volume

  ```go
  m.AddNotes(
      // "Encloses" the octave, so you only have to provide the duration and volume
      mego.OctaveScope(4, func (ctx *mego.OctaveContext) []mego.Note {
          return []Note{
              ctx.A(mego.Quaver, 0.05),
              ctx.B(mego.Quaver, 0.05),
              ctx.D(mego.Quaver, 0.05),
              ctx.FS(mego.Quaver, 0.05),
              ctx.A(mego.Quaver, 0.05),
              // Encloses the volume
              ctx.VolumeScope(0.05, func (ctx *mego.VolumeContext) []mego.Note {
                  return []Note{
                      ctx.A(mego.Quaver),
                      ctx.B(mego.Quaver),
                      ctx.DS(mego.Quaver),
                      ctx.FS(mego.Quaver),
                  }
              })
          }
      })
  )
  ```

- [x] Rename to a proper name (?)
- [ ] Make examples
- [ ] Implement a simple format ~~- mp3 /~~ wav
- [x] Define all notes in all octaves a seperate file
- [ ] GUI + keybinds
- [ ] Play multiple notes at once
  - `f(x) + g(x)` where `f` and `g` produce different notes, should result in them playing at the same time
- [ ] Support for merging melodies / have melodies with differing BPM

  ```go
  m := mego.NewMelody(...)
  m.AddNotes(...)
  n := mego.NewMelody(...)
  n.AddNotes(...)
  o := m.Merge(n) // combines both melodies and makes a longer one
  ```

## How to play music

```console
ffplay megolovania.bin -autoexit -showmode 1 -f f64le -ar 48000
```

## Shoutouts

[viktordanov](https://github.com/viktordanov)

## Resources Used

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
