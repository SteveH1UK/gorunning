package utils

import "testing"

func TestOKDates(t *testing.T) {

	items := []string{"18-12-2013", "01-12-2010", "20-02-1922"}
	for _, item := range items {
		if !ValidateDate(item) {
			t.Error("invalid date " + item)
		}
	}
}

func TestBadDates(t *testing.T) {

	items := []string{"cake", "01/02/2000", "1-2-2010", "30-02-2000", "01-13-2000"}
	for _, item := range items {
		if ValidateDate(item) {
			t.Error("invalid date not detected " + item)
		}
	}
}
