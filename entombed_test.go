package main

import (
	"testing"
)

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

type renderLineTestCase struct {
	row      uint
	expected string
}

func TestRenderLine(t *testing.T) {
	tests := []renderLineTestCase{
		{
			row:      0b10101010,
			expected: "XX X X X  X X X XX",
		},
		{
			row:      0b11111111,
			expected: "XXXXXXXXXXXXXXXXXX",
		},
		{
			row:      0b00000000,
			expected: "X                X",
		},
		{
			row:      0b00100010,
			expected: "X  X   X  X   X  X",
		},
	}

	for _, test := range tests {
		result := RenderLine(test.row)
		if result != test.expected {
			t.Errorf("RenderLine(%b), expected: %s, got: %s", test.row, test.expected, result)
		}
	}
}

func TestGenerateRow(t *testing.T) {
	rows := []uint{128}
	var err error

	for i := 0; i < 10; i++ {
		rows, err = GenerateRow(rows)
	}

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	rowsLength := len(rows)
	expectedLength := 11
	if rowsLength != expectedLength {
		t.Errorf("Slice should have only %d rows, got %d",
			expectedLength, rowsLength)
	}

	for i := 0; i < 10; i++ {
		rows, err = GenerateRow(rows)
	}

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	rowsLength = len(rows)
	expectedLength = 12
	if rowsLength != expectedLength {
		t.Errorf("Slice should have only %d rows, got %d",
			expectedLength, rowsLength)
	}

	rows2 := []uint{}

	for i := 0; i < 20; i++ {
		rows2, err = GenerateRow(rows2)
	}

	if err == nil {
		t.Errorf("An empty slice is not a valid parameter")
	}
}
