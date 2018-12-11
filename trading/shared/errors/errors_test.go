package errors

import (
	"testing"

	"twiggg.com/lk/trading/shared/langs/v1.0.0"
)

func TestGetErrorMsg(t *testing.T) {
	tests := []struct {
		code   int
		i      int
		period int
		name   string
		f      string
		args   string
		lang   string
		msg    string
	}{
		{ErrEmptySerie, 0, 5, "", "", "", langs.FR, "la série est vide"},
		{ErrEmptySerie, 0, 5, "", "", "", langs.ENG, "empty serie"},
		{ErrEmptySerie, 0, 5, "", "", "", "esp", "empty serie"},
		{ErrInvalidIndex, 5, 5, "", "", "", langs.FR, "indice '5' invalide"},
		{ErrNotEnoughPoints, 5, 5, "", "", "", langs.FR, "nombre de points insuffisant (<5)"},
		{ErrFailedSelectByName, 5, 5, "mean", "", "", langs.FR, "impossible de sélectionner 'mean'"},
	}

	for i, test := range tests {
		msg := GetErrorMsg(test.code, test.i, test.period, test.name, test.f, test.args, test.lang)
		if msg != test.msg {
			t.Errorf("test %d: expected '%s' received '%s'", i, test.msg, msg)
		}
	}
}
