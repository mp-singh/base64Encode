package Decode

import (
	"fmt"
	"strings"
)

type Decode struct {
	Input             string
	Decoded           string
	binaryString      []string
	withSinglePadding bool
	withDoublePadding bool
}

//e.convertStringToBinary()
//e.splitBinaryStringIntoBits()
//e.convertBitsToDecimals()
//e.encode()
//e.addPadding()
func (d *Decode) base64Decode() {
	d.removePadding()
}

func (d *Decode) removePadding() {
	if d.Input[len(d.Input)-2:] == "==" {
		d.Decoded = fmt.Sprintf(d.Input[:len(d.Input)-2])
		return
	}
	if d.Input[len(d.Input)-1:] == "=" {
		d.Decoded = fmt.Sprintf(d.Input[:len(d.Input)-1])
		return
	}
}

//char := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/", "")
//for _, c := range e.decimalIntArray {
//e.Encoded = fmt.Sprintf("%s%s", e.Encoded, char[c])
//}

func (d *Decode) decode() {
	char := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/", "")
	for _, c := range d.Decoded {
		d.binaryString = append(d.binaryString, fmt.Sprintf("%s%s", d.binaryString, char[c]))
	}
}

func (d *Decode) convertDecimalToBinary() {


}

func (d *Decode) convertBinaryToString() {

}
