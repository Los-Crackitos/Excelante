package headers

import (
	"fmt"
	"testing"
)


func TestGetNextColIndex(t *testing.T) {
	index := 36
	if index >= len(columnsIndex) { // If there are more than 25 headers, header index seems to be AA, AB, AC, AD...
		index = index - len(columnsIndex)
		t.Logf(fmt.Sprintf("A%s", columnsIndex[index]))
		return
		//return fmt.Sprintf("A%s", columnsIndex[index])
	}
	t.Logf(fmt.Sprintf("A%s", columnsIndex[index]))
	return
	//return fmt.Sprintf(columnsIndex[index])
}
