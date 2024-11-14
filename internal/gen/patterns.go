package main

import (
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"strings"
)

var (
	idMd    = parseDatePattern("Md")
	idMMd   = parseDatePattern("MMd")
	idMdd   = parseDatePattern("Mdd")
	idMMdd  = parseDatePattern("MMdd")
	idyM    = parseDatePattern("yM")
	idyMM   = parseDatePattern("yMM")
	idyyyyM = parseDatePattern("yyyyM")
)

type DatePatternItem struct {
	Value   string
	Literal bool
}

type DatePattern []DatePatternItem

func (p DatePattern) String() string {
	var sb strings.Builder

	for _, v := range p {
		escape := func() bool {
			if !v.Literal {
				return false
			}

			for _, c := range v.Value {
				if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
					return true
				}
			}

			return false
		}()

		if escape {
			sb.WriteByte('\'')
		}

		sb.WriteString(v.Value)

		if escape {
			sb.WriteByte('\'')
		}
	}

	return sb.String()
}

func (p DatePattern) Copy() DatePattern {
	c := make(DatePattern, len(p))

	copy(c, p)

	return c
}

func (p DatePattern) MonthLen(cmp int) bool {
	return len(p.Month()) == cmp
}

func (p DatePattern) Month() string {
	i := p.IndexOfMonth()
	if i < 0 {
		return ""
	}

	return p[i].Value
}

func (p DatePattern) DayLen(cmp int) bool {
	return len(p.Day()) == cmp
}

func (p DatePattern) Day() string {
	i := p.IndexOfDay()
	if i < 0 {
		return ""
	}

	return p[i].Value
}

func (p DatePattern) IndexOfMonth() int {
	for i, v := range p {
		if v.Literal {
			continue
		}

		// all patterns starting with 'M' or 'L' are months
		if v.Value[0] == 'M' || v.Value[0] == 'L' {
			return i
		}
	}

	return -1
}

func (p DatePattern) IndexOfDay() int {
	for i, v := range p {
		if v.Literal {
			continue
		}

		if v.Value[0] == 'd' {
			return i
		}
	}

	return -1
}

func (p DatePattern) Score(id DatePattern) int {
	m, d := len(id.Month()), len(id.Day())
	mn, dn := len(p.Month()), len(p.Day())

	switch {
	default:
		return 0
	case mn == m && dn == d:
		return m + d
	case mn == m:
		return m
	case dn == d:
		return d
	}
}

func (p DatePattern) ReplaceMonth(v string) {
	p[p.IndexOfMonth()].Value = v
}

func (p DatePattern) ReplaceDay(v string) {
	p[p.IndexOfDay()].Value = v
}

type DatePatterns []DatePattern

func (p DatePatterns) FindClosest(id DatePattern) DatePattern {
	closest := p[0]

	for _, v := range p[1:] {
		if closest.Score(id) <= v.Score(id) {
			closest = v
		}
	}

	return closest.Copy()
}

func parseDatePattern(pattern string) DatePattern {
	if pattern == "" {
		return nil
	}

	var (
		last            rune
		literal, quoted bool
		elem            strings.Builder
		elements        = make(DatePattern, 0, 1)
	)

	write := func(r rune, asLiteral bool) {
		if literal && asLiteral {
			elem.WriteRune(r)
			last = r

			return
		}

		if !asLiteral && r == last {
			elem.WriteRune(r)

			return
		}

		if elem.Len() > 0 {
			elements = append(elements, DatePatternItem{Value: elem.String(), Literal: literal})
		}

		elem.Reset()
		elem.WriteRune(r)

		last = r
		literal = asLiteral
	}

	for i, r := range pattern {
		if i == 0 {
			last = r

			if r == '\'' && len(pattern) > 1 {
				quoted = true
				continue
			}

			elem.WriteRune(r)
			literal = !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z')

			continue
		}

		switch {
		default:
			write(r, true)
		case r == '\'':
			quoted = !quoted

			if last != r {
				last = r
				continue
			}

			write(r, true)

			last = 0
		case !quoted && ('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z'):
			write(r, false)
		}
	}

	if elem.Len() > 0 {
		elements = append(elements, DatePatternItem{Value: elem.String(), Literal: literal})
	}

	return elements
}

