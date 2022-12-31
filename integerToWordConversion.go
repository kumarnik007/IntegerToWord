package main

import "strings"

// Constants for Conversion
const EMPTY string = ""

var LessThanTwenty = []string{EMPTY, "One ", "Two ", "Three ", "Four ", "Five ",
	"Six ", "Seven ", "Eight ", "Nine ", "Ten ", "Eleven ",
	"Twelve ", "Thirteen ", "Fourteen ", "Fifteen ",
	"Sixteen ", "Seventeen ", "Eighteen ", "Nineteen "}

var Tenths = []string{EMPTY, EMPTY, "Twenty ", "Thirty ", "Forty ", "Fifty ",
	"Sixty ", "Seventy ", "Eighty ", "Ninety "}

// Functions for Conversion
func convert2digit(inputNumber int, inputSuffix string) string {
	if inputNumber == 0 {
		return EMPTY
	}

	if inputNumber > 19 {
		return Tenths[inputNumber/10] + LessThanTwenty[inputNumber%10] + inputSuffix
	}

	return LessThanTwenty[inputNumber] + inputSuffix
}

func convert3digit(inputNumber int, inputSuffix string) string {
	if inputNumber == 0 {
		return EMPTY
	}

	if inputNumber > 99 {
		if ((inputNumber % 100) > 9) &&
			((inputNumber % 100) < 20) {
			return (LessThanTwenty[inputNumber/100] + "Hundred " + convert2digit((inputNumber%100), EMPTY) + inputSuffix)
		}

		return (LessThanTwenty[inputNumber/100] + "Hundred " + Tenths[(inputNumber/10)%10] + LessThanTwenty[inputNumber%10] + inputSuffix)
	}

	return convert2digit(inputNumber, inputSuffix)
}

func NumberToWords(inputNumber int) (string, error) {
	if inputNumber == 0 {
		return "Zero", nil
	}

	if 0 > inputNumber {
		return EMPTY, NewHTTPError(nil, 400, "Bad Input - Negative Integer is not accepted")
	}

	if 999999999 < inputNumber {
		return EMPTY, NewHTTPError(nil, 400, "Bad Input - Value greater than 999999999 cannot be processed")
	}

	var resultWord string = EMPTY

	resultWord = convert2digit((inputNumber % 100), "")
	resultWord = convert3digit(((inputNumber/100)%10), "Hundred ") + resultWord
	resultWord = convert3digit(((inputNumber/1000)%1000), "Thousand ") + resultWord
	resultWord = convert3digit(((inputNumber/1000000)%1000), "Million ") + resultWord

	return strings.TrimRight(resultWord, " "), nil
}
