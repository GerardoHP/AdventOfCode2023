package Days_test

import (
	"github.com/GerardoHP/AdventOfCode2023/Days"
	"testing"
)

func TestDay1(t *testing.T) {
	// Arrange
	codes := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expected := 142

	// Act
	result := Days.Day1(codes)

	// Assert
	if result != expected {
		t.Fatalf("expected %v, but found %v \n", expected, result)
	}
}

func TestDay1_part2(t *testing.T) {
	// Arrange
	codes := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := 281

	// Act
	result := Days.Day1(codes)

	// Assert
	if result != expected {
		t.Fatalf("expected %v, but found %v \n", expected, result)
	}
}
