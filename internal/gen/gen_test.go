package main

import (
	"slices"
	"testing"
)

func TestSplitDatePattern(t *testing.T) {
	t.Parallel()

	cases := []struct {
		pattern  string
		elements []datePatternElement
	}{
		{
			pattern:  "",
			elements: []datePatternElement{},
		},
		{
			pattern: "d",
			elements: []datePatternElement{
				{value: "d", literal: false},
			},
		},
		{
			pattern: "'", // NOTE(mvilks): should be invalid pattern but we don't care
			elements: []datePatternElement{
				{value: "'", literal: true},
			},
		},
		{
			pattern: "yMd",
			elements: []datePatternElement{
				{value: "y", literal: false},
				{value: "M", literal: false},
				{value: "d", literal: false},
			},
		},
		{
			pattern: "yyyy.MM.dd.",
			elements: []datePatternElement{
				{value: "yyyy", literal: false},
				{value: ".", literal: true},
				{value: "MM", literal: false},
				{value: ".", literal: true},
				{value: "dd", literal: false},
				{value: ".", literal: true},
			},
		},
		{
			pattern: "y. 'g.' M",
			elements: []datePatternElement{
				{value: "y", literal: false},
				{value: ". g. ", literal: true},
				{value: "M", literal: false},
			},
		},
		{
			pattern: "d''M''y. 'a''b' E",
			elements: []datePatternElement{
				{value: "d", literal: false},
				{value: "'", literal: true},
				{value: "M", literal: false},
				{value: "'", literal: true},
				{value: "y", literal: false},
				{value: ". a'b ", literal: true},
				{value: "E", literal: false},
			},
		},
		{
			pattern: "G y. 'gada' d. MMM, E – G y. 'gada' d. MMM, E",
			elements: []datePatternElement{
				{value: "G", literal: false},
				{value: " ", literal: true},
				{value: "y", literal: false},
				{value: ". gada ", literal: true},
				{value: "d", literal: false},
				{value: ". ", literal: true},
				{value: "MMM", literal: false},
				{value: ", ", literal: true},
				{value: "E", literal: false},
				{value: " – ", literal: true},
				{value: "G", literal: false},
				{value: " ", literal: true},
				{value: "y", literal: false},
				{value: ". gada ", literal: true},
				{value: "d", literal: false},
				{value: ". ", literal: true},
				{value: "MMM", literal: false},
				{value: ", ", literal: true},
				{value: "E", literal: false},
			},
		},
		{
			pattern: "EEEE, 'ngày' dd MMMM 'năm' y G",
			elements: []datePatternElement{
				{value: "EEEE", literal: false},
				{value: ", ngày ", literal: true},
				{value: "dd", literal: false},
				{value: " ", literal: true},
				{value: "MMMM", literal: false},
				{value: " năm ", literal: true},
				{value: "y", literal: false},
				{value: " ", literal: true},
				{value: "G", literal: false},
			},
		},
		{
			pattern: "'Ngày' dd",
			elements: []datePatternElement{
				{value: "Ngày ", literal: true},
				{value: "dd", literal: false},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.pattern, func(t *testing.T) {
			t.Parallel()

			elems := splitDatePattern(test.pattern)
			if !slices.Equal(elems, test.elements) {
				t.Errorf("want %v, got %v", test.elements, elems)
			}
		})
	}
}
