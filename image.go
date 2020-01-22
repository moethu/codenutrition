package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"strings"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

// Default colors for badges
var black color.RGBA = color.RGBA{0, 0, 0, 255}
var white color.RGBA = color.RGBA{255, 255, 255, 255}
var red color.RGBA = color.RGBA{220, 53, 69, 255}
var yellow color.RGBA = color.RGBA{225, 193, 7, 255}
var green color.RGBA = color.RGBA{40, 167, 69, 255}
var grey color.RGBA = color.RGBA{108, 117, 125, 255}

// getColor returns badge background and foreground colors according to content
func getColor(code string) (color.RGBA, color.RGBA) {
	if strings.Contains(code, "+") {
		return green, white
	}
	if strings.Contains(code, "-") {
		return red, white
	}
	if strings.Contains(code, "!") {
		return yellow, black
	}
	return grey, white
}

// roundRect renders a filled rounded rectangle (badge)
func roundRect(path *draw2dimg.GraphicContext, x1, y1, x2, y2, arcLeft, arcRight float64, color color.RGBA) {
	path.SetFillColor(color)
	path.BeginPath()
	arcLeft = arcLeft / 2
	arcRight = arcRight / 2
	path.MoveTo(x1, y1+arcLeft)
	path.QuadCurveTo(x1, y1, x1+arcLeft, y1)
	path.LineTo(x2-arcRight, y1)
	path.QuadCurveTo(x2, y1, x2, y1+arcRight)
	path.LineTo(x2, y2-arcRight)
	path.QuadCurveTo(x2, y2, x2-arcRight, y2)
	path.LineTo(x1+arcLeft, y2)
	path.QuadCurveTo(x1, y2, x1, y2-arcLeft)
	path.Close()
	path.Fill()
}

// writeString renders a string and returns it's width in pixels
func writeString(gc *draw2dimg.GraphicContext, text string, fontsize, x, y float64, color color.RGBA) float64 {
	draw2d.SetFontFolder(operatingDir + "/fonts/")
	gc.SetFontData(draw2d.FontData{Name: "luxisr", Family: draw2d.FontFamilyMono, Style: draw2d.FontStyleNormal})
	gc.SetFillColor(color)
	gc.SetFontSize(fontsize)
	width := gc.FillStringAt(text, x, y)
	return width
}

// renderSegment renders a badge segment
func renderSegment(gc *draw2dimg.GraphicContext, text string, x, top, bottom, rLeft, rRight, fontsize, offset, baseline float64) float64 {
	bgcol, fgcol := getColor(text)
	textWidth := writeString(gc, text, fontsize, float64(x)+offset, baseline, fgcol)
	roundRect(gc, float64(x), top, float64(x)+textWidth+(2*offset), bottom, rLeft, rRight, bgcol)
	writeString(gc, text, fontsize, float64(x)+offset, baseline, fgcol)
	return textWidth + (2 * offset)
}

// createBadge creates a badge image from an array of code values
func createBadge(code []string) []byte {
	img := image.NewRGBA(image.Rect(0, 0, 300, 20))
	gc := draw2dimg.NewGraphicContext(img)

	spacing := 1.0      // spacing in between badges
	x := spacing        // initial x offset
	offset := 3.0       // text offset in badge
	radiusLeft := 0.0   // corner radius
	radiusRight := 0.0  // corner radius
	baseline := 14.0    // text baseline
	fontsize := 10.0    // fontsize
	badgeTop := 1.0     // badge top y
	badgeBottom := 18.0 // badge bottom y
	radius := 6.0

	width := renderSegment(gc, " codenutrition ", x, badgeTop, badgeBottom, radius, radiusRight, fontsize, offset, baseline)
	x = x + width + spacing

	for i, segment := range code {
		if len(segment) > 0 {
			if i == len(code)-1 {
				radiusRight = radius
			}
			width := renderSegment(gc, segment, x, badgeTop, badgeBottom, radiusLeft, radiusRight, fontsize, offset, baseline)
			x = x + width + spacing
		}
	}

	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
