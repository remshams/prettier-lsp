package analysis

import (
	"testing"
)

func TestGetTextMetadataText(t *testing.T) {
	text := `First Line
  Second Line
  Third Line
  Fourth Line`

	metadata := GetTextMetadata(text)
	if metadata.LastLineNumber != 3 {
		t.Fatalf("Got %d as last line number, should be 3", metadata.LastLineCharacterNumber)
	}
	// 2: 2 tab indents
	// 11: Fourth Line
	if metadata.LastLineCharacterNumber != (11 + 2) {
		t.Fatalf("Got %d as last character number, should be 11", metadata.LastLineCharacterNumber)
	}
}

func TestGetTextMetadataEmptyText(t *testing.T) {
	metadata := GetTextMetadata("")

	if metadata.LastLineNumber != 0 {
		t.Fatalf("Got %d as last line number, should be 0", metadata.LastLineCharacterNumber)
	}
	if metadata.LastLineCharacterNumber != 0 {
		t.Fatalf("Got %d as last character number, should be 0", metadata.LastLineCharacterNumber)
	}
}

func TestGetTextMetadataSingleLineText(t *testing.T) {
	metadata := GetTextMetadata("Fourth Line")

	if metadata.LastLineNumber != 0 {
		t.Fatalf("Got %d as last line number, should be 0", metadata.LastLineNumber)
	}
	if metadata.LastLineCharacterNumber != 11 {
		t.Fatalf("Got %d as last character number, should be 11", metadata.LastLineCharacterNumber)
	}
}
