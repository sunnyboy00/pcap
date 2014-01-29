package pcap

import (
	"bytes"
	"syscall"
	"testing"
)

func TestDecodeSimpleTcpPacket(t *testing.T) {
	p := &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x20, 0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49,
			0x08, 0x00, 0x45, 0x00, 0x01, 0xa4, 0x39, 0xdf, 0x40, 0x00, 0x40, 0x06,
			0x55, 0x5a, 0xac, 0x11, 0x51, 0x49, 0xad, 0xde, 0xfe, 0xe1, 0xc5, 0xf7,
			0x00, 0x50, 0xc5, 0x7e, 0x0e, 0x48, 0x49, 0x07, 0x42, 0x32, 0x80, 0x18,
			0x00, 0x73, 0xab, 0xb1, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x03, 0x77,
			0x37, 0x9c, 0x42, 0x77, 0x5e, 0x3a, 0x47, 0x45, 0x54, 0x20, 0x2f, 0x20,
			0x48, 0x54, 0x54, 0x50, 0x2f, 0x31, 0x2e, 0x31, 0x0d, 0x0a, 0x48, 0x6f,
			0x73, 0x74, 0x3a, 0x20, 0x77, 0x77, 0x77, 0x2e, 0x66, 0x69, 0x73, 0x68,
			0x2e, 0x63, 0x6f, 0x6d, 0x0d, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
			0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x6b, 0x65, 0x65, 0x70, 0x2d, 0x61,
			0x6c, 0x69, 0x76, 0x65, 0x0d, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x2d, 0x41,
			0x67, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x4d, 0x6f, 0x7a, 0x69, 0x6c, 0x6c,
			0x61, 0x2f, 0x35, 0x2e, 0x30, 0x20, 0x28, 0x58, 0x31, 0x31, 0x3b, 0x20,
			0x4c, 0x69, 0x6e, 0x75, 0x78, 0x20, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34,
			0x29, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x57, 0x65, 0x62, 0x4b, 0x69,
			0x74, 0x2f, 0x35, 0x33, 0x35, 0x2e, 0x32, 0x20, 0x28, 0x4b, 0x48, 0x54,
			0x4d, 0x4c, 0x2c, 0x20, 0x6c, 0x69, 0x6b, 0x65, 0x20, 0x47, 0x65, 0x63,
			0x6b, 0x6f, 0x29, 0x20, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2f, 0x31,
			0x35, 0x2e, 0x30, 0x2e, 0x38, 0x37, 0x34, 0x2e, 0x31, 0x32, 0x31, 0x20,
			0x53, 0x61, 0x66, 0x61, 0x72, 0x69, 0x2f, 0x35, 0x33, 0x35, 0x2e, 0x32,
			0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x3a, 0x20, 0x74, 0x65,
			0x78, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x2c, 0x61, 0x70, 0x70, 0x6c,
			0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x68, 0x74, 0x6d,
			0x6c, 0x2b, 0x78, 0x6d, 0x6c, 0x2c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
			0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x6d, 0x6c, 0x3b, 0x71, 0x3d,
			0x30, 0x2e, 0x39, 0x2c, 0x2a, 0x2f, 0x2a, 0x3b, 0x71, 0x3d, 0x30, 0x2e,
			0x38, 0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x45, 0x6e,
			0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x67, 0x7a, 0x69, 0x70,
			0x2c, 0x64, 0x65, 0x66, 0x6c, 0x61, 0x74, 0x65, 0x2c, 0x73, 0x64, 0x63,
			0x68, 0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x4c, 0x61,
			0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x65, 0x6e, 0x2d, 0x55,
			0x53, 0x2c, 0x65, 0x6e, 0x3b, 0x71, 0x3d, 0x30, 0x2e, 0x38, 0x0d, 0x0a,
			0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x43, 0x68, 0x61, 0x72, 0x73,
			0x65, 0x74, 0x3a, 0x20, 0x49, 0x53, 0x4f, 0x2d, 0x38, 0x38, 0x35, 0x39,
			0x2d, 0x31, 0x2c, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x3b, 0x71, 0x3d, 0x30,
			0x2e, 0x37, 0x2c, 0x2a, 0x3b, 0x71, 0x3d, 0x30, 0x2e, 0x33, 0x0d, 0x0a,
			0x0d, 0x0a,
		}}
	p.Decode()
	if !bytes.Equal(p.DestMac, []byte{0, 0, 0x0c, 0x9f, 0xf0, 0x20}) {
		t.Error("Dest mac", p.DestMac)
	}
	if !bytes.Equal(p.SrcMac, []byte{0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49}) {
		t.Error("Src mac", p.SrcMac)
	}
	if len(p.Headers) != 2 {
		t.Fatal("Incorrect number of headers", len(p.Headers))
	}
	ip, ipOk := p.Headers[0].(*IPHdr)
	if !ipOk {
		t.Fatal("First header is not IP header")
	}
	if ip.Version != 4 {
		t.Error("ip Version", ip.Version)
	}
	if ip.Ihl != 5 {
		t.Error("ip raw header length", ip.Ihl)
	}
	if ip.Tos != 0 {
		t.Error("ip TOS", ip.Tos)
	}
	if ip.Length != 420 {
		t.Error("ip Length", ip.Length)
	}
	if ip.ID != 14815 {
		t.Error("ip ID", ip.ID)
	}
	if ip.Flags != 0x02 {
		t.Error("ip Flags", ip.Flags)
	}
	if ip.FragmentOffset != 0 {
		t.Error("ip Fragoffset", ip.FragmentOffset)
	}
	if ip.Ttl != 64 {
		t.Error("ip TTL", ip.Ttl)
	}
	if ip.Protocol != 6 {
		t.Error("ip Protocol", ip.Protocol)
	}
	if ip.Checksum != 0x555A {
		t.Error("ip Checksum", ip.Checksum)
	}
	if !bytes.Equal(ip.SrcIP, []byte{172, 17, 81, 73}) {
		t.Error("ip Src", ip.SrcIP)
	}
	if !bytes.Equal(ip.DestIP, []byte{173, 222, 254, 225}) {
		t.Error("ip Dest", ip.DestIP)
	}
	tcp, tcpOk := p.Headers[1].(*TCPHdr)
	if !tcpOk {
		t.Fatal("Second header is not TCP header")
	}
	if tcp.SrcPort != 50679 {
		t.Error("tcp srcport", tcp.SrcPort)
	}
	if tcp.DestPort != 80 {
		t.Error("tcp destport", tcp.DestPort)
	}
	if tcp.Seq != 0xc57e0e48 {
		t.Error("tcp seq", tcp.Seq)
	}
	if tcp.Ack != 0x49074232 {
		t.Error("tcp ack", tcp.Ack)
	}
	if tcp.DataOffset != 8 {
		t.Error("tcp dataoffset", tcp.DataOffset)
	}
	if tcp.Flags != 0x18 {
		t.Error("tcp flags", tcp.Flags)
	}
	if tcp.Window != 0x73 {
		t.Error("tcp window", tcp.Window)
	}
	if tcp.Checksum != 0xabb1 {
		t.Error("tcp checksum", tcp.Checksum)
	}
	if tcp.Urgent != 0 {
		t.Error("tcp urgent", tcp.Urgent)
	}
	if string(p.Payload) != "GET / HTTP/1.1\r\nHost: www.fish.com\r\nConnection: keep-alive\r\nUser-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/15.0.874.121 Safari/535.2\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip,deflate,sdch\r\nAccept-Language: en-US,en;q=0.8\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.3\r\n\r\n" {
		t.Error("--- PAYLOAD STRING ---\n", string(p.Payload), "\n--- PAYLOAD BYTES ---\n", p.Payload)
	}
}