func yearMonthPatterns(rawyM, rawyMM, rawyyyyM string) (
	patternyM, patternyMM, patternyyyyM DatePattern,
) {
	patternyM = parseDatePattern(rawyM)
	patternyMM = parseDatePattern(rawyMM)
	patternyyyyM = parseDatePattern(rawyyyyM)

	if v := patternyM.Month(); strings.HasPrefix(v, "L") && len(v) <= 2 {
		patternyM.ReplaceMonth(strings.ReplaceAll(v, "L", "M"))
	}

	if v := patternyMM.Month(); strings.HasPrefix(v, "L") && len(v) <= 2 {
		patternyMM.ReplaceMonth(strings.ReplaceAll(v, "L", "M"))
	}

	if v := patternyyyyM.Month(); strings.HasPrefix(v, "L") && len(v) <= 2 {
		patternyyyyM.ReplaceMonth(strings.ReplaceAll(v, "L", "M"))
	}

	var patterns DatePatterns

	if len(patternyM) > 0 {
		patterns = append(patterns, patternyM)
	}

	if len(patternyMM) > 0 {
		patterns = append(patterns, patternyMM)
	}

	if len(patternyyyyM) > 0 {
		patterns = append(patterns, patternyyyyM)
	}

	if len(patternyM) == 0 {
		patternyM = patterns.FindClosest(idyM)
	}

	if len(patternyMM) == 0 {
		patternyMM = patterns.FindClosest(idyMM)
		patternyMM.ReplaceMonth("MM")
	}

	if len(patternyyyyM) == 0 {
		patternyyyyM = patterns.FindClosest(idyyyyM)

		if patternyM.MonthLen(2) {
			patternyyyyM.ReplaceMonth("MM")
		}
	}

	return
}

func monthDayPatterns(Md, MMd, Mdd, MMdd string) (MdPattern, MMdPattern, MddPattern, MMddPattern DatePattern) {
	MdPattern = parseDatePattern(Md)
	MMdPattern = parseDatePattern(MMd)
	MddPattern = parseDatePattern(Mdd)
	MMddPattern = parseDatePattern(MMdd)

	var patterns DatePatterns

	if len(MdPattern) > 0 {
		patterns = append(patterns, MdPattern)
	}

	if len(MMdPattern) > 0 {
		patterns = append(patterns, MMdPattern)
	}

	if len(MddPattern) > 0 {
		patterns = append(patterns, MddPattern)
	}

	if len(MMddPattern) > 0 {
		patterns = append(patterns, MMddPattern)
	}

	//nolint:mnd
	eqIDFmtMonth2D := (len(MMdPattern) == 0 || MMdPattern.MonthLen(2)) &&
		(len(MddPattern) == 0 || MddPattern.MonthLen(1)) &&
		(len(MMddPattern) == 0 || MMddPattern.MonthLen(2))
	idEqMonth := eqIDFmtMonth2D && (len(MdPattern) == 0 || MdPattern.MonthLen(1))

	//nolint:mnd
	idEqDay2D := (len(MMdPattern) == 0 || MMdPattern.DayLen(1)) &&
		(len(MddPattern) == 0 || MddPattern.DayLen(2)) &&
		(len(MMddPattern) == 0 || MMddPattern.DayLen(2))
	idEqDay := idEqDay2D && (len(MdPattern) == 0 || MdPattern.DayLen(1))

	if len(patterns) == 1 && len(MdPattern) > 0 && !idEqMonth && !idEqDay {
		return MdPattern, MdPattern, MdPattern, MdPattern
	}

	if len(MdPattern) == 0 {
		MdPattern = patterns.FindClosest(idMd)

		if idEqMonth {
			MdPattern.ReplaceMonth("M")
		} else {
			MdPattern.ReplaceMonth(patterns[0].Month())
		}

		if idEqDay {
			MdPattern.ReplaceDay("d")
		}
	}

	if len(MMdPattern) == 0 {
		MMdPattern = patterns.FindClosest(idMMd)

		if idEqMonth {
			MMdPattern.ReplaceMonth("MM")
		} else {
			MMdPattern.ReplaceMonth(patterns[0].Month())
		}

		if idEqDay2D {
			MMdPattern.ReplaceDay("d")
		}
	}

	if len(MddPattern) == 0 {
		MddPattern = patterns.FindClosest(idMdd)

		if eqIDFmtMonth2D {
			MddPattern.ReplaceMonth("M")
		} else {
			MddPattern.ReplaceMonth(patterns[0].Month())
		}

		if idEqDay {
			MddPattern.ReplaceDay("dd")
		}
	}

	if len(MMddPattern) == 0 {
		MMddPattern = patterns.FindClosest(idMMdd)

		if idEqMonth {
			MMddPattern.ReplaceMonth("MM")
		} else {
			MddPattern.ReplaceMonth(patterns[0].Month())
		}

		if idEqDay {
			MMddPattern.ReplaceDay("dd")
		}
	}

	return MdPattern, MMdPattern, MddPattern, MMddPattern
}

