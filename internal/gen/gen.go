package main

import (
	_ "embed"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"iter"
	"log/slog"
	"maps"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"text/template"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/cldr"
)

//go:embed cldr_data.go.tmpl
var cldrDataTemplate string

// LocaleLookup maps a shared property (the key) to a list of locales that share that property.
type LocaleLookup map[string][]string

type Generator struct {
	cldr *cldr.CLDR
}

type Conf struct {
	cldrDir    string
	out        string
	saveMerged bool
}

func Gen(conf Conf, log *slog.Logger) error {
	var g Generator

	if err := g.load(conf.cldrDir, log); err != nil {
		return err
	}

	data, err := g.process(conf, log)
	if err != nil {
		return err
	}

	if err := g.write(conf.out, data); err != nil {
		return err
	}

	return nil
}

func (g *Generator) load(dir string, log *slog.Logger) error {
	var (
		d   cldr.Decoder
		err error
	)

	now := time.Now()

	defer func() {
		log.Debug("loading", "duration", time.Since(now))
	}()

	d.SetDirFilter("main", "supplemental")

	g.cldr, err = d.DecodePath(dir)
	if err != nil {
		return fmt.Errorf(`decode CLDR at path "%s": %w`, dir, err)
	}

	return nil
}

func (g *Generator) process(conf Conf, log *slog.Logger) (*TemplateData, error) {
	g.filterApproved()
	g.merge(log)

	if conf.saveMerged {
		if err := g.saveMerged(conf.out); err != nil {
			return nil, err
		}
	}

	defaultNumberingSystems := g.defaultNumberingSystems()
	calendarPreferences := g.calendarPreferences()

	return &TemplateData{
		Eras:                    g.eras(calendarPreferences),
		CalendarPreferences:     calendarPreferences,
		NumberingSystems:        g.numberingSystems(defaultNumberingSystems),
		NumberingSystemIota:     g.numberingSystemsIota(defaultNumberingSystems),
		DefaultNumberingSystems: defaultNumberingSystems,
		Fields:                  g.fields(),
		Months:                  g.months(),
	}, nil
}

func (g *Generator) write(out string, data *TemplateData) error {
	dataTpl, err := template.New("cldr_data").Funcs(template.FuncMap{
		"join":  strings.Join,
		"sub":   func(a, b int) int { return a - b },
		"title": title,
	}).Parse(cldrDataTemplate)
	if err != nil {
		return fmt.Errorf("parse cldr_data: %w", err)
	}

	name := path.Join(out, "cldr_data.go")

	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("create file %s: %w", name, err)
	}

	defer f.Close()

	if err := dataTpl.Execute(f, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	return nil
}

func (g *Generator) saveMerged(out string) error {
	for _, locale := range g.cldr.Locales() {
		name := path.Join(out, ".cldr_merged", locale+".xml")

		f, err := os.Create(name)
		if err != nil {
			return fmt.Errorf("create %s: %w", name, err)
		}

		defer f.Close()

		ldml := g.cldr.RawLDML(locale)
		enc := xml.NewEncoder(f)
		enc.Indent("", "\t")

		if err = enc.Encode(ldml); err != nil {
			return fmt.Errorf("save to %s: %w", name, err)
		}
	}

	return nil
}