// Makes sure packet payload doesn't display the 6 trailing null of this packet
// as part of the payload.  They're actually the ethernet trailer.
func TestDecodeSmallTcpPacketHasEmptyPayload(t *testing.T) {
	p := &Packet{
		// This packet is only 54 bits (an empty TCP RST), thus 6 trailing null
		// bytes are added by the ethernet layer to make it the minimum packet size.
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49, 0xb8, 0xac, 0x6f, 0x92, 0xd5, 0xbf,
			0x08, 0x00, 0x45, 0x00, 0x00, 0x28, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06,
			0x3f, 0x9f, 0xac, 0x11, 0x51, 0xc5, 0xac, 0x11, 0x51, 0x49, 0x00, 0x63,
			0x9a, 0xef, 0x00, 0x00, 0x00, 0x00, 0x2e, 0xc1, 0x27, 0x83, 0x50, 0x14,
			0x00, 0x00, 0xc3, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}}
	p.Decode()
	if p.Payload == nil {
		t.Error("Nil payload")
	}
	if len(p.Payload) != 0 {
		t.Error("Non-empty payload:", p.Payload)
	}
}

func TestDecodeMaliciousIPHeaderLength(t *testing.T) {
	p := &Packet{
		// This packet is only 54 bits (an empty TCP RST), thus 6 trailing null
		// bytes are added by the ethernet layer to make it the minimum packet size.
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49, 0xb8, 0xac, 0x6f, 0x92, 0xd5, 0xbf,
			0x08, 0x00, 0x4f, 0x00, 0x00, 0x28, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06,
			0x3f, 0x9f, 0xac, 0x11, 0x51, 0xc5, 0xac, 0x11, 0x51, 0x49, 0x00, 0x63,
			0x9a, 0xef, 0x00, 0x00, 0x00, 0x00, 0x2e, 0xc1, 0x27, 0x83, 0x50, 0x14,
			0x00, 0x00, 0xc3, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}}
	p.Decode()
}