func BuildFmtMD(md, mmd, mdd, mmdd string, log *slog.Logger) string {
	mdP, mmdP, mddP, mmddP := monthDayPatterns(md, mmd, mdd, mmdd)

	log.Debug("infer MD patterns", "Md", mdP.String(), "MMd", mmdP.String(), "Mdd", mddP.String(), "MMdd", mmddP.String())

	groups := GroupLayouts(
		Layout{ID: parseDatePattern("Md"), Pattern: mdP},
		Layout{ID: parseDatePattern("MMd"), Pattern: mmdP},
		Layout{ID: parseDatePattern("Mdd"), Pattern: mddP},
		Layout{ID: parseDatePattern("MMdd"), Pattern: mmddP},
	)

	var sb strings.Builder

	writePattern := func(pattern DatePattern, group LayoutGroup) {
		for i, v := range pattern {
			if i > 0 {
				sb.WriteString(" + ")
			}

			switch v.Value {
			default:
				sb.WriteString(`"` + v.Value + `"`)
			case "M", "L":
				if group.FmtTypeMonth == FmtTypeNumericOnly || group.Expr != "" {
					sb.WriteString("fmtMonth(m, MonthNumeric)")
				} else {
					sb.WriteString("fmtMonth(m, cmp.Or(opts.Month, MonthNumeric))")
				}
			case "MM", "LL":
				if group.FmtTypeMonth == FmtType2DigitOnly || group.Expr != "" {
					sb.WriteString("fmtMonth(m, Month2Digit)")
				} else {
					sb.WriteString("fmtMonth(m, cmp.Or(opts.Month, Month2Digit))")
				}
			case "MMM":
				sb.WriteString(`fmtMonth(m, opts.Month)`)
			case "MMMMM":
				sb.WriteString("fmtMonth(m, opts.Month)")
			case "d":
				if group.FmtTypeDay == FmtTypeNumericOnly || group.Expr != "" {
					sb.WriteString("fmtDay(d, DayNumeric)")
				} else {
					sb.WriteString("fmtDay(d, cmp.Or(opts.Day, DayNumeric))")
				}
			case "dd":
				if group.FmtTypeDay == FmtType2DigitOnly || group.Expr != "" {
					sb.WriteString("fmtDay(d, Day2Digit)")
				} else {
					sb.WriteString("fmtDay(d, cmp.Or(opts.Day, Day2Digit))")
				}
			}
		}
	}

	switch len(groups) {
	case 1:
		group := groups[0]

		switch group.Layouts[0].Pattern.Month() {
		case "MMM":
			sb.WriteString("fmtMonth = fmtMonthName(locale.String(), \"stand-alone\", \"abbreviated\")\n")
		case "MMMMM":
			sb.WriteString("fmtMonth = fmtMonthName(locale.String(), \"stand-alone\", \"narrow\")\n")
		}

		sb.WriteString("return ")

		writePattern(group.Layouts[0].Pattern, group)
	case 2: //nolint:mnd
		for _, group := range groups {
			if group.Expr != "" {
				sb.WriteString("if " + group.Expr + " {\n\treturn ")
				writePattern(group.Layouts[0].Pattern, group)
				sb.WriteString("\n}\n")

				continue
			}

			for _, layout := range group.Layouts {
				if layout.ID.String() != "MMdd" {
					continue
				}

				sb.WriteString("return ")

				writePattern(layout.Pattern, group)
			}
		}
	}

	return sb.String()
}

func (l Layout) equalFlow(other Layout) bool {
	if len(l.Pattern) != len(other.Pattern) {
		return false
	}

	if l.Pattern.String() == other.Pattern.String() {
		return true
	}

	eqMonth := func(otherMonth string) bool {
		month := l.Pattern.Month()

		if l.ID.Month() == month {
			return other.ID.Month() == otherMonth
		}

		return len(month) == len(otherMonth)
	}

	eqDay := func(otherDay string) bool {
		day := l.Pattern.Day()

		if l.ID.Day() == day {
			return other.ID.Day() == otherDay
		}

		return len(day) == len(otherDay)
	}

	for i, item := range l.Pattern {
		otherItem := other.Pattern[i]

		if item.Literal != otherItem.Literal {
			return false
		}

		switch item.Value {
		default:
			if item.Value != otherItem.Value {
				return false
			}
		case "M", "MM":
			if !eqMonth(otherItem.Value) {
				return false
			}
		case "d", "dd":
			if !eqDay(otherItem.Value) {
				return false
			}
		}
	}

	return true
}

type Layout struct {
	ID      DatePattern
	Pattern DatePattern
}

type FmtType int

func (t FmtType) String() string {
	switch t {
	default:
		return "Unknown"
	case FmtTypeSame:
		return "Same"
	case FmtTypeNumericOnly:
		return "NumericOnly"
	case FmtType2DigitOnly:
		return "2DigitOnly"
	}
}

