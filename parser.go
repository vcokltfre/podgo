package podgo

import (
	"errors"
)

var (
	ErrInvalidPronounsRecord = errors.New("invalid pronouns record")
	ErrUnknownTag            = errors.New("unknown tag in pronouns record")
	ErrNoValidPronouns       = errors.New("no valid pronouns found")
)

func isAnyRecord(record []token) (bool, error) {
	if len(record) == 1 {
		return record[0].typ == tokenTypeAny, nil
	}

	for _, tok := range record {
		if tok.typ == tokenTypeAny {
			return false, ErrContextualAny
		}
	}

	return false, nil
}

func isNoneRecord(record []token) (bool, error) {
	if len(record) == 1 {
		return record[0].typ == tokenTypeNone, nil
	}

	for _, tok := range record {
		if tok.typ == tokenTypeNone {
			return false, ErrContextualNone
		}
	}

	return false, nil
}

func parsePronounSpec(tokens []token) (*Pronoun, error) {
	pronouns := []string{}
	tags := []string{}

	pronounsDone := false

	for _, tok := range tokens {
		if tok.typ == tokenTypeSep {
			pronounsDone = true
			continue
		}

		if tok.typ == tokenTypePronounSep && pronounsDone {
			return nil, ErrInvalidPronounsRecord
		}

		if tok.typ == tokenTypeValue {
			if pronounsDone {
				tags = append(tags, tok.value)
			} else {
				pronouns = append(pronouns, tok.value)
			}
		}
	}

	if len(pronouns) < 2 || len(pronouns) > 5 {
		return nil, ErrInvalidPronounsRecord
	}

	var p *Pronoun
	switch len(pronouns) {
	case 2:
		p = &Pronoun{
			Subject: pronouns[0],
			Object:  pronouns[1],
		}
	case 3:
		p = &Pronoun{
			Subject:              pronouns[0],
			Object:               pronouns[1],
			PossessiveDeterminer: pronouns[2],
		}
	case 4:
		p = &Pronoun{
			Subject:              pronouns[0],
			Object:               pronouns[1],
			PossessiveDeterminer: pronouns[2],
			PossessivePronoun:    pronouns[3],
		}
	case 5:
		p = &Pronoun{
			Subject:              pronouns[0],
			Object:               pronouns[1],
			PossessiveDeterminer: pronouns[2],
			PossessivePronoun:    pronouns[3],
			Reflexive:            pronouns[4],
		}
	}

	if p.Subject == "it" && p.Object == "its" {
		p.Object = "it"
		p.PossessiveDeterminer = "its"
		p.PossessivePronoun = "its"
		p.Reflexive = "itself"
	}

	for _, tag := range tags {
		switch tag {
		case "preferred":
			p.Preferred = true
		case "plural":
			p.Plural = true
		default:
			return nil, ErrUnknownTag
		}
	}

	return p, nil
}

func parsePronounsRecords(records []string, strict bool) (*Pronouns, error) {
	pronouns := &Pronouns{}

	for _, record := range records {
		tokens, err := tokenise(record)
		if err != nil {
			if !strict {
				continue
			}

			return nil, err
		}

		if len(tokens) == 0 {
			continue
		}

		isAny, err := isAnyRecord(tokens)
		if err != nil {
			if !strict {
				continue
			}

			return nil, err
		}

		if isAny {
			pronouns.Any = true
			continue
		}

		isNone, err := isNoneRecord(tokens)
		if err != nil {
			if !strict {
				continue
			}

			return nil, err
		}

		if isNone {
			return &Pronouns{None: true}, nil
		}

		p, err := parsePronounSpec(tokens)
		if err != nil {
			if !strict {
				continue
			}

			return nil, err
		}

		pronouns.Accept = append(pronouns.Accept, *p)
	}

	if pronouns.Any && len(pronouns.Accept) == 0 {
		pronouns.Accept = []Pronoun{
			{
				Subject:              "they",
				Object:               "them",
				PossessiveDeterminer: "their",
				PossessivePronoun:    "theirs",
				Reflexive:            "themself",
				Plural:               true,
			},
		}
	}

	if len(pronouns.Accept) == 0 && !pronouns.Any && !pronouns.None {
		return nil, ErrNoValidPronouns
	}

	pronouns.condense()

	return pronouns, nil
}
