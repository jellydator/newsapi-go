package newsapi

import (
	"net/url"
	"strconv"
	"time"
)

// All available sort keys.
const (
	SortByRelevancy   SortBy = "relevancy"
	SortByPopularity  SortBy = "popularity"
	SortByPublishedAt SortBy = "publishedAt"
)

// All available search keys.
const (
	SearchInTitle       SearchIn = "title"
	SearchInDescription SearchIn = "description"
	SearchInContent     SearchIn = "content"
)

// All available languages.
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

// All available categories.
const (
	CategoryBusiness      Category = "business"
	CategoryEntertainment Category = "entertainment"
	CategoryGeneral       Category = "general"
	CategoryHealth        Category = "health"
	CategoryScience       Category = "science"
	CategorySports        Category = "sports"
	CategoryTechnology    Category = "technology"
)

// All available countries.
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

// SortBy determines newsapi result order.
type SortBy string

// isValid checks if sort key is valid.
func (sb SortBy) isValid() bool {
	switch sb {
	case SortByRelevancy,
		SortByPopularity,
		SortByPublishedAt:

		return true
	}

	return false
}

// SearchIn determines in which article part newsapi should search.
type SearchIn string

// isValid checks if search key is valid.
func (si SearchIn) isValid() bool {
	switch si {
	case SearchInTitle,
		SearchInDescription,
		SearchInContent:

		return true
	}

	return false
}

// Language determines the language of the source.
type Language string

// isValid checks if language is valid.
func (l Language) isValid() bool {
	switch l {
	case LanguageArabic,
		LanguageGerman,
		LanguageEnglish,
		LanguageSpanish,
		LanguageFrench,
		LanguageHebrew,
		LanguageItalian,
		LanguageDutch,
		LanguageNorwegian,
		LanguagePortugese,
		LanguageRussian,
		LanguageSami,
		LanguageUrdu,
		LanguageChinese:

		return true
	}

	return false
}

// Category determines the category of the source.
type Category string

// isValid checks if category is valid.
func (c Category) isValid() bool {
	switch c {
	case CategoryBusiness,
		CategoryEntertainment,
		CategoryGeneral,
		CategoryHealth,
		CategoryScience,
		CategorySports,
		CategoryTechnology:

		return true
	}

	return false
}

// Country determines the origin of the source.
type Country string

// isValid checks if country is valid.
func (c Country) isValid() bool {
	switch c {
	case CountryUnitedArabEmirates,
		CountryArgentina,
		CountryAustria,
		CountryAustralia,
		CountryBelgium,
		CountryBulgaria,
		CountryBrazil,
		CountryCanada,
		CountrySwitzerland,
		CountryChina,
		CountryColombia,
		CountryCuba,
		CountryCzechia,
		CountryGermany,
		CountryEgypt,
		CountryFrance,
		CountryUnitedKingdom,
		CountryGreece,
		CountryHonkKong,
		CountryHungary,
		CountryIndonesia,
		CountryIreland,
		CountryIsrael,
		CountryIndia,
		CountryItaly,
		CountryJapan,
		CountryKorea,
		CountryLithuania,
		CountryLatvia,
		CountryMorocco,
		CountryMexico,
		CountryMalaysia,
		CountryNigeria,
		CountryNetherlands,
		CountryNorway,
		CountryNewZealand,
		CountryPhilippines,
		CountryPoland,
		CountryPortugal,
		CountryRomania,
		CountrySerbia,
		CountryRussia,
		CountrySaudiArabia,
		CountrySweden,
		CountrySingapore,
		CountrySlovenia,
		CountrySlovakia,
		CountryThailand,
		CountryTurkey,
		CountryTaiwan,
		CountryUkraine,
		CountryUnitedStates,
		CountryVenezuela,
		CountrySouthAfrica:

		return true
	}

	return false
}

// SourceID contains identifying information of a news publisher.
type SourceID struct {
	// ID is the identifier of the news source.
	ID string `json:"id"`

	// Name of the news source.
	Name string `json:"name"`
}

// Source contains information about news publisher.
type Source struct {
	SourceID

	// Description specifies short introduction about the news source.
	Description string `json:"description"`

	// URL specifies the url of the news source.
	URL string `json:"url"`

	// Category specifies the type of news that this source produces.
	Category Category `json:"category"`

	// Language specifies the language in which the news source
	// publishes.
	Language Language `json:"language"`

	// Country specifies the country in which the news source publishes.
	Country Country `json:"country"`
}

