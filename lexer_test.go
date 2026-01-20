package podgo

import "testing"

func TestLexerEmptyRecord(t *testing.T) {
	record := ""
	tokens, err := tokenise(record)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(tokens) != 0 {
		t.Fatalf("Expected 0 tokens, got %d", len(tokens))
	}
}

func TestLexerCommentOnlyRecord(t *testing.T) {
	record := "# This is a comment"
	tokens, err := tokenise(record)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(tokens) != 0 {
		t.Fatalf("Expected 0 tokens, got %d", len(tokens))
	}
}

func TestLexerAny(t *testing.T) {
	record := "*"
	tokens, err := tokenise(record)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("Expected 1 token, got %d", len(tokens))
	}
	if tokens[0].typ != tokenTypeAny {
		t.Fatalf("Expected token type %d, got %d", tokenTypeAny, tokens[0].typ)
	}
}

func TestLexerNone(t *testing.T) {
	record := "!"
	tokens, err := tokenise(record)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("Expected 1 token, got %d", len(tokens))
	}
	if tokens[0].typ != tokenTypeNone {
		t.Fatalf("Expected token type %d, got %d", tokenTypeNone, tokens[0].typ)
	}
}

func TestLexerValuesAndSeparators(t *testing.T) {
	record := "they/them/their/theirs/themself;preferred"
	tokens, err := tokenise(record)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedTypes := []tokenType{
		tokenTypeValue,
		tokenTypePronounSep,
		tokenTypeValue,
		tokenTypePronounSep,
		tokenTypeValue,
		tokenTypePronounSep,
		tokenTypeValue,
		tokenTypePronounSep,
		tokenTypeValue,
		tokenTypeSep,
		tokenTypeValue,
	}
	expectedValues := []string{
		"they",
		"/",
		"them",
		"/",
		"their",
		"/",
		"theirs",
		"/",
		"themself",
		";",
		"preferred",
	}
	if len(tokens) != len(expectedTypes) {
		t.Fatalf("Expected %d tokens, got %d", len(expectedTypes), len(tokens))
	}
	for i, expectedType := range expectedTypes {
		if tokens[i].typ != expectedType {
			t.Fatalf("At token %d, expected type %d, got %d", i, expectedType, tokens[i].typ)
		}

		if tokens[i].value != expectedValues[i] {
			t.Fatalf("At token %d, expected value '%s', got '%s'", i, expectedValues[i], tokens[i].value)
		}
	}
}

func TestLexerValuesAndSeparatorsWithSpacesAndComment(t *testing.T) {
	record := " they/ them / their   / the irs   / themself ; preferred # comment"
	tokens, err := tokenise(record)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedTypes := []tokenType{
		tokenTypeValue,
		tokenTypePronounSep,
		tokenTypeValue,
		tokenTypePronounSep,
		tokenTypeValue,
		tokenTypePronounSep,
		tokenTypeValue,
		tokenTypePronounSep,
		tokenTypeValue,
		tokenTypeSep,
		tokenTypeValue,
	}
	expectedValues := []string{
		"they",
		"/",
		"them",
		"/",
		"their",
		"/",
		"the irs",
		"/",
		"themself",
		";",
		"preferred",
	}

	if len(tokens) != len(expectedTypes) {
		t.Fatalf("Expected %d tokens, got %d", len(expectedTypes), len(tokens))
	}

	for i, expectedType := range expectedTypes {
		if tokens[i].typ != expectedType {
			t.Fatalf("At token %d, expected type %d, got %d", i, expectedType, tokens[i].typ)
		}

		if tokens[i].value != expectedValues[i] {
			t.Fatalf("At token %d, expected value '%s', got '%s'", i, expectedValues[i], tokens[i].value)
		}
	}
}

func TestLexerInvalidCharacter(t *testing.T) {
	record := "they/them/their/theirs/themself!;preferred"
	_, err := tokenise(record)
	if err == nil {
		t.Fatalf("Expected error for invalid character, got nil")
	}
}

func TestLexerMixedRecord(t *testing.T) {
	record := "!*;they/  them\t;preferred"
	tokens, err := tokenise(record)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedTypes := []tokenType{
		tokenTypeNone,
		tokenTypeAny,
		tokenTypeSep,
		tokenTypeValue,
		tokenTypePronounSep,
		tokenTypeValue,
		tokenTypeSep,
		tokenTypeValue,
	}
	expectedValues := []string{
		"!",
		"*",
		";",
		"they",
		"/",
		"them",
		";",
		"preferred",
	}

	if len(tokens) != len(expectedTypes) {
		t.Fatalf("Expected %d tokens, got %d", len(expectedTypes), len(tokens))
	}

	for i, expectedType := range expectedTypes {
		if tokens[i].typ != expectedType {
			t.Fatalf("At token %d, expected type %d, got %d", i, expectedType, tokens[i].typ)
		}

		if tokens[i].value != expectedValues[i] {
			t.Fatalf("At token %d, expected value '%s', got '%s'", i, expectedValues[i], tokens[i].value)
		}
	}
}

func TestLexerCondensesSeparators(t *testing.T) {
	record := "a;;;b"
	tokens, err := tokenise(record)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedTypes := []tokenType{
		tokenTypeValue,
		tokenTypeSep,
		tokenTypeValue,
	}
	expectedValues := []string{
		"a",
		";",
		"b",
	}

	if len(tokens) != len(expectedTypes) {
		t.Fatalf("Expected %d tokens, got %d", len(expectedTypes), len(tokens))
	}

	for i, expectedType := range expectedTypes {
		if tokens[i].typ != expectedType {
			t.Fatalf("At token %d, expected type %d, got %d", i, expectedType, tokens[i].typ)
		}

		if tokens[i].value != expectedValues[i] {
			t.Fatalf("At token %d, expected value '%s', got '%s'", i, expectedValues[i], tokens[i].value)
		}
	}
}
