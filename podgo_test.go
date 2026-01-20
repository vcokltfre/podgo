package podgo

import "testing"

type testCase struct {
	inputs   []string
	expected Pronouns
	fail     bool
	strict   bool
}

func (t *testCase) match(pronouns *Pronouns) bool {
	if t.expected.Any != pronouns.Any {
		return false
	}
	if t.expected.None != pronouns.None {
		return false
	}
	if len(t.expected.Accept) != len(pronouns.Accept) {
		return false
	}
	for i := range t.expected.Accept {
		exp := t.expected.Accept[i]
		got := pronouns.Accept[i]
		if exp.Preferred != got.Preferred ||
			exp.Plural != got.Plural ||
			exp.Subject != got.Subject ||
			exp.Object != got.Object ||
			exp.PossessiveDeterminer != got.PossessiveDeterminer ||
			exp.PossessivePronoun != got.PossessivePronoun ||
			exp.Reflexive != got.Reflexive {
			return false
		}
	}

	if pref := pronouns.Preferred(); pref != nil {
		expPref := t.expected.Preferred()
		if expPref == nil {
			return false
		}
		if expPref.Preferred != pref.Preferred ||
			expPref.Plural != pref.Plural ||
			expPref.Subject != pref.Subject ||
			expPref.Object != pref.Object ||
			expPref.PossessiveDeterminer != pref.PossessiveDeterminer ||
			expPref.PossessivePronoun != pref.PossessivePronoun ||
			expPref.Reflexive != pref.Reflexive {
			return false
		}
	} else if t.expected.Preferred() != nil {
		return false
	}

	return true
}

var cases = []testCase{
	{
		inputs: []string{"they/them"},
		expected: Pronouns{
			Any:  false,
			None: false,
			Accept: []Pronoun{
				{
					Preferred: false,
					Plural:    false,
					Subject:   "they",
					Object:    "them",
				},
			},
		},
		strict: true,
	},
	{
		inputs: []string{"she/her/her/hers/herself;preferred"},
		expected: Pronouns{
			Any:  false,
			None: false,
			Accept: []Pronoun{
				{
					Preferred:            true,
					Plural:               false,
					Subject:              "she",
					Object:               "her",
					PossessiveDeterminer: "her",
					PossessivePronoun:    "hers",
					Reflexive:            "herself",
				},
			},
		},
		strict: true,
	},
	{
		inputs: []string{"it/its"},
		expected: Pronouns{
			Any:  false,
			None: false,
			Accept: []Pronoun{
				{
					Preferred:            false,
					Plural:               false,
					Subject:              "it",
					Object:               "it",
					PossessiveDeterminer: "its",
					PossessivePronoun:    "its",
					Reflexive:            "itself",
				},
			},
		},
		strict: true,
	},
	{
		inputs: []string{"*"},
		expected: Pronouns{
			Any:  true,
			None: false,
			Accept: []Pronoun{
				{
					Preferred:            false,
					Plural:               true,
					Subject:              "they",
					Object:               "them",
					PossessiveDeterminer: "their",
					PossessivePronoun:    "theirs",
					Reflexive:            "themself",
				},
			},
		},
		strict: true,
	},
	{
		inputs: []string{"!"},
		expected: Pronouns{
			Any:    false,
			None:   true,
			Accept: []Pronoun{},
		},
		strict: true,
	},
	{
		inputs: []string{"!*"},
		fail:   true,
		strict: true,
	},
	{
		inputs: []string{"they/them/their/theirs/themself;preferred", "it/its", "invalid;record"},
		expected: Pronouns{
			Any:  false,
			None: false,
			Accept: []Pronoun{
				{
					Preferred:            true,
					Plural:               false,
					Subject:              "they",
					Object:               "them",
					PossessiveDeterminer: "their",
					PossessivePronoun:    "theirs",
					Reflexive:            "themself",
				},
				{
					Preferred:            false,
					Plural:               false,
					Subject:              "it",
					Object:               "it",
					PossessiveDeterminer: "its",
					PossessivePronoun:    "its",
					Reflexive:            "itself",
				},
			},
		},
		strict: false,
	},
	{
		inputs: []string{
			"*",
			"she/her/her/hers/herself",
			"they/them/their/theirs/themself;preferred",
			"he/him/his/his/himself;preferred",
		},
		expected: Pronouns{
			Any:  true,
			None: false,
			Accept: []Pronoun{
				{
					Preferred:            false,
					Plural:               false,
					Subject:              "she",
					Object:               "her",
					PossessiveDeterminer: "her",
					PossessivePronoun:    "hers",
					Reflexive:            "herself",
				},
				{
					Preferred:            true,
					Plural:               false,
					Subject:              "they",
					Object:               "them",
					PossessiveDeterminer: "their",
					PossessivePronoun:    "theirs",
					Reflexive:            "themself",
				},
				{
					Preferred:            true,
					Plural:               false,
					Subject:              "he",
					Object:               "him",
					PossessiveDeterminer: "his",
					PossessivePronoun:    "his",
					Reflexive:            "himself",
				},
			},
		},
		strict: true,
	},
	{
		inputs: []string{
			"they/them/their/theirs/themself",
			"they/them/their/theirs/themself;preferred",
			"she/her;plural",
			"she/her/her;preferred",
		},
		expected: Pronouns{
			Any:  false,
			None: false,
			Accept: []Pronoun{
				{
					Preferred:            true,
					Plural:               false,
					Subject:              "they",
					Object:               "them",
					PossessiveDeterminer: "their",
					PossessivePronoun:    "theirs",
					Reflexive:            "themself",
				},
				{
					Preferred:            true,
					Plural:               true,
					Subject:              "she",
					Object:               "her",
					PossessiveDeterminer: "her",
				},
			},
		},
		strict: true,
	},
}

func TestManyCases(t *testing.T) {
	for i, c := range cases {
		pronouns, err := parsePronounsRecords(c.inputs, c.strict)

		if c.fail {
			if err == nil {
				t.Fatalf("Case %d: Expected failure but got success", i)
			}
			continue
		}

		if err != nil {
			t.Fatalf("Case %d: Unexpected error: %v", i, err)
		}

		if !c.match(pronouns) {
			t.Fatalf("Case %d: Mismatched pronouns.\nExpected: %+v\nGot:      %+v", i, c.expected, pronouns)
		}
	}
}
