package newsapi

import (
	"net/url"
	"strconv"
	"time"
)

// SourceID contains identifying information for news publisher.
type SourceID struct {
	// ID is the identifier of the news source.
	ID string `json:"id"`

	// Name of the news source.
	Name string `json:"name"`
}

// Source contains information about news publisher.
type Source struct {
	SourceID

	// Description of the news source.
	Description string `json:"description"`

	// URL of the news source home page.
	URL string `json:"url"`

	// Category specifies the type of news that this source produces.
	Category Category `json:"category"`

	// Language specifies the langue the news source writes in.
	Language Language `json:"language"`

	// Country specifies the country the news source is based in.
	Country Country `json:"country"`
}

// Article contains information for a specific news publisher article.
type Article struct {
	// SourceID specifies source identifying information.
	Source SourceID `json:"source"`

	// Author of the article.
	Author string `json:"author"`

	// Title or headline of the article.
	Title string `json:"title"`

	// Description or snippet of the article.
	Description string `json:"description"`

	// URL to the article.
	URL string `json:"url"`

	// URLToImage specifies url to the article image.
	URLToImage string `json:"urlToImage"`

	// PublishedAt specifies the date and time this arcticle was published.
	PublishedAt time.Time `json:"publishedAt"`

	// Content is an unformatted text of the article. This is truncated
	// to 200 chars.
	Content string `json:"content"`
}

// SourceParams contains source endpoint filters.
type SourceParams struct {
	// Categories filters sources by categories. If left empty all
	// categories are used.
	Categories []Category

	// Languages filters sources by languages. If left empty all
	// languages are used.
	Languages []Language

	// Countries filters sources by countries. If left empty all
	// countries are used.
	Countries []Country
}

// Validate checks for params validity and whether they can be used
// together.
func (sr *SourceParams) Validate() error {
	for _, category := range sr.Categories {
		if !category.Valid() {
			return ErrInvalidCategory
		}
	}

	for _, language := range sr.Languages {
		if !language.Valid() {
			return ErrInvalidLanguage
		}
	}

	for _, country := range sr.Countries {
		if !country.Valid() {
			return ErrInvalidCountry
		}
	}

	return nil
}

// rawQuery constructs raw query out of params.
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
	// Query filters articles text based on specified a query. Unlike
	// Everything endpoint it doesn't allow for an advanced search, so
	// basic keywords or phrases should be used.
	// Query has a maximum length of 500 characters.
	Query string

	// Category filters top headlines by category. If left empty all
	// categories are used.
	Category Category

	// Language filters top headlines by language. If left empty all
	// languages are used.
	Language Language

	// Country filters top headlines by country. If left empty all
	// countries are used.
	Country Country

	// Sources filters news publishers. List of available sources can be
	// fetched using Sources method on a client or by looking at the
	// sources index here:
	// https://newsapi.org/sources
	Sources []string

	// PageSize specifies total number of results to return per page.
	// 20 is default, 100 is the maximum.
	PageSize uint

	// Page allows to page through results if the total results found
	// is greater than the page size.
	Page uint
}

// Validate checks for params validity and whether they can be used
// together.
func (thp *TopHeadlinesParams) Validate() error {
	if len(thp.Query) > 500 {
		return ErrInvalidQueryLength
	}

	if thp.Category != "" && !thp.Category.Valid() {
		return ErrInvalidCategory
	}

	if thp.Language != "" && !thp.Language.Valid() {
		return ErrInvalidLanguage
	}

	if thp.Country != "" && !thp.Country.Valid() {
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

// rawQuery constructs raw query out of params.
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
type EverythingParams struct {
	// Query filters articles text based on a specified query. It allows
	// for an advanced search:
	// * Surround phrases with quotes (") for exact match.
	// * Prepend words or phrases that must appear with a plus (+)
	//   symbol. Eg: +bitcoin.
	// * Prepend words or phrases that must not appear with a minus (-)
	//   symbol. Eg: -bitcoin.
	// * Alternatively you can use AND / OR / NOT keywords, and
	//   optionally group these with paranthesis. Eg: crypto AND
	//   (ethereum OR litecoin) NOT bitcoin.
	// Original documentation can be found here:
	// https://newsapi.org/docs/endpoints/everything
	// Query has a maximum length of 500 characters.
	Query string

	// QueryInTitle filters article title based on a specified query.
	QueryInTitle string

	// SearchIn specifies which fields should be searched for query.
	SearchIn SearchIn

	// Sources filters news publishers. List of available sources can be
	// fetched using Sources method on a client or by looking at the
	// sources index here:
	// https://newsapi.org/sources
	// 20 is the maximum sources allowed.
	Sources []string

	// Domains allows to restrict the search to the specified domains.
	Domains []string

	// ExcludeDomains allows to remove results that contain specified
	// domains.
	ExcludeDomains []string

	// From is a date and time for the oldest allowed article.
	From time.Time

	// To is a date and time for the newest allowed article.
	To time.Time

	// Languages filters sources by languages. If left empty all
	// languages are used.
	Language Language

	// SortBy specifies an order in which articles should be sorted.
	SortBy SortBy

	// PageSize specifies total number of results to return per page.
	// 20 is default, 100 is the maximum.
	PageSize uint

	// Page allows to page through results if the total results found
	// is greater than the page size.
	Page uint
}

// Validate checks for params validity and whether they can be used
// together.
func (ep *EverythingParams) Validate() error {
	if len(ep.Query) > 500 {
		return ErrInvalidQueryLength
	}

	if ep.SearchIn != "" && !ep.SearchIn.Valid() {
		return ErrInvalidSearchIn
	}

	if len(ep.Sources) > 20 {
		return ErrTooManySources
	}

	if !ep.From.IsZero() && !ep.To.IsZero() && ep.From.After(ep.To) {
		return ErrInvalidFromTime
	}

	if ep.Language != "" && !ep.Language.Valid() {
		return ErrInvalidLanguage
	}

	if ep.SortBy != "" && !ep.SortBy.Valid() {
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

// rawQuery constructs raw query out of params.
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
