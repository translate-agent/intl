package intl

import (
	"go.expect.digital/intl/internal/cldr"
	"go.expect.digital/intl/internal/symbols"
	"golang.org/x/text/language"
)

//nolint:cyclop
func seqEraYearDay(locale language.Tag, opts Options) *symbols.Seq {
	lang, script, _ := locale.Raw()
	region, _ := locale.Region()
	seq := symbols.NewSeq(locale)
	era := opts.Era.symbol()
	year := seqYear(locale, opts)
	day := seqDay(locale, opts.Day)

	switch lang {
	case cldr.TH, cldr.SHN:
		if region == cldr.RegionTH {
			seq.AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		} else {
			seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.LRC, cldr.MZN, cldr.PS, cldr.CKB:
		if region == cldr.RegionIR || region == cldr.RegionAF {
			seq.AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		} else {
			seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.BE, cldr.RU:
		seq.AddSeq(year).Add(' ', symbols.Txt00, ' ', era, ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.KK:
		if script == cldr.Arab {
			seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		} else {
			seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txtж, '.', ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.KY:
		seq.Add(era, ' ').AddSeq(year).Add('-', symbols.Txtж, '.', ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.HY:
		seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txtթ, '.', ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.TT:
		seq.Add(era, ' ').AddSeq(year).Add(' ', symbols.Txt05, ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.SAH:
		seq.AddSeq(year).Add(' ', symbols.Txtс, '.', ' ', era, ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.LT:
		seq.AddSeq(year).Add(' ', 'm', '.', ' ', era, ' ', '(', symbols.DayUnit, ':', ' ', symbols.Symbol_dd, ')')
	case cldr.AGQ, cldr.AK, cldr.AS, cldr.ASA, cldr.AZ, cldr.BAS, cldr.BEM, cldr.BEZ, cldr.BGC, cldr.BHO, cldr.BM, cldr.BO,
		cldr.CE, cldr.CGG, cldr.CSW, cldr.CV, cldr.DAV, cldr.DJE, cldr.DOI, cldr.DUA, cldr.DYO, cldr.DZ, cldr.EBU,
		cldr.EU, cldr.EWO, cldr.FUR, cldr.FY, cldr.GAA, cldr.GSW, cldr.GU, cldr.GUZ, cldr.GV, cldr.HA, cldr.HU, cldr.IG,
		cldr.II, cldr.JGO, cldr.JMC, cldr.JV, cldr.KAB, cldr.KAM, cldr.KDE, cldr.KHQ, cldr.KI, cldr.KL, cldr.KLN, cldr.KN,
		cldr.KO, cldr.KOK, cldr.KSB, cldr.KSF, cldr.KSH, cldr.KU, cldr.KW, cldr.LAG, cldr.LG, cldr.LIJ, cldr.LKT, cldr.LMO,
		cldr.LN, cldr.LO, cldr.LU, cldr.LUO, cldr.LUY, cldr.LV, cldr.MAS, cldr.MER, cldr.MFE, cldr.MG, cldr.MGH,
		cldr.MGO, cldr.ML, cldr.MN, cldr.MNI, cldr.MR, cldr.MT, cldr.MUA, cldr.MY, cldr.NAQ, cldr.ND, cldr.NDS, cldr.NE,
		cldr.NMG, cldr.NNH, cldr.NQO, cldr.NSO, cldr.NUS, cldr.NYN, cldr.OC, cldr.OM, cldr.OS, cldr.PA, cldr.PCM, cldr.PMS,
		cldr.PRG, cldr.QU, cldr.RAJ, cldr.RN, cldr.ROF, cldr.RWK, cldr.SAQ, cldr.SAT, cldr.SBP, cldr.SEH, cldr.SES,
		cldr.SG, cldr.SHI, cldr.SI, cldr.SN, cldr.ST, cldr.SZL, cldr.TA, cldr.TE, cldr.TEO, cldr.TK, cldr.TN,
		cldr.TR, cldr.TWQ, cldr.TZM, cldr.VAI, cldr.VMW, cldr.VUN, cldr.WAE, cldr.XOG, cldr.YAV, cldr.YI, cldr.YO, cldr.ZA,
		cldr.ZGH, cldr.ZU:
		seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.UZ:
		if region == cldr.RegionAF {
			_, _, rawRegion := locale.Raw()
			if rawRegion != cldr.RegionAF {
				seq.AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')

				break
			}
		}

		seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.FF:
		if script != cldr.Adlm {
			seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.KS:
		if script == cldr.Deva {
			seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.SD:
		if script != cldr.Deva {
			seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.BRX, cldr.JA, cldr.YUE, cldr.ZH:
		seq.Add(era).AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.SE:
		if region != cldr.RegionFI {
			seq.Add(era, ' ').AddSeq(year).Add(' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
		}
	case cldr.BA:
		seq.Add(era, ' ').AddSeq(year).
			Add(' ', symbols.TxtCyrillicShortI, '.', ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	case cldr.BUA, cldr.TOK, cldr.TYV, cldr.XH:
		seq.Add(era, ' ', opts.Year.symbol(), ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	}

	if seq.Empty() {
		seq.AddSeq(year).Add(' ', era, ' ', '(', symbols.DayUnit, ':', ' ').AddSeq(day).Add(')')
	}

	return seq
}
