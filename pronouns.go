package podgo

type Pronoun struct {
	Preferred            bool
	Plural               bool
	Subject              string
	Object               string
	PossessiveDeterminer string
	PossessivePronoun    string
	Reflexive            string
}

func (p *Pronoun) isSubsetOf(other *Pronoun) bool {
	ownValues := []string{p.Subject, p.Object}
	if p.PossessiveDeterminer != "" {
		ownValues = append(ownValues, p.PossessiveDeterminer)
	}
	if p.PossessivePronoun != "" {
		ownValues = append(ownValues, p.PossessivePronoun)
	}
	if p.Reflexive != "" {
		ownValues = append(ownValues, p.Reflexive)
	}

	otherValues := []string{other.Subject, other.Object}
	if other.PossessiveDeterminer != "" {
		otherValues = append(otherValues, other.PossessiveDeterminer)
	}
	if other.PossessivePronoun != "" {
		otherValues = append(otherValues, other.PossessivePronoun)
	}
	if other.Reflexive != "" {
		otherValues = append(otherValues, other.Reflexive)
	}

	if len(ownValues) > len(otherValues) {
		return false
	}

	for i, v := range ownValues {
		if v != otherValues[i] {
			return false
		}
	}

	return true
}

func condense(p1, p2 Pronoun) []Pronoun {
	if p1.isSubsetOf(&p2) {
		if p1.Preferred {
			(&p2).Preferred = true
		}
		if p1.Plural {
			(&p2).Plural = true
		}

		return []Pronoun{p2}
	} else if p2.isSubsetOf(&p1) {
		if p2.Preferred {
			(&p1).Preferred = true
		}
		if p2.Plural {
			(&p1).Plural = true
		}

		return []Pronoun{p1}
	} else {
		return []Pronoun{p1, p2}
	}
}

type Pronouns struct {
	Any    bool
	None   bool
	Accept []Pronoun
}

func (p *Pronouns) condense() {
	final := []Pronoun{}

	for i, pa := range p.Accept {
		for j, pb := range final {
			if i == j {
				continue
			}

			condensed := condense(pa, pb)
			if len(condensed) == 1 {
				final[j] = condensed[0]
				goto nextOuter
			}
		}

		final = append(final, pa)

	nextOuter:
	}

	p.Accept = final
}

func (p *Pronouns) Preferred() *Pronoun {
	if p.None {
		return nil
	}

	if len(p.Accept) == 0 {
		return nil
	}

	preferred := []*Pronoun{}
	for i := range p.Accept {
		if p.Accept[i].Preferred {
			preferred = append(preferred, &p.Accept[i])
		}
	}

	if len(preferred) == 0 {
		return &p.Accept[0]
	}

	return preferred[0]
}
