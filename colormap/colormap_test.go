package colormap

import "testing"

import "image/color"

func assert(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Value was incorrect, got: %v, expected: %v.", actual, expected)
	}
}
func TestHexColor(t *testing.T) {
	s := hexColor(color.Black)
	assert(t, "#000000", s)
	s = hexColor(color.RGBA{255, 0, 0, 0})
	assert(t, "#ff0000", s)
}

func TestLightenColor(t *testing.T) {
	c := lightenColor(color.RGBA{255, 0, 0, 0}, 0.5)
	assert(t, uint8(255), c.R)
	assert(t, uint8(127), c.G)
	assert(t, uint8(127), c.B)

	c = lightenColor(color.RGBA{255, 0, 0, 0}, 1.0)
	assert(t, uint8(255), c.R)
	assert(t, uint8(255), c.G)
	assert(t, uint8(255), c.B)
}

func TestReverseArray(t *testing.T) {
	a := []string{"1", "2", "3", "4"}
	b := reverseArray(a)
	assert(t, b[0], "4")
	assert(t, b[1], "3")
	assert(t, b[2], "2")
	assert(t, b[3], "1")
}

func TestColorGradient(t *testing.T) {
	colormap := make(map[string]string)
	values := []string{"A++", "A+", "A", "A-", "A--", "A---"}

	getColorGradient(colormap, values, "+", color.RGBA{255, 0, 0, 0}, false)
	assert(t, "#ff0000", colormap["A++"])
	assert(t, "#ff3f3f", colormap["A+"])

	getColorGradient(colormap, values, "-", color.RGBA{255, 0, 0, 0}, true)
	assert(t, "#ff5555", colormap["A-"])
	assert(t, "#ff2a2a", colormap["A--"])
	assert(t, "#ff0000", colormap["A---"])
}

func TestLoad(t *testing.T) {
	Load("../static/spectrum.json")
	if len(colorMap) == 0 {
		t.Error("Could not load spectrum.json")
	}
	assert(t, "#6c757d", Get("A+"))
	assert(t, "#28a745", Get("O++"))
}
