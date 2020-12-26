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
	}

	for _, test := range tests {
		result := isZeroPresent(test.lastRows)
		if result != test.expected {
			t.Errorf("isZeroPresent(%v), expected: %t, got: %t", test.lastRows, test.expected, result)
		}
	}
}
