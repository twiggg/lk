package format

import (
	"math"
	"testing"
)

func TestFormatFixedPrecision(t *testing.T) {
	tests := []struct {
		f         float64
		f2        float64
		precision int
	}{
		{0.123456789123456, 0.1235, 4},
		{7065.98143333, 7065.981, 3},
		{7065.98143333, 7066.0, 1},
	}

	for ind, test := range tests {
		res := FloatToFixedPrecision(test.f, test.precision)
		if math.Abs(test.f2-res) > 0.0001*math.Abs(test.f2) {
			t.Errorf("test %d: expected %f received %f\n", ind, test.f2, res)
		}
	}
}

func TestParagraph(t *testing.T) {
	tests := []struct {
		p   string
		res string
	}{
		{
			"   compute the runnning slope of the serie in /min( linear)ok+1,blablabla",
			"Compute the runnning slope of the serie in /min ( linear ) ok+1, blablabla",
		},
	}

	for ind, test := range tests {
		res := Paragraph(test.p)
		if test.res != res {
			t.Errorf("test %d: expected %s received %s\n", ind, test.res, res)
		}
	}
}
