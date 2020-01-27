package badge

import "testing"

func assert(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Value was incorrect, got: %v, expected: %v.", actual, expected)
	}
}
func TestCleanString(t *testing.T) {
	res := cleanString("A+")
	assert(t, "A", res)
	res = cleanString("A++")
	assert(t, "A", res)
	res = cleanString("A!")
	assert(t, "A", res)
	res = cleanString("A--")
	assert(t, "A", res)
}

func TestWidthCalculation(t *testing.T) {
	w := calulateTextWidth("A")
	assert(t, 14.0, w)
	w = calulateTextWidth("AB")
	assert(t, 28.0, w)
}
