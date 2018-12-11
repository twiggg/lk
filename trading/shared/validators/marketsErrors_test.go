package validators

import (
	"testing"
)

func TestListValidCandleIntervals(t *testing.T) {
	tests := []struct {
		exch   string
		valids []string
	}{
		{
			Bittrex, []string{"030", "060", "1440"},
		},
	}

	/*
		fmt.Printf("ints bittrex candle intervals: %v\n", extractInts(validBittrexCandlesInterval))
		fmt.Printf("ints binance candle intervals: %v\n", extractInts(validBinanceCandlesInterval))
		fmt.Printf("ints gdax candle intervals: %v\n", extractInts(validGdaxCandlesInterval))
	*/
	for ind, test := range tests {
		valids := listValidCandlesIntervals(test.exch)
		l1 := len(valids)
		l2 := len(test.valids)
		if l1 != l2 {
			t.Errorf("test %d: expected %d valid candle intervals received %d\n", ind, l2, l1)
			continue
		}
		for i := range valids {
			if test.valids[i] != valids[i] {
				t.Errorf("test %d: interval %d -> expected '%s' received '%s'\n", ind, i, test.valids[i], valids[i])
			}
		}
	}
}
