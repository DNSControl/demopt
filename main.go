package main

import (
	"fmt"

	"codeberg.org/miekg/dns"
	privatetypes "github.com/DNSControl/demopt/pt"
	privatetypesrdata "github.com/DNSControl/demopt/pt/rdata"
)

func main() {
	rd, _ := dns.NewData(dns.TypeMX, "10 mx.miek.nl.")
	fmt.Printf("MX: %s\n", rd.String())

	rdalias1 := privatetypesrdata.ALIAS{Target: "foo.com."}
	fmt.Printf("ALIAS1: %s\n", rdalias1.String())

	rdalias2, _ := dns.NewData(privatetypes.TypeALIAS, "foo.com.")
	fmt.Printf("ALIAS2: %s\n", rdalias2.String())

}
