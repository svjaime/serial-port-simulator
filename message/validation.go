package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//Validate checks header and checksum validation
func (h *Handler) validate() error {

	err := h.validateHeader()

	if err != nil {
		return err
	}

	err = h.validateChecksum()

	if err != nil {
		return err
	}

	err = h.validateDataLen()

	if err != nil {
		return err
	}

	return nil
}

//data length validation
func (h *Handler) validateDataLen() error {
	if h.dataLen() < 0 || uint16(h.dataLen()) != uint16(len(h.data())) {
		return fmt.Errorf("invalid data length")
	}
	return nil
}

//header validation
func (h *Handler) validateHeader() error {

	if h.header()[0] != 0xAA || h.header()[1] != 0xAA {
		return fmt.Errorf("invalid header")
	}

	return nil
}

//checksum validation
func (h *Handler) validateChecksum() error {
	var cs uint16
	for i, b := range h.Bs {
		if i < len(h.Bs)-2 {
			cs += uint16(b)
		}
	}

	expectedCs, err := toUint16(h.checksum())
	if err != nil {
		return fmt.Errorf("failed to convert checksum slice to uint16: %v", err)
	}

	if cs != expectedCs {
		return fmt.Errorf("invalid checksum")
	}

	return nil
}

//converts byte slice to uint16 (helper func)
func toUint16(bs []byte) (uint16, error) {
	var res uint16
	bBuf := bytes.NewReader(bs)
	err := binary.Read(bBuf, binary.BigEndian, &res)
	return res, err
}
