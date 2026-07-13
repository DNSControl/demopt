package privatetypesrdata

import (
	"fmt"
	"strings"

	dnsv2 "codeberg.org/miekg/dns"
)

type ALIAS struct {
	Target string
}

func (rd ALIAS) Len() int {
	return len(rd.String())
}

func (rd ALIAS) String() string {
	parts := make([]string, 0, 1)
	parts = append(parts, rd.Target)
	return strings.Join(parts, " ")
}

func MakeALIAS(origin string, _ map[string]string, args ...any) (dnsv2.RDATA, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("ALIAS expects 1 arguments, got %d: %+v", len(args), args)
	}
	return ALIAS{
		Target: args[0].(string),
	}, nil
}
