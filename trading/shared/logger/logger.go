package logger

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//Infof write an info level log. //if isAppengine, an appengine context from request is mandatory
func Infof(r *http.Request /* ctx context.Context,*/, isAppengine bool, format string, values ...interface{}) {
	if isAppengine {
		actx := appengine.NewContext(r)
		log.Infof(actx, format, values)

	} else {
		/*logger, _ := zap.NewProduction()
		sugar := logger.Sugar()
		sugar.Infof(format, values...)
		*/
		logrus.Infof(format, values)
	}
}
