package board

import (
	"fmt"
	"testing"
)

func TestCreateCellContent(t *testing.T) {
	for i := 1; i <= 9; i++ {
		expectedContent := fmt.Sprintf(" %v ", i)
		actualContent := createCellContent(i)
		if expectedContent != actualContent {
			t.Errorf("\nExpected cell content: %v\nActual cell content: %v", expectedContent, actualContent)
		}
	}
}