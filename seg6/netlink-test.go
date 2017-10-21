package main

import (
	"fmt"
//	"syscall"

	"github.com/coreswitch/netlink"
//	"github.com/coreswitch/netlink/nl"
)

func main() {
	do_RouteList()
	//do_ListRoutes(netlink.FAMILY_V4)
}

/*
// change netlink.deserializeRoute to capital D to use this.
func do_ListRoutes(family int) {
	req := nl.NewNetlinkRequest(syscall.RTM_GETROUTE, syscall.NLM_F_DUMP)
	infmsg := nl.NewIfInfomsg(family)
	req.AddData(infmsg)

	msgs, err := req.Execute(syscall.NETLINK_ROUTE, syscall.RTM_NEWROUTE)
	if err != nil {
		fmt.Println("Error: executing req.")
		return
	}

	//	 type RtMsg struct {
	//		Family   uint8
	//		Dst_len  uint8
	//		Src_len  uint8
	//		Tos      uint8
	//		Table    uint8
	//		Protocol uint8
	//		Scope    uint8
	//		Type     uint8
	//		Flags    uint32
	//	}

	for _, m := range msgs {
//		msg := nl.DeserializeRtMsg(m)
//		if msg.Flags&syscall.RTM_F_CLONED != 0 {
			// Ignore cloned routes
//			continue
//		}
		//if msg.Table != syscall.RT_TABLE_MAIN {
		//    if filter == nil || filter != nil && filterMask&RT_FILTER_TABLE == 0 {
		//        // Ignore non-main tables
		//        continue
		//    }
		//}
		route, err := netlink.deserializeRoute(m)
		if err != nil {
			fmt.Println("Error: deserializeRoute(m)")
			return
		}
		fmt.Printf("%s\n", route)
	}
}
*/
func do_RouteList() {
	link, err := netlink.LinkByName("enp0s3")
	routes, err := netlink.RouteList(link, netlink.FAMILY_V4)
	if err != nil {
		// t.Fatal(err)
		fmt.Println("t.Fatal(err)\n")
		return
	}
	for index, route := range routes {
		fmt.Printf("index:%d | %s\n", index, route)
	}
}
