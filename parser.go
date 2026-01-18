package podgo

import (
	"errors"
	"strings"
)

var ErrInvalidPronounsRecord = errors.New("invalid pronouns record")

func parseNormalPronounsRecord(record string) (*Pronoun, error) {
	parts := strings.Split(record, "/")
	if len(parts) < 2 || len(parts) > 5 {
		return nil, ErrInvalidPronounsRecord
	}

	pronouns := &Pronoun{}
	switch len(parts) {
	case 2:
		pronouns.Subject = parts[0]
		pronouns.Object = parts[1]
	case 3:
		pronouns.Subject = parts[0]
		pronouns.Object = parts[1]
		pronouns.PossessiveDeterminer = parts[2]
	case 4:
		pronouns.Subject = parts[0]
		pronouns.Object = parts[1]
		pronouns.PossessiveDeterminer = parts[2]
		pronouns.PossessivePronoun = parts[3]
	case 5:
		pronouns.Subject = parts[0]
		pronouns.Object = parts[1]
		pronouns.PossessiveDeterminer = parts[2]
		pronouns.PossessivePronoun = parts[3]
		pronouns.Reflexive = parts[4]
	}

	return pronouns, nil
}

func parseWildCardPronounsRecord(record string) (*Pronoun, error) {
	parts := strings.Split(record, ";")
	if len(parts) > 2 {
		return nil, ErrInvalidPronounsRecord
	}

	if len(parts) == 1 {
		if strings.TrimSpace(parts[0]) != "*" {
			return nil, ErrInvalidPronounsRecord
		}

		return &Pronoun{
			Subject:              "they",
			Object:               "them",
			PossessiveDeterminer: "their",
			PossessivePronoun:    "theirs",
			Reflexive:            "themself",
		}, nil
	}

	return parseNormalPronounsRecord(strings.TrimSpace(parts[1]))
}

func parsePronounsRecords(records []string, skipParseFails bool) (*Pronouns, error) {
	pronouns := &Pronouns{}

	for _, rec := range records {
		if strings.Contains(rec, "*") {
			pronoun, err := parseWildCardPronounsRecord(rec)
			if err != nil {
				if skipParseFails {
					continue
				}

				return nil, err
			}

			pronouns.Any = true
			pronouns.Accept = append(pronouns.Accept, *pronoun)

			continue
		}

		pronoun, err := parseNormalPronounsRecord(rec)
		if err != nil {
			if skipParseFails {
				continue
			}

			return nil, err
		}

		pronouns.Accept = append(pronouns.Accept, *pronoun)
	}

	return pronouns, nil
}
