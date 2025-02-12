package cldr

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"slices"
	"strings"
	"sync"
)

func DecodePath(dir string) (*CLDR, error) {
	commonMain := os.DirFS(path.Join(dir, "common/main"))

	const numberOfLocales = 1067
	cldr := CLDR{
		ldml: make(map[string]*LDML, numberOfLocales),
	}

	files, err := fs.Glob(commonMain, "*.xml")
	if err != nil {
		return nil, fmt.Errorf("get all filenames in common/main: %w", err)
	}

	type result struct {
		err    error
		ldml   *LDML
		locale string
	}

	var wg sync.WaitGroup

	wg.Add(len(files))

	ch := make(chan result)

	for _, file := range files {
		go func() {
			defer wg.Done()

			b, fErr := fs.ReadFile(commonMain, file)
			if fErr != nil {
				ch <- result{err: fmt.Errorf("read %s: %w", file, fErr)}
				return
			}

			var ldml LDML

			if fErr = xml.Unmarshal(b, &ldml); fErr != nil {
				ch <- result{err: fmt.Errorf("unmarshal %s: %w", file, fErr)}
				return
			}

			cleanLDML(&ldml)

			locale, _ := strings.CutSuffix(file, ".xml")

			ch <- result{locale: locale, ldml: &ldml}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		if r.err != nil {
			err = errors.Join(err, r.err)
		} else {
			cldr.ldml[r.locale] = r.ldml
		}
	}

	if err != nil {
		return nil, err
	}

	locales := make([]string, 0, len(cldr.ldml))

	locales = append(locales, "root")

	for v := range cldr.ldml {
		if v != "root" {
			locales = append(locales, v)
		}
	}

	slices.Sort(locales[1:])

	cldr.locales = locales

	b, err := fs.ReadFile(os.DirFS(path.Join(dir, "common/supplemental")), "supplementalData.xml")
	if err != nil {
		return nil, fmt.Errorf("read common/supplemental/supplementalData.xml: %w", err)
	}

	if err = xml.Unmarshal(b, &cldr.supplemental); err != nil {
		return nil, fmt.Errorf("decode common/supplemental/supplementalData.xml: %w", err)
	}

	b, err = fs.ReadFile(os.DirFS(path.Join(dir, "common/supplemental")), "numberingSystems.xml")
	if err != nil {
		return nil, fmt.Errorf("read common/supplemental/numberingSystems.xml: %w", err)
	}

	if err = xml.Unmarshal(b, &cldr.supplemental); err != nil {
		return nil, fmt.Errorf("decode common/supplemental/numberingSystems.xml: %w", err)
	}

	return &cldr, nil
}

// cleanLDML removes unnecessary data from the LDML struct.
func cleanLDML(ldml *LDML) {
	if numbers := ldml.Numbers; numbers != nil {
		numbers.DefaultNumberingSystem = filter(numbers.DefaultNumberingSystem, func(v *Common) bool {
			return v.isContributedOrApproved()
		})
	}

	dates := ldml.Dates
	if dates == nil || dates.Calendars == nil {
		return
	}

	for _, calendar := range dates.Calendars.Calendar {
		if calendar.Months != nil {
			for _, monthContext := range calendar.Months.MonthContext {
				for _, monthWidth := range monthContext.MonthWidth {
					monthWidth.Month = filter(monthWidth.Month, func(v *Month) bool {
						return v.isContributedOrApproved()
					})
				}
			}
		}

		if calendar.DateTimeFormats != nil {
			for _, dateTimeFormat := range calendar.DateTimeFormats.AvailableFormats {
				dateTimeFormat.DateFormatItem = filter(dateTimeFormat.DateFormatItem, func(v *DateFormatItem) bool {
					return v.isContributedOrApproved()
				})
			}
		}

		if eras := calendar.Eras; eras != nil {
			eras.EraAbbr = cleanEra(eras.EraAbbr)
			eras.EraNames = cleanEra(eras.EraNames)
			eras.EraNarrow = cleanEra(eras.EraNarrow)

			if eras.EraAbbr == nil && eras.EraNames == nil && eras.EraNarrow == nil {
				calendar.Eras = nil
			}
		}

		if dates.Fields != nil {
			for i := len(dates.Fields.Field) - 1; i >= 0; i-- {
				dates.Fields.Field = filter(dates.Fields.Field, func(v *Field) bool {
					return len(v.DisplayName) > 0 && v.DisplayName[0].isContributedOrApproved()
				})
			}
		}
	}
}

func filter[T any](s []T, f func(T) bool) []T {
	var r []T

	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func cleanEra(era *Era) *Era {
	if era == nil {
		return nil
	}

	era.Era = filter(era.Era, func(v *Common) bool {
		return v.isContributedOrApproved() && v.Alt == ""
	})

	if len(era.Era) == 0 {
		return nil
	}

	return era
}
