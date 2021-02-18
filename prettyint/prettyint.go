package prettyint

import (
	"bytes"
	"fmt"
	"strconv"
)

type Int int

func (i Int) Format(state fmt.State, verb rune) {
	var r bytes.Buffer
	s := strconv.Itoa(int(i))

	if s[0] == '-' {
		r.WriteString(s[:1])
		s = s[1:]
	}
	m := len(s) % 3
	r.WriteString(s[:m])
	for i, c := range s[m:] {
		if i%3 == 0 {
			if m != i {
				r.WriteRune(',')
			}
		}
		r.WriteRune(c)
	}
	state.Write(r.Bytes())
}
