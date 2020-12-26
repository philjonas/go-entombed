package entombed

import "testing"

func TestGetRandomBit(t *testing.T) {
	result := getRandomBit()
	if result > 1 {
		t.Errorf("getRandomBit() should return 0 or 1, got: %d", result)
	}
}

type isZeroPresentTestCase struct {
	lastRows []uint
	expected bool
}

func TestIsZeroPresent(t *testing.T) {
	tests := []isZeroPresentTestCase{
		{
			lastRows: []uint{0, 1, 2},
			expected: true,
		},
		{
			lastRows: []uint{1, 2},
			expected: false,
		},
		{
			lastRows: []uint{0},
			expected: true,
		},
		{
			lastRows: []uint{},
			expected: false,
		},
	}

	for _, test := range tests {
		result := isZeroPresent(test.lastRows)
		if result != test.expected {
			t.Errorf("isZeroPresent(%v), expected: %t, got: %t", test.lastRows, test.expected, result)
		}
	}
}

type doTheMagicBitNowTestCase struct {
	lasttwo    uint
	threeabove uint
	expected   int
}

func TestDoTheMagicBitNow(t *testing.T) {
	tests := []doTheMagicBitNowTestCase{
		{
			lasttwo:    0b11,
			threeabove: 0b111,
			expected:   noWall,
		},
		{
			lasttwo:    0b00,
			threeabove: 0b000,
			expected:   isWall,
		},
		{
			lasttwo:    0b00,
			threeabove: 0b110,
			expected:   randomWall,
		},
	}

	for _, test := range tests {
		result := doTheMagicBitNow(test.lasttwo, test.threeabove)
		if result != test.expected {
			t.Errorf("doTheMagicBitNow(%b, %b), expected: %d, got: %d",
				test.lasttwo, test.threeabove, test.expected, result)
		}
	}
}
