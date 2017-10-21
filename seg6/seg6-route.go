package main

import (
	"fmt"
	"net"
	"syscall"

	"github.com/coreswitch/netlink"
	"github.com/coreswitch/netlink/nl"
)

func main() {
	do_f1()
}

/*** TEST PROCEDURE ***

ip route add 10.0.0.22/32 nexthop encap seg6 mode encap segs fc00:a000::1 dev lo
> 10.0.0.22  encap seg6 mode encap segs 1 [ fc00:a000::1 ] dev lo scope link
ip route add 10.0.0.22/32 encap seg6 mode encap segs fc00:a000::1 dev lo
> 10.0.0.22  encap seg6 mode encap segs 1 [ fc00:a000::1 ] dev lo scope link
*/

func do_f1() {
	//link, err := netlink.LinkByName("lo")
	link, err := netlink.LinkByName("enp0s3")
	if err != nil {
		fmt.Println("Error: netlink.LinkByName lo")
		return
	}
	if err := netlink.LinkSetUp(link); err != nil {
		fmt.Println("Error: LinkSetUp(link)")
		return
	}
	dst := &net.IPNet{
		IP:   net.IPv4(10, 0, 0, 22),
		Mask: net.CIDRMask(32, 32),
	}

	var sl []net.IP
	sl = append(sl, net.ParseIP("::"))
	sl = append(sl, net.ParseIP("fc00:a000::2"))
	sl = append(sl, net.ParseIP("fc00:a000::1"))
	//e := &netlink.SEG6Encap{ Mode: nl.SEG6_IPTUN_MODE_ENCAP }
	e := &netlink.SEG6Encap{ Mode: nl.SEG6_IPTUN_MODE_INLINE }
	e.Srh.Segments = sl
	fmt.Printf("%s\n", e.String())
	flags := int(0)
	route := netlink.Route{
		LinkIndex: link.Attrs().Index,
		Dst: dst,
		Flags: flags,
		Type: syscall.RTN_UNICAST,
		Encap: e,
	}
	if err := netlink.RouteAdd(&route); err != nil {
		fmt.Println("Error: RouteAdd(&route)")
		return
	}

	routes, err := netlink.RouteList(link, netlink.FAMILY_V4)
	for index, route := range routes {
		fmt.Printf("index:%d | %s\n", index, route)
	}

}
