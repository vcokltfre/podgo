package podgo

import (
	"net"
)

func GetPronouns(domain string, fail bool) (*Pronouns, error) {
	pronouns := &Pronouns{}

	records, err := net.LookupTXT("primary.pronouns." + domain)
	if err == nil {
		p, err := parsePronounsRecords(records)
		if err != nil {
			return nil, err
		}

		pronouns = p
	}

	records, err = net.LookupTXT("pronouns." + domain)
	if err != nil {
		return nil, err
	}

	p, err := parsePronounsRecords(records)
	if err != nil {
		return nil, err
	}

	if p.Any {
		pronouns.Any = true
	}

	pronouns.Accept = append(pronouns.Accept, p.Accept...)

	return pronouns, nil
}
