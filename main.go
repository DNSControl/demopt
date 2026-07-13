package main

import (
	"fmt"

	dnsv2 "codeberg.org/miekg/dns"
	privatetypes "github.com/DNSControl/demopt/pt"
	privatetypesrdata "github.com/DNSControl/demopt/pt/rdata"
)

func main() {
	fmt.Printf("DEBUG: dnsv2.StringToType[%q] = %d\n", "ALIAS", dnsv2.StringToType["ALIAS"])
	fmt.Printf("DEBUG: dnsv2.TypeToString[%d] = %s\n", privatetypes.TypeALIAS, dnsv2.TypeToString[privatetypes.TypeALIAS])

	rd, _ := dnsv2.NewData(dnsv2.TypeMX, "10 mx.miek.nl.")
	fmt.Printf("MX: %s\n", rd.String())

	rdalias1 := privatetypesrdata.ALIAS{Target: "foo.com."}
	fmt.Printf("ALIAS1: %s\n", rdalias1.String())

	rdalias2, _ := dnsv2.NewData(privatetypes.TypeALIAS, "foo.com.")
	fmt.Printf("ALIAS2: .String() %s\n", rdalias2.String())
	fmt.Printf("ALIAS2: (type) %T\n", rdalias2)
}
