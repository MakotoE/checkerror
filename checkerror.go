package checkerror

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Check(t *testing.T, expectError bool, err error, msgAndArgs ...interface{}) {
	if expectError {
		assert.NotNil(t, err, msgAndArgs...)
	} else {
		assert.Nil(t, err, msgAndArgs...)
	}
}