func TestDecodeTruncatedUpperLayer(t *testing.T) {
	// TCP
	p := &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x20, 0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49,
			0x08, 0x00, 0x45, 0x00, 0x01, 0xa4, 0x39, 0xdf, 0x40, 0x00, 0x40, syscall.IPPROTO_TCP,
			0x55, 0x5a, 0xac, 0x11, 0x51, 0x49, 0xad, 0xde, 0xfe, 0xe1, 0xc5,
		}}
	p.Decode()

	// ICMP
	p = &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x20, 0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49,
			0x08, 0x00, 0x45, 0x00, 0x01, 0xa4, 0x39, 0xdf, 0x40, 0x00, 0x40, syscall.IPPROTO_ICMP,
			0x55, 0x5a, 0xac, 0x11, 0x51, 0x49, 0xad, 0xde, 0xfe, 0xe1, 0xc5,
		}}
	p.Decode()

	// UDP
	p = &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x20, 0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49,
			0x08, 0x00, 0x45, 0x00, 0x01, 0xa4, 0x39, 0xdf, 0x40, 0x00, 0x40, syscall.IPPROTO_UDP,
			0x55, 0x5a, 0xac, 0x11, 0x51, 0x49, 0xad, 0xde, 0xfe, 0xe1, 0xc5,
		}}
	p.Decode()
}

func TestDecodeMaliciousTCPDataOffset(t *testing.T) {
	p := &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x20, 0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49,
			0x08, 0x00, 0x45, 0x00, 0x01, 0xa4, 0x39, 0xdf, 0x40, 0x00, 0x40, 0x06,
			0x55, 0x5a, 0xac, 0x11, 0x51, 0x49, 0xad, 0xde, 0xfe, 0xe1, 0xc5, 0xf7,
			0x00, 0x50, 0xc5, 0x7e, 0x0e, 0x48, 0x49, 0x07, 0x42, 0x32, 0xf0, 0x18,
			0x00, 0x73, 0xab, 0xb1, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x03, 0x77,
			0x37, 0x9c, 0x42, 0x77, 0x5e, 0x3a,
		}}
	p.Decode()
}

