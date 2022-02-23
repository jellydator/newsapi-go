package newsapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_SortBy_isValid(t *testing.T) {
	for _, sortBy := range []SortBy{
		SortByRelevancy,
		SortByPopularity,
		SortByPublishedAt,
	} {
		assert.True(t, sortBy.isValid())
	}

	sortBy := SortBy("test")
	assert.False(t, sortBy.isValid())
}

func Test_SearchIn_isValid(t *testing.T) {
	for _, searchIn := range []SearchIn{
		SearchInTitle,
		SearchInDescription,
		SearchInContent,
	} {
		assert.True(t, searchIn.isValid())
	}

	searchIn := SearchIn("test")
	assert.False(t, searchIn.isValid())
}

func Test_Language_isValid(t *testing.T) {
	for _, language := range []Language{
		LanguageArabic,
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
		LanguageChinese,
	} {
		assert.True(t, language.isValid())
	}

	language := Language("test")
	assert.False(t, language.isValid())
}

func Test_Category_isValid(t *testing.T) {
	for _, category := range []Category{
		CategoryBusiness,
		CategoryEntertainment,
		CategoryGeneral,
		CategoryHealth,
		CategoryScience,
		CategorySports,
		CategoryTechnology,
	} {
		assert.True(t, category.isValid())
	}

	category := Category("test")
	assert.False(t, category.isValid())
}

func Test_Country_isValid(t *testing.T) {
	for _, country := range []Country{
		CountryUnitedArabEmirates,
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
		CountrySouthAfrica,
	} {
		assert.True(t, country.isValid())
	}

	country := Country("test")
	assert.False(t, country.isValid())
}

func Test_SourceParams_validate(t *testing.T) {
	tests := map[string]struct {
		Params SourceParams
		Err    error
	}{
		"Invalid category": {
			Params: SourceParams{
				Categories: []Category{
					"123",
				},
			},
			Err: ErrInvalidCategory,
		},
		"Invalid language": {
			Params: SourceParams{
				Languages: []Language{
					"123",
				},
			},
			Err: ErrInvalidLanguage,
		},
		"Invalid country": {
			Params: SourceParams{
				Countries: []Country{
					"123",
				},
			},
			Err: ErrInvalidCountry,
		},
		"Valid parameters": {
			Params: SourceParams{
				Categories: []Category{
					CategoryBusiness,
					CategoryEntertainment,
				},
				Languages: []Language{
					LanguageItalian,
					LanguageSpanish,
				},
				Countries: []Country{
					CountryArgentina,
					CountryAustria,
				},
			},
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Err, test.Params.validate())
		})
	}
}

func Test_SourceParams_rawQuery(t *testing.T) {
	assert.Equal(t, "", (&SourceParams{}).rawQuery())
	assert.Equal(
		t,
		"category=business&category=entertainment&country=ar&country=at&language=it&language=es",
		(&SourceParams{
			Categories: []Category{
				CategoryBusiness,
				CategoryEntertainment,
			},
			Languages: []Language{
				LanguageItalian,
				LanguageSpanish,
			},
			Countries: []Country{
				CountryArgentina,
				CountryAustria,
			},
		}).rawQuery(),
	)
}

func Test_TopHeadlinesParams_validate(t *testing.T) {
	tests := map[string]struct {
		Params TopHeadlinesParams
		Err    error
	}{
		"Invalid query": {
			Params: TopHeadlinesParams{
				Query: string(make([]byte, 501)),
			},
			Err: ErrInvalidQueryLength,
		},
		"Invalid category": {
			Params: TopHeadlinesParams{
				Category: Category("test"),
			},
			Err: ErrInvalidCategory,
		},
		"Invalid language": {
			Params: TopHeadlinesParams{
				Language: Language("test"),
			},
			Err: ErrInvalidLanguage,
		},
		"Invalid country": {
			Params: TopHeadlinesParams{
				Country: Country("test"),
			},
			Err: ErrInvalidCountry,
		},
		"Incompatible params": {
			Params: TopHeadlinesParams{
				Country: CountryArgentina,
				Sources: []string{
					"test",
				},
			},
			Err: ErrIncompatibleParams,
		},
		"Invalid page size": {
			Params: TopHeadlinesParams{
				PageSize: 101,
			},
			Err: ErrInvalidPageSize,
		},
		"Params scope too broad": {
			Params: TopHeadlinesParams{},
			Err:    ErrParamsScopeTooBroad,
		},
		"Valid parameters": {
			Params: TopHeadlinesParams{
				Sources: []string{
					"test",
				},
			},
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Err, test.Params.validate())
		})
	}
}

