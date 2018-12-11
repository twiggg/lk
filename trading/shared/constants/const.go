package constants

//length constants
const (
	ConvLength          = 250
	MaxPeriod           = 90
	DftPeriod           = 12
	DftMacdFastPeriod   = 12
	DftMacdSlowPeriod   = 26
	DftSignalLinePeriod = 9
	DftMovePct          = 1.0
	MinMovePct          = 0.25
	MaxMovePct          = 100.0
	DftTolPct           = 0.0
	MinTolPct           = 0.0
	MaxTolPct           = 5.0
)

//sequences and function calls
const (
	CallsSeparator       = "/"
	NameAndArgsSeparator = "_"
	ArgsSeparator        = ","
	SequenceNameSep      = ":"
)

//operators constants
const (
	OperGreater      = ">"
	OperGreaterEqual = ">="
	OperLower        = "<"
	OperLowerEqual   = "<="
	OperEqual        = "="
)

//to return fixed float precision
const (
	Digits8 = float64(100000000)
)

//FormatDigits8 formats a number with 8 digits precision
func FormatDigits8(f float64) float64 {
	res := float64(int64(f*Digits8)) / Digits8
	return res
}

//mode constants
const (
	ModeSma  = "sma"
	ModeSmma = "smma"
	ModeEma  = "ema"
)
