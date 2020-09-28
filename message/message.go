//Package message handles messaging
package message

//Handler struct represents the massage standard structure
type Handler struct {
	Bs []byte
}

//Getters

func (h *Handler) header() []byte {
	return h.Bs[:2]
}

func (h *Handler) source() []byte {
	return h.Bs[2:4]
}

func (h *Handler) dest() []byte {
	return h.Bs[4:6]
}

func (h *Handler) control() byte {
	return h.Bs[6]
}

func (h *Handler) function() byte {
	return h.Bs[7]
}

func (h *Handler) dataLen() byte {
	return h.Bs[8]
}

func (h *Handler) data() []byte {
	return h.Bs[9 : len(h.Bs)-2]
}

func (h *Handler) checksum() []byte {
	return h.Bs[len(h.Bs)-2:]
}
