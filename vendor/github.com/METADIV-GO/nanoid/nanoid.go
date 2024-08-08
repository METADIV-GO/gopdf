package nanoid

import (
	"strings"

	gonanoid "github.com/matoous/go-nanoid"
)

type Opts struct {
	Numbers   bool
	Lowercase bool
	Uppercase bool
	/*
		Symbols are "_-"
	*/
	Symbols bool
	/*
		LookAlike are "1lI0Oouv5Ss"
	*/
	ExcludeAlike bool
	Length       int
}

/*
New returns a random string with the given options.
*/
func New(opts Opts) string {
	chars := ""
	if opts.Numbers {
		chars += CHARACTERS_NUMBERS
	}
	if opts.Lowercase {
		chars += CHARACTERS_LOWERCASE
	}
	if opts.Uppercase {
		chars += CHARACTERS_UPPERCASE
	}
	if opts.Symbols {
		chars += CHARACTERS_SYMBOLS
	}
	if opts.ExcludeAlike {
		for _, c := range CHARACTERS_LOOK_ALIKE {
			chars = strings.ReplaceAll(chars, string(c), "")
		}
	}
	if chars == "" {
		panic("no chars")
	}
	if opts.Length == 0 {
		panic("length is zero")
	}
	s, err := gonanoid.Generate(chars, opts.Length)
	if err != nil {
		panic(err)
	}
	return s
}

/*
NewSafe quickly returns a random string with the safe options.
*/
func NewSafe() string {
	return New(Opts{
		Numbers:      true,
		Lowercase:    true,
		Uppercase:    true,
		Symbols:      false,
		ExcludeAlike: true,
		Length:       SAFE_LENGTH,
	})
}
