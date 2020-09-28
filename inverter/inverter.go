//Package inverter stores and retrieve the inverter info
package inverter

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
)

//Inverter represents the inverter
type Inverter struct {
	Info info `json:"info"`
	Bms  bms  `json:"bms"`
}

//Active stores the current (last loaded) inverter information
var Active Inverter

//Validate does the Inverter full validation
func (i *Inverter) Validate() error {
	err := i.Info.validate()
	if err != nil {
		return fmt.Errorf("Inverter Info error: %v", err)
	}
	err = i.Bms.validate()
	if err != nil {
		return fmt.Errorf("Bms error: %v", err)
	}
	return nil
}

//helper func - to handle numeric fields
func numerics(str string, multiplier float64) (string, error) {
	v, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return "", err
	}
	vstring := strconv.Itoa(int(v * multiplier))
	return vstring, nil
}

//helper func - to convert strings into byte slices (ascii representation), with padding
func strToBytes(str string, size byte, padding func([]byte, byte) []byte) []byte {
	bs := []byte(str)

	return padding(bs, size)
}

//helper func - to convert numeric fields into byte slices
func numToBytes(str string, multiplier float64, size byte) ([]byte, error) {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, err
	}
	v = v * multiplier

	buff := make([]byte, size)

	switch size {
	case 2:
		binary.BigEndian.PutUint16(buff, uint16(v))
	case 4:
		binary.BigEndian.PutUint32(buff, uint32(v))
	default:
		binary.BigEndian.PutUint32(buff, uint32(v))
	}

	return buff, nil
}

//helper func - adds padding to the right
func paddingRight(bs []byte, size byte) []byte {
	r := bs

	for len(r) < int(size) {
		r = append(r, 0x20)
	}
	return r
}

//helper func - adds padding to the left
func paddingLeft(bs []byte, size byte) []byte {
	r := []byte{}

	for len(r) < int(size)-len(bs) {
		r = append(r, 0x20)
	}

	r = append(r, bs...)
	return r
}

//helper func - size checker for numeric fields
func checkSize(str string, multiplier float64, maxSize byte) bool {

	if len(str) <= 0 {
		return false
	}

	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}

	v *= multiplier

	if multiplier == tempFactor {

		switch maxSize {
		case 2:
			return v <= math.MaxInt16 && v >= math.MinInt16
		case 4:
			return v <= math.MaxInt32 && v >= math.MinInt32
		default:
			return false
		}

	} else {

		switch maxSize {
		case 2:
			return v <= math.MaxUint16 && v >= 0
		case 4:
			return v <= math.MaxUint32 && v >= 0
		default:
			return false
		}
	}
}
