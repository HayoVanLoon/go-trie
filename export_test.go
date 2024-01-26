package trie

func New2(cs map[rune]Trie, values []string) Trie {
	t := Trie{
		children: cs,
		valueMap: make(map[string]bool),
		values:   values,
	}
	for _, v := range values {
		t.valueMap[v] = true
	}
	return t
}
