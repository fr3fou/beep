# megolovania

Megalovania in Go

Run with `ffplay megolovania.bin -autoexit -showmode 1 -f f64le -ar 48000` or by running `make` or just `make play` if the file already exists

## TODO

- [ ] Refactor to idiomatic go code
- [ ] Rename to a proper name
- [ ] Make examples
- [ ] Implement a simple format - mp3 / wav / something simpler?
- [ ] Define all notes in all octaves a seperate file

## Shoutouts

[viktordanov](https://github.com/viktordanov)

## Resources Used

- https://pages.mtu.edu/~suits/NoteFreqCalcs.html
- https://medium.com/@MicroPyramid/understanding-audio-quality-bit-rate-sample-rate-14286953d71f
- https://www.youtube.com/watch?v=FYTZkE5BZ-0
- https://www.reddit.com/r/explainlikeimfive/comments/4d4krv/eli5_what_is_the_difference_between_sample_rate/
- https://www.geogebra.org/graphing
- https://golang.org/pkg/encoding/binary
- https://www.wikiwand.com/en/A440_(pitch_standard)
- https://mathbitsnotebook.com/Algebra1/FunctionGraphs/FNGContinuousDiscrete.html
- https://xiph.org/video/vid2.shtml
- https://wiki.multimedia.cx/index.php/PCM
