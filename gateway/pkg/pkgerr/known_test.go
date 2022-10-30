package pkgerr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	originErrorMsg = "complex error happened"
	simpleErrorMsg = "error ocurred"
	wrapErrorMsg   = "wrapping"
)

func TestKnown(t *testing.T) {
	err := genWrappedError()
	require.Error(t, err)

	kerr := AsKnown(err)
	require.NotNil(t, kerr, "error should be known error")

	assert.Equal(t, fmt.Sprintf("%s: %s: %s", wrapErrorMsg, simpleErrorMsg, originErrorMsg), err.Error())
	assert.Equal(t, fmt.Sprintf("%s: %s", simpleErrorMsg, originErrorMsg), kerr.Error())
	assert.Equal(t, simpleErrorMsg, kerr.SimpleError())
}

func genWrappedError() error {
	return fmt.Errorf("%s: %w", wrapErrorMsg, genKnownError())
}

func genKnownError() error {
	return Known{
		Errno:  ErrnoInternal,
		Origin: genOriginError(),
		Simple: fmt.Errorf(simpleErrorMsg),
	}
}

func genOriginError() error {
	return fmt.Errorf(originErrorMsg)
}
