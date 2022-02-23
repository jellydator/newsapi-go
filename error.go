package newsapi

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidQueryLength is returned whenever query length exceeds
	// maximum characters size, which is 500.
	ErrInvalidQueryLength = errors.New("query exceeds 500 character limit")

	// ErrInvalidSearchIn is returned whenever search in type has a value
	// that is not predefined list.
	ErrInvalidSearchIn = errors.New("invalid search key")

	// ErrInvalidFromTime is returned whenever from and to times are used
	// and to is earlier than the from time.
	ErrInvalidFromTime = errors.New("from time cannot be after to time")

	// ErrInvalidCategory is returned whenever category type has a value
	// that is not in the predefined list.
	ErrInvalidCategory = errors.New("invalid category")

	// ErrInvalidLanguage is returned whenever language type has a value
	// that is not in the predefined list.
	ErrInvalidLanguage = errors.New("invalid language")

	// ErrInvalidCountry is returned whenever country type has a value
	// that is not in the predefined list.
	ErrInvalidCountry = errors.New("invalid country")

	// ErrInvalidSortBy is returned whenever sort by type has a value
	// that is not in the predefined list.
	ErrInvalidSortBy = errors.New("invalid sort key")

	// ErrInvalidPageSize is returned whenever page size exceeds a
	// maximum of 100 entries.
	ErrInvalidPageSize = errors.New("page size exceeds 100 entries limit")

	// ErrIncompatibleParams is returned whenever in TopHeadlinesParams
	// parameters sources is used along with country or category.
	ErrIncompatibleParams = errors.New("country/category parameter cannot be used along with sources parameter")

	// ErrTooManySources is returned whenever sources list exceeds 20
	// entries.
	ErrTooManySources = errors.New("sources exceeds 20 entries limit")

	// ErrParamsScopeTooBroad is returned when the scope of parameters is
	// too broad.
	ErrParamsScopeTooBroad = errors.New("scope of parameters is too broad")
)

// Error contains newsapi error information.
type Error struct {
	// HTTPCode specifies the response status code.
	HTTPCode int

	// APICode specifies the error code returned from newsapi.
	APICode string

	// Message specifies the error message returned from newsapi.
	Message string
}

// Error implements error interface and returns formatted error message.
func (e *Error) Error() string {
	return fmt.Sprintf(
		`message: %q (http code: "%d"; api code: %q)`,
		e.Message,
		e.HTTPCode,
		e.APICode,
	)
}
