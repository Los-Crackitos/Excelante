package headers

import (
	"testing"
)

func TestGetNextColIndex(t *testing.T) {
	index := 1234
	if index < 1 {
	}
	var col string
	for index > 0 {
		col = string((index-1)%26+65) + col
		index = (index - 1) / 26
	}
	t.Errorf(col)
	return
}
