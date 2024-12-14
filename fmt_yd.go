package intl

import "golang.org/x/text/language"

func fmtYearDayGregorian(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	lang, script, _ := locale.Raw()
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)
	fmtDay := fmtDayGregorian(locale, digits)

	switch lang {
	case bg:
		if opts.Year == Year2Digit && opts.Day == DayNumeric {
			return func(y, d int) string {
				return layoutYear(fmtYear(y, opts.Year)) + " " + fmtDay(d, opts.Day)
			}
		}

		dayName := unitName(locale).Day

		return func(y, d int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " (" + dayName + ": " + fmtDay(d, opts.Day) + ")"
		}
	case en:
		if opts.Year == Year2Digit && opts.Day == DayNumeric {
			return func(y, d int) string {
				return fmtDay(d, opts.Day) + " " + layoutYear(fmtYear(y, opts.Year))
			}
		}

		dayName := unitName(locale).Day

		return func(y, d int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " (" + dayName + ": " + fmtDay(d, opts.Day) + ")"
		}
	case hi:
		if opts.Year == Year2Digit && opts.Day == DayNumeric {
			if script == latn {
				return func(y, d int) string {
					return fmtDay(d, opts.Day) + " " + layoutYear(fmtYear(y, opts.Year))
				}
			}

			return func(y, d int) string {
				return layoutYear(fmtYear(y, opts.Year)) + " " + fmtDay(d, opts.Day)
			}
		}

		dayName := unitName(locale).Day

		return func(y, d int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " (" + dayName + ": " + fmtDay(d, opts.Day) + ")"
		}
	}

	if opts.Year == Year2Digit && opts.Day == DayNumeric {
		return func(y, d int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " " + fmtDay(d, opts.Day)
		}
	}

	dayName := unitName(locale).Day

	return func(y, d int) string {
		return layoutYear(fmtYear(y, opts.Year)) + " (" + dayName + ": " + fmtDay(d, opts.Day) + ")"
	}
}

func fmtYearDayPersian(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	lang, _, region := locale.Raw()
	layoutYear := fmtYearGregorian(locale)
	fmtYear := fmtYear(digits)
	fmtDay := fmtDayGregorian(locale, digits)

	switch lang {
	case fa:
		if opts.Year == Year2Digit && opts.Day == DayNumeric {
			return func(y, d int) string {
				return layoutYear(fmtYear(y, opts.Year)) + " " + fmtDay(d, opts.Day)
			}
		}

		dayName := unitName(locale).Day

		return func(y, d int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " (" + dayName + ": " + fmtDay(d, opts.Day) + ")"
		}
	case uz:
		if region != regionAF {
			break
		}

		if opts.Year == Year2Digit && opts.Day == DayNumeric {
			return func(y, d int) string {
				return layoutYear(fmtYear(y, opts.Year)) + " " + fmtDay(d, opts.Day)
			}
		}

		dayName := unitName(locale).Day

		return func(y, d int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " (" + dayName + ": " + fmtDay(d, opts.Day) + ")"
		}
	}

	if opts.Year == Year2Digit && opts.Day == DayNumeric {
		return func(y, d int) string {
			return fmtEra(locale, EraNarrow) + " " + layoutYear(fmtYear(y, opts.Year)) + " " + fmtDay(d, opts.Day)
		}
	}

	dayName := unitName(locale).Day

	return func(y, d int) string {
		return fmtEra(locale, EraNarrow) + " " + layoutYear(fmtYear(y, opts.Year)) +
			" (" + dayName + ": " + fmtDay(d, opts.Day) + ")"
	}
}

func fmtYearDayBuddhist(locale language.Tag, digits digits, opts Options) func(y, d int) string {
	layoutYear := fmtYearBuddhist(locale, EraNarrow)
	fmtYear := fmtYear(digits)
	fmtDay := fmtDayBuddhist(locale, digits)

	if opts.Year == Year2Digit && opts.Day == DayNumeric {
		return func(y, d int) string {
			return layoutYear(fmtYear(y, opts.Year)) + " " + fmtDay(d, opts.Day)
		}
	}

	dayName := unitName(locale).Day

	return func(y, d int) string {
		return layoutYear(fmtYear(y, opts.Year)) + " (" + dayName + ": " + fmtDay(d, opts.Day) + ")"
	}
}
