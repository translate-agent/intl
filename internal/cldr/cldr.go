package cldr

type CLDR struct {
	supplemental *Supplemental
	ldml         map[string]*LDML
	locales      []string
}

func (c *CLDR) Locales() []string {
	return c.locales
}

func (c *CLDR) RawLDML(locale string) *LDML {
	return c.ldml[locale]
}

func (c *CLDR) Supplemental() *Supplemental {
	return c.supplemental
}

// Common holds several of the most common attributes and sub elements
// of an XML element.
type Common struct {
	Type     string `xml:"type,attr,omitempty"`
	Draft    string `xml:"draft,attr,omitempty"`
	Alt      string `xml:"alt,attr,omitempty"`
	CharData string `xml:",chardata"`
}

func (c Common) isContributedOrApproved() bool {
	return c.Draft == "" || c.Draft == "contributed"
}

type LDML struct {
	Identity *Identity `xml:"identity"`
	Numbers  *Numbers  `xml:"numbers"`
	Dates    *Dates    `xml:"dates"`
	Common
}

type Identity struct {
	Language  *Language `xml:"language"`
	Script    *Common   `xml:"script"`
	Territory *Common   `xml:"territory"`
	Common
}

type Language struct {
	Type string `xml:"type,attr"`
	Common
}

type Script struct {
	Type string `xml:"type,attr"`
	Common
}

type Territory struct {
	Type string `xml:"type,attr"`
	Common
}

type Variant struct {
	Type string `xml:"type,attr"`
	Common
}

type Numbers struct {
	Common
	DefaultNumberingSystem []*Common `xml:"defaultNumberingSystem"`
}

type Dates struct {
	Calendars *Calendars `xml:"calendars"`
	Fields    *Fields    `xml:"fields"`
	Common
}

type Calendars struct {
	Common
	Calendar []*Calendar `xml:"calendar"`
}

type Calendar struct {
	Alias           *Alias           `xml:"alias"`
	Months          *Months          `xml:"months"`
	DateTimeFormats *DateTimeFormats `xml:"dateTimeFormats"`
	Eras            *Eras            `xml:"eras"`
	Common
}

type Months struct {
	Common
	MonthContext []*MonthContext `xml:"monthContext"`
}

type MonthContext struct {
	Common
	MonthWidth []*MonthWidth `xml:"monthWidth"`
}

type MonthWidth struct {
	Common
	Month []*Month `xml:"month"`
}

type Month struct {
	Common
}

type DateTimeFormats struct {
	Common
	Alias            *Alias              `xml:"alias"`
	AvailableFormats []*AvailableFormats `xml:"availableFormats"`
}

type AvailableFormats struct {
	Common
	DateFormatItem []*DateFormatItem `xml:"dateFormatItem"`
}

type DateFormatItem struct {
	Common
	ID string `xml:"id,attr"`
}

type Eras struct {
	EraNames  *Era `xml:"eraNames"`
	EraAbbr   *Era `xml:"eraAbbr"`
	EraNarrow *Era `xml:"eraNarrow"`
	Common
}

type Era struct {
	Common
	Era []*Common `xml:"era"`
}

type Fields struct {
	Common
	Field []*Field `xml:"field"`
}

type Field struct {
	Common
	DisplayName []*DisplayName `xml:"displayName"`
}

type DisplayName struct {
	Common
	Count string `xml:"count,attr"`
}

type Supplemental struct {
	Common
	CalendarPreferenceData *CalendarPreferenceData `xml:"calendarPreferenceData"`
	NumberingSystems       *NumberingSystems       `xml:"numberingSystems"`
	ParentLocales          []*ParentLocales        `xml:"parentLocales"`
}

type ParentLocales struct {
	Common
	Component    string          `xml:"component,attr"`
	ParentLocale []*ParentLocale `xml:"parentLocale"`
}

type ParentLocale struct {
	Common
	Parent  string `xml:"parent,attr"`
	Locales string `xml:"locales,attr"`
}

type Alias struct {
	Common
	Source string `xml:"source,attr"`
	Path   string `xml:"path,attr"`
}

type CalendarPreferenceData struct {
	Common
	CalendarPreference []*CalendarPreference `xml:"calendarPreference"`
}

type CalendarPreference struct {
	Common
	Territories string `xml:"territories,attr"`
	Ordering    string `xml:"ordering,attr"`
}

type NumberingSystems struct {
	Common
	NumberingSystem []*NumberingSystem `xml:"numberingSystem"`
}

type NumberingSystem struct {
	Common
	ID     string `xml:"id,attr"`
	Type   string `xml:"type,attr"`
	Digits string `xml:"digits,attr"`
}
