package Days

import "testing"

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
	result := Day1(codes)

	// Assert
	if result != expected {
		t.Fatalf("expected %v, but found %v \n", expected, result)
	}
}
