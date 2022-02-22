package newsapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error_Error(t *testing.T) {
	err := Error{
		StatusCode: 500,
		Code:       "123",
		Message:    "321",
	}

	assert.Equal(t, `statusCode: "500", code: "123", message: "321"`, err.Error())
}