func Test_TopHeadlinesParams_rawQuery(t *testing.T) {
	assert.Equal(t, "", (&TopHeadlinesParams{}).rawQuery())
	assert.Equal(
		t,
		"category=business&country=ar&language=ar&page=5&pageSize=10&q=123",
		(&TopHeadlinesParams{
			Query:    "123",
			Category: CategoryBusiness,
			Country:  CountryArgentina,
			Language: LanguageArabic,
			PageSize: 10,
			Page:     5,
		}).rawQuery(),
	)
	assert.Equal(
		t,
		"language=ar&page=5&pageSize=10&q=123&sources=test&sources=test2",
		(&TopHeadlinesParams{
			Query:    "123",
			Sources:  []string{"test", "test2"},
			Language: LanguageArabic,
			PageSize: 10,
			Page:     5,
		}).rawQuery(),
	)
}

func Test_EverythingParams_validate(t *testing.T) {
	tests := map[string]struct {
		Params EverythingParams
		Err    error
	}{
		"Invalid query": {
			Params: EverythingParams{
				Query: string(make([]byte, 501)),
			},
			Err: ErrInvalidQueryLength,
		},
		"Invalid search in": {
			Params: EverythingParams{
				SearchIn: SearchIn("test"),
			},
			Err: ErrInvalidSearchIn,
		},
		"Invalid sources": {
			Params: EverythingParams{
				Sources: make([]string, 21),
			},
			Err: ErrTooManySources,
		},
		"Invalid from to times": {
			Params: EverythingParams{
				From: time.Now(),
				To:   time.Now().Add(-time.Minute),
			},
			Err: ErrInvalidFromTime,
		},
		"Invalid language": {
			Params: EverythingParams{
				Language: Language("test"),
			},
			Err: ErrInvalidLanguage,
		},
		"Invalid sort by": {
			Params: EverythingParams{
				SortBy: SortBy("test"),
			},
			Err: ErrInvalidSortBy,
		},
		"Invalid page size": {
			Params: EverythingParams{
				PageSize: 101,
			},
			Err: ErrInvalidPageSize,
		},
		"Params scope too broad": {
			Params: EverythingParams{},
			Err:    ErrParamsScopeTooBroad,
		},
		"Valid parameters": {
			Params: EverythingParams{
				Sources: []string{
					"test",
				},
			},
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Err, test.Params.validate())
		})
	}
}

func Test_EverythingParams_rawQuery(t *testing.T) {
	tstamp := time.Date(2022, 02, 22, 22, 22, 22, 22, time.UTC)

	assert.Equal(t, "", (&EverythingParams{}).rawQuery())
	assert.Equal(
		t,
		"domains=test.com&domains=test2.com&excludeDomains=tes4.com&excludeDomains=test3.com&from=2022-02-22T22%3A22%3A22&language=hr&page=3&pageSize=50&q=123&qInTitle=312&searchIn=content&sortBy=publishedAt&sources=test&sources=test2&to=2022-02-22T22%3A23%3A22",
		(&EverythingParams{
			Query:          "123",
			QueryInTitle:   "312",
			SearchIn:       SearchInContent,
			Sources:        []string{"test", "test2"},
			Domains:        []string{"test.com", "test2.com"},
			ExcludeDomains: []string{"tes4.com", "test3.com"},
			From:           tstamp,
			To:             tstamp.Add(time.Minute),
			Language:       LanguageHebrew,
			SortBy:         SortByPublishedAt,
			PageSize:       50,
			Page:           3,
		}).rawQuery(),
	)
}
