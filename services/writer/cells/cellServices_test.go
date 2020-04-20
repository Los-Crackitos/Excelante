package cells

import (
	"fmt"
	"testing"
)

func TestGetCellCoordonates(t *testing.T) {
	columnIndex := 12
	lineIndex := 2
	if columnIndex < 1 {
		t.Errorf("index (%d) must be positive and superior than 1", columnIndex)
		return
	}
	var col string
	for columnIndex > 0 {
		col = string((columnIndex-1)%26+65) + col
		columnIndex = (columnIndex - 1) / 26
	}
	t.Errorf(fmt.Sprintf("%s%d", col, lineIndex))
	return

}
