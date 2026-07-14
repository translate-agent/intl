package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func seqEraMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	region, _ := locale.Region()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()

	switch lang {
	case cldr.TH:
		seq.Add(era, ' ').AddSeq(seqMonthDay(locale, opts))
	case cldr.SHN:
		seq.Add(era, ' ', symbols.Symbol_MM, '-', symbols.Symbol_dd)
	case cldr.FA, cldr.PS:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(era, ' ', opts.Month.symbolFormat(), '/', opts.Day.symbol())
		} else {
			seq.Add(era, ' ', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.MZN, cldr.CKB:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(era, ' ', opts.Month.symbolFormat(), '-', opts.Day.symbol())
		} else {
			seq.Add(era, ' ', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.UZ:
		if region == cldr.RegionAF {
			seq.Add(era, ' ', opts.Month.symbolFormat(), '-', opts.Day.symbol())
		} else {
			seq.Add(era, ' ', symbols.Symbol_dd, '/', symbols.Symbol_MM)
		}
	case cldr.AF, cldr.AS, cldr.IA, cldr.KY, cldr.MI, cldr.RM, cldr.RW, cldr.TG, cldr.WO:
		seq.Add(era, ' ', symbols.Symbol_dd, '-', symbols.Symbol_MM)
	case cldr.SD:
		if script == cldr.Deva {
			seq.Add(era, ' ', opts.Month.symbolFormat(), '/', opts.Day.symbol())
			break
		}

		fallthrough
	case cldr.BGC, cldr.BHO, cldr.BO, cldr.BUA, cldr.CE, cldr.CSW, cldr.EO, cldr.GAA, cldr.GV, cldr.KL, cldr.KSH,
		cldr.KW, cldr.LIJ, cldr.LKT, cldr.LMO, cldr.MGO, cldr.MT, cldr.NDS, cldr.NNH, cldr.NE, cldr.NQO, cldr.NSO, cldr.OC,
		cldr.PMS, cldr.PRG, cldr.QU, cldr.RAJ, cldr.SAH, cldr.SAT, cldr.SN, cldr.ST, cldr.SZL, cldr.TN,
		cldr.TOK, cldr.TYV, cldr.VMW, cldr.XH, cldr.YI, cldr.ZA, cldr.ZU:
		seq.Add(era, ' ', symbols.Symbol_MM, '-', symbols.Symbol_dd)
	case cldr.LT:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
		}

		fallthrough
	case cldr.DZ, cldr.SI:
		seq.Add(era, ' ', opts.Month.symbolFormat(), '-', opts.Day.symbol())
	case cldr.NL:
		if region == cldr.RegionBE {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
			break
		}

		fallthrough
	case cldr.FY, cldr.UG:
		seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbolFormat())
	case cldr.OR:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(era, ' ', opts.Month.symbolFormat(), '/', opts.Day.symbol())
		} else {
			seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbolFormat())
		}
	case cldr.MS:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbolFormat())
		} else {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		}
	case cldr.SE:
		if region == cldr.RegionFI {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		} else {
			seq.Add(era, ' ', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.KN, cldr.MR, cldr.VI:
		if !opts.Month.numeric() || !opts.Day.numeric() {
			seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbolFormat())
		} else {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		}
	case cldr.TI:
		seq.Add(era, ' ', symbols.Symbol_d, '/', symbols.Symbol_M)
	case cldr.FF:
		if script == cldr.Adlm {
			seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbolFormat())
		} else {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		}
	case cldr.BN, cldr.CCP, cldr.GU, cldr.TA, cldr.TE:
		if opts.Month.twoDigit() || opts.Day.twoDigit() {
			seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbolFormat())
		} else {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		}
	case cldr.KK:
		if script == cldr.Arab {
			seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbolFormat())

			break
		}

		fallthrough
	case cldr.AZ, cldr.BA, cldr.FO, cldr.HY, cldr.KU, cldr.OS, cldr.TK, cldr.TT, cldr.UK:
		seq.Add(era, ' ', symbols.Symbol_dd, '.', symbols.Symbol_MM)
	case cldr.SQ:
		seq.Add(era, ' ', symbols.Symbol_d, '.', symbols.Symbol_M)
	case cldr.BG, cldr.PL:
		seq.Add(era, ' ', opts.Day.symbol(), '.', symbols.Symbol_MM)
	case cldr.BE, cldr.DA, cldr.ET, cldr.HE, cldr.IE, cldr.JGO, cldr.KA:
		seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbolFormat())
	case cldr.MK:
		seq.Add(era, ' ', symbols.Symbol_d, '.', opts.Month.symbolFormat())
	case cldr.NB, cldr.NN, cldr.NO:
		seq.Add(era, ' ', symbols.Symbol_d, '.', symbols.Symbol_M, '.')
	case cldr.LV:
		seq.Add(era, ' ', symbols.Symbol_dd, '.', symbols.Symbol_MM, '.')
	case cldr.SR:
		seq.Add(era, ' ', opts.Day.symbol(), '.')

		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(' ')
		}

		seq.Add(opts.Month.symbolFormat(), '.')
	case cldr.HR:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(era, ' ', symbols.Symbol_dd, '.', ' ', symbols.Symbol_MM, '.')

			break
		}

		fallthrough
	case cldr.CS, cldr.SK, cldr.SL:
		seq.Add(era, ' ', opts.Day.symbol(), '.', ' ', opts.Month.symbolFormat(), '.')
	case cldr.RO, cldr.RU:
		if opts.Month.numeric() && opts.Day.numeric() {
			seq.Add(era, ' ', symbols.Symbol_dd, '.', symbols.Symbol_MM)
		} else {
			seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbolFormat())
		}
	case cldr.DE:
		if opts.Month.twoDigit() && opts.Day.numeric() {
			seq.Add(era, ' ', symbols.Symbol_dd, '.', opts.Month.symbolFormat(), '.')

			break
		}

		fallthrough
	case cldr.DSB, cldr.FI, cldr.GSW, cldr.HSB, cldr.LB, cldr.IS, cldr.SMN:
		seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbolFormat(), '.')
	case cldr.HU, cldr.KO:
		seq.Add(era, ' ', opts.Month.symbolFormat(), '.', ' ', opts.Day.symbol(), '.')
	case cldr.WAE:
		seq.Add(era, ' ', opts.Day.symbol(), '.', ' ', symbols.Symbol_MMM)
	case cldr.BS:
		if script == cldr.Cyrl {
			seq.Add(era, ' ', symbols.Symbol_dd, '.', symbols.Symbol_MM, '.')
		} else {
			seq.Add(era, ' ', symbols.Symbol_d, '.', ' ', symbols.Symbol_M, '.')
		}
	case cldr.OM:
		if opts.Month.twoDigit() || opts.Day.twoDigit() {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		} else {
			seq.Add(era, ' ', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		}
	case cldr.KS:
		if script == cldr.Deva {
			seq.Add(era, ' ', symbols.Symbol_MM, '-', symbols.Symbol_dd)

			break
		}

		fallthrough
	case cldr.AK, cldr.ASA, cldr.BEM, cldr.BLO, cldr.BEZ, cldr.BRX, cldr.CEB, cldr.CGG, cldr.CHR, cldr.DAV, cldr.EBU,
		cldr.EE, cldr.EU, cldr.FIL, cldr.GUZ, cldr.HA, cldr.KAM, cldr.KDE, cldr.KLN, cldr.TEO, cldr.VAI, cldr.JA,
		cldr.JMC, cldr.KI, cldr.KSB, cldr.LAG, cldr.LG, cldr.LUO, cldr.LUY, cldr.MAS, cldr.MER, cldr.NAQ, cldr.ND,
		cldr.NYN, cldr.ROF, cldr.RWK, cldr.SAQ, cldr.SBP, cldr.SO, cldr.TZM, cldr.VUN, cldr.XOG, cldr.YUE:
		seq.Add(era, ' ', opts.Month.symbolFormat(), '/', opts.Day.symbol())
	case cldr.MN:
		seq.Add(era, ' ', symbols.Symbol_LLLLL, '/', symbols.Symbol_dd)
	case cldr.ZH:
		if region == cldr.RegionHK || region == cldr.RegionMO {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		} else {
			seq.Add(era, ' ', opts.Month.symbolFormat())

			if region == cldr.RegionSG {
				seq.Add('-')
			} else {
				seq.Add('/')
			}

			seq.Add(opts.Day.symbol())
		}
	case cldr.FR:
		switch region {
		case cldr.RegionCA:
			if opts.Month.twoDigit() || opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			seq.Add(era, ' ', opts.Month.symbolFormat(), '-', opts.Day.symbol())
		case cldr.RegionCH:
			if opts.Month.numeric() && opts.Day.numeric() {
				seq.Add(era, ' ', symbols.Symbol_dd, '.', symbols.Symbol_MM, '.')
			} else {
				seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbolFormat())
			}
		default:
			seq.Add(era, ' ', symbols.Symbol_dd, '/', symbols.Symbol_MM)
		}
	case cldr.BR, cldr.GA, cldr.IT, cldr.JV, cldr.KKJ, cldr.SC, cldr.SYR, cldr.VEC:
		seq.Add(era, ' ', symbols.Symbol_dd, '/', symbols.Symbol_MM)
	case cldr.PCM:
		seq.Add(era, ' ', opts.Day.symbol(), ' ', '/', opts.Month.symbolFormat())
	case cldr.SV:
		if opts.Month.twoDigit() && opts.Day.numeric() {
			opts.Month = MonthNumeric
		}

		if region == cldr.RegionFI {
			seq.Add(era, ' ', opts.Day.symbol(), '.', opts.Month.symbolFormat())
		} else {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		}
	case cldr.KEA, cldr.PT:
		if opts.Month.numeric() && opts.Day.numeric() {
			opts.Month = Month2Digit
			opts.Day = Day2Digit
		}

		seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
	case cldr.HI:
		if script != cldr.Latn {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		} else {
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), ' ', era)
		}
	case cldr.AR:
		seq.Add(era, ' ', opts.Day.symbol(), symbols.Txt02, opts.Month.symbolFormat())
	case cldr.LRC:
		switch region {
		case cldr.RegionIR, cldr.RegionAF:
			seq.Add(era, ' ', opts.Month.symbolFormat(), '-', opts.Day.symbol())
		case cldr.RegionIQ:
			seq.Add(era, ' ', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		default:
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
		}
	case cldr.EN:
		switch {
		case region == cldr.RegionUS ||
			region == cldr.RegionAS ||
			region == cldr.RegionBI ||
			region == cldr.RegionGU ||
			region == cldr.RegionJP ||
			region == cldr.RegionMH ||
			region == cldr.RegionMP ||
			region == cldr.RegionPH ||
			region == cldr.RegionPR ||
			region == cldr.RegionUM ||
			region == cldr.RegionVI ||
			region == cldr.RegionZZ:
			seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), ' ', era)
		case region == cldr.RegionAU ||
			region == cldr.RegionBE ||
			region == cldr.RegionIE ||
			region == cldr.RegionNZ ||
			region == cldr.RegionZW:
			seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), ' ', era)
		case region == cldr.RegionCA:
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			seq.Add(opts.Month.symbolFormat(), '-', opts.Day.symbol(), ' ', era)
		case region == cldr.RegionCH:
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			seq.Add(opts.Day.symbol(), '.', opts.Month.symbolFormat(), ' ', era)
		case region == cldr.RegionZA && (!opts.Month.twoDigit() || !opts.Day.twoDigit()):
			seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd, ' ', era)
		case script == cldr.Shaw || script == cldr.Dsrt:
			seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), ' ', era)
		default:
			if opts.Month.numeric() && opts.Day.numeric() {
				opts.Month = Month2Digit
				opts.Day = Day2Digit
			}

			seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), ' ', era)
		}
	case cldr.ES:
		seq.Add(era, ' ')

		switch region {
		default:
			seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
		case cldr.RegionUS, cldr.RegionMX:
			seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat())
		case cldr.RegionCL:
			if opts.Month.twoDigit() {
				seq.Add(symbols.Symbol_d, '/', symbols.Symbol_M)
			} else {
				seq.Add(symbols.Symbol_dd, '-', symbols.Symbol_MM)
			}
		case cldr.RegionPA, cldr.RegionPR:
			if opts.Month.numeric() {
				seq.Add(symbols.Symbol_MM, '/', symbols.Symbol_dd)
			} else {
				if opts.Month.twoDigit() {
					opts.Month = MonthNumeric
					opts.Day = DayNumeric
				} else {
					opts.Month = Month2Digit
					opts.Day = Day2Digit
				}

				seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat())
			}
		}
	case cldr.II:
		seq.Add(era, ' ', symbols.Symbol_MM, symbols.Txt03, symbols.Symbol_dd, symbols.Txtꑍ)
	case cldr.KOK:
		seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
	case cldr.KAA, cldr.MHN:
		seq.Add(opts.Month.symbolFormat(), ' ', opts.Day.symbol(), ' ', era)
	case cldr.CV:
		if opts.Month.numeric() || opts.Day.numeric() {
			seq.Add(era, ' ', symbols.Symbol_MM, '.', symbols.Symbol_dd)
		} else {
			seq.Add(era, ' ', opts.Month.symbolFormat(), '.', opts.Day.symbol())
		}

	}

	if seq.Empty() {
		seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat())
	}

	return seq
}
