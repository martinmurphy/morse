package morse

import (
	"testing"
)

// TestCharToCode calls morse.CharToCode with a character, checking
// for a valid return value.
func TestCharToCode(t *testing.T) {
	character := "A"
	want := ".-"
	code, err := CharToCode(character)
	if want != code || err != nil {
		t.Fatalf(`CharToCode = %q, %v, want match for %#q, nil`, code, err, want)
	}
}

// TestCharToCodeEmptyChar calls morse.CharToCode with an empty character, checking
// for an error return value.
func TestCharToCodeEmptyChar(t *testing.T) {
	character := ""
	code, err := CharToCode(character)
	if code != "" || err == nil {
		t.Fatalf(`CharToCode = %q, %v, want "", error`, code, err)
	}
}

// TestCharToCodeUnknownChar calls morse.CharToCode with an unknown character, checking
// for an error return value.
func TestCharToCodeUnknownChar(t *testing.T) {
	character := "ZZ"
	code, err := CharToCode(character)
	if code != "" || err == nil {
		t.Fatalf(`CharToCode = %q, %v, want "", error`, code, err)
	}
}

// TestCharToCodeUsingProwords calls morse.CharToCode with all known prowords, checking
// for valid return values.
func TestCharToCodeUsingProwords(t *testing.T) {
	prowords := []string{"<BT>", "<AR>", "<BK>", "<SK>"}
	for _, value := range prowords {
		code, err := CharToCode(value)
		if code == "" || err != nil {
			t.Fatalf(`CharToCode(%q) = %q, %v; do not want error`, value, code, err)
		}

	}
}

// TestCodeToChar calls morse.CodeToChar with a morse string character, checking
// for a valid return value.
func TestCodeToChar(t *testing.T) {
	code := "-..."
	want := "B"
	text, err := CodeToChar(code)
	if want != text || err != nil {
		t.Fatalf(`CodeToChar = %q, %v, want match for %#q, nil`, text, err, want)
	}
}

// TestCodeToCharEmptyChar calls morse.CodeToChar with an empty character, checking
// for an error return value.
func TestCodeToCharEmptyChar(t *testing.T) {
	code := ""
	character, err := CodeToChar(code)
	if character != "" || err == nil {
		t.Fatalf(`CodeToChar = %q, %v, want "", error`, character, err)
	}
}

// TestCodeToCharUnknownChar calls morse.CodeToChar with an unknown code sequence, checking
// for an error return value.
func TestCodeToCharUnknownCCode(t *testing.T) {
	code := "ZZ"
	character, err := CodeToChar(code)
	if character != "" || err == nil {
		t.Fatalf(`CodeToChar = %q, %v, want "", error`, character, err)
	}
}

// TestCodeToCharUsingProwords calls morse.CodeToChar with all known prowords, checking
// for valid return values.
func TestCodeToCharUsingProwords(t *testing.T) {
	prowords := []string{"-...-", ".-.-.", "-...-.-", "...-.-"}
	for _, value := range prowords {
		text, err := CodeToChar(value)
		if text == "" || err != nil {
			t.Fatalf(`CodeToChar(%q) = %q, %v; do not want error`, value, text, err)
		}

	}
}

// TestCharToCode calls morse.CharToCode with a character, checking
// for a valid return value.
func TestStringToCode(t *testing.T) {
	strings := []struct{ input, expected string }{
		{"", ""},
		{"A", ".- "},
		{"HELLO", ".... . .-.. .-.. --- "},
		{"HELLO WORLD 123", ".... . .-.. .-.. ---   .-- --- .-. .-.. -..   .---- ..--- ...-- "},
	}
	for _, tester := range strings {
		code, err := StringToCode(tester.input)
		if tester.expected != code || err != nil {
			t.Fatalf(`StringToCode(%q) = %q, %v, expected %q, nil`, tester.input, code, err, tester.expected)
		}

	}
}

// TestStringToCodeSlice calls morse.StringToCodeWordSlice with sentences, checking
// for a valid return value.
func TestStringToCodeSlice(t *testing.T) {
	strings := []struct{ input string, expected []string }{
		{"", {""}},
		{"A", {".- "}},
		{"HELLO", {".... . .-.. .-.. --- "}},
		{"HELLO WORLD 123", {".... . .-.. .-.. --- ", ".-- --- .-. .-.. -.. ",".---- ..--- ...-- "}},
	}
	for _, tester := range strings {
		code, err := StringToCodeWordSlice(tester.input)
		if tester.expected != code || err != nil {
			t.Fatalf(`StringToCodeWordSlice(%q) = %q, %v, expected %q, nil`, tester.input, code, err, tester.expected)
		}

	}
}
