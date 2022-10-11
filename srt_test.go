package gosrt

import "testing"

func TestSRTMessageNumberField_Write(t *testing.T) {
	field := SRTMessageNumberField{
		isFirst:       true,
		isLast:        false,
		inOrder:       false,
		encrypted:     0,
		retransmitted: false,
		number:        0b11_1010_1010_1010_1010_1010_1010,
	}
	exp := []byte{
		0, 0, 0, 0, // Packet sequence number
		0b1000_0011,
		0b1010_1010,
		0b1010_1010,
		0b1010_1010,
	}
	got := make([]byte, 8)
	field.Write(got)
	if got[4] != exp[4] {
		t.Fatalf("Not the same bytes: %b != %b", got[4], exp[4])
	}
	if got[5] != exp[5] {
		t.Fatalf("Not the same bytes: %b != %b", got[5], exp[5])
	}
	if got[6] != exp[6] {
		t.Fatalf("Not the same bytes: %b != %b", got[6], exp[6])
	}
	if got[7] != exp[7] {
		t.Fatalf("Not the same bytes: %b != %b", got[7], exp[7])
	}
}
