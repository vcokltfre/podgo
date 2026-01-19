package podgo

import (
	"errors"
	"strings"
)

var (
	ErrContextualAny  = errors.New("any with context is not allowed")
	ErrContextualNone = errors.New("none with context is not allowed")
	ErrUnexpectedChar = errors.New("unexpected character in pronouns record")
)

type tokenType int

const (
	tokenTypeAny tokenType = iota
	tokenTypeNone
	tokenTypeSep
	tokenTypePronounSep
	tokenTypeValue
)

type token struct {
	typ   tokenType
	value string
}

func tokenise(record string) ([]token, error) {
	tokens := []token{}

	current := ""

	for _, c := range strings.ToLower(record) {
		if c == '#' {
			break
		}

		if c == '*' {
			if current != "" {
				return nil, ErrContextualAny
			}

			tokens = append(tokens, token{typ: tokenTypeAny, value: "*"})
			continue
		}

		if c == '!' {
			if current != "" {
				return nil, ErrContextualNone
			}

			tokens = append(tokens, token{typ: tokenTypeNone, value: "!"})
			continue
		}

		if c == ';' || c == '/' {
			if current != "" {
				tokens = append(tokens, token{typ: tokenTypeValue, value: strings.TrimSpace(current)})
				current = ""
			}
			var typ tokenType
			if c == ';' {
				typ = tokenTypeSep
			} else {
				typ = tokenTypePronounSep
			}
			tokens = append(tokens, token{typ: typ, value: string(c)})
			continue
		}

		if c == '\t' || c == '\n' || c == '\r' {
			continue
		}

		// if (c < 'a' || c > 'z') && c != ' ' {
		// 	return nil, ErrUnexpectedChar
		// }

		current += string(c)
	}

	if current != "" {
		tokens = append(tokens, token{typ: tokenTypeValue, value: strings.TrimSpace(current)})
	}

	out := []token{}
	prev := tokenTypeSep

	for _, tok := range tokens {
		if tok.typ == tokenTypeSep && prev == tokenTypeSep {
			continue
		}
		out = append(out, tok)
		prev = tok.typ
	}

	return out, nil
}
