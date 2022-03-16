package morse

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// CharToCode - convert a character to a morse code representation
func CharToCode(ch string) (string, error) {
	if ch == "" {
		return "", errors.New("empty character")
	}
	morse, known := characters[strings.ToUpper(ch)]
	if !known {
		return "", errors.New("unknown character" + ch)
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

func getMorseCharOrProsign(text string) (string, int) {
	var retString strings.Builder
	var retBytes int
	var inProSign bool

	for {
		runeValue, runeWidth := utf8.DecodeRuneInString(text[retBytes:])
		if runeValue == utf8.RuneError {
			break
		}
		retBytes += runeWidth
		morseChar := string(runeValue)
		retString.WriteString(morseChar)
		if morseChar == "<" {
			inProSign = true
			continue
		}
		if inProSign && morseChar != ">" {
			continue
		}
		break
	}
	return retString.String(), retBytes
}

func StringToCode(text string) (string, error) {
	var retval strings.Builder
	for i := 0; i < len(text); {
		morseChar, charWidth := getMorseCharOrProsign(text[i:])
		i += charWidth
		if morseChar == " " { // space between words
			retval.WriteString("  ")
		} else {
			code, err := CharToCode(morseChar)
			if err != nil {
				return "", err
			}
			retval.WriteString(code)
			retval.WriteString(" ")
		}
	}
	return retval.String(), nil
}

func CodeToString(code string) (text string, err error) {
	var retval strings.Builder
	characters := strings.Split(strings.TrimSpace(code), " ")
	charText := ""
	for _, ch := range characters {
		charText, err = CodeToChar(ch)
		if err != nil {
			return
		}
		retval.WriteString(charText)
	}
	return retval.String(), nil
}

// StringToCodeWordSlice - convert each word of text to marose and return as a slice of strings
func StringToCodeWordSlice(text string) ([]string, error) {
	var retval []string = make([]string, 0, 10)
	words := strings.Split(text, " ")
	for _, word := range words {
		code, err := StringToCode(word)
		if err != nil {
			return make([]string, 0), err
		}
		retval = append(retval, code)
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
	":":    "---...",
	"<BT>": "-...-",
	"<AR>": ".-.-.",
	"<BK>": "-...-.-",
	"<SK>": "...-.-",
}
