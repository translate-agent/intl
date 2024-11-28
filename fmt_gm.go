package intl

import (
	"time"

	"golang.org/x/text/language"
)

func fmtEraMonthGregorian(locale language.Tag, digits digits, opts Options) func(m time.Month) string {
	lang, _ := locale.Base()

	era := fmtEra(locale, opts.Era)
	fmtMonth := fmtMonth(digits)

	switch lang {
	default:
		return func(m time.Month) string {
			return era + " " + fmtMonth(m, opts.Month)
		}
	case lv:
		// era=long,month=numeric,out=mūsu ērā 1
		// era=long,month=2-digit,out=mūsu ērā (mēnesis: 01)
		// era=short,month=numeric,out=m.ē. (mēnesis: 1)
		// era=short,month=2-digit,out=m.ē. (mēnesis: 01)
		// era=narrow,month=numeric,out=m.ē. 1
		// era=narrow,month=2-digit,out=m.ē. 01
		if opts.Era == EraShort ||
			opts.Era == EraLong && opts.Month == Month2Digit {
			return func(m time.Month) string {
				return era + " (mēnesis: " + fmtMonth(m, opts.Month) + ")"
			}
		}

		return func(m time.Month) string {
			return era + " " + fmtMonth(m, opts.Month)
		}
	}
}
