package main

import (
	"fmt"

	dnsv2 "codeberg.org/miekg/dns"
	privatetypes "github.com/DNSControl/demopt/pt"
	privatetypesrdata "github.com/DNSControl/demopt/pt/rdata"
)

func main() {
	// Demonstrate that the private type is registered and can be used with the dnsv2 package.
	fmt.Printf("DEBUG: dnsv2.StringToType[%q] = %d\n", "ALIAS", dnsv2.StringToType["ALIAS"])
	fmt.Printf("DEBUG: dnsv2.TypeToString[%d] = %s\n", privatetypes.TypeALIAS, dnsv2.TypeToString[privatetypes.TypeALIAS])

	// Public types work:
	rd, _ := dnsv2.NewData(dnsv2.TypeMX, "10 mx.miek.nl.")
	fmt.Printf("MX: NewData %s\n", rd.String())
	rr, _ := dnsv2.New("foo.example.com. IN MX 10 mx.miek.nl.")
	fmt.Printf("MX: New %s\n", rr.String())

	// Private types work for dns.New()
	rdalias1 := privatetypesrdata.ALIAS{Target: "foo.com."}
	fmt.Printf("ALIAS: Constructor %s\n", rdalias1.String())
	rr2, _ := dnsv2.New("foo.example.com. IN alias foo.com.")
	fmt.Printf("ALIAS: New %s\n", rr2.String())

	// Private types DON'T work for dns.NewData()
	rdalias2, _ := dnsv2.NewData(privatetypes.TypeALIAS, "foo.com.")
	fmt.Printf("ALIAS: NewData .String() %s\n", rdalias2.String())
	fmt.Printf("ALIAS: NewData (type) %T\n", rdalias2)
}
