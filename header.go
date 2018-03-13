package triforce

import "encoding/binary"

type Header struct {
	Data []byte
}

func (h *Header) TriangleCount() uint32 {
	return binary.LittleEndian.Uint32(h.Data[80:84])
}
