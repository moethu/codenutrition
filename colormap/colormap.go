package colormap

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"github.com/moethu/codenutrition/configuration"
)

// Default colors for badges
var black color.RGBA = color.RGBA{0, 0, 0, 255}
var white color.RGBA = color.RGBA{255, 255, 255, 255}
var red color.RGBA = color.RGBA{220, 53, 69, 255}
var yellow color.RGBA = color.RGBA{225, 193, 7, 255}
var green color.RGBA = color.RGBA{40, 167, 69, 255}
var grey color.RGBA = color.RGBA{108, 117, 125, 255}

// colormap holding hex colors for each key
var colorMap map[string]string

// hexColor returns an HTML hex-representation of c. The alpha channel is dropped
// and precision is truncated to 8 bits per channel
func hexColor(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}

// lightenColor calculates a lighter color based on color intensity factor f
func lightenColor(inColor color.RGBA, f float64) color.RGBA {
	return color.RGBA{
		uint8(math.Min(255, float64(inColor.R)+255*f)),
		uint8(math.Min(255, float64(inColor.G)+255*f)),
		uint8(math.Min(255, float64(inColor.B)+255*f)),
		255,
	}
}

// getColorGradient adds a color gradient for a set of values to a colormap
// it filters values by a char like +,- or ! and uses a basecolor to start from.
// reverse will reverse the color gradient
func getColorGradient(colormap map[string]string, values []string, char string, basecolor color.RGBA, reverse bool) map[string]string {
	var data []string

	for _, value := range values {
		if strings.Contains(value, char) {
			data = append(data, value)
		}
	}
	if reverse {
		data = reverseArray(data)
	}
	factor := 0.0 + (0.5 / float64(len(data)))
	for i, value := range data {
		colormap[value] = hexColor(lightenColor(basecolor, float64(i)*factor))
	}

	return colormap
}

// reverseArray creates an array in reverse order
func reverseArray(numbers []string) []string {
	newNumbers := make([]string, len(numbers))
	for i, j := 0, len(numbers)-1; i <= j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}
	return newNumbers
}

// populateColorMap populates a colormap by string values
func populateColorMap(colormap map[string]string, values []string) map[string]string {
	for _, value := range values {
		if strings.Contains(value, "+") {
			getColorGradient(colormap, values, "+", green, false)
		} else if strings.Contains(value, "-") {
			getColorGradient(colormap, values, "-", red, true)
		} else if strings.Contains(value, "!") {
			colormap[value] = hexColor(yellow)
		} else {
			colormap[value] = hexColor(grey)
		}
	}
	return colormap
}

// Load loads the colormap based on the configuration files
// needs to be called first to initialize the colormap
func Load() {
	data := configuration.ReadConfigJson()
	colormap := make(map[string]string)
	for _, category := range data {
		var values []string
		for _, c := range category.Content {
			values = append(values, c.ID)
		}
		populateColorMap(colormap, values)
	}
	colorMap = colormap
}

// Get returns a color from the colormap
func Get(key string) string {
	if val, ok := colorMap[key]; ok {
		return val
	}

	// default to grey
	return hexColor(grey)
}
