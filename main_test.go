package main

import (
	"fmt"
	"testing"
)

func TestCardAt(t *testing.T) {
	testcases := []struct {
		number   int
		expected string
	}{
		{-1, "Incorrect input number."},
		{0, "2C"},
		{1, "3C"},
		{2, "4C"},
		{3, "5C"},
		{4, "6C"},
		{5, "7C"},
		{6, "8C"},
		{7, "9C"},
		{8, "0C"},
		{9, "JC"},
		{10, "QC"},
		{11, "KC"},
		{12, "AC"},
		{13, "2D"},
		{14, "3D"},
		{15, "4D"},
		{16, "5D"},
		{17, "6D"},
		{18, "7D"},
		{19, "8D"},
		{20, "9D"},
		{21, "0D"},
		{22, "JD"},
		{23, "QD"},
		{24, "KD"},
		{25, "AD"},
		{26, "2H"},
		{27, "3H"},
		{28, "4H"},
		{29, "5H"},
		{30, "6H"},
		{31, "7H"},
		{32, "8H"},
		{33, "9H"},
		{34, "0H"},
		{35, "JH"},
		{36, "QH"},
		{37, "KH"},
		{38, "AH"},
		{39, "2S"},
		{40, "3S"},
		{41, "4S"},
		{42, "5S"},
		{43, "6S"},
		{44, "7S"},
		{45, "8S"},
		{46, "9S"},
		{47, "0S"},
		{48, "JS"},
		{49, "QS"},
		{50, "KS"},
		{51, "AS"},
		{52, "Incorrect input number."},
	}

	for _, tc := range testcases {
		description := fmt.Sprintf("Should got %s when number is %d", tc.expected, tc.number)
		t.Run(description, func(t *testing.T) {
			result := cardAt(tc.number)
			if tc.expected != result {
				t.Errorf("got %s, expect %s", result, tc.expected)
			}
		})
	}
}
