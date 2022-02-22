package newsapi

// SortBy specifies available newsapi sorting options.
type SortBy string

const (
	// SortByRelevancy sorts articles based on the query. Articles
	// most closest to the query comes first.
	SortByRelevancy SortBy = "relevancy"

	// SortByPopularity sorts articles based on news publisher and source
	// popularity.
	SortByPopularity SortBy = "popularity"

	// SortByPublishedAt sorts articles based on their publish time.
	SortByPublishedAt SortBy = "publishedAt"
)

// _validSortBy contains keys for predefined sort by values.
var _validSortBy = map[SortBy]struct{}{
	SortByRelevancy:   {},
	SortByPopularity:  {},
	SortByPublishedAt: {},
}

// Valid checks if search in exists in predefined search in map.
func (sb SortBy) Valid() bool {
	_, ok := _validSortBy[sb]
	return ok
}

// SearchIn specifies available newsapi search attributes in an article.
type SearchIn string

const (
	// SearchInTitle searches given query in a title.
	SearchInTitle SearchIn = "title"

	// SearchInDescription searches given query in a description.
	SearchInDescription SearchIn = "description"

	// SearchInContent searches given query in a content.
	SearchInContent SearchIn = "content"
)

// _validSearchIn contains keys for predefined search in values.
var _validSearchIn = map[SearchIn]struct{}{
	SearchInTitle:       {},
	SearchInDescription: {},
	SearchInContent:     {},
}

// Valid checks if search in exists in predefined search in map.
func (si SearchIn) Valid() bool {
	_, ok := _validSearchIn[si]
	return ok
}

// Language specifies available articles languages.
type Language string

const (
	LanguageArabic    Language = "ar"
	LanguageGerman    Language = "de"
	LanguageEnglish   Language = "en"
	LanguageSpanish   Language = "es"
	LanguageFrench    Language = "fr"
	LanguageHebrew    Language = "hr"
	LanguageItalian   Language = "it"
	LanguageDutch     Language = "nl"
	LanguageNorwegian Language = "no"
	LanguagePortugese Language = "pt"
	LanguageRussian   Language = "ru"
	LanguageSami      Language = "se"
	LanguageUrdu      Language = "ud"
	LanguageChinese   Language = "zh"
)

// _validLanguage contains keys for predefined languages.
var _validLanguage = map[Language]struct{}{
	LanguageArabic:    {},
	LanguageGerman:    {},
	LanguageEnglish:   {},
	LanguageSpanish:   {},
	LanguageFrench:    {},
	LanguageHebrew:    {},
	LanguageItalian:   {},
	LanguageDutch:     {},
	LanguageNorwegian: {},
	LanguagePortugese: {},
	LanguageRussian:   {},
	LanguageSami:      {},
	LanguageUrdu:      {},
	LanguageChinese:   {},
}

// Valid checks if language exists in predefined languages map.
func (l Language) Valid() bool {
	_, ok := _validLanguage[l]
	return ok
}

// Category specifies available articles categories.
type Category string

const (
	CategoryBusiness      Category = "business"
	CategoryEntertainment Category = "entertainment"
	CategoryGeneral       Category = "general"
	CategoryHealth        Category = "health"
	CategoryScience       Category = "science"
	CategorySports        Category = "sports"
	CategoryTechnology    Category = "technology"
)

// _validCategory contains keys for predefined categories.
var _validCategory = map[Category]struct{}{
	CategoryBusiness:      {},
	CategoryEntertainment: {},
	CategoryGeneral:       {},
	CategoryHealth:        {},
	CategoryScience:       {},
	CategorySports:        {},
	CategoryTechnology:    {},
}

// Valid checks if category exists in predefined categories map.
func (c Category) Valid() bool {
	_, ok := _validCategory[c]
	return ok
}

// Country specifies available articles countries.
type Country string