func TestDecodeLinuxCooked(t *testing.T) {
	// IPv4 and a UDP packet.
	p := &Packet{
		DatalinkType: DLTLINUXSSL,
		Data: []byte{
			0x00, 0x04, 0xff, 0xfe, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x45, 0x00, 0x00, 0x30, 0x3f, 0xbc, 0x00, 0x00, 0x40, 0x11, 0x59, 0xb1, 0x2e, 0xf6, 0x22, 0xa4, 0x5c, 0x81, 0x33, 0x35, 0x1a, 0xe1, 0xd3, 0x99, 0x00, 0x1c, 0x8d, 0xe4, 0x21, 0x00, 0x80, 0x98, 0x45, 0x23, 0xde, 0x2a, 0xca, 0x0d, 0xc1, 0x5c, 0x06, 0x3f, 0xfa, 0x89, 0x65, 0xe9, 0xeb, 0x02,
		},
	}
	p.Decode()
	if p.DestMac != nil {
		t.Error("Dest mac", p.DestMac)
	}
	if p.SrcMac != nil {
		t.Error("Src mac", p.SrcMac)
	}
	if len(p.Headers) != 2 {
		t.Fatal("Incorrect number of headers", len(p.Headers))
	}

	ip, ipOk := p.Headers[0].(*IPHdr)
	if !ipOk {
		t.Fatal("First header is not an IP header")
	}
	if ip.Version != 4 {
		t.Error("ip Version", ip.Version)
	}
	if ip.Ihl != 5 {
		t.Error("ip raw header length", ip.Ihl)
	}

	udp, udpOk := p.Headers[1].(*UDPHdr)
	if !udpOk {
		t.Fatal("Second header is not a UDP header")
	}
	if udp.SrcPort != 6881 {
		t.Error("udp srcport", udp.SrcPort)
	}
	if udp.DestPort != 54169 {
		t.Error("udp destport", udp.DestPort)
	}

	// Leftover payload (so this is UDP payload)
	if len(p.Payload) != 20 {
		t.Error("Wrong payload length")
	}
}

func TestDecodeRaw(t *testing.T) {
	// IPv4 and a UDP packet in a RAW socket.
	p := &Packet{
		DatalinkType: DLTRAW,
		Data: []byte{
			0x45, 0x00, 0x00, 0x2c, 0x71, 0x7f, 0x00, 0x00, 0x01, 0x11, 0x0f, 0xaf, 0x2e, 0xf6, 0x29, 0x9c, 0xe0, 0x00, 0x00, 0x01, 0xf9, 0x15, 0x21, 0xa4, 0x00, 0x18, 0x10, 0xde, 0x50, 0x4e, 0x4a, 0x42, 0x01, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	}
	p.Decode()
	if p.DestMac != nil {
		t.Error("Dest mac", p.DestMac)
	}
	if p.SrcMac != nil {
		t.Error("Src mac", p.SrcMac)
	}
	if len(p.Headers) != 2 {
		t.Fatal("Incorrect number of headers", len(p.Headers))
	}

	ip, ipOk := p.Headers[0].(*IPHdr)
	if !ipOk {
		t.Fatal("First header is not an IP header")
	}
	if ip.Version != 4 {
		t.Error("ip Version", ip.Version)
	}
	if ip.Ihl != 5 {
		t.Error("ip raw header length", ip.Ihl)
	}

	udp, udpOk := p.Headers[1].(*UDPHdr)
	if !udpOk {
		t.Fatal("Second header is not a UDP header")
	}
	if udp.SrcPort != 63765 {
		t.Error("udp srcport", udp.SrcPort)
	}
	if udp.DestPort != 8612 {
		t.Error("udp destport", udp.DestPort)
	}

	// Leftover payload (so this is UDP payload)
	if len(p.Payload) != 16 {
		t.Error("Wrong payload length")
	}
}

func TestDecodeICMPv6(t *testing.T) {
	// ICMPv6 neighbor solicitation
	p := &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x33, 0x33, 0xff, 0x0e, 0x04, 0x63, 0x08, 0x96, 0xd7, 0x07, 0x93, 0x0d, 0x86, 0xdd, 0x60, 0x00, 0x00, 0x00, 0x00, 0x20, 0x3a, 0xff, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x96, 0xd7, 0xff, 0xfe, 0x07, 0x93, 0x0d, 0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xff, 0x0e, 0x04, 0x63, 0x87, 0x00, 0xef, 0x25, 0x00, 0x00, 0x00, 0x00, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x3e, 0x84, 0xff, 0xfe, 0x0e, 0x04, 0x63, 0x01, 0x01, 0x08, 0x96, 0xd7, 0x07, 0x93, 0x0d,
		},
	}
	p.Decode()
	if !bytes.Equal(p.DestMac, []byte{0x33, 0x33, 0xff, 0x0e, 0x04, 0x63}) {
		// 33:33:ff:0e:04:63
		t.Error("Dest mac", p.DestMac)
	}
	if !bytes.Equal(p.SrcMac, []byte{0x08, 0x96, 0xd7, 0x07, 0x93, 0x0d}) {
		// 08:96:d7:07:93:0d
		t.Error("Src mac", p.SrcMac)
	}
	if len(p.Headers) != 2 {
		t.Fatal("Incorrect number of headers", len(p.Headers))
	}

	ip, ipOk := p.Headers[0].(*IP6Hdr)
	if !ipOk {
		t.Fatal("First header is not an IPv6 header")
	}
	if ip.Version != 6 {
		t.Error("ip Version", ip.Version)
	}
	if ip.Length != 32 {
		t.Error("ipv6 payload length", ip.Length)
	}
	if ip.SrcAddr() != "fd00::a96:d7ff:fe07:930d" {
		t.Error("ipv6 src address", ip.SrcIP)
	}
	if ip.DestAddr() != "ff02::1:ff0e:463" {
		t.Error("ipv6 dest address", ip.DestIP)
	}
	if ip.Fragmented() {
		t.Error("not fragmented")
	}

	icmp, icmpOk := p.Headers[1].(*ICMPv6Hdr)
	if !icmpOk {
		t.Fatal("Second header is not an ICMPv6 header")
	}
	if icmp.Type != 135 {
		t.Error("ICMPv6 type", icmp.Type)
	}
	if icmp.Code != 0 {
		t.Error("ICMPv6 code", icmp.Code)
	}

	// Leftover payload (so this is ICMPv6's payload)
	if len(p.Payload) != 32-8 {
		t.Error("Wrong ICMPv6 payload length", len(p.Payload))
	}
}

