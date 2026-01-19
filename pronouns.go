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

type Pronouns struct {
	Any    bool
	None   bool
	Accept []Pronoun
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
