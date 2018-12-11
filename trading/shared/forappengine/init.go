package forappengine

import (
	lg "twiggg.com/lk/logger/v1.0.0"
)

var log lg.Logger

func init() {
	log = lg.Appengine()
}