func TestDecodeIPv6FragmentFirst(t *testing.T) {
	// IPv6 packet with 'fragment' extended header.
	// This is the first fragment, so we can analyze the payload.
	// (packets are generated with `ping6 -p DEADBEEF -s 40000 somehost`)
	p := &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x80, 0xee, 0x73, 0x83, 0x58, 0x8f, 0x1c, 0x3e, 0x84, 0x0e, 0x04,
			0x63, 0x86, 0xdd, 0x60, 0x00, 0x00, 0x00, 0x05, 0xb0, 0x2c, 0xff,
			0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x3e, 0x84,
			0xff, 0xfe, 0x0e, 0x04, 0x63, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x1d, 0xb4, 0x9c, 0x8c, 0x8d, 0x9c, 0xe2, 0xd5, 0x3a,
			0x00, 0x00, 0x01, 0x63, 0x21, 0x8d, 0x29, 0x81, 0x00, 0x7d, 0x18,
			0x20, 0xda, 0x00, 0x01, 0x7e, 0xde, 0xe7, 0x52, 0x00, 0x00, 0x00,
			0x00, 0xab, 0x8c, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0xde, 0xad,
			0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef,
			// ... and quite a few more deadbeefs.
		},
	}
	p.Decode()
	if !bytes.Equal(p.SrcMac, []byte{0x1c, 0x3e, 0x84, 0x0e, 0x04, 0x63}) {
		// 1c:3e:84:0e:04:63
		t.Errorf("Src mac %v", p.SrcMac)
	}
	if !bytes.Equal(p.DestMac, []byte{0x80, 0xee, 0x73, 0x83, 0x58, 0x8f}) {
		// 80:ee:73:83:58:8f
		t.Errorf("Dest mac %v", p.DestMac)
	}
	if len(p.Headers) != 2 {
		t.Fatal("Incorrect number of headers", len(p.Headers))
	}

	ip, ipOk := p.Headers[0].(*IP6Hdr)
	if !ipOk {
		t.Fatal("First header is not an IPv6 header")
	}
	if ip.Version != 6 {
		t.Error("ip Version", ip.Version)
	}
	if ip.Len() != 1496 {
		// We chopped the test bytes above.
		t.Error("ipv6 length", ip.Len())
	}
	if ip.PayloadLen() != 1448 {
		// We chopped the test bytes above.
		t.Error("ipv6 payload length", ip.Length)
	}
	if ip.SrcAddr() != "fd00::1e3e:84ff:fe0e:463" {
		t.Error("ipv6 src address", ip.SrcIP)
	}
	if ip.DestAddr() != "fd00::1db4:9c8c:8d9c:e2d5" {
		t.Error("ipv6 dest address", ip.DestIP)
	}
	if !ip.Fragmented() {
		t.Error("not fragmented")
	}

	icmp, icmpOk := p.Headers[1].(*ICMPv6Hdr)
	if !icmpOk {
		t.Fatal("Second header is not an ICMPv6 header")
	}
	// ping reply
	if icmp.Type != 129 {
		t.Error("ICMPv6 type", icmp.Type)
	}
	if icmp.Code != 0 {
		t.Error("ICMPv6 code", icmp.Code)
	}

	// Leftover of our chopped payload
	if len(p.Payload) != 32-8 {
		t.Error("Wrong ICMPv6 payload length", len(p.Payload))
	}
}

