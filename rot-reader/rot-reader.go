package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("ÖLbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func (rot rot13Reader) Read(b []byte) (int, error) {

	n, err := rot.r.Read(b)
	if err != nil {
		return n, err
	}

	//for i := 0; i < n; i++ {
	for index, byteWert := range b[:n] {
		runeWert, _ := utf8.DecodeRune([]byte{byteWert})
		stringWert := fmt.Sprintf("%c: ", runeWert)
		konvWert := strings.Map(convert(runeWert), stringWert)
		fmt.Println("Konv: ", konvWert)
		b[index] = []byte(konvWert)[0]

	}
	//TODO Hier muss ein BREAK rein, da er sonst 21x konvertiert
	//break

	//}

	return n, nil
}

func convert(r rune) func(r rune) rune {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}

		return r
	}
	return rot13
}
