package textwrap

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/bee_raw.txt
var beeRaw string

//go:embed testdata/bee_16.txt
var bee16 string

//go:embed testdata/bee_8.txt
var bee8 string

func TestRegularWrap(t *testing.T) {
	assert := assert.New(t)
	result, err := Wrap(beeRaw, 16)
	if assert.Nil(err, "bee16 shouldn't error") {
		assert.Equal(bee16, result, fmt.Sprintf("%s\n\n===should equal===\n\n%s", result, bee16))
	}

	result, err = Wrap(beeRaw, 8)
	if assert.Nil(err, "bee8 shouldn't error") {
		assert.Equal(bee8, result, fmt.Sprintf("%s\n\n===should equal===\n\n%s", result, bee8))
	}

	_, err = WrapCustom("abc", 3, "   ", "\n")
	assert.NotNil(err, "delimiter larger than size should error")
}
