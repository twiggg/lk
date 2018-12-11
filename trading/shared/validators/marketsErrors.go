package validators

import (
	"fmt"

	"twiggg.com/lk/trading/shared/format"
	langs "twiggg.com/lk/trading/shared/langs/v1.0.0"
)

//NewErrorInvalidCandlesInterval error
func NewErrorInvalidCandlesInterval(lang string, exch string, minutes int) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("taille de chandelles '%dmn' invalide, valeurs permises %s", minutes, format.ConcatenateStrings(listValidCandlesIntervals(exch), "[", "]", ", "))
	case langs.ENG:
		return fmt.Errorf("invalid candles interval '%d mn', expected values %s", minutes, format.ConcatenateStrings(listValidCandlesIntervals(exch), "[", "]", ", "))
	default:
		return fmt.Errorf("invalid candles interval '%d mn', expected values %s", minutes, format.ConcatenateStrings(listValidCandlesIntervals(exch), "[", "]", ", "))
	}
}

//NewErrorCandlesIntervalTooLow error
func NewErrorCandlesIntervalTooLow(lang string, interval int, mini int) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("taille de chandelles '%dmn' invalide, doit être supérieure ou égale à %dmn", interval, mini)
	case langs.ENG:
		return fmt.Errorf("candles interval '%dmn' too short, should be at least %dmn", interval, mini)
	default:
		return fmt.Errorf("candles interval '%dmn' too short, should be at least %dmn", interval, mini)
	}
}

//NewErrorInvalidPair error
func NewErrorInvalidPair(lang string, exch string, pair string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("la paire '%s' est invalide pour la plateforme d'échange %s, veuillez vous référer à la documentation pour connaitre les valeurs permises", pair, exch)
	case langs.ENG:
		return fmt.Errorf("invalid pair '%s' for exchange platform '%s', see documentation for valid pairs", pair, exch)
	default:
		return fmt.Errorf("invalid pair '%s' for exchange platform '%s', see documentation for valid pairs", pair, exch)
	}
}

//NewErrorInvalidWatchInterval error
func NewErrorInvalidWatchInterval(lang string, minutes int) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("interval de test '%dmn' invalide, valeurs permises %s", minutes, format.ConcatenateStrings(listValidWatchIntervals(), "[", "]", ", "))
	case langs.ENG:
		return fmt.Errorf("invalid tests interval '%d mn', expected values %s", minutes, format.ConcatenateStrings(listValidWatchIntervals(), "[", "]", ", "))
	default:
		return fmt.Errorf("invalid tests interval '%d mn', expected values %s", minutes, format.ConcatenateStrings(listValidWatchIntervals(), "[", "]", ", "))
	}
}

//NewErrorWatchIntervalTooLow error
func NewErrorWatchIntervalTooLow(lang string, interval int, mini int) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("fréquence de surveillance '%dmn' trop courte, doit être supérieure ou égale à %dmn", interval, mini)
	case langs.ENG:
		return fmt.Errorf("test frequency '%dmn' too short, should be at least %dmn", interval, mini)
	default:
		return fmt.Errorf("test frequency '%dmn' too short, should be at least %dmn", interval, mini)
	}
}

//NewErrorInvalidExchange error
func NewErrorInvalidExchange(lang string, exch string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("plateforme d'échange '%s' invalide, valeurs permises %s", exch, format.ConcatenateStrings(listValidExchanges(), "[", "]", ", "))
	case langs.ENG:
		return fmt.Errorf("invalid exchange platform '%s', expected values %s", exch, format.ConcatenateStrings(listValidExchanges(), "[", "]", ", "))
	default:
		return fmt.Errorf("invalid exchange platform '%s', expected values %s", exch, format.ConcatenateStrings(listValidExchanges(), "[", "]", ", "))
	}
}

//NewErrorInvalidMarketFormat error
func NewErrorInvalidMarketFormat(lang string, instru, format string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("market name '%s' invalide, format attendu '%s'", instru, format)
	case langs.ENG:
		return fmt.Errorf("invalid market name '%s', expected format '%s'", instru, format)
	default:
		return fmt.Errorf("invalid market name '%s', expected format '%s'", instru, format)
	}
}

//NewErrorInvalidPairFormat error
func NewErrorInvalidPairFormat(lang string) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("paire invalide: format attendu '<BASE>%s<QUOTE>', par exemple 'NEO%sBTC'", PairSeparator, PairSeparator)
	case langs.ENG:
		return fmt.Errorf("invalid pair: expected format '<BASE>%s<QUOTE>', such as 'NEO%sBTC'", PairSeparator, PairSeparator)
	default:
		return fmt.Errorf("invalid pair: expected format '<BASE>-%sQUOTE>', such as 'NEO%sBTC'", PairSeparator, PairSeparator)
	}
}
