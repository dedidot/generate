package stringer

import (
	"testing"
)

func TestGenerateSting(t *testing.T) {
	generateNow := RandomStr(28, "nozero")

	var expectedResult = "dkdhsk"

	if generateNow != expectedResult {
		t.Fatalf("expected %s but Got %s", expectedResult, generateNow)
	}
}