// Article contains information of a specific news publisher article.
type Article struct {
	// SourceID specifies source identifying information.
	Source SourceID `json:"source"`

	// Author specifies the author of the article.
	Author string `json:"author"`

	// Title specifies the title/headline of the article.
	Title string `json:"title"`

	// Description specifies short summary of the article.
	Description string `json:"description"`

	// URL specifies the url of the article.
	URL string `json:"url"`

	// URLToImage specifies url to the article image.
	URLToImage string `json:"urlToImage"`

	// PublishedAt specifies the date and time at which the article
	// was published.
	PublishedAt time.Time `json:"publishedAt"`

	// Content specifies an unformatted text of the article. It is
	// truncated to 200 chars.
	Content string `json:"content"`
}

// SourceParams contains source endpoint filters.
type SourceParams struct {
	// Categories is used to filter sources by categories. If left empty
	// all categories are used.
	Categories []Category

	// Languages is used to filter sources by languages. If left empty
	// all languages are used.
	Languages []Language

	// Countries is used to filter sources by countries. If left empty
	// all countries are used.
	Countries []Country
}

// validate validates parameters and their compatibility.
func (sr *SourceParams) validate() error {
	for _, category := range sr.Categories {
		if !category.isValid() {
			return ErrInvalidCategory
		}
	}

	for _, language := range sr.Languages {
		if !language.isValid() {
			return ErrInvalidLanguage
		}
	}

	for _, country := range sr.Countries {
		if !country.isValid() {
			return ErrInvalidCountry
		}
	}

	return nil
}

// rawQuery constructs a raw query from parameters.
func (sr *SourceParams) rawQuery() string {
	q := make(url.Values)

	for _, category := range sr.Categories {
		q.Add("category", string(category))
	}

	for _, language := range sr.Languages {
		q.Add("language", string(language))
	}

	for _, country := range sr.Countries {
		q.Add("country", string(country))
	}

	return q.Encode()
}

// TopHeadlinesParams contains top headlines endpoint filters.
type TopHeadlinesParams struct {
	// Query is used to filter articles' text. Unlike Everything endpoint
	// parameters it doesn't allow for an advanced search, so only basic
	// keywords or phrases should be used. Query has a maximum length of
	// 500 characters.
	Query string

	// Category is used to filter articles by category. If left
	// empty all categories are used.
	Category Category

	// Language is used to filter articles by language. If
	// left empty all languages are used.
	Language Language

	// Country is used to filter articles by country. If left empty
	// all countries are used.
	Country Country

	// Sources is used to filter news publishers. The list of available
	// sources can be retrieved using Sources method on a client or by
	// looking at the sources index here:
	// https://newsapi.org/sources
	Sources []string

	// PageSize specifies the total number of results to return per page.
	// 20 is default, 100 is the maximum.
	PageSize uint

	// Page pages through results if the total results found is greater
	// than the page size.
	Page uint
}

// validate validates parameters and their compatibility.
func (thp *TopHeadlinesParams) validate() error {
	if len(thp.Query) > 500 {
		return ErrInvalidQueryLength
	}

	if thp.Category != "" && !thp.Category.isValid() {
		return ErrInvalidCategory
	}

	if thp.Language != "" && !thp.Language.isValid() {
		return ErrInvalidLanguage
	}

	if thp.Country != "" && !thp.Country.isValid() {
		return ErrInvalidCountry
	}

	if len(thp.Sources) != 0 &&
		(thp.Country != "" || thp.Category != "") {

		return ErrIncompatibleParams
	}

	if thp.PageSize > 100 {
		return ErrInvalidPageSize
	}

	if thp.Query == "" &&
		thp.Category == "" &&
		thp.Language == "" &&
		thp.Country == "" &&
		len(thp.Sources) == 0 {

		return ErrParamsScopeTooBroad
	}

	return nil
}

// rawQuery constructsa a raw query from parameters.
func (thp *TopHeadlinesParams) rawQuery() string {
	q := make(url.Values)

	if thp.Query != "" {
		q.Add("q", thp.Query)
	}

	if thp.Category != "" {
		q.Add("category", string(thp.Category))
	}

	if thp.Country != "" {
		q.Add("country", string(thp.Country))
	}

	if thp.Language != "" {
		q.Add("language", string(thp.Language))
	}

	for _, source := range thp.Sources {
		q.Add("sources", source)
	}

	if thp.PageSize != 0 {
		q.Add("pageSize", strconv.Itoa(int(thp.PageSize)))
	}

	if thp.Page != 0 {
		q.Add("page", strconv.Itoa(int(thp.Page)))
	}

	return q.Encode()
}

