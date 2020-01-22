package main

import (
	"fmt"
	"strings"
)

// standard colors for svg
var svgGreen = "#28A745"
var svgGrey = "#6C757D"
var svgYellow = "#FFC107"
var svgRed = "#DC3545"

// getSVGColor returns badge background and foreground colors according to content
func getSVGColor(code string) string {
	if strings.Contains(code, "+") {
		return svgGreen
	}
	if strings.Contains(code, "-") {
		return svgRed
	}
	if strings.Contains(code, "!") {
		return svgYellow
	}
	return svgGrey
}

// calulateTextWidth calculates text width for a string
// TODO: calculation is not considering rendered character widths - needs to be improved
func calulateTextWidth(text string) float64 {
	return float64(len(text)) * 12.0
}

// generateSVGSegment generates a SVG segment by position and text
func generateSVGSegment(text string, x float64) (string, string, float64) {
	w := calulateTextWidth(text)
	bgcolor := getSVGColor(text)
	background := fmt.Sprintf(`<path fill="%s" d="M%f 0h%fv20H%fz"/>`, bgcolor, x, x+w, x)
	offset := w / 2
	foreground := fmt.Sprintf(`<text x="%f" y="15" fill="#010101" fill-opacity=".3">%s</text><text x="%f" y="14">%s</text>`, x+offset, text, x+offset, text)
	return background, foreground, w
}

// generateSVGBadge generates a full SVG badge with all segments
// generates SVG without much overhead (no XML or SVG library) since most of the svg content stays the same.
// Only segments are changing.
func generateSVGBadge(width float64, background, foreground string) string {
	return fmt.Sprintf(`
	<svg xmlns="http://www.w3.org/2000/svg" width="%f" height="20">
		<linearGradient id="b" x2="0" y2="100%%">
			<stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
			<stop offset="1" stop-opacity=".1"/>
		</linearGradient>
		<mask id="a">
			<rect width="%f" height="20" rx="3" fill="#fff"/>
		</mask>
		<g mask="url(#a)">
			<path fill="#555" d="M0 0h90v20H0z"/>
			%s
			<path fill="url(#b)" d="M0 0h%fv20H0z"/>
		</g>
		<g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
			<text x="45" y="15" fill="#010101" fill-opacity=".3">codenutrition</text>
			<text x="45" y="14">codenutrition</text>
			%s
		</g>
	</svg>`, width, width, background, width, foreground)
}

// getSVG returns an svg badge
func getSVG(code []string) []byte {
	foreground := ""
	background := ""
	x := 90.0
	for _, segment := range code {
		if len(segment) > 0 {
			bg, fg, w := generateSVGSegment(segment, x)
			x = x + w
			foreground = foreground + fg
			background = background + bg
		}
	}
	return []byte(generateSVGBadge(x, background, foreground))
}