const (
	FmtTypeSame FmtType = iota
	FmtTypeNumericOnly
	FmtType2DigitOnly
)

type LayoutGroup struct {
	Expr                     string
	Layouts                  []Layout
	FmtTypeMonth, FmtTypeDay FmtType
}

func GroupLayouts(layouts ...Layout) []LayoutGroup {
	if len(layouts) == 0 {
		return nil
	}

	byPatterns := groupByPatterns(layouts)

	switch len(byPatterns) {
	case 1:
		// all have identical formatting pattern
		return postProcessGroups([]LayoutGroup{{
			Layouts: layouts,
		}})
	case 2:
		values := slices.Collect(maps.Values(byPatterns))
		if len(values[0]) > len(values[1]) {
			values[0], values[1] = values[1], values[0]
		}

		if len(values[0]) == 1 {
			// all have identical formatting pattern, except the first one
			return postProcessGroups([]LayoutGroup{
				{
					Layouts: values[0],
				},
				{
					Layouts: values[1],
				},
			})
		}
	}

	eqIDFmtMonth := make([]bool, len(layouts))
	eqIDFmtDay := make([]bool, len(layouts))
	eqPattern := slices.Repeat([]int{-1}, len(layouts))

	for i, layout := range layouts {
		eqIDFmtMonth[i] = layout.ID.Month() == layout.Pattern.Month()
		eqIDFmtDay[i] = layout.ID.Day() == layout.Pattern.Day()
	}

	for i, layout := range layouts {
		if eqPattern[i] >= 0 {
			continue
		}

		eqPattern[i] = i

		for j := i + 1; j < len(layouts); j++ {
			if layout.equalFlow(layouts[j]) && eqIDFmtMonth[i] == eqIDFmtMonth[j] && eqIDFmtDay[i] == eqIDFmtDay[j] {
				eqPattern[j] = i
			}
		}
	}

	toGroupIdx := func() []int {
		v := make([]int, len(eqPattern))
		copy(v, eqPattern)

		slices.Sort(v)
		return slices.Compact(v)
	}()

	groups := make([]LayoutGroup, len(toGroupIdx))

	for i, layout := range layouts {
		groupIdx := slices.Index(toGroupIdx, eqPattern[i])
		groups[groupIdx].Layouts = append(groups[groupIdx].Layouts, layout)
	}

	return postProcessGroups(groups)
}

func postProcessGroups(groups []LayoutGroup) []LayoutGroup {
	slices.SortFunc(groups, func(a, b LayoutGroup) int { return len(a.Layouts) - len(b.Layouts) })

	for i := range groups {
		groups[i].FmtTypeMonth = func() FmtType {
			if len(groups[i].Layouts) > 1 {
				m := groups[i].Layouts[0].Pattern.Month()
				same := !slices.ContainsFunc(groups[i].Layouts[1:], func(l Layout) bool {
					return l.Pattern.Month() != m
				})

				if same {
					switch len(m) {
					case 1:
						return FmtTypeNumericOnly
					case 2:
						return FmtType2DigitOnly
					}
				}
			}

			different := slices.ContainsFunc(groups[i].Layouts, func(l Layout) bool {
				return l.ID.Month() != l.Pattern.Month()
			})

			if !different && len(groups[i].Layouts) > 1 {
				return FmtTypeSame
			}

			if groups[i].Layouts[0].Pattern.MonthLen(1) {
				return FmtTypeNumericOnly
			}

			return FmtType2DigitOnly
		}()

		groups[i].FmtTypeDay = func() FmtType {
			different := slices.ContainsFunc(groups[i].Layouts, func(l Layout) bool {
				return l.ID.Day() != l.Pattern.Day()
			})

			if !different && len(groups[i].Layouts) > 1 {
				return FmtTypeSame
			}

			if groups[i].Layouts[0].Pattern.DayLen(1) {
				return FmtTypeNumericOnly
			}

			return FmtType2DigitOnly
		}()

		if len(groups[i].Layouts) == 1 {
			groups[i].Expr = func() string {
				optMonth := "MonthNumeric"
				optDay := "DayNumeric"

				if groups[i].Layouts[0].ID.MonthLen(2) {
					optMonth = "Month2Digit"
				}

				if groups[i].Layouts[0].ID.DayLen(2) {
					optDay = "Day2Digit"
				}

				return fmt.Sprintf("opts.Month == %s && opts.Day == %s", optMonth, optDay)
			}()
		}
	}

	return groups
}

func groupByPatterns(layouts []Layout) map[string][]Layout {
	result := make(map[string][]Layout, len(layouts))

	for _, layout := range layouts {
		result[layout.Pattern.String()] = append(result[layout.Pattern.String()], layout)
	}

	return result
}
