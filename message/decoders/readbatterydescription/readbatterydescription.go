//Package readbatterydescription provides the necessary decoder to retrieve battery info
package readbatterydescription

import (
	"ubiwhere.com/serial-port-simulator/inverter"
)

//Decoder to retrieve batery description (empty struct)
type Decoder struct {
	//r io.ReadCloser
}

//NewDecoder returns a decoder instance for reading battery description
func NewDecoder() Decoder {
	return Decoder{}
}

//Decode returns the battery description
func (d Decoder) Decode() ([]byte, error) {

	info, err := inverter.Active.Bms.Info.GetInfo()
	if err != nil {
		return nil, err
	}

	return info, nil
}