func TestDecodeIPv6FragmentEtc(t *testing.T) {
	// IPv6 packet with 'fragment' extended header.
	// This fragment is the final one, so we can't look at the protocol
	// (packets are generated with `ping6 -p DEADBEEF -s 40000 somehost`)
	p := &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x80, 0xee, 0x73, 0x83, 0x58, 0x8f, 0x1c, 0x3e, 0x84, 0x0e, 0x04,
			0x63, 0x86, 0xdd, 0x60, 0x00, 0x00, 0x00, 0x03, 0x98, 0x2c, 0xff,
			0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x3e, 0x84,
			0xff, 0xfe, 0x0e, 0x04, 0x63, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x1d, 0xb4, 0x9c, 0x8c, 0x8d, 0x9c, 0xe2, 0xd5, 0x3a,
			0x00, 0x98, 0xb8, 0x63, 0x21, 0x8d, 0x29, 0xde, 0xad, 0xbe, 0xef,
			0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef,
			// ... and quite a few more deadbeefs.
		},
	}
	p.Decode()
	if !bytes.Equal(p.SrcMac, []byte{0x1c, 0x3e, 0x84, 0x0e, 0x04, 0x63}) {
		// 1c:3e:84:0e:04:63
		t.Errorf("Src mac %v", p.SrcMac)
	}
	if !bytes.Equal(p.DestMac, []byte{0x80, 0xee, 0x73, 0x83, 0x58, 0x8f}) {
		// 80:ee:73:83:58:8f
		t.Errorf("Dest mac %v", p.DestMac)
	}
	// Not the first fragment, so we can't look at the payload
	if len(p.Headers) != 2 {
		t.Fatal("Incorrect number of headers", len(p.Headers))
	}

	// First header. IPv6.
	ip, ipOk := p.Headers[0].(*IP6Hdr)
	if !ipOk {
		t.Fatal("First header is not an IPv6 header")
	}
	if ip.Version != 6 {
		t.Error("ip Version", ip.Version)
	}
	if ip.Len() != 960 {
		// We chopped the test packet above.
		t.Error("ipv6 length", ip.Len())
	}
	if ip.PayloadLen() != 920-8 {
		// We chopped the test packet above.
		t.Error("ipv6 payload length", ip.PayloadLen())
	}
	if ip.SrcAddr() != "fd00::1e3e:84ff:fe0e:463" {
		t.Error("ipv6 src address", ip.SrcIP)
	}
	if ip.DestAddr() != "fd00::1db4:9c8c:8d9c:e2d5" {
		t.Error("ipv6 dest address", ip.DestIP)
	}
	if !ip.Fragmented() {
		t.Error("not fragmented")
	}
	if ip.FragmentOffset != 4887 {
		t.Error("Wrong raw fragment offset", ip.FragmentOffset)
	}

	// Fragment.
	f, fOk := p.Headers[1].(*Fragment)
	if !fOk {
		t.Fatal("Second part is not an Fragment")
	}
	if f.Length != 12 {
		// We truncated the testbytes above
		t.Error("fragment length", f.Length)
	}
	if f.ProtocolID != syscall.IPPROTO_ICMPV6 {
		t.Error("fragment protocol id", f.ProtocolID)
	}
}

