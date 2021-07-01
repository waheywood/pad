package pad_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waheywood/pad"
)

func TestLeftPad(t *testing.T) {
	tests := map[string]struct {
		stringToPad        string
		pad                string
		numberOfTimesToPad int
		expectedString     string
	}{
		"pad once": {
			stringToPad:        "pad",
			pad:                " ",
			numberOfTimesToPad: 1,
			expectedString:     " pad",
		},
		"pad twice": {
			stringToPad:        "pad",
			pad:                " ",
			numberOfTimesToPad: 2,
			expectedString:     "  pad",
		},
		"pad 0 times": {
			stringToPad:        "pad",
			pad:                " ",
			numberOfTimesToPad: 0,
			expectedString:     "pad",
		},
		"pad -1 times": {
			stringToPad:        "pad",
			pad:                " ",
			numberOfTimesToPad: -1,
			expectedString:     "pad",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resultString := pad.LeftPad(test.stringToPad, test.pad, test.numberOfTimesToPad)
			assert.Equal(t, test.expectedString, resultString)
		})
	}
}

func TestRightPad(t *testing.T) {
	tests := map[string]struct {
		stringToPad        string
		pad                string
		numberOfTimesToPad int
		expectedString     string
	}{
		"pad once": {
			stringToPad:        "pad",
			pad:                " ",
			numberOfTimesToPad: 1,
			expectedString:     "pad ",
		},
		"pad twice": {
			stringToPad:        "pad",
			pad:                " ",
			numberOfTimesToPad: 2,
			expectedString:     "pad  ",
		},
		"pad 0 times": {
			stringToPad:        "pad",
			pad:                " ",
			numberOfTimesToPad: 0,
			expectedString:     "pad",
		},
		"pad -1 times": {
			stringToPad:        "pad",
			pad:                " ",
			numberOfTimesToPad: -1,
			expectedString:     "pad",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resultString := pad.RightPad(test.stringToPad, test.pad, test.numberOfTimesToPad)
			assert.Equal(t, test.expectedString, resultString)
		})
	}
}
