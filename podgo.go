package podgo

import (
	"context"
	"net"
)

var resolver = net.Resolver{}

func GetPronouns(records []string, strict bool) (*Pronouns, error) {
	return parsePronounsRecords(records, strict)
}

func GetPronounsResolved(ctx context.Context, domain string, strict bool) (*Pronouns, error) {
	records, err := resolver.LookupTXT(ctx, "pronouns."+domain)
	if err != nil {
		return nil, err
	}

	return parsePronounsRecords(records, strict)
}
