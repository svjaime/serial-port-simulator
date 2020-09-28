package message

import (
	"encoding/binary"
	"fmt"
)

func getResponse(bs []byte) ([]byte, error) {

	msg := Handler{bs}

	err := msg.validate()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	resp := []byte{0xAA, 0xAA}
	resp = append(resp, msg.dest()...)
	resp = append(resp, msg.source()...)

	data, err := decodeRequest(msg)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	resp = append(resp, data...)
	resp = appendChecksum(resp)

	r := Handler{resp}
	err = r.validate()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return resp, nil
}

//appends checksum to a byte slice (helper func)
func appendChecksum(bs []byte) []byte {

	var cs uint16
	for _, b := range bs {
		cs += uint16(b)
	}

	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(cs))

	res := append(bs, b...)
	return res
}
