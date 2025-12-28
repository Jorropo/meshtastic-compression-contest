package main

import (
	"log"

	"github.com/Jorropo/meshtastic-compression-contest/unishox2"
	"google.golang.org/protobuf/encoding/protowire"
)

const (
	TEXT_MESSAGE_APP = 1
	COMPRESSED       = 7
)

type compressorNum uint8

const (
	TEXT_MESSAGE_APP_COMPRESSOR compressorNum = 1
)

func explodePacketForPortnumPayloadSubstitution(substitute func(portnum uint64, payload []byte) (newPortnum uint64, newPayload []byte, changed bool)) compressor {
	return func(packet []byte) []byte {
		var portnum uint64
		var payload []byte
		var startOfPortnum, endOfPortnum, startOfPayload, endOfPayload int
		var parseCursor int
		for parseCursor < len(packet) {
			num, typ, n := protowire.ConsumeTag(packet[parseCursor:])
			if n < 0 {
				panic("invalid packet")
			}
			begin := parseCursor
			parseCursor += n

			switch num {
			case 1, 2:
				switch num {
				case 1:
					// portnum
					if typ != protowire.VarintType {
						panic("invalid portnum field type")
					}
					v, n := protowire.ConsumeVarint(packet[parseCursor:])
					if n < 0 {
						panic("invalid portnum value")
					}
					if uint64(uint32(v)) != v {
						panic("portnum value out of range")
					}
					portnum = v
					parseCursor += n
					startOfPortnum = begin
					endOfPortnum = parseCursor
				case 2:
					// payload
					if typ != protowire.BytesType {
						panic("invalid payload field type")
					}
					v, n := protowire.ConsumeBytes(packet[parseCursor:])
					if n < 0 {
						panic("invalid payload value")
					}
					payload = v
					parseCursor += n
					startOfPayload = begin
					endOfPayload = parseCursor
				}
			default:
				// Unknown field, skip it
				n := protowire.ConsumeFieldValue(num, typ, packet[parseCursor:])
				if n < 0 {
					panic("invalid field value")
				}
				parseCursor += n
			}
		}

		if portnum == 0 || payload == nil {
			panic("missing required fields")
		}

		newPortnum, newPayload, changed := substitute(portnum, payload)
		if !changed {
			return packet
		}

		newPacket := make([]byte, 0, 256)
		startFirst, endFirst, startLast, endLast := startOfPortnum, endOfPortnum, startOfPayload, endOfPayload
		if startFirst > startLast {
			startFirst, endFirst, startLast, endLast = startLast, endLast, startFirst, endFirst
		}
		newPacket = append(newPacket, packet[:startFirst]...)
		newPacket = append(newPacket, packet[endFirst:startLast]...)
		newPacket = append(newPacket, packet[endLast:]...)

		newPacket = protowire.AppendTag(newPacket, 1, protowire.VarintType)
		newPacket = protowire.AppendVarint(newPacket, uint64(newPortnum))

		newPacket = protowire.AppendTag(newPacket, 2, protowire.BytesType)
		newPacket = protowire.AppendBytes(newPacket, newPayload)

		return newPacket
	}
}

func compressPerPortnumTuned(portnum uint64, payload []byte) (uint64, []byte, bool) {
	var compressedPortnum compressorNum
	var compressedPayload []byte
	switch portnum {
	case TEXT_MESSAGE_APP:
		portnum = COMPRESSED
		compressedPortnum = TEXT_MESSAGE_APP_COMPRESSOR
		var err error
		compressedPayload, err = unishox2.CompressSimple(payload)
		if err != nil {
			log.Fatalf("unishox2 compression failed: %v", err)
		}
	default:
		return 0, nil, false
	}

	return COMPRESSED, append([]byte{byte(compressedPortnum)}, compressedPayload...), true
}