//nolint:cyclop,gocognit
func (g *Generator) filterApproved() {
	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Numbers != nil {
			var defaultNumberingSystem []*cldr.Common

			for _, v := range ldml.Numbers.DefaultNumberingSystem {
				if isContributedOrApproved(v.Draft) {
					defaultNumberingSystem = append(defaultNumberingSystem, v)
				}
			}

			ldml.Numbers.DefaultNumberingSystem = defaultNumberingSystem
		}

		if ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		for _, calendar := range ldml.Dates.Calendars.Calendar {
			if calendar.Months != nil {
				for _, monthContext := range calendar.Months.MonthContext {
					for _, monthWidth := range monthContext.MonthWidth {
						var months []*struct {
							cldr.Common
							Yeartype string `xml:"yeartype,attr"`
						}

						for _, month := range monthWidth.Month {
							if isContributedOrApproved(month.Draft) {
								months = append(months, month)
							}
						}

						monthWidth.Month = months
					}
				}
			}

			if calendar.DateTimeFormats != nil {
				for _, dateTimeFormat := range calendar.DateTimeFormats.AvailableFormats {
					var dateFormatItems []*struct {
						cldr.Common
						Id    string `xml:"id,attr"` //nolint:revive,stylecheck
						Count string `xml:"count,attr"`
					}

					for _, dateFormatItem := range dateTimeFormat.DateFormatItem {
						if isContributedOrApproved(dateFormatItem.Draft) {
							dateFormatItems = append(dateFormatItems, dateFormatItem)
						}
					}

					dateTimeFormat.DateFormatItem = dateFormatItems
				}
			}

			//nolint:nestif
			if eras := calendar.Eras; eras != nil {
				if v := eras.EraAbbr; v != nil {
					for i := len(v.Era) - 1; i >= 0; i-- {
						if era := v.Era[i]; !isContributedOrApproved(era.Draft) || era.Alt != "" {
							v.Era = slices.Delete(v.Era, i, i+1)
						}
					}

					if len(v.Era) == 0 {
						eras.EraAbbr = nil
					}
				}

				if v := eras.EraNames; v != nil {
					for i := len(v.Era) - 1; i >= 0; i-- {
						if era := v.Era[i]; !isContributedOrApproved(era.Draft) || era.Alt != "" {
							v.Era = slices.Delete(v.Era, i, i+1)
						}
					}

					if len(v.Era) == 0 {
						eras.EraNames = nil
					}
				}

				if v := calendar.Eras.EraNarrow; v != nil {
					for i := len(v.Era) - 1; i >= 0; i-- {
						if era := v.Era[i]; !isContributedOrApproved(era.Draft) || era.Alt != "" {
							v.Era = slices.Delete(v.Era, i, i+1)
						}
					}

					if len(v.Era) == 0 {
						eras.EraNarrow = nil
					}
				}

				if eras.EraAbbr == nil && eras.EraNames == nil && eras.EraNarrow == nil {
					calendar.Eras = nil
				}
			}
		}

		if ldml.Dates.Fields != nil {
			for i := len(ldml.Dates.Fields.Field) - 1; i >= 0; i-- {
				field := ldml.Dates.Fields.Field[i]
				if len(field.DisplayName) == 0 || !isContributedOrApproved(field.DisplayName[0].Draft) {
					ldml.Dates.Fields.Field = slices.Delete(ldml.Dates.Fields.Field, i, i+1)
				}
			}
		}
	}
}

func (g *Generator) merge(log *slog.Logger) {
	g.mergeAliases()
	g.mergeParent(log)
	g.mergeLocal(log)

	root := g.cldr.RawLDML("root")

	// merge parent to child
	for _, parentLocale := range g.cldr.Supplemental().ParentLocales.ParentLocale {
		// ignore, cldr package does NOT have the attribute "component"
		// <parentLocales component="collations">
		// 	<parentLocale parent="sr_ME" locales="sr_Cyrl_ME"/>
		// 	<parentLocale parent="zh_Hant" locales="yue yue_Hant"/>
		// 	<parentLocale parent="zh_Hans" locales="yue_CN yue_Hans yue_Hans_CN"/>
		// </parentLocales>
		if slices.Contains([]string{"sr_ME", "zh_Hant", "zh_Hans"}, parentLocale.Parent) {
			continue
		}

		parent := g.cldr.RawLDML(parentLocale.Parent)

		// merge root to parent
		merge(parent, root, log)

		for _, locale := range strings.Split(parentLocale.Locales, " ") {
			child := g.cldr.RawLDML(locale)

			if child == nil {
				continue
			}

			merge(child, parent, log)
		}
	}

	// merge root to language

	for _, locale := range g.cldr.Locales() {
		if strings.ContainsRune(locale, '_') {
			continue
		}

		ldml := g.cldr.RawLDML(locale)

		merge(ldml, root, log)
	}

	// merge language to territory
	for _, locale := range g.cldr.Locales() {
		if !strings.ContainsRune(locale, '_') {
			continue
		}

		ldml := g.cldr.RawLDML(locale)

		parts := strings.Split(locale, "_")
		parts = parts[:len(parts)-1]

		fallback := g.cldr.RawLDML(strings.Join(parts, "_"))

		merge(ldml, fallback, log)
	}
}

