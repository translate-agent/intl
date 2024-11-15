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

func (p DatePattern) copy() DatePattern {
	c := make(DatePattern, len(p))

	copy(c, p)

	return c
}

func (p DatePattern) monthLen(n int) bool {
	return len(p.month()) == n
}

func (p DatePattern) month() string {
	i := p.indexOfMonth()
	if i < 0 {
		return ""
	}

	return p[i].Value
}

func (p DatePattern) dayLen(n int) bool {
	return len(p.day()) == n
}

func (p DatePattern) day() string {
	i := p.indexOfDay()
	if i < 0 {
		return ""
	}

	return p[i].Value
}

func (p DatePattern) indexOfMonth() int {
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

func (p DatePattern) indexOfDay() int {
	for i, v := range p {
		if v.Literal {
			continue
		}

		// all patterns starting with 'd' are days
		if v.Value[0] == 'd' {
			return i
		}
	}

	return -1
}

// weight returns weight of the pattern relatively to id.
// Higher weight means more similar patterns.
func (p DatePattern) weight(id DatePattern) int {
	m, d := len(id.month()), len(id.day())
	mn, dn := len(p.month()), len(p.day())

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

func (p DatePattern) replaceMonth(v string) {
	p[p.indexOfMonth()].Value = v
}

func (p DatePattern) replaceDay(v string) {
	p[p.indexOfDay()].Value = v
}

type DatePatterns []DatePattern

func (p DatePatterns) findClosest(id DatePattern) DatePattern {
	closest := p[0]

	for _, v := range p[1:] {
		if closest.weight(id) <= v.weight(id) {
			closest = v
		}
	}

	return closest.copy()
}

func parseDatePattern(format string) DatePattern {
	if format == "" {
		return nil
	}

	var (
		last            rune
		literal, quoted bool
		sb              strings.Builder
		pattern         = make(DatePattern, 0, 1)
	)

	write := func(r rune, asLiteral bool) {
		if literal && asLiteral {
			sb.WriteRune(r)
			last = r

			return
		}

		if !asLiteral && r == last {
			sb.WriteRune(r)

			return
		}

		if sb.Len() > 0 {
			pattern = append(pattern, DatePatternItem{Value: sb.String(), Literal: literal})
		}

		sb.Reset()
		sb.WriteRune(r)

		last = r
		literal = asLiteral
	}

	for i, r := range format {
		if i == 0 {
			last = r

			if r == '\'' && len(format) > 1 {
				quoted = true
				continue
			}

			sb.WriteRune(r)
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

	if sb.Len() > 0 {
		pattern = append(pattern, DatePatternItem{Value: sb.String(), Literal: literal})
	}

	return pattern
}

func yearMonthPatterns(
	formatyM, formatyMM, formatyyyyM string,
) (patternyM, patternyMM, patternyyyyM DatePattern) {
	patternyM = parseDatePattern(formatyM)
	patternyMM = parseDatePattern(formatyMM)
	patternyyyyM = parseDatePattern(formatyyyyM)

	if v := patternyM.month(); strings.HasPrefix(v, "L") && len(v) <= 2 {
		patternyM.replaceMonth(strings.ReplaceAll(v, "L", "M"))
	}

	if v := patternyMM.month(); strings.HasPrefix(v, "L") && len(v) <= 2 {
		patternyMM.replaceMonth(strings.ReplaceAll(v, "L", "M"))
	}

	if v := patternyyyyM.month(); strings.HasPrefix(v, "L") && len(v) <= 2 {
		patternyyyyM.replaceMonth(strings.ReplaceAll(v, "L", "M"))
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
		patternyM = patterns.findClosest(idyM)
	}

	if len(patternyMM) == 0 {
		patternyMM = patterns.findClosest(idyMM)
		patternyMM.replaceMonth("MM")
	}

	if len(patternyyyyM) == 0 {
		patternyyyyM = patterns.findClosest(idyyyyM)

		if patternyM.monthLen(2) { //nolint:mnd
			patternyyyyM.replaceMonth("MM")
		}
	}

	return patternyM, patternyMM, patternyyyyM
}

//nolint:gocognit,cyclop
func monthDayPatterns(
	formatMd, formatMMd, formatMdd, formatMMdd string,
) (patternMd, patternMMd, patternMdd, patternMMdd DatePattern) {
	patternMd = parseDatePattern(formatMd)
	patternMMd = parseDatePattern(formatMMd)
	patternMdd = parseDatePattern(formatMdd)
	patternMMdd = parseDatePattern(formatMMdd)

	var patterns DatePatterns

	if len(patternMd) > 0 {
		patterns = append(patterns, patternMd)
	}

	if len(patternMMd) > 0 {
		patterns = append(patterns, patternMMd)
	}

	if len(patternMdd) > 0 {
		patterns = append(patterns, patternMdd)
	}

	if len(patternMMdd) > 0 {
		patterns = append(patterns, patternMMdd)
	}

	//nolint:mnd
	eqMonth2D := (len(patternMMd) == 0 || patternMMd.monthLen(2)) &&
		(len(patternMdd) == 0 || patternMdd.monthLen(1)) &&
		(len(patternMMdd) == 0 || patternMMdd.monthLen(2))
	eqMonth := eqMonth2D && (len(patternMd) == 0 || patternMd.monthLen(1))

	//nolint:mnd
	eqDay2D := (len(patternMMd) == 0 || patternMMd.dayLen(1)) &&
		(len(patternMdd) == 0 || patternMdd.dayLen(2)) &&
		(len(patternMMdd) == 0 || patternMMdd.dayLen(2))
	eqDay := eqDay2D && (len(patternMd) == 0 || patternMd.dayLen(1))

	if len(patterns) == 1 && len(patternMd) > 0 && !eqMonth && !eqDay {
		return patternMd, patternMd, patternMd, patternMd
	}

	if len(patternMd) == 0 {
		patternMd = patterns.findClosest(idMd)

		if eqMonth {
			patternMd.replaceMonth("M")
		} else {
			patternMd.replaceMonth(patterns[0].month())
		}

		if eqDay {
			patternMd.replaceDay("d")
		}
	}

	if len(patternMMd) == 0 {
		patternMMd = patterns.findClosest(idMMd)

		if eqMonth {
			patternMMd.replaceMonth("MM")
		} else {
			patternMMd.replaceMonth(patterns[0].month())
		}

		if eqDay2D {
			patternMMd.replaceDay("d")
		}
	}

	if len(patternMdd) == 0 {
		patternMdd = patterns.findClosest(idMdd)

		if eqMonth2D {
			patternMdd.replaceMonth("M")
		} else {
			patternMdd.replaceMonth(patterns[0].month())
		}

		if eqDay {
			patternMdd.replaceDay("dd")
		}
	}

	if len(patternMMdd) == 0 {
		patternMMdd = patterns.findClosest(idMMdd)

		if eqMonth {
			patternMMdd.replaceMonth("MM")
		} else {
			patternMdd.replaceMonth(patterns[0].month())
		}

		if eqDay {
			patternMMdd.replaceDay("dd")
		}
	}

	return patternMd, patternMMd, patternMdd, patternMMdd
}

//nolint:gocognit
func buildFmtMD(formatMd, formatMMd, formatMdd, formatMMdd string, log *slog.Logger) string {
	patternMd, patternMMd, patternMdd, patternMMdd := monthDayPatterns(formatMd, formatMMd, formatMdd, formatMMdd)

	log.Debug("infer MD patterns",
		"Md", patternMd.String(),
		"MMd", patternMMd.String(),
		"Mdd", patternMdd.String(),
		"MMdd", patternMMdd.String())

	groups := groupLayouts(
		Layout{ID: idMd, Pattern: patternMd},
		Layout{ID: idMMd, Pattern: patternMMd},
		Layout{ID: idMdd, Pattern: patternMdd},
		Layout{ID: idMMdd, Pattern: patternMMdd},
	)

	var sb strings.Builder

	writePattern := func(pattern DatePattern, group LayoutGroup) {
		sb.WriteString("return ")

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

		switch group.Layouts[0].Pattern.month() {
		case "MMM":
			sb.WriteString("fmtMonth = fmtMonthName(locale.String(), \"stand-alone\", \"abbreviated\");")
		case "MMMMM":
			sb.WriteString("fmtMonth = fmtMonthName(locale.String(), \"stand-alone\", \"narrow\");")
		}

		sb.WriteString(" return func(m time.Month, d int) string { ")
		writePattern(group.Layouts[0].Pattern, group)
		sb.WriteString(" }")
	case 2: //nolint:mnd
		sb.WriteString("return func(m time.Month, d int) string { ")

		for _, group := range groups {
			if group.Expr != "" {
				sb.WriteString("if " + group.Expr + " { ")
				writePattern(group.Layouts[0].Pattern, group)
				sb.WriteString(" }; ")

				continue
			}

			for _, layout := range group.Layouts {
				if layout.ID.String() != "MMdd" {
					continue
				}

				writePattern(layout.Pattern, group)
			}
		}

		sb.WriteString(" }")
	}

	return sb.String()
}

// equalFlow returns true if all layouts have the same formatting pattern.
func (l Layout) equalFlow(other Layout) bool {
	if len(l.Pattern) != len(other.Pattern) {
		return false
	}

	if l.Pattern.String() == other.Pattern.String() {
		return true
	}

	eqMonth := func(otherMonth string) bool {
		month := l.Pattern.month()

		if l.ID.month() == month {
			return other.ID.month() == otherMonth
		}

		return len(month) == len(otherMonth)
	}

	eqDay := func(otherDay string) bool {
		day := l.Pattern.day()

		if l.ID.day() == day {
			return other.ID.day() == otherDay
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

// LayoutGroup contains similar formatting patterns for Golang code generation.
type LayoutGroup struct {
	Expr                     string
	Layouts                  []Layout
	FmtTypeMonth, FmtTypeDay FmtType
}

// groupLayouts groups layouts by their formatting patterns. It aims to reduce redundancy in generated Go code by
// identifying common patterns. The function analyzes the provided layouts, grouping them based on similarities in
// their formatting. The output is a slice of LayoutGroup structs, each representing a distinct group of layouts
// with similar formatting characteristics.
func groupLayouts(layouts ...Layout) []LayoutGroup {
	if len(layouts) == 0 {
		return nil
	}

	byPatterns := groupByPatterns(layouts)

	switch len(byPatterns) {
	case 1:
		// all have identical formatting pattern
		return processLayoutGroups([]LayoutGroup{{Layouts: layouts}})
	case 2: //nolint:mnd
		values := slices.Collect(maps.Values(byPatterns))
		if len(values[0]) > len(values[1]) {
			values[0], values[1] = values[1], values[0]
		}

		if len(values[0]) == 1 {
			// all have identical formatting pattern, except the first one
			return processLayoutGroups([]LayoutGroup{{Layouts: values[0]}, {Layouts: values[1]}})
		}
	}

	eqIDFmtMonth := make([]bool, len(layouts))
	eqIDFmtDay := make([]bool, len(layouts))
	eqPattern := slices.Repeat([]int{-1}, len(layouts))

	for i, layout := range layouts {
		eqIDFmtMonth[i] = layout.ID.month() == layout.Pattern.month()
		eqIDFmtDay[i] = layout.ID.day() == layout.Pattern.day()
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

	return processLayoutGroups(groups)
}

func processLayoutGroups(groups []LayoutGroup) []LayoutGroup {
	slices.SortFunc(groups, func(a, b LayoutGroup) int { return len(a.Layouts) - len(b.Layouts) })

	for i := range groups {
		groups[i].FmtTypeMonth = func() FmtType {
			if len(groups[i].Layouts) > 1 {
				m := groups[i].Layouts[0].Pattern.month()
				same := !slices.ContainsFunc(groups[i].Layouts[1:], func(l Layout) bool {
					return l.Pattern.month() != m
				})

				if same {
					switch len(m) {
					case 1:
						return FmtTypeNumericOnly
					case 2: //nolint:mnd
						return FmtType2DigitOnly
					}
				}
			}

			different := slices.ContainsFunc(groups[i].Layouts, func(l Layout) bool {
				return l.ID.month() != l.Pattern.month()
			})

			switch {
			default:
				return FmtType2DigitOnly
			case !different && len(groups[i].Layouts) > 1:
				return FmtTypeSame
			case groups[i].Layouts[0].Pattern.monthLen(1):
				return FmtTypeNumericOnly
			}
		}()

		groups[i].FmtTypeDay = func() FmtType {
			different := slices.ContainsFunc(groups[i].Layouts, func(l Layout) bool {
				return l.ID.day() != l.Pattern.day()
			})

			switch {
			default:
				return FmtType2DigitOnly
			case !different && len(groups[i].Layouts) > 1:
				return FmtTypeSame
			case groups[i].Layouts[0].Pattern.dayLen(1):
				return FmtTypeNumericOnly
			}
		}()

		if len(groups[i].Layouts) == 1 {
			optMonth := "MonthNumeric"
			optDay := "DayNumeric"

			if groups[i].Layouts[0].ID.monthLen(2) { //nolint:mnd
				optMonth = "Month2Digit"
			}

			if groups[i].Layouts[0].ID.dayLen(2) { //nolint:mnd
				optDay = "Day2Digit"
			}

			groups[i].Expr = fmt.Sprintf("opts.Month == %s && opts.Day == %s", optMonth, optDay)
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
