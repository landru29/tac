package base45_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/landru29/tac/internal/base45"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		stream, err := base45.Decode(strings.NewReader("%69 VD92EX"))
		assert.NoError(t, err)
		if !assert.NotNil(t, stream) {
			return
		}

		d, err := ioutil.ReadAll(stream)
		assert.NoError(t, err)

		assert.Equal(t, []byte("Hello!!"), d)
	})
}
