package text2img

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type converter struct {
	Font            *truetype.Font
	ForegroundColor *image.Uniform
	BackgroundColor *image.Uniform
	DPI             float64
	FontSize        float64
}

func NewConverter(fontPath string) (*converter, error) {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return nil, fmt.Errorf("read font file: %w", err)
	}

	parsedFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("parse font: %w", err)
	}

	return &converter{
		Font: parsedFont,
		// TODO: make these attributes configurable
		FontSize:        40.0,
		BackgroundColor: image.Black,
		ForegroundColor: image.White,
		DPI:             60.0,
	}, nil
}

func (c *converter) Convert(text string) (string, error) {
	var imgHeight, imgWidth int
	// Draw the background
	background := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(background, background.Bounds(), c.BackgroundColor, image.ZP, draw.Src)

	ctx := freetype.NewContext()
	ctx.SetDPI(c.DPI) // screen resolution in Dots Per Inch
	ctx.SetFont(c.Font)
	ctx.SetFontSize(c.FontSize) // font size in points
	ctx.SetClip(background.Bounds())
	ctx.SetDst(background)
	ctx.SetSrc(c.ForegroundColor)

	// pt := freetype.Pt(10, 10+int(ctx.PointToFixed(c.FontSize)>>6))

	return "", errors.New("foo")
}
