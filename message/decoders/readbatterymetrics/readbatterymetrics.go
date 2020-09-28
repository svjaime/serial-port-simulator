//Package readbatterymetrics provides the necessary decoder to retrieve battery metrics
package readbatterymetrics

import (
	"ubiwhere.com/serial-port-simulator/inverter"
)

//Decoder to retrieve batery metrics (empty struct)
type Decoder struct {
	//r io.ReadCloser
}

//NewDecoder returns a decoder instance for reading battery metrics
func NewDecoder() Decoder {
	return Decoder{}
}

//Decode returns the battery metrics
func (d Decoder) Decode() ([]byte, error) {

	data, err := inverter.Active.Bms.Data.GetData()
	if err != nil {
		return nil, err
	}

	return data, nil
}
