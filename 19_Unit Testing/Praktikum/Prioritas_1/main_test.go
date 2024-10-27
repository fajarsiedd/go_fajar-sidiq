package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangleArea(t *testing.T) {
	testCases := []struct {
		length   int
		width    int
		expected string
	}{
		{3, 5, "odd rectangle"},
		{7, 4, "even rectangle"},
		{9, 7, "odd rectangle"},
		{6, 4, "even rectangle"},
	}

	for _, tc := range testCases {
		result := RectangleArea(tc.length, tc.width)

		assert.Equal(t, tc.expected, result)
	}

}

func TestRectanglePerimeter(t *testing.T) {
	testCases := []struct {
		length   int
		width    int
		expected bool
	}{
		{10, 5, true},
		{7, 3, true},
		{12, 7, false},
		{6, 3, false},
	}

	for _, tc := range testCases {
		result := RectanglePerimeter(tc.length, tc.width)

		assert.Equal(t, tc.expected, result)
	}
}

func TestSquareArea(t *testing.T) {
	testCases := []struct {
		s        int
		expected string
	}{
		{6, "even square"},
		{9, "odd square"},
		{2, "even square"},
		{5, "odd square"},
	}

	for _, tc := range testCases {
		result := SquareArea(tc.s)

		assert.Equal(t, tc.expected, result)
	}
}

func TestSquarePerimeter(t *testing.T) {
	testCases := []struct {
		s        int
		expected bool
	}{
		{15, true},
		{9, false},
		{10, true},
		{7, false},
	}

	for _, tc := range testCases {
		result := SquarePerimeter(tc.s)

		assert.Equal(t, tc.expected, result)
	}
}
