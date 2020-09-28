package message

import (
	"fmt"

	"ubiwhere.com/serial-port-simulator/message/decoders/readbatterydescription"
	"ubiwhere.com/serial-port-simulator/message/decoders/readbatterymetrics"
	"ubiwhere.com/serial-port-simulator/message/decoders/readinverterinfo"
)

type decoder interface {
	Decode() ([]byte, error)
}

func decodeRequest(msg Handler) ([]byte, error) {

	decoder, err := getDecoder(msg.control(), msg.function())
	if err != nil {
		return nil, err
	}

	resp := []byte{msg.control(), (msg.function() + 0x80)}

	data, err := decoder.Decode()
	if err != nil {
		return nil, err
	}

	resp = append(resp, byte(len(data)))
	resp = append(resp, data...)

	return resp, nil
}

func getDecoder(ctrlCode, fnCode byte) (decoder, error) {
	switch ctrlCode {
	case 0x01:
		switch fnCode {
		case 0x03:
			return readinverterinfo.NewDecoder(), nil
		case 0x09:
			return readbatterydescription.NewDecoder(), nil
		case 0x0A:
			return readbatterymetrics.NewDecoder(), nil
		}
	default:
		return nil, fmt.Errorf("Failed to get ... ")
	}
	return nil, fmt.Errorf("Failed to get ... ")
}
