package validators

import (
	binance "twiggg.com/lk/trading/core/wrappers/binance/v1.0.0"
	bittrex "twiggg.com/lk/trading/core/wrappers/bittrex/v1.0.0"
	gdax "twiggg.com/lk/trading/core/wrappers/gdax/v1.0.0"
)

//constants
const (
	OneDayMinutes       = 1440
	SixteenHoursMinutes = 960
	TwelveHoursMinutes  = 720
	EightHoursMinutes   = 480
	SixHoursMinutes     = 360
	FourHoursMinutes    = 240
	ThreeHoursMinutes   = 180
	TwoHoursMinutes     = 120
	OneHourMinutes      = 60
	ThirtyMinutes       = 30
	FifteenMinutes      = 15
)

var validExch = map[string]struct{}{Bittrex: {}, Binance: {}, Gdax: {}}
var validCandlesInterval = map[int]string{FifteenMinutes: "15", ThirtyMinutes: "30", OneHourMinutes: "60", OneDayMinutes: "1440"}
var validBittrexCandlesInterval = map[int]string{ThirtyMinutes: bittrex.ThirtyMinShort, OneHourMinutes: bittrex.OneHourShort, OneDayMinutes: bittrex.OneDayShort}
var validBinanceCandlesInterval = map[int]string{FifteenMinutes: binance.FifteenMin, OneHourMinutes: binance.OneHour, OneDayMinutes: binance.OneDay}
var validGdaxCandlesInterval = map[int]string{FifteenMinutes: gdax.FifteenMin, OneHourMinutes: gdax.OneHour, OneDayMinutes: gdax.OneDay}
var validCandlesStringInterval = map[string]int{"15m": FifteenMinutes, "15mn": FifteenMinutes, "15min": FifteenMinutes, "30m": ThirtyMinutes, "30mn": ThirtyMinutes, "30min": ThirtyMinutes, "60m": OneHourMinutes, "60mn": OneHourMinutes, "60min": OneHourMinutes, "1d": OneDayMinutes, "1day": OneDayMinutes}
var validWatchIntervals = map[int]string{ /*FifteenMinutes: "15",*/ ThirtyMinutes: "30", OneHourMinutes: "60", ThreeHoursMinutes: "180", SixHoursMinutes: "360", TwelveHoursMinutes: "720", OneDayMinutes: "1440"}
var validWatchDays = map[int]string{-1: "permanent", 5: "5", 10: "10", 30: "30", 60: "60", 120: "120", 180: "180", 365: "365"}
