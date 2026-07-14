package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop,gocognit
func seqEraYearMonthDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	region, _ := locale.Region()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	year := opts.Year.symbol()
	month := symbols.Symbol_MM
	day := symbols.Symbol_dd

	switch lang {
	case cldr.TH:
		if region == cldr.RegionTH {
			seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/').AddSeq(seqYear(locale, opts))
		}
	case cldr.SHN:
		if region == cldr.RegionTH {
			if opts.Era.narrow() {
				seq.AddSeq(seqYear(locale, opts)).Add('-', opts.Month.symbolFormat(), '-', opts.Day.symbol())
			} else {
				seq.AddSeq(seqYear(locale, opts)).Add('-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
			}
		}
	case cldr.FA:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', opts.Year.symbol(), ' ', era)
		}
	case cldr.LRC, cldr.MZN:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.Add(era, ' ', opts.Year.symbol(), '-', opts.Month.symbolFormat(), '-', opts.Day.symbol())
		}
	case cldr.PS:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			if !opts.Era.narrow() {
				seq.Add(era, ' ', opts.Year.symbol(), '-', opts.Month.symbolFormat(), '-', opts.Day.symbol())
			} else {
				seq.Add(era, ' ', opts.Year.symbol(), '/', opts.Month.symbolFormat(), '/', opts.Day.symbol())
			}
		} else {
			seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
		}
	case cldr.CKB:
		if region == cldr.RegionIR {
			seq.Add(era, ' ', opts.Year.symbol(), '-', opts.Month.symbolFormat(), '-', opts.Day.symbol())
		}
	case cldr.EN:
		switch region {
		default:
			if script == cldr.Shaw || script == cldr.Dsrt || script.String() == "" {
				seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
			} else {
				seq.Add(day, '/', month, '/', year, ' ', era)
			}
		case cldr.RegionAE, cldr.RegionAS, cldr.RegionBI, cldr.RegionCA, cldr.RegionGU, cldr.RegionMH, cldr.RegionMP,
			cldr.RegionPH, cldr.RegionPR, cldr.RegionUM, cldr.RegionUS,
			cldr.RegionVI:
			seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
		case cldr.RegionCH:
			seq.Add(opts.Day.symbol(), '.', opts.Month.symbolFormat(), '.', year, ' ', era)
		case cldr.RegionZZ:
			seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
		case cldr.RegionJP:
			seq.Add(year, '/', month, '/', day, ' ', era)
		}
	case cldr.BRX, cldr.LV, cldr.MNI:
		seq.Add(era, ' ', day, '-', month, '-', year)
	case cldr.DA, cldr.DSB, cldr.HSB, cldr.IE, cldr.KA, cldr.NB, cldr.NN, cldr.NO, cldr.SQ:
		seq.Add(opts.Day.symbol(), '.', opts.Month.symbolFormat(), '.', year, ' ', era)
	case cldr.MK:
		seq.Add(opts.Day.symbol(), '.', opts.Month.symbolFormat(), '.', year, ' ', symbols.Txt00, ' ', era)
	case cldr.ET, cldr.PL:
		seq.Add(opts.Day.symbol(), '.', month, '.', year, ' ', era)
	case cldr.DE, cldr.FO, cldr.RO, cldr.RU:
		seq.Add(day, '.', month, '.', year, ' ', era)
	case cldr.BE:
		seq.Add(day, '.', month, '.', opts.Year.symbol(), ' ', era)
	case cldr.SR:
		seq.Add(opts.Day.symbol(), '.', month, '.', year, '.', ' ', era)
	case cldr.BG:
		seq.Add(day, '.', month, '.', year, ' ', symbols.Txt00, ' ', era)
	case cldr.FI:
		seq.Add(opts.Month.symbolFormat(), '.', opts.Day.symbol(), '.', year, ' ', era)
	case cldr.FR:
		if region == cldr.RegionCA {
			seq.Add(year, '-', month, '-', day, ' ', era)
		} else {
			seq.Add(day, '/', month, '/', year, ' ', era)
		}
	case cldr.AM, cldr.AS, cldr.ES, cldr.GD, cldr.GL, cldr.HE, cldr.EL, cldr.ID, cldr.IS, cldr.NL, cldr.SCN, cldr.SU,
		cldr.SW, cldr.TA, cldr.XNR, cldr.UR, cldr.VI, cldr.YO:
		seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year, ' ', era)
	case cldr.BR, cldr.GA, cldr.IT, cldr.KEA, cldr.PT, cldr.SC, cldr.SYR, cldr.VEC:
		seq.Add(day, '/', month, '/', year, ' ', era)
	case cldr.CEB, cldr.CHR, cldr.CY, cldr.BLO, cldr.FIL, cldr.KAA, cldr.MHN, cldr.NE, cldr.OR, cldr.SD,
		cldr.SO, cldr.TI, cldr.ZU:
		seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
	case cldr.AR, cldr.IA, cldr.BN, cldr.CA, cldr.MAI, cldr.RM, cldr.RW, cldr.UK, cldr.WO:
		seq.Add(day, '-', month, '-', year, ' ', era)
	case cldr.SV:
		if region == cldr.RegionAX || region == cldr.RegionFI {
			seq.Add(opts.Day.symbol(), '.', opts.Month.symbolFormat(), '.', year, ' ', era)
			break
		}

		fallthrough
	case cldr.LT, cldr.TG:
		seq.Add(year, '-', month, '-', day, ' ', era)
	case cldr.BS:
		if script != cldr.Cyrl {
			seq.Add(opts.Day.symbol(), '.', ' ', opts.Month.symbolFormat(), '.', ' ', year, '.', ' ', era)
		}
	case cldr.FF:
		if script == cldr.Adlm {
			seq.Add(opts.Day.symbol(), '-', opts.Month.symbolFormat(), '-', year, ' ', era)
		}
	case cldr.KS:
		if script != cldr.Deva {
			seq.Add(opts.Month.symbolFormat(), '/', opts.Day.symbol(), '/', year, ' ', era)
		}
	case cldr.UZ:
		if region == cldr.RegionAF {
			seq.Add(era, ' ', opts.Year.symbol(), '-', symbols.Symbol_MM, '-', symbols.Symbol_dd)
		} else if script != cldr.Cyrl {
			seq.Add(day, '.', month, '.', year, ' ', era)
		}
	case cldr.AZ:
		if script != cldr.Cyrl {
			seq.Add(era, ' ', opts.Day.symbol(), ' ', opts.Month.symbolFormat(), ' ', year)
		}
	case cldr.KU:
		seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year)
	case cldr.HU:
		seq.Add(era, ' ', year, '.', ' ', month, '.', ' ', day, '.')
	case cldr.CS, cldr.SK, cldr.SL:
		seq.Add(opts.Day.symbol(), '.', ' ', opts.Month.symbolFormat(), '.', ' ', year, ' ', era)
	case cldr.HR:
		seq.Add(opts.Day.symbol(), '.', ' ', opts.Month.symbolFormat(), '.', ' ', year, '.', ' ', era)
	case cldr.HI:
		if script == cldr.Latn {
			seq.Add(day, '/', month, '/', opts.Year.symbol(), ' ', era)
		} else {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year)
		}
	case cldr.ZH:
		if script == cldr.Hant {
			seq.Add(era, ' ', year, '/', opts.Month.symbolFormat(), '/', opts.Day.symbol())
		} else {
			seq.Add(era, year, '-', month, '-', day)
		}
	case cldr.KXV:
		if script != cldr.Deva && script != cldr.Orya && script != cldr.Telu {
			seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year)
		}
	case cldr.JA:
		seq.Add(era, year, '/', opts.Month.symbolFormat(), '/', opts.Day.symbol())
	case cldr.KO, cldr.ML, cldr.MY:
		seq.Add(era, ' ', year, '/', opts.Month.symbolFormat(), '/', opts.Day.symbol())
	case cldr.MR, cldr.QU:
		seq.Add(era, ' ', opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year)
	case cldr.TO:
		seq.Add(day, ' ', month, ' ', year, ' ', era)
	case cldr.KK:
		if script == cldr.Arab {
			seq.Add(era, ' ', opts.Day.symbol(), '-', opts.Month.symbolFormat(), '-', year)
		} else {
			seq.Add(day, '-', month, '-', era, ' ', year)
		}
	case cldr.LO:
		seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', era, ' ', year)
	case cldr.PA:
		if script != cldr.Arab {
			seq.Add(opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', era, ' ', year)
		}
	case cldr.KOK:
		if script == cldr.Latn {
			seq.Add(opts.Day.symbol(), '-', opts.Month.symbolFormat(), '-', year, ' ', era)
		}
	case cldr.BA:
		seq.Add(era, ' ', day, '.', month, '.', opts.Year.symbol())
	case cldr.CV:
		seq.Add(era, ' ', year, '.', month, '.', day)
	case cldr.EO:
		seq.Add(opts.Day.symbol(), ' ', symbols.Symbol_MMM, ' ', year, ' ', era)
	case cldr.HY, cldr.TK, cldr.TR:
		seq.Add(era, ' ', day, '.', month, '.', year)
	case cldr.YUE:
		seq.Add(era, opts.Day.symbol(), '/', opts.Month.symbolFormat(), '/', year)
	}

	if seq.Empty() {
		seq.Add(era, ' ', year, '-', month, '-', day)
	}

	return seq
}
