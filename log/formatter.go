package log

import (
	"fmt"
	"regexp"

	"github.com/sirupsen/logrus"
)

type SafeJSONFormatter struct {
	senstiveFields []string
	logrus.JSONFormatter
}

func (sf *SafeJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	for _, field := range sf.senstiveFields {
		r, err := regexp.Compile(fmt.Sprintf(`%s:".*?"`, field))
		if err == nil {
			entry.Message = r.ReplaceAllLiteralString(entry.Message, fmt.Sprintf(`%s:******`, field))
		}
	}

	return sf.JSONFormatter.Format(entry)

}
