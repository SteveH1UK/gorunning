package utils

import "testing"

func TestContainsAllowedCharactersOnlyPasses(t *testing.T) {

	items := []string{"mfmfm", "amfmfm", "an-b_1", "_A-zeto_z"}
	for _, item := range items {
		if !ContainsAllowedCharactersOnly(item) {
			t.Error("Incorrectly rejected " + item)
		}
	}
}

func TestContainsAllowedCharactersOnlyFailures(t *testing.T) {

	items := []string{"#123d", "a?R", "£Jot", "_*e"}
	for _, item := range items {
		if ContainsAllowedCharactersOnly(item) {
			t.Error("Incorrectly rejected " + item)
		}
	}
}
