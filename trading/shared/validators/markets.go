package validators

import (
	"fmt"
	"sort"
	"strings"

	binance "twiggg.com/lk/trading/core/wrappers/binance/v1.0.0"
	"twiggg.com/lk/trading/core/wrappers/bittrex/v1.0.0"
	gdax "twiggg.com/lk/trading/core/wrappers/gdax/v1.0.0"
	"twiggg.com/lk/trading/shared/format"
)

//constants
const (
	Bittrex       = "bittrex"
	Gdax          = "coinbase-pro"
	Binance       = "binance"
	PairSeparator = "-"
)

//IsValidExch checks if exch is a valid exchange name
func IsValidExch(exch string) bool {
	_, ok := validExch[exch]
	return ok
}

/*
//ValidateMarketFormat checks if the pair has valid format depending on exchange platform
func ValidateMarketFormat(lang string, exch, pair string) (string, string, error) {
	var err error
	base, quote := "", ""
	switch exch {
	case Bittrex:
		base, quote, err = bittrex.SplitMarketName(pair)
		if err != nil {
			return base, quote, NewErrorInvalidMarketFormat(lang, pair, bittrex.MarketNameFormat())
		}
	case Binance:
		base, quote, err = binance.SplitMarketName(pair)
		if err != nil {
			return base, quote, NewErrorInvalidMarketFormat(lang, pair, binance.MarketNameFormat())
		}
	case Gdax:
		base, quote, err = gdax.SplitMarketName(pair)
		if err != nil {
			return base, quote, NewErrorInvalidMarketFormat(lang, pair, gdax.MarketNameFormat())
		}
	default:
		return "", "", NewErrorInvalidExchange(lang, exch)
	}
	return base, quote, nil
}
*/

//FormatPairForExch formats the pair name based on exchange api format
func FormatPairForExch(exch, base, quote string) string {
	switch exch {
	case Bittrex:
		return bittrex.MarketName(base, quote)
	case Binance:
		return binance.MarketName(base, quote)
	case Gdax:
		return gdax.MarketName(base, quote)
	default:
		return fmt.Sprintf("%s-%s", base, quote)
	}
}

func listValidExchanges() []string {
	list := make([]string, len(validExch))
	i := 0
	for key := range validExch {
		list[i] = key
		i++
	}
	return list
}

//validateCandlesInterval checks if the candles interval is valid
func validateCandlesInterval(lang string, exch string, candleMinutes int) error {
	var err error
	itv := ""
	switch exch {
	case Bittrex:
		itv = bittrex.TranslateInterval(candleMinutes)
		itv, err = bittrex.ValidateInterval(itv)
		if err != nil {
			return NewErrorInvalidCandlesInterval(lang, exch, candleMinutes)
		}
	case Binance:
		itv = binance.TranslateInterval(candleMinutes)
		itv, err = binance.ValidateInterval(itv)
		if err != nil {
			return NewErrorInvalidCandlesInterval(lang, exch, candleMinutes)
		}
	case Gdax:
		itv = gdax.TranslateInterval(candleMinutes)
		itv, err = gdax.ValidateInterval(itv)
		if err != nil {
			return NewErrorInvalidCandlesInterval(lang, exch, candleMinutes)
		}
	default:
		return NewErrorInvalidExchange(lang, exch)
	}
	return nil
}

//IsValidWatchInterval validator
func IsValidWatchInterval(minutes int) bool {
	_, ok := validWatchIntervals[minutes]
	return ok
}

//IsValidWatchDur validator
func IsValidWatchDur(ndays int) bool {
	_, ok := validWatchDays[ndays]
	return ok
}

func listValidWatchIntervals() []string {
	return format.StringsFromInts(extractInts(validWatchIntervals))
}

//ListValidWatchMn returns the list of valid watchMn intervals (in minutes)
func ListValidWatchMn() []int {
	//exch = strings.ToLower(exch)
	res := []int{}
	/*switch exch {
	case Bittrex:
		for i := range validBittrexCandlesInterval {
			res = append(res, i)
		}
	case Binance:
		for i := range validBinanceCandlesInterval {
			res = append(res, i)
		}
	case Gdax:
		for i := range validGdaxCandlesInterval {
			res = append(res, i)
		}
	default:
		for i := range validCandlesInterval {
			res = append(res, i)
		}
	}
	*/
	for i := range validWatchIntervals {
		res = append(res, i)
	}
	sort.Ints(res)
	return res
}
func listValidWatchDays() []string {
	return format.StringsFromInts(extractInts(validWatchDays))
}

//ListValidWatchDays returns the list of valid watch duration (in days)
func ListValidWatchDays() []int {
	res := []int{}
	for i := range validWatchDays {
		res = append(res, i)
	}
	sort.Ints(res)
	return res
}

