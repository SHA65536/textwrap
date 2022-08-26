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
	result, _ := Wrap(beeRaw, 16)
	assert.Equal(bee16, result, fmt.Sprintf("%s\n\n===should equal===\n\n%s", result, bee16))

	result, _ = Wrap(beeRaw, 8)
	assert.Equal(bee8, result, fmt.Sprintf("%s\n\n===should equal===\n\n%s", result, bee8))
}
