package gosrt

type Encryption uint8

const (
	NotEncrypted Encryption = iota
	EvenKey
	OddKey
	Control
)

// SRTMessageNumberField corresponds to the https://datatracker.ietf.org/doc/html/draft-sharabayko-srt-01#section-3.1
// section of the RFC. Only if packet type is Data packet
type SRTMessageNumberField struct {
	// isFirst [0] P bit: indicate whether it is the first packet
	isFirst bool

	// isLast [1] P bit: indicate whether it is the last packet
	isLast bool

	// inOrder [2] packet should be received in order or not
	inOrder bool

	// encrypted [3-4] Key-based encryption flag
	encrypted Encryption

	// retransmitted [5] Retransmitted Packet Flag
	retransmitted bool

	// number [6 - 32] : 26 bits. The sequential number of consecutive data packets that form a message
	number uint32
}

// Write writes the Message Number Field into the packet. We assume that there is enough space
// within packet.
func (s SRTMessageNumberField) Write(packet []byte) {
	var nmb = s.number
	var header uint32
	if s.isFirst {
		header |= 0b1000_0000
	}
	if s.isLast {
		header |= 0b0100_0000
	}
	if s.inOrder {
		header |= 0b0010_0000
	}
	switch s.encrypted {
	case NotEncrypted:
		// do nothing
	case EvenKey:
		header |= 0b0000_1000
	case OddKey:
		header |= 0b0001_0000
	case Control:
		header |= 0b0001_1000
	}
	if s.retransmitted {
		header |= 0b0000_0100
	}
	nmb |= header << 24
	packet[4] = byte(nmb >> 24)
	packet[5] = byte(nmb >> 16)
	packet[6] = byte(nmb >> 8)
	packet[7] = byte(nmb)
}

type SRTWriter struct {
}

func (s *SRTWriter) Write(data []byte) {

}
