package podgo

type Pronoun struct {
	Subject              string
	Object               string
	PossessiveDeterminer string
	PossessivePronoun    string
	Reflexive            string
}

type Pronouns struct {
	Any    bool
	Accept []Pronoun
}
