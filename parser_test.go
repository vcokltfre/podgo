package podgo

import "testing"

func TestParserValidAny(t *testing.T) {
	record := "*"
	pronouns, err := parsePronounsRecords([]string{record}, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !pronouns.Any {
		t.Fatalf("Expected Any to be true")
	}
	if pronouns.None {
		t.Fatalf("Expected None to be false")
	}
	if len(pronouns.Accept) != 1 {
		t.Fatalf("Expected 1 accepted pronouns, got %d", len(pronouns.Accept))
	}
}

func TestParserInvalidAny(t *testing.T) {
	records := []string{
		"*;extra",
		"*/them/their/theirs/themself",
	}
	for _, record := range records {
		_, err := parsePronounsRecords([]string{record}, true)
		if err == nil {
			t.Fatalf("Expected error for record: %s", record)
		}
	}
}

func TestParserValidNone(t *testing.T) {
	record := "!"
	pronouns, err := parsePronounsRecords([]string{record}, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if pronouns.Any {
		t.Fatalf("Expected Any to be false")
	}
	if !pronouns.None {
		t.Fatalf("Expected None to be true")
	}
	if len(pronouns.Accept) != 0 {
		t.Fatalf("Expected 0 accepted pronouns, got %d", len(pronouns.Accept))
	}
}

func TestParserInvalidNone(t *testing.T) {
	records := []string{
		"!;extra",
		"!/them/their/theirs/themself",
	}
	for _, record := range records {
		_, err := parsePronounsRecords([]string{record}, true)
		if err == nil {
			t.Fatalf("Expected error for record: %s", record)
		}
	}
}

func TestParserInvalidPronounSpec(t *testing.T) {
	record := "they/them;tag/invalid"
	_, err := parsePronounsRecords([]string{record}, true)
	if err == nil {
		t.Fatalf("Expected error for invalid pronoun spec")
	}
}

func TestParserValidBasicPronouns(t *testing.T) {
	record := "they/them"
	pronouns, err := parsePronounsRecords([]string{record}, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if pronouns.Any {
		t.Fatalf("Expected Any to be false")
	}
	if pronouns.None {
		t.Fatalf("Expected None to be false")
	}
	if len(pronouns.Accept) != 1 {
		t.Fatalf("Expected 1 accepted pronouns, got %d", len(pronouns.Accept))
	}

	p := pronouns.Accept[0]
	if p.Subject != "they" || p.Object != "them" {
		t.Fatalf("Unexpected pronouns values: %+v", p)
	}
}

func TestParserValidFullPronouns(t *testing.T) {
	record := "she/her/her/hers/herself;preferred"
	pronouns, err := parsePronounsRecords([]string{record}, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if pronouns.Any {
		t.Fatalf("Expected Any to be false")
	}
	if pronouns.None {
		t.Fatalf("Expected None to be false")
	}
	if len(pronouns.Accept) != 1 {
		t.Fatalf("Expected 1 accepted pronouns, got %d", len(pronouns.Accept))
	}

	p := pronouns.Accept[0]
	if p.Subject != "she" || p.Object != "her" || p.PossessiveDeterminer != "her" ||
		p.PossessivePronoun != "hers" || p.Reflexive != "herself" {
		t.Fatalf("Unexpected pronouns values: %+v", p)
	}

	if !p.Preferred {
		t.Fatalf("Expected Preferred to be true")
	}
}

func TestParserConvertsItPronouns(t *testing.T) {
	record := "it/its"
	pronouns, err := parsePronounsRecords([]string{record}, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if pronouns.Any {
		t.Fatalf("Expected Any to be false")
	}
	if pronouns.None {
		t.Fatalf("Expected None to be false")
	}
	if len(pronouns.Accept) != 1 {
		t.Fatalf("Expected 1 accepted pronouns, got %d", len(pronouns.Accept))
	}

	p := pronouns.Accept[0]
	if p.Subject != "it" || p.Object != "it" || p.PossessiveDeterminer != "its" ||
		p.PossessivePronoun != "its" || p.Reflexive != "itself" {
		t.Fatalf("Unexpected pronouns values: %+v", p)
	}
}

func TestParserMultipleRecords(t *testing.T) {
	records := []string{
		"she/her/her/hers/herself;preferred",
		"they/them",
		"*",
	}
	pronouns, err := parsePronounsRecords(records, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !pronouns.Any {
		t.Fatalf("Expected Any to be true")
	}
	if pronouns.None {
		t.Fatalf("Expected None to be false")
	}
	if len(pronouns.Accept) != 2 {
		t.Fatalf("Expected 2 accepted pronouns, got %d", len(pronouns.Accept))
	}
}

func TestParserIgnoresInvalidRecordsInNonStrictMode(t *testing.T) {
	records := []string{
		"she/her/her/hers/herself;preferred",
		"they/them",
		"*/extra",
	}
	pronouns, err := parsePronounsRecords(records, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if pronouns.Any {
		t.Fatalf("Expected Any to be false")
	}
	if pronouns.None {
		t.Fatalf("Expected None to be false")
	}
	if len(pronouns.Accept) != 2 {
		t.Fatalf("Expected 2 accepted pronouns, got %d", len(pronouns.Accept))
	}
}

func TestParserUsesAllDefault(t *testing.T) {
	records := []string{"*"}
	pronouns, err := parsePronounsRecords(records, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !pronouns.Any {
		t.Fatalf("Expected Any to be true")
	}
	if pronouns.None {
		t.Fatalf("Expected None to be false")
	}
	if len(pronouns.Accept) != 1 {
		t.Fatalf("Expected 1 accepted pronouns, got %d", len(pronouns.Accept))
	}

	p := pronouns.Accept[0]
	if p.Subject != "they" || p.Object != "them" || p.PossessiveDeterminer != "their" ||
		p.PossessivePronoun != "theirs" || p.Reflexive != "themself" {
		t.Fatalf("Unexpected pronouns values: %+v", p)
	}
}
