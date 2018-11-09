package validation

import (
	"testing"
)

func TestOKDates(t *testing.T) {

	items := []string{"01-12-1999", "20-02-1922"}
	for _, item := range items {
		if !IsOldEnough(item, 16) {
			t.Error("Should be old enough " + item)
		}
	}
}
