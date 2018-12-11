package format

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/jinzhu/now"

	"twiggg.com/lk/trading/shared/langs"
)

//Market returns the name from base and quote currencies
func Market(base, quote string) string {
	return fmt.Sprintf("%s-%s", base, quote)
}

//CompileMarketName compiles an market ID from its definition
func CompileMarketName(exch, pair string, candlesMn int) string {
	return fmt.Sprintf("%s_%s_%d", exch, pair, candlesMn)
}

//DecomposeMarketName extracts an market's definition from compiled name
func DecomposeMarketName(lang string, keyString string) (string, string, int, error) {
	parts := strings.Split(keyString, "_")
	if len(parts) != 3 {
		switch lang {
		case langs.FR:
			return "", "", 0, fmt.Errorf("clé d'instrument '%s' invalide: doit être au format 'exchange_instrument_candlesMn'", keyString)
		case langs.ENG:
			return "", "", 0, fmt.Errorf("invalid instrument keyString '%s': must have format 'exch_instrument_candlesMn'", keyString)
		default:
			return "", "", 0, fmt.Errorf("invalid instrument keyString '%s': must have format 'exch_instrument_candlesMn'", keyString)
		}

	}
	candlesMn, err := strconv.Atoi(parts[2])
	if err != nil {
		switch lang {
		case langs.FR:
			return "", "", 0, fmt.Errorf("clé d'instrument '%s' invalide: candlesMn doit être un entier positif", keyString)
		case langs.ENG:
			return "", "", 0, fmt.Errorf("invalid instrument keyString '%s': candlesMn must be a positive integer", keyString)
		default:
			return "", "", 0, fmt.Errorf("invalid instrument keyString '%s': candlesMn must be a positive integer", keyString)
		}

	}
	if candlesMn < 0 {
		switch lang {
		case langs.FR:
			return "", "", 0, fmt.Errorf("clé d'instrument '%s' invalide: candlesMn doit être un entier positif", keyString)
		case langs.ENG:
			return "", "", 0, fmt.Errorf("invalid instrument keyString '%s': candlesMn must be a positive integer", keyString)
		default:
			return "", "", 0, fmt.Errorf("invalid instrument keyString '%s': candlesMn must be a positive integer", keyString)
		}
	}
	return parts[0], parts[1], candlesMn, nil
}

//DayStart computes the 00.00 of t in UTC
func DayStart(t time.Time) time.Time {
	return now.New(t.UTC()).BeginningOfDay()
}

//DayEnd computes the 23.59 of t in UTC
func DayEnd(t time.Time) time.Time {
	return now.New(t.UTC()).EndOfDay()
}

//NextDayEnd computes next day's end in UTC
func NextDayEnd(t time.Time) time.Time {
	return now.New(t.UTC().Add(time.Hour * 24)).EndOfDay()
}

//NextDayStart computes next day's start in UTC
func NextDayStart(t time.Time) time.Time {
	return now.New(t.UTC().Add(time.Hour * 24)).BeginningOfDay()
}

//MillisTimestamp returns the timestamp in ms
func MillisTimestamp(t time.Time) int64 {
	nano := int64(t.Nanosecond())
	sec := t.Unix()
	ms := sec*1000 + nano/1000000
	return ms
}

//constant
const (
	DateLayout = "January _2, 2006" //"Jan _2, 2006"
	FormatDate = "02/01/2006"
)

//FromMillisTimestamp returns the time from a timestamp in ms
func FromMillisTimestamp(millis int64) time.Time {
	sec := int64(millis / int64(1000))
	nano := int64(millis-sec*1000) * 1000000
	return time.Unix(sec, nano)
}

//IntString takes an int and returns a string representing that int with at list n digits
func IntString(val int64, n int) string {
	if n <= 0 {
		return fmt.Sprintf("%d", val)
	}
	ref := int64(math.Pow10(n))
	if val >= ref {
		return fmt.Sprintf("%d", val)
	}
	val += ref
	return string(fmt.Sprintf("%d", val)[1:])
}

//DateYYYYMMDD returns a date from time
func DateYYYYMMDD(t time.Time, utc bool) int64 {
	if utc {
		t = t.UTC()
	}
	y, m, d := t.Date()
	date := int64(y)*10000 + int64(m)*100 + int64(d)
	return int64(date)
}

//FromYYYYMMDD extracts time from a date string
func FromYYYYMMDD(date string) (time.Time, error) {
	//expInt := y*10000 + m*100 + 01
	return time.Parse("20060102", date)
}

