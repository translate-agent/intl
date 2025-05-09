package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func seqEraMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, region := locale.Raw()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()

	switch lang {
	default:
		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
	case cldr.AF, cldr.AS, cldr.IA, cldr.KY, cldr.MI, cldr.RM, cldr.TG, cldr.WO:
		return seq.Add(era, ' ', Day2Digit.symbol(), '-', Month2Digit.symbol("format"))
	case cldr.SD:
		if script == cldr.Deva {
			return seq.Add(era, ' ', opts.Month.symbol("format"), '/', opts.Day.symbol())
		}

		fallthrough
	case cldr.BGC, cldr.BHO, cldr.BO, cldr.CE, cldr.CKB, cldr.CSW, cldr.EO, cldr.GAA, cldr.GV, cldr.KL, cldr.KSH,
		cldr.KW, cldr.LIJ, cldr.LKT, cldr.LMO, cldr.MGO, cldr.MT, cldr.NDS, cldr.NNH, cldr.NE, cldr.NQO, cldr.NSO, cldr.OC,
		cldr.PRG, cldr.PS, cldr.QU, cldr.RAJ, cldr.RW, cldr.SAH, cldr.SAT, cldr.SN, cldr.ST, cldr.SZL, cldr.TN, cldr.TOK,
		cldr.VMW, cldr.YI, cldr.ZA, cldr.ZU:
		return seq.Add(era, ' ', Month2Digit.symbol("format"), '-', Day2Digit.symbol())
	case cldr.LT:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
		}

		fallthrough
	case cldr.DZ, cldr.SI:
		return seq.Add(era, ' ', opts.Month.symbol("format"), '-', opts.Day.symbol())
	case cldr.NL:
		if region == cldr.RegionBE {
			return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
		}

		fallthrough
	case cldr.FY, cldr.UG:
		return seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbol("format"))
	case cldr.OR:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(era, ' ', opts.Month.symbol("format"), '/', opts.Day.symbol())
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbol("format"))
	case cldr.MS:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbol("format"))
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
	case cldr.SE:
		if region == cldr.RegionFI {
			return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
		}

		return seq.Add(era, ' ', Month2Digit.symbol("format"), '-', Day2Digit.symbol())
	case cldr.KN, cldr.MR, cldr.VI:
		if !opts.Month.numeric() || !opts.Day.numeric() {
			return seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbol("format"))
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
	case cldr.TI:
		return seq.Add(era, ' ', DayNumeric.symbol(), '/', MonthNumeric.symbol("format"))
	case cldr.FF:
		if script == cldr.Adlm {
			return seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbol("format"))
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
	case cldr.BN, cldr.CCP, cldr.GU, cldr.TA, cldr.TE:
		if opts.Month.twoDigit() || opts.Day.twoDigit() {
			return seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbol("format"))
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
	case cldr.AZ, cldr.CV, cldr.FO, cldr.HY, cldr.KK, cldr.KU, cldr.OS, cldr.TK, cldr.TT, cldr.UK:
		return seq.Add(era, ' ', Day2Digit.symbol(), '.', Month2Digit.symbol("format"))
	case cldr.SQ:
		return seq.Add(era, ' ', DayNumeric.symbol(), '.', MonthNumeric.symbol("format"))
	case cldr.BG, cldr.PL:
		return seq.Add(era, ' ', opts.Day.symbol(), '.', Month2Digit.symbol("format"))
	case cldr.BE, cldr.DA, cldr.ET, cldr.HE, cldr.IE, cldr.JGO, cldr.KA:
		return seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbol("format"))
	case cldr.MK:
		return seq.Add(era, ' ', DayNumeric.symbol(), '.', opts.Month.symbol("format"))
	case cldr.NB, cldr.NN, cldr.NO:
		return seq.Add(era, ' ', DayNumeric.symbol(), '.', MonthNumeric.symbol("format"), '.')
	case cldr.LV:
		return seq.Add(era, ' ', Day2Digit.symbol(), '.', Month2Digit.symbol("format"), '.')
	case cldr.SR:
		seq.Add(era, ' ', opts.Day.symbol(), '.')

		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(' ')
		}

		return seq.Add(opts.Month.symbol("format"), '.')
	case cldr.HR:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(era, ' ', Day2Digit.symbol(), '.', ' ', Month2Digit.symbol("format"), '.')
		}

		fallthrough
	case cldr.CS, cldr.SK, cldr.SL:
		return seq.Add(era, ' ', opts.Day.symbol(), '.', ' ', opts.Month.symbol("format"), '.')
	case cldr.RO, cldr.RU:
		if opts.Month.numeric() && opts.Day.numeric() {
			return seq.Add(era, ' ', Day2Digit.symbol(), '.', Month2Digit.symbol("format"))
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbol("format"))
	case cldr.DE, cldr.DSB, cldr.FI, cldr.GSW, cldr.HSB, cldr.LB, cldr.IS, cldr.SMN:
		return seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbol("format"), '.')
	case cldr.HU, cldr.KO:
		return seq.Add(era, ' ', opts.Month.symbol("format"), '.', ' ', opts.Day.symbol(), '.')
	case cldr.WAE:
		return seq.Add(era, ' ', opts.Day.symbol(), '.', ' ', MonthLong.symbol("format"))
	case cldr.BS:
		if script == cldr.Cyrl {
			return seq.Add(era, ' ', Day2Digit.symbol(), '.', Month2Digit.symbol("format"), '.')
		}

		return seq.Add(era, ' ', DayNumeric.symbol(), '.', ' ', MonthNumeric.symbol("format"), '.')
	case cldr.OM:
		if opts.Month.twoDigit() || opts.Day.twoDigit() {
			return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
		}

		return seq.Add(era, ' ', Month2Digit.symbol("format"), '-', Day2Digit.symbol())
	case cldr.KS:
		if script == cldr.Deva {
			return seq.Add(era, ' ', Month2Digit.symbol("format"), '-', Day2Digit.symbol())
		}

		fallthrough
	case cldr.AK, cldr.ASA, cldr.BEM, cldr.BLO, cldr.BEZ, cldr.BRX, cldr.CEB, cldr.CGG, cldr.CHR, cldr.DAV, cldr.EBU,
		cldr.EE, cldr.EU, cldr.FIL, cldr.GUZ, cldr.HA, cldr.KAM, cldr.KDE, cldr.KLN, cldr.TEO, cldr.VAI, cldr.JA,
		cldr.JMC, cldr.KI, cldr.KSB, cldr.LAG, cldr.LG, cldr.LUO, cldr.LUY, cldr.MAS, cldr.MER, cldr.NAQ, cldr.ND,
		cldr.NYN, cldr.ROF, cldr.RWK, cldr.SAQ, cldr.SBP, cldr.SO, cldr.TZM, cldr.VUN, cldr.XH, cldr.XOG, cldr.YUE:
		return seq.Add(era, ' ', opts.Month.symbol("format"), '/', opts.Day.symbol())
	case cldr.MN:
		return seq.Add(era, ' ', MonthNarrow.symbol("stand-alone"), '/', Day2Digit.symbol())
	case cldr.ZH:
		if region == cldr.RegionHK || region == cldr.RegionMO {
			return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
		}

		seq.Add(era, ' ', opts.Month.symbol("format"))

		if region == cldr.RegionSG {
			seq.Add('-')
		} else {
			seq.Add('/')
		}

		return seq.Add(opts.Day.symbol())
	case cldr.FR:
		if region == cldr.RegionCA {
			if opts.Month.twoDigit() || opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return seq.Add(era, ' ', opts.Month.symbol("format"), '-', opts.Day.symbol())
		}

		if region == cldr.RegionCH {
			if opts.Month.numeric() && opts.Day.numeric() {
				return seq.Add(era, ' ', Day2Digit.symbol(), '.', Month2Digit.symbol("format"), '.')
			}

			return seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbol("format"))
		}

		fallthrough
	case cldr.BR, cldr.GA, cldr.IT, cldr.JV, cldr.KKJ, cldr.SC, cldr.SYR, cldr.VEC, cldr.UZ:
		return seq.Add(era, ' ', Day2Digit.symbol(), '/', Month2Digit.symbol("format"))
	case cldr.PCM:
		return seq.Add(era, ' ', opts.Day.symbol(), ' ', '/', opts.Month.symbol("format"))
	case cldr.SV:
		if opts.Month.twoDigit() && opts.Day.numeric() {
			opts.Month = MonthNumeric
		}

		if region == cldr.RegionFI {
			return seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbol("format"))
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
	case cldr.KEA, cldr.PT:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
	case cldr.HI:
		if script != cldr.Latn {
			return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
		}

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return seq.Add(opts.Day.symbol(), '/', opts.Month.symbol("format"), ' ', era)
	case cldr.AR:
		return seq.Add(era, ' ', opts.Day.symbol(), symbols.Txt02, opts.Month.symbol("format"))
	case cldr.LRC:
		if region == cldr.RegionIQ {
			return seq.Add(era, ' ', Month2Digit.symbol("format"), '-', Day2Digit.symbol())
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
	case cldr.EN:
		switch region {
		case cldr.RegionUS, cldr.RegionAS, cldr.RegionBI, cldr.RegionPH, cldr.RegionPR, cldr.RegionUM, cldr.RegionVI:
			return seq.Add(opts.Month.symbol("format"), '/', opts.Day.symbol(), ' ', era)
		case cldr.RegionAU, cldr.RegionBE, cldr.RegionIE, cldr.RegionNZ, cldr.RegionZW:
			return seq.Add(opts.Day.symbol(), '/', opts.Month.symbol("format"), ' ', era)
		case cldr.RegionGU, cldr.RegionMH, cldr.RegionMP, cldr.RegionZZ:
			return seq.Add(opts.Month.symbol("format"), '/', opts.Day.symbol(), ' ', era)
		case cldr.RegionCA:
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return seq.Add(opts.Month.symbol("format"), '-', opts.Day.symbol(), ' ', era)
		case cldr.RegionCH:
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return seq.Add(opts.Day.symbol(), '.', opts.Month.symbol("format"), ' ', era)
		case cldr.RegionZA:
			if !opts.Month.twoDigit() || !opts.Day.twoDigit() {
				return seq.Add(Month2Digit.symbol("format"), '/', Day2Digit.symbol(), ' ', era)
			}
		}

		if script == cldr.Shaw || script == cldr.Dsrt {
			return seq.Add(opts.Month.symbol("format"), '/', opts.Day.symbol(), ' ', era)
		}

		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		return seq.Add(opts.Day.symbol(), '/', opts.Month.symbol("format"), ' ', era)
	case cldr.ES:
		seq.Add(era, ' ')

		switch region {
		default:
			return seq.Add(DayNumeric.symbol(), '/', MonthNumeric.symbol("format"))
		case cldr.RegionUS, cldr.RegionMX:
			return seq.Add(opts.Day.symbol(), '/', opts.Month.symbol("format"))
		case cldr.RegionCL:
			if opts.Month.twoDigit() {
				return seq.Add(DayNumeric.symbol(), '/', MonthNumeric.symbol("format"))
			}

			return seq.Add(Day2Digit.symbol(), '-', Month2Digit.symbol("format"))
		case cldr.RegionPA, cldr.RegionPR:
			if opts.Month.numeric() {
				return seq.Add(Month2Digit.symbol("format"), '/', Day2Digit.symbol())
			}

			if opts.Month.twoDigit() {
				opts.Month = MonthNumeric
				opts.Day = DayNumeric
			} else {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			return seq.Add(opts.Day.symbol(), '/', opts.Month.symbol("format"))
		}
	case cldr.II:
		return seq.Add(era, ' ', Month2Digit.symbol("format"), symbols.Txt03, Day2Digit.symbol(), symbols.Txtꑍ)
	case cldr.KOK:
		if script != cldr.Latn {
			return seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbol("format"))
		}

		return seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbol("format"))
	case cldr.KAA, cldr.MHN:
		return seq.Add(opts.Month.symbol("format"), ' ', opts.Day.symbol(), ' ', era)
	}
}

func seqEraMonthDayPersian(locale language.Tag, opts Options) *symbols.Seq {
	seq := symbols.NewSeq(locale)
	lang, _ := locale.Base()

	seq.Add(opts.Era.symbol(), ' ', opts.Month.symbol("format"))

	switch lang {
	default:
		seq.Add('-')
	case cldr.FA, cldr.PS:
		seq.Add('/')
	}

	return seq.Add(opts.Day.symbol())
}

func seqEraMonthDayBuddhist(locale language.Tag, opts Options) *symbols.Seq {
	return symbols.NewSeq(locale).Add(opts.Era.symbol(), ' ').AddSeq(seqMonthDayBuddhist(locale, opts))
}