func (g *Generator) mergeAliases() {
	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		for i, calendar := range ldml.Dates.Calendars.Calendar {
			if calendar.Alias != nil {
				calendarType := strings.Split(calendar.Alias.Path, "'")[1]
				calendar = findCalendar(ldml, calendarType)
				ldml.Dates.Calendars.Calendar[i] = calendar

				continue
			}

			if calendar.DateTimeFormats == nil || calendar.DateTimeFormats.Alias == nil {
				continue
			}

			// example: ../../calendar[@type='generic']/dateTimeFormats
			calendarType := strings.Split(calendar.DateTimeFormats.Alias.Path, "'")[1]

			calendar.DateTimeFormats = findCalendar(ldml, calendarType).DateTimeFormats
		}
	}
}

func (g *Generator) mergeParent(log *slog.Logger) {
	log = log.With("func", "mergeParent")

	for _, locale := range g.cldr.Locales() {
		if locale == "root" {
			continue
		}

		// main language, skip it
		parts := strings.Split(locale, "_")
		if len(parts) == 1 {
			continue
		}

		// <lang>_<script>_<region>
		// * use <lang>, if <lang>_<script> or <lang>_<region>
		// * use <lang>-<script>, if <lang>_<script>_<region>
		parentLocale := parts[0]
		if len(parts) == 3 { //nolint:mnd
			parentLocale += "_" + parts[1]
		}

		parent := g.cldr.RawLDML(parentLocale)
		if parent.Dates == nil {
			continue
		}

		ldml := g.cldr.RawLDML(locale)

		if ldml.Dates == nil {
			ldml.Dates = deepCopy(parent.Dates)
			continue
		}

		mergeFields(ldml, parent)

		if parent.Dates.Calendars != nil {
			parentGregorian := findCalendar(parent, "gregorian")

			if ldml.Dates.Calendars == nil {
				continue
			}

			logger := log.With("locale", locale)

			calendar := findCalendar(ldml, "gregorian")
			if calendar == nil {
				logger.Debug("copy gregorian calendar")

				ldml.Dates.Calendars.Calendar = append(ldml.Dates.Calendars.Calendar, deepCopy(parentGregorian))

				continue
			}

			mergeCalendar(calendar, parentGregorian, logger)
		}
	}
}

func (g *Generator) mergeLocal(log *slog.Logger) {
	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Identity.Language.Type == "root" || ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		// merge generic calendar to persian or buddhist calendar
		generic := findCalendar(ldml, "generic")

		if generic == nil || generic.DateTimeFormats == nil {
			continue
		}

		for _, calendar := range ldml.Dates.Calendars.Calendar {
			if !slices.Contains([]string{"persian", "buddhist"}, calendar.Type) {
				continue
			}

			if calendar.DateTimeFormats == nil {
				calendar.DateTimeFormats = deepCopy(generic.DateTimeFormats)
			}

			mergeCalendar(calendar, generic, log.With("locale", locale))
		}
	}
}

// findCalendar returns *cldr.Calendar by its type if found. Otherwise, returns nil.
func findCalendar(ldml *cldr.LDML, calendarType string) *cldr.Calendar {
	for _, v := range ldml.Dates.Calendars.Calendar {
		if v.Type == calendarType {
			return v
		}
	}

	return nil
}

// containsDateFormatItem returns true if calendar contains dateFormatItem with given id.
func containsDateFormatItem(calendar *cldr.Calendar, id string) bool {
	for _, v := range calendar.DateTimeFormats.AvailableFormats[0].DateFormatItem {
		if v.Id == id {
			return true
		}
	}

	return false
}

var supportedCalendarTypes = []string{"gregorian", "persian", "buddhist"}

