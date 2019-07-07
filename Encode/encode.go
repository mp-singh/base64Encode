package Encode

import (
	"fmt"
	"strconv"
	"strings"
)

type Encoding struct {
	Input             string
	Encoded           string
	binaryString      string
	binaryStringArray []string
	decimalIntArray   []int64
	addSinglePadding  bool
	addDoublePadding  bool
}

func (e *Encoding) convertStringToBinary() {
	for _, char := range e.Input {
		e.binaryString = fmt.Sprintf("%s%.8b", e.binaryString, char)
	}
}

func (e *Encoding) Base64Encode() {
	e.convertStringToBinary()
	e.splitBinaryStringIntoBits()
	e.convertBitsToDecimals()
	e.encode()
	e.addPadding()
}

func (e *Encoding) encode() {
	char := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/", "")
	for _, c := range e.decimalIntArray {
		e.Encoded = fmt.Sprintf("%s%s", e.Encoded, char[c])
	}
}

func (e *Encoding) addPadding() {
	if e.addDoublePadding {
		e.Encoded = fmt.Sprintf("%s%s", e.Encoded, "==")
	}
	if e.addSinglePadding {
		e.Encoded = fmt.Sprintf("%s%s", e.Encoded, "=")
	}
}

func (e *Encoding) convertBitsToDecimals() {
	for _, binary := range e.binaryStringArray {
		decimal, _ := strconv.ParseInt(binary, 2, 0)
		e.decimalIntArray = append(e.decimalIntArray, decimal)
	}
}

func (e *Encoding) splitBinaryStringIntoBits() {
	// split the binary string into 6 bit slices and add extra bits for padding at the end if needed
	for i := 6; i < len(e.binaryString); i += 6 {
		e.binaryStringArray = append(e.binaryStringArray, e.binaryString[i-6:i])
		if len(e.binaryString[i:]) == 6 {
			e.binaryStringArray = append(e.binaryStringArray, e.binaryString[i:])
		}
		if len(e.binaryString[i:]) < 6 {
			e.addExtraBits(i)
		}
	}
}

func (e *Encoding) addExtraBits(i int) {
	var s string
	if len(e.binaryString[i:]) == 2 {
		s = fmt.Sprintf("%s0000", e.binaryString[i:])
		e.addDoublePadding = true
	} else if len(e.binaryString[i:]) == 4 {
		s = fmt.Sprintf("%s00", e.binaryString[i:])
		e.addSinglePadding = true
	}
	e.binaryStringArray = append(e.binaryStringArray, s)
}
