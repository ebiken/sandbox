package main

import (
    "net"
    "fmt"
)

type IPv4Range struct {
    sipStart    net.IP
    dipStart    net.IP
    sipEnd      net.IP
    dipEnd      net.IP
    sip         net.IP
    dip         net.IP
}

func (v IPv4Range) next() {

    for i := 0; i < 4; i++ {
        if v.sip[15-i] >= v.sipEnd[15-i] {
            v.sip[15-i] = v.sipStart[15-i]
        } else {
            v.sip[15-i]++
            return
        }
    }
    for i := 0; i < 4; i++ {
        if v.dip[15-i] >= v.dipEnd[15-i] {
            v.dip[15-i] = v.dipStart[15-i]
        } else {
            v.dip[15-i]++
            return
        }
    }
}

func main() {

    //var xip IPv4Range
    //xip.sipStart = net.ParseIP("1.2.3.4")
    //xip.sip = net.ParseIP("1.2.3.4")
    //xip.sipEnd = net.ParseIP("2.3.4.4")
    //xip.dipStart = net.ParseIP("11.12.13.14")
    //xip.dip = net.ParseIP("11.12.13.14")
    //xip.dipEnd = net.ParseIP("12.13.14.15")

    xip := IPv4Range{
        sipStart: net.ParseIP("1.2.3.4"),
        sip: net.ParseIP("1.2.3.4"),
        sipEnd: net.ParseIP("2.3.4.4"),
        dipStart: net.ParseIP("11.12.13.14"),
        dip: net.ParseIP("11.12.13.14"),
        dipEnd: net.ParseIP("12.13.14.15"),
    }

    for i := 0; i < 30; i++ {
        fmt.Println("sip | dip:", xip.sip, xip.dip)
        xip.next()
    }

    /*
    var sip net.IP
    sip = net.ParseIP("1.2.3.4")
    fmt.Println("sip:", sip)
    fmt.Println("sip[0]:", net.IPv4(10,20,30,40))
    for i := 0; i < 4; i++ {
        fmt.Println("sip[x]: byte", i, sip[12+i])
    }

    xip := [4]byte{11,22,33,44}
    fmt.Println("xip:", xip)
    fmt.Println("xip:", net.IPv4(xip[0],xip[1],xip[2],xip[3]))
    */
}
