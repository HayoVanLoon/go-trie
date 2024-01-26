// Package trie implements a simple tree.
package trie

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Trie struct {
	children map[rune]Trie
	valueMap map[string]bool
	values   []string
}

func (t Trie) String() string {
	b := new(strings.Builder)
	for s := range t.valueMap {
		b.WriteString(s + ",")
	}
	for k, v := range t.children {
		_, _ = fmt.Fprintf(b, "%s:%v", string(k), v)
	}
	return fmt.Sprintf("(%s)", b.String())
}

func Add(t Trie, s string) Trie {
	return add(t, s, 0)
}

func add(t Trie, s string, depth int) Trie {
	if s == "" {
		return t
	}
	if _, ok := t.valueMap[s]; !ok {
		t.values = append(t.values, s)
		t.valueMap[s] = true
	} else {
		return t
	}
	if depth == len(s) {
		return t
	}
	r, n := utf8.DecodeRuneInString(s[depth:])
	if r == utf8.RuneError {
		panic("could not decode rune in " + s)
	}
	v, ok := t.children[r]
	if !ok {
		v = New()
	}
	v = add(v, s, depth+n)
	t.children[r] = add(v, s, depth+n)
	return t
}

func Get(t Trie, s string) []string {
	return get(t, s, 0)
}

func get(t Trie, s string, depth int) []string {
	if len(s) == depth {
		if len(t.values) == 0 {
			return nil
		}
		cp := make([]string, len(t.values))
		copy(cp, t.values)
		return cp
	}
	r, n := utf8.DecodeRuneInString(s[depth:])
	if r == utf8.RuneError {
		panic("could not decode rune in " + s)
	}
	v, ok := t.children[r]
	if !ok {
		return nil
	}
	return get(v, s, depth+n)
}

func Has(t Trie, s string) bool {
	return has(t, s, 0)
}

func has(t Trie, s string, depth int) bool {
	if len(s) == depth {
		return len(t.values) > 0
	}
	r, n := utf8.DecodeRuneInString(s[depth:])
	if r == utf8.RuneError {
		panic("could not decode rune in " + s)
	}
	v, ok := t.children[r]
	return ok && has(v, s, depth+n)
}

// New creates a new, empty trie.
func New() Trie {
	return Trie{children: make(map[rune]Trie), valueMap: make(map[string]bool)}
}
