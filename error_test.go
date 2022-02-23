package newsapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error_Error(t *testing.T) {
	err := &Error{
		HTTPCode: 500,
		APICode:  "123",
		Message:  "321",
	}

	assert.EqualError(t, err, `message: "321" (http code: "500"; api code: "123")`)
}