func TestDecodeIPFragmentFirst(t *testing.T) {
	// IPv4 fragmented packet.
	// This is the first fragment, so we can analyze the payload.
	// (packets are generated with `ping -p DEADBEEF -s 40000 somehost`)
	p := &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x1c, 0x3e, 0x84, 0x0e, 0x04, 0x63, 0x80, 0xee, 0x73, 0x83, 0x58,
			0x8f, 0x08, 0x00, 0x45, 0x00, 0x05, 0xdc, 0xa2, 0xe4, 0x20, 0x00,
			0x40, 0x01, 0x2c, 0x69, 0xc0, 0xa8, 0x02, 0x6f, 0xc0, 0xa8, 0x02,
			0x14, 0x08, 0x00, 0xc7, 0x5f, 0x2e, 0x2f, 0x00, 0x01, 0x9f, 0xf1,
			0xe7, 0x52, 0x00, 0x00, 0x00, 0x00, 0x06, 0xc5, 0x0e, 0x00, 0x00,
			0x00, 0x00, 0x00, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef,
			0xde, 0xad, 0xbe, 0xef,
			// ... and quite a few more deadbeefs.
		},
	}
	p.Decode()
	if !bytes.Equal(p.SrcMac, []byte{0x80, 0xee, 0x73, 0x83, 0x58, 0x8f}) {
		// 80:ee:73:83:58:8f
		t.Errorf("Src mac %v", p.SrcMac)
	}
	if !bytes.Equal(p.DestMac, []byte{0x1c, 0x3e, 0x84, 0x0e, 0x04, 0x63}) {
		// 1c:3e:84:0e:04:63
		t.Errorf("Dest mac %v", p.DestMac)
	}
	if len(p.Headers) != 2 {
		t.Fatal("Incorrect number of headers", len(p.Headers))
	}

	ip, ipOk := p.Headers[0].(*IPHdr)
	if !ipOk {
		t.Fatal("First header is not an IPv4 header")
	}
	if ip.Version != 4 {
		t.Error("ip Version", ip.Version)
	}
	if ip.Len() != 1500 {
		// We chopped the test bytes above.
		t.Error("ip length", ip.Length)
	}
	if ip.PayloadLen() != 1480 {
		// We chopped the test bytes above.
		t.Error("ip payload length", ip.PayloadLen())
	}
	if ip.SrcAddr() != "192.168.2.111" {
		t.Error("ip src address", ip.SrcAddr())
	}
	if ip.DestAddr() != "192.168.2.20" {
		t.Error("ip dest address", ip.DestAddr())
	}
	if !ip.Fragmented() {
		t.Error("not fragmented")
	}

	icmp, icmpOk := p.Headers[1].(*ICMPHdr)
	if !icmpOk {
		t.Fatal("Second header is not an ICMP header")
	}
	// ping request
	if icmp.Type != 8 {
		t.Error("ICMP type", icmp.Type)
	}
	if icmp.Code != 0 {
		t.Error("ICMP code", icmp.Code)
	}

	// Leftover of our chopped payload
	if len(p.Payload) != 28 {
		t.Error("Wrong ICMP payload length", len(p.Payload))
	}
}