// EverythingParams contains everything endpoint filters.
// Original documentation can be found here:
// https://newsapi.org/docs/endpoints/everything
type EverythingParams struct {
	// Query is used to filter articles' text. An advanced search is
	// allowed:
	// * Phrases that must be matched exactly as is should be surrounded
	//   with quotes ("). e.g "my short phrase"
	// * Words or phrases that must appear in an article should be
	//   prefixed with a plus (+) sign. e.g +bitcoin.
	// * Words or phrases that must not appear in an article should be
	//   prefixed with a minus (-) sign. e.g -bitcoin.
	// * "AND", "OR" and "NOT" keywords can be used to group several
	//   queries. Parenthesis can be used to create subgroups. e.g:
	//   crypto AND (ethereum OR litecoin) NOT bitcoin.
	// Query has a maximum length of 500 characters.
	Query string

	// QueryInTitle is used to filter article title. Unlike query
	// parameter it doesn't allow for an advanced search, so only basic
	// keywords or phrases should be used.
	QueryInTitle string

	// SearchIn specifies which part of an article should be searched.
	SearchIn SearchIn

	// Sources is used to filter news publishers. The list of available
	// sources can be retrieved using Sources method on a client or by
	// looking at the sources index here:
	// https://newsapi.org/sources
	// 20 is the maximum sources allowed.
	Sources []string

	// Domains is used to restrict the search to the specified domains.
	Domains []string

	// ExcludeDomains is used to remove results that contain specified
	// domains.
	ExcludeDomains []string

	// From is a date and time for the oldest allowed article.
	From time.Time

	// To is a date and time for the newest allowed article.
	To time.Time

	// Language is used to filter articles by language. If
	// left empty all languages are used.
	Language Language

	// SortBy specifies a key by which articles should be sorted.
	SortBy SortBy

	// PageSize specifies the total number of results to return per page.
	// 20 is default, 100 is the maximum.
	PageSize uint

	// Page pages through results if the total results found is greater
	// than the page size.
	Page uint
}

// validate validates parameters and their compatibility.
func (ep *EverythingParams) validate() error {
	if len(ep.Query) > 500 {
		return ErrInvalidQueryLength
	}

	if ep.SearchIn != "" && !ep.SearchIn.isValid() {
		return ErrInvalidSearchIn
	}

	if len(ep.Sources) > 20 {
		return ErrTooManySources
	}

	if !ep.From.IsZero() && !ep.To.IsZero() && ep.From.After(ep.To) {
		return ErrInvalidFromTime
	}

	if ep.Language != "" && !ep.Language.isValid() {
		return ErrInvalidLanguage
	}

	if ep.SortBy != "" && !ep.SortBy.isValid() {
		return ErrInvalidSortBy
	}

	if ep.PageSize > 100 {
		return ErrInvalidPageSize
	}

	if ep.Query == "" &&
		ep.QueryInTitle == "" &&
		len(ep.Sources) == 0 &&
		len(ep.Domains) == 0 {

		return ErrParamsScopeTooBroad
	}

	return nil
}

// rawQuery constructsa a raw query from parameters.
func (ep *EverythingParams) rawQuery() string {
	q := make(url.Values)

	if ep.Query != "" {
		q.Add("q", ep.Query)
	}

	if ep.QueryInTitle != "" {
		q.Add("qInTitle", ep.QueryInTitle)
	}

	if ep.SearchIn != "" {
		q.Add("searchIn", string(ep.SearchIn))
	}

	for _, source := range ep.Sources {
		q.Add("sources", source)
	}

	for _, domain := range ep.Domains {
		q.Add("domains", domain)
	}

	for _, domain := range ep.ExcludeDomains {
		q.Add("excludeDomains", domain)
	}

	if !ep.From.IsZero() {
		q.Add("from", ep.From.Format("2006-01-02T15:04:05"))
	}

	if !ep.To.IsZero() {
		q.Add("to", ep.To.Format("2006-01-02T15:04:05"))
	}

	if ep.Language != "" {
		q.Add("language", string(ep.Language))
	}

	if ep.SortBy != "" {
		q.Add("sortBy", string(ep.SortBy))
	}

	if ep.PageSize != 0 {
		q.Add("pageSize", strconv.Itoa(int(ep.PageSize)))
	}

	if ep.Page != 0 {
		q.Add("page", strconv.Itoa(int(ep.Page)))
	}

	return q.Encode()
}