func supportedCalendars(calendars []*cldr.Calendar) iter.Seq[*cldr.Calendar] {
	return func(yield func(*cldr.Calendar) bool) {
		for _, v := range calendars {
			if slices.Contains(supportedCalendarTypes, v.Type) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// merge copies particular src values to dst.
func merge(dst, src *cldr.LDML, log *slog.Logger) {
	if src.Dates == nil || src.Dates.Calendars == nil {
		return
	}

	if dst.Dates == nil {
		dst.Dates = deepCopy(src.Dates)
	}

	mergeFields(dst, src)

	if dst.Dates.Calendars == nil {
		dst.Dates.Calendars = deepCopy(src.Dates.Calendars)
	}

	if len(dst.Dates.Calendars.Calendar) == 0 {
		dst.Dates.Calendars.Calendar = deepCopy(src.Dates.Calendars.Calendar)
	}

	for parentCalendar := range supportedCalendars(src.Dates.Calendars.Calendar) {
		calendar := findCalendar(dst, parentCalendar.Type)
		if calendar == nil {
			calendar = deepCopy(parentCalendar)
			dst.Dates.Calendars.Calendar = append(dst.Dates.Calendars.Calendar, calendar)

			continue
		}

		mergeCalendar(calendar, parentCalendar, log)
	}
}

func mergeCalendar(dst, src *cldr.Calendar, log *slog.Logger) {
	log.Debug("merge calendars", "dst", dst.Type, "src", src.Type)

	switch dst.DateTimeFormats {
	default:
		if dst.DateTimeFormats.AvailableFormats == nil && src.DateTimeFormats != nil {
			dst.DateTimeFormats.AvailableFormats = deepCopy(src.DateTimeFormats.AvailableFormats)
		}

		if src.DateTimeFormats != nil {
			for _, availableFormats := range src.DateTimeFormats.AvailableFormats {
				for _, dateFormatItem := range availableFormats.DateFormatItem {
					if containsDateFormatItem(dst, dateFormatItem.Id) {
						continue
					}

					// NOTE(jhorsts): Why the first AvailableFormats? I don't remember.
					dst.DateTimeFormats.AvailableFormats[0].DateFormatItem = append(
						dst.DateTimeFormats.AvailableFormats[0].DateFormatItem,
						deepCopy(dateFormatItem))
				}
			}
		}
	case nil:
		dst.DateTimeFormats = deepCopy(src.DateTimeFormats)
	}

	// months
	if src.Months != nil {
		if dst.Months == nil {
			dst.Months = deepCopy(src.Months)
		}

		if dst.Months != nil && len(dst.Months.MonthContext) == 0 && len(src.Months.MonthContext) > 0 {
			dst.Months.MonthContext = deepCopy(src.Months.MonthContext)
		}
	}

	// eras
	if src.Eras != nil {
		if dst.Eras == nil {
			dst.Eras = deepCopy(src.Eras)
		}
	}
}

func mergeFields(dst, src *cldr.LDML) {
	if src.Dates == nil || src.Dates.Fields == nil {
		return
	}

	if dst.Dates == nil {
		dst.Dates = deepCopy(src.Dates)
		return
	}

	if dst.Dates.Fields == nil {
		dst.Dates.Fields = deepCopy(src.Dates.Fields)
		return
	}

	if len(dst.Dates.Fields.Field) == 0 {
		dst.Dates.Fields.Field = deepCopy(src.Dates.Fields.Field)
		return
	}

	for _, field := range src.Dates.Fields.Field {
		found := func() bool {
			for _, v := range dst.Dates.Fields.Field {
				if v.Type == field.Type {
					return true
				}
			}

			return false
		}()

		if !found {
			dst.Dates.Fields.Field = append(dst.Dates.Fields.Field, deepCopy(field))
		}
	}
}

func (g *Generator) calendarPreferences() CalendarPreferences {
	calendarPreferences := g.cldr.Supplemental().CalendarPreferenceData.CalendarPreference
	preferences := make(CalendarPreferences, 0, len(calendarPreferences))

	// calendar preferences
	for _, v := range calendarPreferences {
		preferences = append(preferences, CalendarPreference{
			Regions:   strings.Split(v.Territories, " "),
			Calendars: strings.Split(v.Ordering, " "),
		})
	}

	return preferences
}

func (g *Generator) defaultNumberingSystems() LocaleLookup {
	defaultNumberingSystems := make(LocaleLookup)

	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		if ldml.Numbers == nil || locale == "root" {
			continue
		}

		for _, v := range ldml.Numbers.DefaultNumberingSystem {
			if v.Alt != "" {
				continue
			}

			defaultNumberingSystems[v.CharData] = append(defaultNumberingSystems[v.CharData],
				strings.ReplaceAll(locale, "_", "-"))
		}
	}

	return defaultNumberingSystems
}

func (g *Generator) months() Months { //nolint:gocognit
	months := NewMonths()

	for _, locale := range g.cldr.Locales() {
		ldml := g.cldr.RawLDML(locale)

		locale = strings.ReplaceAll(locale, "_", "-")

		if ldml.Dates == nil || ldml.Dates.Calendars == nil {
			continue
		}

		// month names are available only in gregorian calendars (default)
		calendar := findCalendar(ldml, "gregorian")

		if calendar.Months == nil {
			continue
		}

		for _, monthContext := range calendar.Months.MonthContext {
			for _, monthWidth := range monthContext.MonthWidth {
				if len(monthWidth.Month) == 0 {
					continue
				}

				month := monthWidth.Month[0]

				// skip months with the same digits
				if month.Type == month.CharData && month.CharData == "1" {
					continue
				}

				var monthNames MonthNames

				for _, month = range monthWidth.Month {
					i, err := strconv.Atoi(month.Type)
					if err != nil {
						panic(err)
					}

					i--

					monthNames[i] = month.CharData
				}

				// skip empty names
				if monthNames[0] == "" {
					continue
				}

				i := slices.IndexFunc(months.List, func(names MonthNames) bool {
					for i, v := range names {
						if v != monthNames[i] {
							return false
						}
					}

					return true
				})

				if i == -1 {
					months.List = append(months.List, monthNames)
					i = len(months.List) - 1
				}

				indexes := months.Lookup[locale]
				indexes.Set(monthWidth.Type, monthContext.Type, i)

				// NOTE: fallback "format" context when "stand-alone" not defined
				if monthContext.Type == "format" {
					indexes.Set(monthWidth.Type, "stand-alone", i)
				}

				months.Lookup[locale] = indexes
			}
		}
	}

	return months
}

//nolint:gocognit
func (g *Generator) fields() Fields {
	fields := make(Fields)

	for _, locale := range g.cldr.Locales() {
		if locale == "root" {
			continue
		}

		ldml := g.cldr.RawLDML(locale)

		if ldml.Dates == nil {
			continue
		}

		locale = strings.ReplaceAll(locale, "_", "-")

		if ldml.Dates.Fields != nil {
			for _, field := range ldml.Dates.Fields.Field {
				if len(field.DisplayName) == 0 {
					continue
				}

				f := fields[locale]
				v := field.DisplayName[0].CharData

				switch field.Type {
				default:
					continue
				case "month":
					f.Month = v
				case "month-short":
					f.MonthShort = v
				case "day":
					f.Day = v
				}

				fields[locale] = f
			}
		}
	}

	for _, locale := range []string{"en-Dsrt", "en-Shaw"} {
		f := fields[locale]
		f.Day = "day"
		fields[locale] = f
	}

	// remove the entries if the language has the same values
	for k, v := range fields {
		lang := strings.Split(k, "-")[0]

		if k != lang {
			continue
		}

		for k2, f2 := range fields {
			if lang == k2 {
				continue
			}

			if lang == strings.Split(k2, "-")[0] &&
				v.Month == f2.Month && v.Day == f2.Day {
				delete(fields, k2)
			}
		}
	}

	for locale, f := range fields {
		if f.Month == "" && f.MonthShort != "" {
			f.Month = f.MonthShort
		}

		if f.Month == "" {
			f.Month = "Day"
		}

		if f.Day == "" {
			f.Day = "Day"
		}

		fields[locale] = f
	}

	// Correct the naming! The naming is different in Node.js.

	// year, day formatting
	for _, locale := range []string{"en-Dsrt", "en-Dsrt-US", "en-Shaw", "en-Shaw-GB"} {
		f := fields[locale]
		f.Month = "month"
		f.Day = "day"
		fields[locale] = f
	}

	for _, locale := range []string{
		"az-Cyrl",
		"uz-Arab",
	} {
		f := fields[locale]
		f.Month = "Month"
		f.Day = "Day"
		fields[locale] = f
	}

	f := fields["nn"]
	f.Day = "dag"
	fields["nn"] = f

	f = fields["mn-Mong-MN"]
	f.Month = "сар"
	f.Day = "өдөр"
	fields["mn-Mong-MN"] = f

	for locale, v := range fields {
		// NOTE! all "Day" values at language level can be deleted (manually verified).
		// Correct way is to verify that all scripts and regions have the "Day" value.
		if !strings.Contains(locale, "-") && v.Day == "Day" {
			delete(fields, locale)
		}
	}

	return fields
}

//nolint:cyclop,gocognit
func (g *Generator) eras(calendarPreferences CalendarPreferences) Eras {
	eras := make(Eras)

	for _, locale := range g.cldr.Locales() {
		if locale == "root" {
			continue
		}

		ldml := g.cldr.RawLDML(locale)
		calendar := findCalendar(ldml, calendarPreferences.FindCalendarType(locale))

		if calendar == nil || calendar.Eras == nil {
			continue
		}

		eraType := "0"
		if calendar.Type != "persian" && calendar.Type != "buddhist" {
			eraType = "1"
		}

		locale = strings.ReplaceAll(locale, "_", "-")
		lang, _, region := language.Make(locale).Raw()

		var era Era

		f := func(s string) string {
			if s == "d. C." && locale != "es" && lang.String() == "es" &&
				!slices.Contains([]string{"EA", "ES", "GQ", "IC", "PH"}, region.String()) {
				return "d.C."
			}

			return s
		}

		// narrow
		if calendar.Eras.EraNarrow != nil {
			for _, v := range calendar.Eras.EraNarrow.Era {
				if v.Type == eraType && v.Alt == "" {
					era.Narrow = f(v.CharData)
				}
			}
		}

		// short
		if calendar.Eras.EraAbbr != nil {
			for _, v := range calendar.Eras.EraAbbr.Era {
				if v.Type == eraType && v.Alt == "" {
					era.Short = f(v.CharData)
				}
			}
		}

		// long
		if calendar.Eras.EraNames != nil {
			for _, v := range calendar.Eras.EraNames.Era {
				if v.Type == eraType && v.Alt == "" {
					era.Long = v.CharData
				}
			}
		}

		switch locale {
		case "en-Dsrt", "en-Dsrt-US", "en-Shaw", "en-Shaw-GB":
			era.Narrow = "A"
			era.Short = "AD"
			era.Long = "Anno Domini"
		case "bg", "bg-BG":
			era.Long = "след Христа"
		case "cy", "cy-GB":
			era.Long = "Oed Crist"
		case "es-419":
			era.Long = "después de Cristo"
		case "es-DO":
			era.Narrow = "d.C."
			era.Short = "d.C."
			era.Long = "después de Cristo"
		case "fa", "fa-AF", "fa-IR":
			era.Narrow = "ه\u200d.ش."
			era.Short = "ه\u200d.ش."
			era.Long = "هجری شمسی"
		case "hi-Latn", "hi-Latn-IN":
			era.Narrow = "A"
			era.Short = "AD"
			era.Long = "Anno Domini"
		case "lrc":
			era.Narrow = "AP"
			era.Short = "AP"
			era.Long = "AP"
		case "mk", "mk-MK":
			era.Narrow = "н. е."
			era.Short = "н. е."
			era.Long = "од нашата ера"
		case "mn-Mong-MN":
			era.Narrow = "МЭ"
			era.Short = "МЭ"
			era.Long = "манай эриний"
		case "mzn", "ps", "uz-Arab":
			era.Narrow = "AP"
			era.Short = "AP"
			era.Long = "AP"
		case "th":
			era.Narrow = "พ.ศ."
			era.Short = "พ.ศ."
			era.Long = "พุทธศักราช"
		case "zh-Hant-MO":
			era.Narrow = "公元"
			era.Short = "公元"
			era.Long = "公元"
		}

		if era.Long == "" && era.Narrow == "" && era.Short == "" {
			continue
		}

		if era.Short != "" {
			if era.Narrow == "" {
				era.Narrow = era.Short
			}

			if era.Long == "" {
				era.Long = era.Short
			}
		}

		if era.Narrow == "" {
			era.Narrow = "CE"
		}

		if era.Short == "" {
			era.Short = "CE"
		}

		if era.Long == "" {
			era.Long = "CE"
		}

		eras[locale] = era
	}

	return eras
}

// CLDRDateFormatItem is a copy of CLDR DateFormatItem.
type CLDRDateFormatItem struct {
	cldr.Common
	Id    string //nolint:revive,stylecheck
	Count string
}

func (g *Generator) numberingSystems(defaultNumberingSystems LocaleLookup) []NumberingSystem {
	numberingSystems := make([]NumberingSystem, 0, 12) //nolint:mnd

	ids := slices.Collect(maps.Keys(defaultNumberingSystems))

	for _, v := range g.cldr.Supplemental().NumberingSystems.NumberingSystem {
		if v.Type != "numeric" || !slices.Contains(ids, v.Id) {
			continue
		}

		numberingSystem := NumberingSystem{ID: v.Id}

		var i int

		for _, digit := range v.Digits {
			numberingSystem.Digits[i] = digit
			i++
		}

		numberingSystems = append(numberingSystems, numberingSystem)
	}

	return numberingSystems
}

func (g *Generator) numberingSystemsIota(defaultNumberingSystems LocaleLookup) []string {
	ids := slices.Collect(maps.Keys(defaultNumberingSystems))

	slices.Sort(ids)

	return slices.Compact(ids)
}

type NumberingSystem struct {
	ID     string
	Digits [10]rune
}

type TemplateData struct {
	Eras                    Eras
	Months                  Months
	Fields                  Fields
	DefaultNumberingSystems LocaleLookup
	NumberingSystemIota     []string
	CalendarPreferences     CalendarPreferences
	NumberingSystems        []NumberingSystem
}

// value - locales.
type Months struct {
	// key is locale, value is 6 indexes from [List].
	Lookup map[string]MonthIndexes
	List   []MonthNames
}

func NewMonths() Months {
	return Months{
		Lookup: make(map[string]MonthIndexes),
	}
}

type MonthKey struct {
	Locale       string
	CalendarType string // gregorian, persian or buddhist
	Width        string // wide, narrow, abbreviated
	Context      string // format or stand-alone
}

// MonthIndexes contains indexes for month names in [Months.List]:
//
//	0 - abbreviated, format
//	1 - abbreviated, stand-alone
//	2 - wide, format
//	3 - wide, stand-alone
//	4 - narrow, format
//	5 - narrow, stand-alone
type MonthIndexes [6]int

func (m *MonthIndexes) Set(width, context string, i int) {
	contextCount := 2

	var w, c int

	switch width {
	case "abbreviated":
		w = 0
	case "wide":
		w = 1
	case "narrow":
		w = 2
	}

	switch context {
	case "format":
		c = 0
	case "stand-alone":
		c = 1
	}

	index := w*contextCount + c

	m[index] = i
}

type MonthNames [12]string

func (n MonthNames) String() string {
	return `{"` + strings.Join(n[:], `", "`) + `"}`
}

type Fields map[string]Field

type Field struct {
	Month, MonthShort, Day string
}

// Era contains current era and default calendar only.
type Era struct {
	Narrow, Short, Long string
}

type Eras map[string]Era

type CalendarPreferences []CalendarPreference

func (c CalendarPreferences) FindCalendarType(locale string) string {
	lang, _, region := language.Make(locale).Raw()

	if lang.String() == "az" {
		return "gregorian"
	}

	for _, v := range c {
		if slices.Contains(v.Regions, region.String()) {
			return v.Calendars[0]
		}
	}

	return "gregorian"
}

type CalendarPreference struct {
	Regions   []string
	Calendars []string
}

func Locale(ldml *cldr.LDML) string {
	lang := ldml.Identity.Language.Type

	if ldml.Identity.Script != nil && ldml.Identity.Script.Type != "" {
		lang += "-" + ldml.Identity.Script.Type
	}

	if ldml.Identity.Territory != nil && ldml.Identity.Territory.Type != "" {
		lang += "-" + ldml.Identity.Territory.Type
	}

	return lang
}

//nolint:ireturn
func deepCopy[T any](v T) T {
	var r T

	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(b, &r); err != nil {
		panic(err)
	}

	return r
}

func isContributedOrApproved(draft string) bool {
	return draft == "" || draft == "contributed"
}

func title(s string) string {
	var r string

	for _, v := range strings.Split(s, " ") {
		r += cases.Title(language.English).String(v)
	}

	return strings.ReplaceAll(r, "-", "") // e.g. "islamic - umalqura"
}
