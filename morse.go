package morse

import (
	"errors"
	"strings"
)

// CharToCode - convert a character to a morse code representation
func CharToCode(ch string) (string, error) {
	if ch == "" {
		return "", errors.New("empty character")
	}
	morse, known := characters[ch]
	if !known {
		return "", errors.New("unknown character")
	}
	return morse, nil
}

func CodeToChar(code string) (string, error) {
	if code == "" {
		return "", errors.New("empty code sequence")
	}
	text, known := reverse[code]
	if !known {
		return "", errors.New("unknown morse sequence")
	}
	return text, nil
}

func StringToCode(text string) (string, error) {
	var retval strings.Builder
	for _, ch := range text {
		if ch == ' ' { // space between words
			retval.WriteString("  ")
		} else {
			code, err := CharToCode(string(ch))
			if err != nil {
				return "", err
			}
			retval.WriteString(code)
			retval.WriteString(" ")
		}
	}
	return retval.String(), nil
}

// StringToCodeWordSlice - convert each word of text to marose and return as a slice of strings
func StringToCodeWordSlice(text string) ([]string, error) {
	var retval []string = make([]string)
	words := strings.Split(text, " ")
	for _, word := range text {
		code, err := StringToCode(word)
		if err != nil {
			return "", err
		}
		retval.append(code)
	}
	return retval, nil
}

func init() {
	for key, value := range characters {
		reverse[value] = key
	}
}

var reverse = map[string]string{}

var characters = map[string]string{
	"A":    ".-",
	"B":    "-...",
	"C":    "-.-.",
	"D":    "-..",
	"E":    ".",
	"F":    "..-.",
	"G":    "--.",
	"H":    "....",
	"I":    "..",
	"J":    ".---",
	"K":    "-.-",
	"L":    ".-..",
	"M":    "--",
	"N":    "-.",
	"O":    "---",
	"P":    ".--.",
	"Q":    "--.-",
	"R":    ".-.",
	"S":    "...",
	"T":    "-",
	"U":    "..-",
	"V":    "...-",
	"W":    ".--",
	"X":    "-..-",
	"Y":    "-.--",
	"Z":    "--..",
	"0":    "-----",
	"1":    ".----",
	"2":    "..---",
	"3":    "...--",
	"4":    "....-",
	"5":    ".....",
	"6":    "-....",
	"7":    "--...",
	"8":    "---..",
	"9":    "----.",
	".":    ".-.-.-",
	",":    "--..--",
	"?":    "..--..",
	"/":    "-..-.",
	"<BT>": "-...-",
	"<AR>": ".-.-.",
	"<BK>": "-...-.-",
	"<SK>": "...-.-",
}
