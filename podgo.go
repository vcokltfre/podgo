package podgo

import (
	"net"
)

func dedup(pronouns []Pronoun) []Pronoun {
	seen := make(map[string]struct{})
	result := []Pronoun{}

	for _, p := range pronouns {
		key := p.String()
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, p)
		}
	}

	return result
}

func GetPronouns(domain string, skipParseFails bool) (*Pronouns, error) {
	pronouns := &Pronouns{}

	records, err := net.LookupTXT("primary.pronouns." + domain)
	if err == nil {
		p, err := parsePronounsRecords(records, skipParseFails)
		if err != nil {
			return nil, err
		}

		pronouns = p
	}

	records, err = net.LookupTXT("pronouns." + domain)
	if err != nil {
		return nil, err
	}

	p, err := parsePronounsRecords(records, skipParseFails)
	if err != nil {
		return nil, err
	}

	if p.Any {
		pronouns.Any = true
	}

	pronouns.Accept = append(pronouns.Accept, p.Accept...)

	pronouns.Accept = dedup(pronouns.Accept)

	return pronouns, nil
}
