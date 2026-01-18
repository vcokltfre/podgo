package podgo

import "testing"

func TestBasicRecordParsing(t *testing.T) {
	record := "she/her"
	pronoun, err := parseNormalPronounsRecord(record)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pronoun.Subject != "she" || pronoun.Object != "her" {
		t.Fatalf("Parsed pronouns do not match expected values")
	}
}

func TestInvalidRecordParsing(t *testing.T) {
	record := "invalid/record/with/too/many/parts"
	_, err := parseNormalPronounsRecord(record)
	if err != ErrInvalidPronounsRecord {
		t.Fatalf("Expected ErrInvalidPronounsRecord, got %v", err)
	}
}

func TestRecordParsingWithPossessive(t *testing.T) {
	record := "they/them/their/theirs/themself"
	pronoun, err := parseNormalPronounsRecord(record)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pronoun.Subject != "they" || pronoun.Object != "them" ||
		pronoun.PossessiveDeterminer != "their" ||
		pronoun.PossessivePronoun != "theirs" ||
		pronoun.Reflexive != "themself" {
		t.Fatalf("Parsed pronouns do not match expected values")
	}
}

func TestWildcardRecordParsing(t *testing.T) {
	record := "*"
	pronoun, err := parseWildCardPronounsRecord(record)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pronoun.Subject != "they" || pronoun.Object != "them" ||
		pronoun.PossessiveDeterminer != "their" ||
		pronoun.PossessivePronoun != "theirs" ||
		pronoun.Reflexive != "themself" {
		t.Fatalf("Parsed pronouns do not match expected wildcard values")
	}
}

func TestWildcardRecordParsingWithFallback(t *testing.T) {
	record := "*;she/her"
	pronoun, err := parseWildCardPronounsRecord(record)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pronoun.Subject != "she" || pronoun.Object != "her" {
		t.Fatalf("Parsed pronouns do not match expected fallback values")
	}
}

func TestInvalidWildcardRecordParsing(t *testing.T) {
	record := "*;extra"
	_, err := parseWildCardPronounsRecord(record)
	if err != ErrInvalidPronounsRecord {
		t.Fatalf("Expected ErrInvalidPronounsRecord, got %v", err)
	}
}
