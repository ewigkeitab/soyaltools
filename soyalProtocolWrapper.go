package soyaltools

import (
	"encoding/hex"
	"errors"
	"fmt"
)

// XORgen calc bytes' xor value
func XORgen(o []byte) byte {
	var ib byte = 0xff

	for _, d := range o {
		ib ^= d

	}

	return ib
}

// SUMgen calc sum 8 bit
func SUMgen(o []byte) byte {
	var ib byte = 0
	for _, d := range o {
		ib += d

	}

	return ib
}
func makeXORSUM(b *[]byte) {
	(*b)[len(*b)-2] = XORgen((*b)[2 : len(*b)-2])
	(*b)[len(*b)-1] = SUMgen((*b)[2 : len(*b)-1])
}

// SoyalOpenDoor21h make cmd
func SoyalOpenDoor21h(nid int, d0 byte, wg byte) []byte {

	pro := []byte{0x7e, 0x06, byte(nid), 0x21, d0, wg, 0, 0}
	makeXORSUM(&pro)
	// pro[6] = XORgen(pro[2:6])
	// pro[7] = SUMgen(pro[2:7])
	return pro
}

// SoyalReqNodeCardData87h make cmd
func SoyalReqNodeCardData87h(nid int, nodeAddr int, nodeNums int) []byte {

	pro := []byte{0x7e, 0x07, byte(nid), 0x87, byte(nodeAddr / 256), byte(nodeAddr % 256), byte(nodeNums), 0, 0}
	makeXORSUM(&pro)
	return pro
}

// SoyalReqNodeCardData84h make cmd
func SoyalReqNodeCardData84h(nid int, nodeAddr int, cardTag0, cardTag1 int) ([]byte, error) {
	parabase := "000000005800ffff63011f0000000000"
	tag := fmt.Sprintf("00000000%02x%02x", cardTag0, cardTag1)
	payload, err := hex.DecodeString(tag + parabase + "0000")

	if err == nil {
		pro := []byte{0x7e, 0x00, byte(nid), 0x83, 1, byte(nodeAddr / 256), byte(nodeAddr % 256)}
		pro = append(pro, payload...)
		pro[1] = (byte)(len(pro) - 2)
		makeXORSUM(&pro)
		return pro, nil
	}
	return nil, errors.New("payload error")
}

// SoyalReqNodeCardData85h erase card
func SoyalReqNodeCardData85h(nid int, nodeAddr int) []byte {

	pro := []byte{0x7e, 0x00, byte(nid), 0x85, byte(nodeAddr / 256), byte(nodeAddr % 256), byte(nodeAddr / 256), byte(nodeAddr % 256), 0, 0}

	pro[1] = (byte)(len(pro) - 2)
	makeXORSUM(&pro)
	return pro

}

// SoyalReqUserAlias2Eh
func SoyalReqUserAlias2Eh(nid int, nodeAddr int) []byte {

	pro := []byte{0x7e, 0x07, byte(nid), 0x2e, 0x0, byte(nodeAddr / 256), byte(nodeAddr % 256), 1, 0, 0}
	pro[1] = (byte)(len(pro) - 2)
	makeXORSUM(&pro)
	return pro
}
