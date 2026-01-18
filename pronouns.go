package podgo

import "strings"

type Pronoun struct {
	Subject              string
	Object               string
	PossessiveDeterminer string
	PossessivePronoun    string
	Reflexive            string
}

func (p *Pronoun) String() string {
	parts := []string{p.Subject, p.Object}
	if p.PossessiveDeterminer != "" {
		parts = append(parts, p.PossessiveDeterminer)
	}
	if p.PossessivePronoun != "" {
		parts = append(parts, p.PossessivePronoun)
	}
	if p.Reflexive != "" {
		parts = append(parts, p.Reflexive)
	}
	return strings.Join(parts, "/")
}

type Pronouns struct {
	Any    bool
	Accept []Pronoun
}
