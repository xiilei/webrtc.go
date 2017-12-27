package rtp

// PayloadType is the format of the payload
type PayloadType uint8

// CSRC is the contributing source ID
type CSRC uint32

// RTP is the header of an RTP packet.
type RTP struct {
	Version   uint8
	Padding   uint8
	Extension uint8
	CC        uint8
	Marker    uint8
	Type      PayloadType
	Seq       uint16
	Timestamp uint32
	SSRC      uint32
	CSRCs     []CSRC
}

// PacketType is RTCP packet type
type PacketType uint8

// RTCP is the header of an RTCP packet.
type RTCP struct {
	Version uint8
	Padding uint8
	RC      uint8
	Type    PacketType
	Length  uint16
	SSRC    uint32
}
