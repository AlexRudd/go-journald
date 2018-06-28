package mock

import (
	"fmt"
	"testing"
)

func TestGetJournalEntries(t *testing.T) {
	a := getJournalEntries(10, 0)[5:]
	b := getJournalEntries(5, 5)

	for i := range a {
		if !isEqual(a[i], b[i]) {
			t.Fail()
		}
		fmt.Println(a[i])
	}
}

func isEqual(a, b Entry) bool {
	if len(a) != len(b) {
		return false
	}
	for k, _ := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}