func round(x float64) float64 {
	return float64(int(x + 0.5))
}

//FloatToFixedPrecision formats a float to a fixed precision number
func FloatToFixedPrecision(f float64, precision int) float64 {
	if precision <= 0 {
		precision = 0
	}
	//fmt.Printf("FloatToFixedPrecision: f: %f\n", f)
	exp := math.Pow(10, float64(precision))
	//f = math.Round(f*exp) / exp
	sign := 1.0000000
	if f < 0 {
		sign = -1.0000000
		f = sign * f
	}
	f = round(f*exp) / exp
	f = sign * f
	//fmt.Printf("FloatToFixedPrecision: math.Round: %f\n", f)
	/*s := fmt.Sprintf("%s.%df", "%", precision)
	fmt.Printf("FloatToFixedPrecision: format string: '%s'\n", s)
	s=fmt.Sprintf("%s",)*/
	s := FloatToFixedString(f, precision)
	//fmt.Printf("FloatToFixedPrecision:formated: '%s'\n", s)
	fl, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("FloatToFixedPrecision: failed to parse float: %s\n", err.Error())
		return f
	}
	//fmt.Printf("FloatToFixedPrecision:parsed: %f\n", fl)
	return fl
}

//FloatToFixedString returns the string a formated string. max 8 digits.
func FloatToFixedString(f float64, precision int) string {
	s := fmt.Sprintf("%.8f", f)
	iDot := strings.Index(s, ".")
	s = fmt.Sprintf("%s", s[:iDot+precision+1])
	return s
}

//ToJSONStrings marshals a []string to []byte
func ToJSONStrings(list []string) []byte {
	for i := range list {
		list[i] = fmt.Sprintf(`"%s"`, list[i])
	}
	ConcatenateStrings(list, "[", "]", ",")
	return []byte(ConcatenateStrings(list, "[", "]", ","))
}

//ConcatenateStrings returns a string from a slice of strings
func ConcatenateStrings(list []string, start, end, sep string) string {
	return fmt.Sprintf("%s%s%s", start, strings.Join(list, sep), end)
}

//ConcatenateMapStrings returns a string from a map of strings
func ConcatenateMapStrings(m map[string]string, start, end, sep string) string {
	list := []string{}
	for key, val := range m {
		list = append(list, fmt.Sprintf("%s: %s", key, val))
	}
	return ConcatenateStrings(list, start, end, sep)
}

//WhichPrecision returns the precision of the value (2 for fiat currencies)
func WhichPrecision(quote string) int {
	quote = strings.ToLower(quote)
	switch quote {
	case "EUR":
		return 2
	case "USD":
		return 3
	case "USDT":
		return 8
	case "GBP":
		return 2
	case "JPY":
		return 2
	case "CNH":
		return 2
	case "CNY":
		return 2
	case "KPW":
		return 2
	case "KRW":
		return 2
	case "CAD":
		return 2
	case "MXN":
		return 2
	case "RUB":
		return 2
	case "INR":
		return 2
	case "ILS":
		return 2
	case "BRL":
		return 2
	case "AUD":
		return 2
	case "SGD":
		return 2
	case "TWD":
		return 2
	default:
		return 8
	}
}

//StringsFromInts returns a []string from a []int
func StringsFromInts(vals []int) []string {
	list := make([]string, len(vals))
	//i := 0
	s := ""
	//fmt.Printf("StringsFromInts on %v\n", vals)
	for i := range vals {
		val := vals[i]
		if val < 10 {
			s = fmt.Sprintf("00%d", val)
		} else if val < 100 {
			s = fmt.Sprintf("0%d", val)
		} else {
			s = fmt.Sprintf("%d", val)
		}
		//fmt.Printf("val %d: %d -> '%s'\n", i, val, s)
		list[i] = s
		//i++
	}
	//fmt.Printf("res: %v\n", list)
	return list
}

//IntSliceFromMap returns a []int  a map[int]struct{}
func IntSliceFromMap(m map[int]struct{}) []int {
	list := make([]int, len(m))
	i := 0
	for val := range m {
		list[i] = val
		i++
	}
	return list
}

func needCapital(str string) bool {
	if len(str) != 1 {
		return false
	}
	if str == "." || str == "?" || str == "!" {
		return true
	}
	return false
}

