// Usage: $ go run gopacket-parse.go -r /home/ebiroot/tmp/gtpu.trc
package main

import (
	"github.com/google/gopacket/dumpcommand"
    "github.com/google/gopacket/pcap"
	//"github.com/google/gopacket"
    //"github.com/google/gopacket/examples/util"
	//"fmt"
	"log"
	"flag"
)

var fname = flag.String("r", "", "Filename to read from, overrides -i")

func main() {
	flag.Parse()
	var handle *pcap.Handle
	var err error
	if *fname != "" {
		if handle, err = pcap.OpenOffline(*fname); err != nil {
			log.Fatal("PCAP OpenOffline error:", err)
		}
	} else {
		log.Fatal("no PCAP file name specified.")
	}

	dumpcommand.Run(handle)
}
