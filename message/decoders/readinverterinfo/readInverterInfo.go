//Package readinverterinfo provides the necessary decoder to retrieve the inverter's info
package readinverterinfo

import (
	"ubiwhere.com/serial-port-simulator/inverter"
)

//Decoder to retrieve inverter's info (empty struct)
type Decoder struct {
	//r io.ReadCloser
}

//NewDecoder returns a decoder instance for reading inverter's information
func NewDecoder() Decoder {
	return Decoder{}
}

//Decode returns the inverter's info
func (d Decoder) Decode() ([]byte, error) {

	info, err := inverter.Active.Info.GetInfo()
	if err != nil {
		return nil, err
	}

	return info, nil
}
