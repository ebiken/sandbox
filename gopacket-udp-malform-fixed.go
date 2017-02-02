package main

import (
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
    "log"
    "net"
    "time"
//    "fmt"
)

var (
    device          string = "lo"
    snapshot_len    int32  = 1024
    promiscuous     bool   = false
    err             error
    timeout         time.Duration = 30 * time.Second
    handle          *pcap.Handle
    buffer          gopacket.SerializeBuffer
    options         gopacket.SerializeOptions
    // layer options
    srcMac, dstMac  net.HardwareAddr
    srcIp, dstIp    net.IP
    srcPort, dstPort    int
    count           int
)

func main() {
    // Set other options (false or true)
    options.FixLengths = true
	// TODO: Packet malformed when ComputeChecksum is true !?
    options.ComputeChecksums = true

    // Open device
    handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
    if err != nil { log.Fatal(err) }
    defer handle.Close()

	// Set parameters to build each Layer
	srcPort = 9999
	dstPort = 2152
	srcIp  = net.ParseIP("10.0.0.2")
	dstIp  = net.ParseIP("127.0.0.1")
	srcMac, _ = net.ParseMAC("02:00:00:00:00:01")
	dstMac, _ = net.ParseMAC("06:00:00:00:00:01")

    rawBytes := make([]byte, 16)

    udpLayer := &layers.UDP{
        SrcPort  : layers.UDPPort(srcPort),
        DstPort  : layers.UDPPort(dstPort),
    }
	ipv4Layer := &layers.IPv4{
        Version    : 4, //uint8
        IHL        : 5, //uint8
        TOS        : 0, //uint8
        Id         : 0, //uint16
        Flags      : 0, //IPv4Flag
        FragOffset : 0, //uint16
        TTL        : 255, //uint8
        Protocol   : layers.IPProtocolUDP, //IPProtocol UDP(17)
        SrcIP: srcIp,
        DstIP: dstIp,
    }
    ethernetLayer := &layers.Ethernet{
        SrcMAC: srcMac,
        DstMAC: dstMac,
        EthernetType: 0x800,
    }

	// set ipv4Layer to be used for UDP checksum calculation.
	udpLayer.SetNetworkLayerForChecksum(ipv4Layer)
	//fmt.Println(ethernetLayer, ipv4Layer, udpLayer, gtpLayer)

	// Now actually send packet
    buffer = gopacket.NewSerializeBuffer()
    err = gopacket.SerializeLayers(buffer, options,
        ethernetLayer,
		ipv4Layer,
		udpLayer,
        gopacket.Payload(rawBytes),
    )
    err = handle.WritePacketData( buffer.Bytes() )
    if err != nil {
        log.Fatal(err)
    }
    return
}
