package cell

import (
	"fmt"
	"testing"
)

func TestGenerateCellContent(t *testing.T) {
	for i := 1; i <= 9; i++ {
		expectedContent := fmt.Sprintf(" %v ", i)
		actualContent := GenerateCellContent(i)
		if expectedContent != actualContent {
			t.Errorf("\nExpected cell content: %v\nActual cell content: %v", expectedContent, actualContent)
		}
	}
}