func isPunct(str string) bool {
	if len(str) != 1 {
		return false
	}
	if str == "." || str == "?" || str == "," || str == "!" || str == ":" || str == ";" {
		return true
	}
	return false
}

func isParenth(str string) bool {
	if len(str) != 1 {
		return false
	}
	if str == "(" || str == ")" || str == "[" || str == "]" || str == "{" || str == "}" {
		return true
	}
	return false
}

func isOpenParenth(str string) bool {
	if len(str) != 1 {
		return false
	}
	if str == "(" || str == "[" || str == "{" {
		return true
	}
	return false
}

func isCloseParenth(str string) bool {
	if len(str) != 1 {
		return false
	}
	if str == ")" || str == "]" || str == "}" {
		return true
	}
	return false
}

//Paragraph formats a string to a clean paragraph. Forces capital letters for 1st letters after a point and removes duplicate spaces
func Paragraph(str string) string {
	res := []rune{}
	//iLast := len(str) - 1
	//iFirst := 0
	lastIsSpace := false
	lastIsPunct := false
	space := " "
	dot := "."
	//interro := "?"
	//excla := "!"
	//comma := ","
	//column := ":"
	/*apost := "'"
	paren1 := "("
	paren2 := ")"
	croch1 := "["
	croch2 := "]"
	acco1 := "{"
	acco2 := "}"*/
	ndots := 0
	s := ""
	n := 0
	var r rune
	lastPunct := ""
	//last:=""
	for i := range str {
		s = string(str[i])
		r = unicode.ToLower(rune(str[i]))
		//fmt.Printf("%d '%s'\n", i, s)
		if s == space {
			if lastIsSpace || n == 0 {
				//clear duplicate spaces
				//fmt.Printf("duplicate space\n")
				continue
			}
			lastIsSpace = true
			//fmt.Printf("last is space\n")
			//lastIsPunct = false
		} else if unicode.IsPunct(r) {
			//fmt.Printf("unicode.IsPunct\n")
			/*if s == `/` || s == `\` {
				n++
				res = append(res, r)
				continue
			}*/
			lastPunct = s
			//lastIsPunct = true

			//if s == apost || s== {
			if !isPunct(s) {
				/*lastPunct=s
				lastIsPunct=true
				lastIsSpace=false*/
				lastIsPunct = true
				if isOpenParenth(s) {
					//if !lastIsSpace {
					n++
					res = append(res, rune(space[0]))
					//}
					n++
					res = append(res, r)
					n++
					/*res = append(res, rune(space[0]))
					lastIsSpace = false
					lastIsPunct = false*/
					lastIsSpace = false
					continue
				} else if isCloseParenth(s) {
					n++
					res = append(res, rune(space[0]))
					n++
					res = append(res, r)
					/*lastIsPunct = false
					lastIsSpace = true*/
					lastIsSpace = false
					continue
				} else {
					//fmt.Printf("else\n")

					if s == `/` || s == `\` {
						//fmt.Printf("slash or backslash: '%s'\n", s)
						if lastIsSpace {
							//fmt.Printf("last is space\n")
							n++
							res = append(res, rune(space[0]))
						}
						lastIsPunct = false
						lastIsSpace = false
						lastPunct = ""
					}
					n++
					res = append(res, r)
					lastIsSpace = false
					continue
				}
			}
			//fmt.Printf("is punct: %s\n", s)
			if (s == dot && ndots >= 3) || (s != dot && lastIsPunct) || n == 0 {
				//fmt.Printf("s:'%s' ndots:%d lastIsPunct:%v n:%d\n", s, ndots, lastIsPunct, n)
				//fmt.Printf("jump\n")
				continue
			}

			if s == dot {
				ndots++
			}
			n++
			res = append(res, r)
			lastIsPunct = true
			lastIsSpace = false
			lastPunct = s
		} else {
			ndots = 0
			if lastIsPunct {
				//force space before new first letter
				res = append(res, rune(space[0]))
				//fmt.Printf("isComma: %v, isColumn: %v\n",s==comma, s==column)
				//if lastPunct != comma && lastPunct != column {
				if needCapital(lastPunct) {
					r = unicode.ToTitle(r)
				}
			} else if lastIsSpace {
				res = append(res, rune(space[0]))
			}
			if n == 0 {
				r = unicode.ToTitle(r)
			}
			ndots = 0
			lastIsSpace = false
			lastIsPunct = false
			n++
			res = append(res, r)
		}

	}
	return string(res)
}