func TestDecodeIPFragmentLast(t *testing.T) {
	// IPv4 fragmented packet.
	// This is not the first fragment, so we can not analyze the payload.
	// (packets are generated with `ping -p DEADBEEF -s 40000 somehost`)
	p := &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x1c, 0x3e, 0x84, 0x0e, 0x04, 0x63, 0x80, 0xee, 0x73, 0x83, 0x58,
			0x8f, 0x08, 0x00, 0x45, 0x00, 0x00, 0x44, 0xa2, 0xe4, 0x13, 0x83,
			0x40, 0x01, 0x3e, 0x7e, 0xc0, 0xa8, 0x02, 0x6f, 0xc0, 0xa8, 0x02,
			0x14, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad,
			0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xde,
			0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef,
			0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe,
			0xef, 0xde, 0xad, 0xbe, 0xef,
		},
	}
	p.Decode()
	if !bytes.Equal(p.SrcMac, []byte{0x80, 0xee, 0x73, 0x83, 0x58, 0x8f}) {
		// 80:ee:73:83:58:8f
		t.Errorf("Src mac %v", p.SrcMac)
	}
	if !bytes.Equal(p.DestMac, []byte{0x1c, 0x3e, 0x84, 0x0e, 0x04, 0x63}) {
		// 1c:3e:84:0e:04:63
		t.Errorf("Dest mac %v", p.DestMac)
	}
	if len(p.Headers) != 2 {
		t.Fatal("Incorrect number of headers", len(p.Headers))
	}

	ip, ipOk := p.Headers[0].(*IPHdr)
	if !ipOk {
		t.Fatal("First header is not an IPv4 header")
	}
	if ip.Version != 4 {
		t.Error("ip Version", ip.Version)
	}
	if ip.Len() != 68 {
		t.Error("ip length", ip.Len())
	}
	if ip.PayloadLen() != 48 {
		// We chopped the test bytes above.
		t.Error("ip payload length", ip.PayloadLen())
	}
	if ip.SrcAddr() != "192.168.2.111" {
		t.Error("ip src address", ip.SrcAddr())
	}
	if ip.DestAddr() != "192.168.2.20" {
		t.Error("ip dest address", ip.DestAddr())
	}
	if !ip.Fragmented() {
		t.Error("not fragmented")
	}

	// Fragment.
	f, fOk := p.Headers[1].(*Fragment)
	if !fOk {
		t.Fatal("Second part is not an Fragment")
	}
	if f.Length != 48 {
		// We truncated the testbytes above
		t.Error("fragment length", f.Length)
	}
	if f.ProtocolID != syscall.IPPROTO_ICMP {
		t.Error("fragment protocol id", f.ProtocolID)
	}
}

func TestDecodeIPv6HopByHop(t *testing.T) {
	// IPv6 packet with Hop by Hop extended header
	// Generated with samples/hopbyhop.pkt from
	// https://github.com/karknu/rws
	p := &Packet{
		DatalinkType: DLTEN10MB,
		Data: []byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x11, 0x22, 0x33, 0x44,
			0x55, 0x86, 0xdd, 0x60, 0x00, 0x00, 0x00, 0x00, 0x50, 0x00, 0x40,
			0x02, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x00, 0x10, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x11,
			0x00, 0x01, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x07,
			0x00, 0x48, 0xdb, 0x4d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00,
		},
	}
	p.Decode()
	if !bytes.Equal(p.SrcMac, []byte{0, 0x11, 0x22, 0x33, 0x44, 0x55}) {
		t.Errorf("Src mac %v", p.SrcMac)
	}
	if !bytes.Equal(p.DestMac, []byte{0, 0, 0, 0, 0, 0x01}) {
		t.Errorf("Dest mac %v", p.DestMac)
	}
	if len(p.Headers) != 2 {
		t.Fatal("Incorrect number of headers", len(p.Headers))
	}

	ip, ipOk := p.Headers[0].(*IP6Hdr)
	if !ipOk {
		t.Fatal("First header is not an IPv6 header")
	}
	if ip.Version != 6 {
		t.Error("ip Version", ip.Version)
	}
	if ip.Len() != 120 {
		t.Error("ip length", ip.Len())
	}
	if ip.PayloadLen() != 72 {
		t.Error("ip payload length", ip.PayloadLen())
	}
	if ip.SrcAddr() != "200:1000::1" {
		t.Error("ip src address", ip.SrcAddr())
	}
	if ip.DestAddr() != "200:1000::2" {
		t.Error("ip dest address", ip.DestAddr())
	}
	if ip.Fragmented() {
		t.Error("not fragmented")
	}

	// UDP
	udp, udpOk := p.Headers[1].(*UDPHdr)
	if !udpOk {
		t.Fatal("Second part is not UDP")
	}
	if udp.Length != 72 {
		t.Error("UDP length", udp.Length)
	}
	if udp.SrcPort != 7 {
		t.Error("udp srcport", udp.SrcPort)
	}
	if udp.DestPort != 7 {
		t.Error("udp destport", udp.DestPort)
	}

	// Leftover of payload
	if len(p.Payload) != 64 {
		t.Error("Wrong UDP payload length", len(p.Payload))
	}
}
