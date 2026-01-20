package podgo

import "testing"

func TestBasicPronounSubset(t *testing.T) {
	p1 := &Pronoun{
		Subject:              "they",
		Object:               "them",
		PossessiveDeterminer: "their",
	}
	p2 := &Pronoun{
		Subject:              "they",
		Object:               "them",
		PossessiveDeterminer: "their",
		PossessivePronoun:    "theirs",
		Reflexive:            "themself",
	}

	if !p1.isSubsetOf(p2) {
		t.Errorf("Expected p1 to be subset of p2")
	}

	if p2.isSubsetOf(p1) {
		t.Errorf("Expected p2 to not be subset of p1")
	}
}

func TestIdenticalPronounSubset(t *testing.T) {
	p1 := &Pronoun{
		Subject:              "she",
		Object:               "her",
		PossessiveDeterminer: "her",
		PossessivePronoun:    "hers",
		Reflexive:            "herself",
	}
	p2 := &Pronoun{
		Subject:              "she",
		Object:               "her",
		PossessiveDeterminer: "her",
		PossessivePronoun:    "hers",
		Reflexive:            "herself",
	}

	if !p1.isSubsetOf(p2) {
		t.Errorf("Expected p1 to be subset of p2")
	}

	if !p2.isSubsetOf(p1) {
		t.Errorf("Expected p2 to be subset of p1")
	}
}

func TestPronounBasicCondense(t *testing.T) {
	pronouns := &Pronouns{
		Accept: []Pronoun{
			{
				Subject:              "they",
				Object:               "them",
				PossessiveDeterminer: "their",
			},
			{
				Subject:              "they",
				Object:               "them",
				PossessiveDeterminer: "their",
				PossessivePronoun:    "theirs",
				Reflexive:            "themself",
			},
		},
	}

	pronouns.condense()

	if len(pronouns.Accept) != 1 {
		t.Fatalf("Expected 1 pronoun after condense, got %d", len(pronouns.Accept))
	}

	expected := Pronoun{
		Subject:              "they",
		Object:               "them",
		PossessiveDeterminer: "their",
		PossessivePronoun:    "theirs",
		Reflexive:            "themself",
	}

	if pronouns.Accept[0] != expected {
		t.Errorf("Expected pronoun %+v, got %+v", expected, pronouns.Accept[0])
	}
}

func TestPreferredCondense(t *testing.T) {
	pronouns := &Pronouns{
		Accept: []Pronoun{
			{
				Preferred:            true,
				Subject:              "they",
				Object:               "them",
				PossessiveDeterminer: "their",
			},
			{
				Subject:              "they",
				Object:               "them",
				PossessiveDeterminer: "their",
				PossessivePronoun:    "theirs",
				Reflexive:            "themself",
			},
		},
	}

	pronouns.condense()

	if len(pronouns.Accept) != 1 {
		t.Fatalf("Expected 1 pronoun after condense, got %d", len(pronouns.Accept))
	}

	expected := Pronoun{
		Preferred:            true,
		Subject:              "they",
		Object:               "them",
		PossessiveDeterminer: "their",
		PossessivePronoun:    "theirs",
		Reflexive:            "themself",
	}

	if pronouns.Accept[0] != expected {
		t.Errorf("Expected pronoun %+v, got %+v", expected, pronouns.Accept[0])
	}
}

func TestReversePreferredCondense(t *testing.T) {
	pronouns := &Pronouns{
		Accept: []Pronoun{
			{
				Subject:              "they",
				Object:               "them",
				PossessiveDeterminer: "their",
			},
			{
				Preferred:            true,
				Subject:              "they",
				Object:               "them",
				PossessiveDeterminer: "their",
				PossessivePronoun:    "theirs",
				Reflexive:            "themself",
			},
		},
	}

	pronouns.condense()

	if len(pronouns.Accept) != 1 {
		t.Fatalf("Expected 1 pronoun after condense, got %d", len(pronouns.Accept))
	}

	expected := Pronoun{
		Preferred:            true,
		Subject:              "they",
		Object:               "them",
		PossessiveDeterminer: "their",
		PossessivePronoun:    "theirs",
		Reflexive:            "themself",
	}

	if pronouns.Accept[0] != expected {
		t.Errorf("Expected pronoun %+v, got %+v", expected, pronouns.Accept[0])
	}
}

func TestCondenseManyPronouns(t *testing.T) {
	pronouns := &Pronouns{
		Accept: []Pronoun{
			{
				Subject:              "they",
				Object:               "them",
				PossessiveDeterminer: "their",
			},
			{
				Subject:              "they",
				Object:               "them",
				PossessiveDeterminer: "their",
				PossessivePronoun:    "theirs",
				Reflexive:            "themself",
			},
			{
				Preferred:            true,
				Subject:              "she",
				Object:               "her",
				PossessiveDeterminer: "her",
			},
			{
				Plural:               true,
				Subject:              "she",
				Object:               "her",
				PossessiveDeterminer: "her",
				PossessivePronoun:    "hers",
				Reflexive:            "herself",
			},
		},
	}

	pronouns.condense()

	if len(pronouns.Accept) != 2 {
		t.Fatalf("Expected 2 pronouns after condense, got %d", len(pronouns.Accept))
	}

	expected1 := Pronoun{
		Subject:              "they",
		Object:               "them",
		PossessiveDeterminer: "their",
		PossessivePronoun:    "theirs",
		Reflexive:            "themself",
	}

	expected2 := Pronoun{
		Preferred:            true,
		Plural:               true,
		Subject:              "she",
		Object:               "her",
		PossessiveDeterminer: "her",
		PossessivePronoun:    "hers",
		Reflexive:            "herself",
	}

	if pronouns.Accept[0] != expected1 && pronouns.Accept[1] != expected1 {
		t.Errorf("Expected one pronoun to be %+v", expected1)
	}

	if pronouns.Accept[0] != expected2 && pronouns.Accept[1] != expected2 {
		t.Errorf("Expected one pronoun to be %+v", expected2)
	}
}
