package entities

type Country struct {
	Name         Name         `json:"name"`
	Tld          []string     `json:"tld"`
	Cca2         string       `json:"cca2"`
	Ccn3         string       `json:"ccn3"`
	Cca3         string       `json:"cca3"`
	Cioc         string       `json:"cioc"`
	Independent  bool         `json:"independent"`
	Status       string       `json:"status"`
	UnMember     bool         `json:"unMember"`
	Currencies   Currencies   `json:"currencies"`
	Idd          Idd          `json:"idd"`
	Capital      []string     `json:"capital"`
	AltSpellings []string     `json:"altSpellings"`
	Region       string       `json:"region"`
	Subregion    string       `json:"subregion"`
	Languages    Languages    `json:"languages"`
	Translations Translations `json:"translations"`
	Latlng       []float32    `json:"latlng"`
	Landlocked   bool         `json:"landlocked"`
	Borders      []string     `json:"borders"`
	Area         float64      `json:"area"`
	Demonyms     Demonyms     `json:"demonyms"`
	Flag         string       `json:"flag"`
	Maps         Maps         `json:"maps"`
	Population   int          `json:"population"`
	Gini         Gini         `json:"gini"`
	Fifa         string       `json:"fifa"`
	Car          Car          `json:"car"`
	Timezones    []string     `json:"timezones"`
	Continents   []string     `json:"continents"`
	Flags        Flags        `json:"flags"`
	CoatOfArms   CoatOfArms   `json:"coatOfArms"`
	StartOfWeek  string       `json:"startOfWeek"`
	CapitalInfo  CapitalInfo  `json:"capitalInfo"`
	PostalCode   PostalCode   `json:"postalCode"`
}

type Name struct {
	Common     string     `json:"common"`
	Official   string     `json:"official"`
	NativeName NativeName `json:"nativeName"`
}

type NativeName struct {
	Por Por `json:"por"`
}

type Por struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Currencies struct {
	Brl Brl `json:"BRL"`
}

type Brl struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Idd struct {
	Root     string   `json:"root"`
	Suffixes []string `json:"suffixes"`
}

type Languages struct {
	Por string `json:"por"`
}

type Translations struct {
	Ara TranslationsItem `json:"ara"`
	Ces TranslationsItem `json:"ces"`
	Cym TranslationsItem `json:"cym"`
	Deu TranslationsItem `json:"deu"`
	Est TranslationsItem `json:"est"`
	Fin TranslationsItem `json:"fin"`
	Fra TranslationsItem `json:"fra"`
	Hrv TranslationsItem `json:"hrv"`
	Hun TranslationsItem `json:"hun"`
	Ita TranslationsItem `json:"ita"`
	Jpn TranslationsItem `json:"jpn"`
	Kor TranslationsItem `json:"kor"`
	Nld TranslationsItem `json:"nld"`
	Per TranslationsItem `json:"per"`
	Pol TranslationsItem `json:"pol"`
	Por TranslationsItem `json:"por"`
	Rus TranslationsItem `json:"rus"`
	Slk TranslationsItem `json:"slk"`
	Spa TranslationsItem `json:"spa"`
	Swe TranslationsItem `json:"swe"`
	Urd TranslationsItem `json:"urd"`
	Zho TranslationsItem `json:"zho"`
}

type TranslationsItem struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Demonyms struct {
	Eng DemonymsItem `json:"eng"`
	Fra DemonymsItem `json:"fra"`
}

type DemonymsItem struct {
	F string `json:"f"`
	M string `json:"m"`
}

type Maps struct {
	GoogleMaps     string `json:"googleMaps"`
	OpenStreetMaps string `json:"openStreetMaps"`
}

type Gini struct {
	Num2019 float64 `json:"2019"`
}

type Car struct {
	Signs []string `json:"signs"`
	Side  string   `json:"side"`
}

type Flags struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
}

type CoatOfArms struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
}

type CapitalInfo struct {
	Latlng []float64 `json:"latlng"`
}

type PostalCode struct {
	Format string `json:"format"`
	Regex  string `json:"regex"`
}
