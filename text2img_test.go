//go:test

package text2img_test

import (
	"testing"

	"github.com/bbx-engineer/text2img"
	"github.com/stretchr/testify/assert"
)

func TestBase64Output(t *testing.T) {
	text := "test"

	converter, err := text2img.NewConverter("fixtures/arial-unicode-ms.ttf")
	assert.Nil(t, err)

	base64text, err := converter.Convert(text)
	assert.Nil(t, err)
	assert.Equal(t, "data:image/png;base64,blablabla", base64text)
}