const (
	CountryUnitedArabEmirates Country = "ae"
	CountryArgentina          Country = "ar"
	CountryAustria            Country = "at"
	CountryAustralia          Country = "au"
	CountryBelgium            Country = "be"
	CountryBulgaria           Country = "bg"
	CountryBrazil             Country = "br"
	CountryCanada             Country = "ca"
	CountrySwitzerland        Country = "ch"
	CountryChina              Country = "cn"
	CountryColombia           Country = "co"
	CountryCuba               Country = "cu"
	CountryCzechia            Country = "cz"
	CountryGermany            Country = "de"
	CountryEgypt              Country = "eg"
	CountryFrance             Country = "fr"
	CountryUnitedKingdom      Country = "gb"
	CountryGreece             Country = "gr"
	CountryHonkKong           Country = "hk"
	CountryHungary            Country = "hu"
	CountryIndonesia          Country = "id"
	CountryIreland            Country = "ie"
	CountryIsrael             Country = "il"
	CountryIndia              Country = "in"
	CountryItaly              Country = "it"
	CountryJapan              Country = "jp"
	CountryKorea              Country = "kr"
	CountryLithuania          Country = "lt"
	CountryLatvia             Country = "lv"
	CountryMorocco            Country = "ma"
	CountryMexico             Country = "mx"
	CountryMalaysia           Country = "my"
	CountryNigeria            Country = "ng"
	CountryNetherlands        Country = "nl"
	CountryNorway             Country = "no"
	CountryNewZealand         Country = "nz"
	CountryPhilippines        Country = "ph"
	CountryPoland             Country = "pl"
	CountryPortugal           Country = "pt"
	CountryRomania            Country = "ro"
	CountrySerbia             Country = "rs"
	CountryRussia             Country = "ru"
	CountrySaudiArabia        Country = "sa"
	CountrySweden             Country = "se"
	CountrySingapore          Country = "sg"
	CountrySlovenia           Country = "si"
	CountrySlovakia           Country = "sk"
	CountryThailand           Country = "th"
	CountryTurkey             Country = "tr"
	CountryTaiwan             Country = "tw"
	CountryUkraine            Country = "ua"
	CountryUnitedStates       Country = "us"
	CountryVenezuela          Country = "ve"
	CountrySouthAfrica        Country = "za"
)

// _validCountry contains keys for predefined countries.
var _validCountry = map[Country]struct{}{
	CountryUnitedArabEmirates: {},
	CountryArgentina:          {},
	CountryAustria:            {},
	CountryAustralia:          {},
	CountryBelgium:            {},
	CountryBulgaria:           {},
	CountryBrazil:             {},
	CountryCanada:             {},
	CountrySwitzerland:        {},
	CountryChina:              {},
	CountryColombia:           {},
	CountryCuba:               {},
	CountryCzechia:            {},
	CountryGermany:            {},
	CountryEgypt:              {},
	CountryFrance:             {},
	CountryUnitedKingdom:      {},
	CountryGreece:             {},
	CountryHonkKong:           {},
	CountryHungary:            {},
	CountryIndonesia:          {},
	CountryIreland:            {},
	CountryIsrael:             {},
	CountryIndia:              {},
	CountryItaly:              {},
	CountryJapan:              {},
	CountryKorea:              {},
	CountryLithuania:          {},
	CountryLatvia:             {},
	CountryMorocco:            {},
	CountryMexico:             {},
	CountryMalaysia:           {},
	CountryNigeria:            {},
	CountryNetherlands:        {},
	CountryNorway:             {},
	CountryNewZealand:         {},
	CountryPhilippines:        {},
	CountryPoland:             {},
	CountryPortugal:           {},
	CountryRomania:            {},
	CountrySerbia:             {},
	CountryRussia:             {},
	CountrySaudiArabia:        {},
	CountrySweden:             {},
	CountrySingapore:          {},
	CountrySlovenia:           {},
	CountrySlovakia:           {},
	CountryThailand:           {},
	CountryTurkey:             {},
	CountryTaiwan:             {},
	CountryUkraine:            {},
	CountryUnitedStates:       {},
	CountryVenezuela:          {},
	CountrySouthAfrica:        {},
}

// Valid checks if country exists in predefined countries map.
func (c Country) Valid() bool {
	_, ok := _validCountry[c]
	return ok
}
