package newsapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SortBy_Valid(t *testing.T) {
	for sortBy := range _validSortBy {
		assert.True(t, sortBy.Valid())
	}

	sortBy := SortBy("test")
	assert.False(t, sortBy.Valid())
}

func Test_SearchIn_Valid(t *testing.T) {
	for searchIn := range _validSearchIn {
		assert.True(t, searchIn.Valid())
	}

	searchIn := SearchIn("test")
	assert.False(t, searchIn.Valid())
}

func Test_Language_Valid(t *testing.T) {
	for language := range _validLanguage {
		assert.True(t, language.Valid())
	}

	language := Language("test")
	assert.False(t, language.Valid())
}

func Test_Category_Valid(t *testing.T) {
	for category := range _validCategory {
		assert.True(t, category.Valid())
	}

	category := Category("test")
	assert.False(t, category.Valid())
}

func Test_Country_Valid(t *testing.T) {
	for country := range _validCountry {
		assert.True(t, country.Valid())
	}

	country := Country("test")
	assert.False(t, country.Valid())
}