//IsValidCandlesInterval checks if valid interval
func IsValidCandlesInterval(exch string, minutes int) bool {
	switch exch {
	case Bittrex:
		_, ok := validBittrexCandlesInterval[minutes]
		return ok
	case Binance:
		_, ok := validBinanceCandlesInterval[minutes]
		return ok
	case Gdax:
		_, ok := validGdaxCandlesInterval[minutes]
		return ok
	default:
		_, ok := validCandlesInterval[minutes]
		return ok
	}

}

func extractInts(m map[int]string) []int {
	//fmt.Printf("map: %+v\n", m)
	list := make([]int, len(m))
	i := 0
	for k := range m {
		//fmt.Printf("add val %d: '%d'", i, k)
		list[i] = k
		i++
	}
	//fmt.Printf("res: %v\n", list)
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	return list
}

func listValidCandlesIntervals(exch string) []string {
	switch exch {
	case Bittrex:
		return format.StringsFromInts(extractInts(validBittrexCandlesInterval))
	case Binance:
		return format.StringsFromInts(extractInts(validBinanceCandlesInterval))
	case Gdax:
		return format.StringsFromInts(extractInts(validGdaxCandlesInterval))
	default:
		return format.StringsFromInts(extractInts(validCandlesInterval))
	}
}

//ListValidCandlesMn returns the this of valid candle intervals in minutes
func ListValidCandlesMn(exch string) []int {
	switch exch {
	case Bittrex:
		return extractInts(validBittrexCandlesInterval)
	case Binance:
		return extractInts(validBinanceCandlesInterval)
	case Gdax:
		return extractInts(validGdaxCandlesInterval)
	default:
		return extractInts(validCandlesInterval)
	}
}

//TranslateCandlesMn translates the candles minutes in exchange specific string
func TranslateCandlesMn(lang string, exch string, candlesMn int) (string, error) {
	switch exch {
	case Bittrex:
		str, ok := validBittrexCandlesInterval[candlesMn]
		if ok {
			return str, nil
		}
		return "", NewErrorInvalidCandlesInterval(lang, exch, candlesMn)
	case Binance:
		str, ok := validBinanceCandlesInterval[candlesMn]
		if ok {
			return str, nil
		}
		return "", NewErrorInvalidCandlesInterval(lang, exch, candlesMn)
	case Gdax:
		str, ok := validGdaxCandlesInterval[candlesMn]
		if ok {
			return str, nil
		}
		return "", NewErrorInvalidCandlesInterval(lang, exch, candlesMn)
	default:
		return "", NewErrorInvalidExchange(lang, exch)
	}

}

//TranslatePair generates the symbol
func TranslatePair(base, quote, exch string) string {
	exch = strings.ToLower(exch)
	switch exch {
	case Bittrex:
		return strings.ToLower(bittrex.MarketName(base, quote))
	case Binance:
		return strings.ToLower(binance.MarketName(base, quote))
	case Gdax:
		return strings.ToLower(gdax.MarketName(base, quote))
	default:
		return ""
	}
}

//ParseExchangeSymbol parses the symbol depending on the exchange's format. it returns base,quote,error
func ParseExchangeSymbol(lang, m, exch string) (string, string, error) {
	exch = strings.ToLower(exch)
	m = strings.ToLower(m)
	switch exch {
	case Bittrex:
		return bittrex.SplitMarketName(m)
	case Binance:
		return binance.SplitMarketName(m)
	case Gdax:
		return gdax.SplitMarketName(m)
	default:
		return "", "", NewErrorInvalidExchange(lang, exch)
	}
}

//WriteSymbol generate a standardized symbol designation from base and quote currencies
func WriteSymbol(base, quote string) string {
	return strings.ToLower(fmt.Sprintf("%s%s%s", base, PairSeparator, quote))
}

//ParsePair extract base and quote currencies from the pair. Valid format: BASE-QUOTE.
func ParsePair(lang string, pair string) (string, string, error) {
	parts := strings.Split(pair, PairSeparator)
	if len(parts) == 2 {
		base := strings.ToUpper(parts[0])
		quote := strings.ToUpper(parts[1])
		return base, quote, nil
	}

	return "", "", NewErrorInvalidPairFormat(lang)
}

//IsValidPair checks if the pair is valid. Option 1 -> all markets from the exchange, Option 2 -> restrict to markets to gain controle on the computation load.
func IsValidPair(exch string, pair string) bool {
	pair = strings.ToLower(pair)
	switch exch {
	case Bittrex:
		_, ok := validBittrexPairs[pair]
		return ok
	case Binance:
		_, ok := validBinancePairs[pair]
		return ok
	case Gdax:
		_, ok := validGdaxPairs[pair]
		return ok
	default:
		return false
	}
}
