package langs

import (
	"strings"
)

//language constants
const (
	FR  = "fr"
	ENG = "eng"
	Dft = "eng"
)

//Validate validates the lang value
func Validate(lang string) string {
	lang = strings.ToLower(lang)
	switch lang {
	case FR:
		return lang
	case ENG:
		return lang
	default:
		return Dft
	}
}
