package checkerror

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Check(t *testing.T, expectError bool, err error) {
	if expectError {
		assert.NotNil(t, err)
	} else {
		assert.Nil(t, err)
	}
}
