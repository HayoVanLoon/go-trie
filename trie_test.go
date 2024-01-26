package trie_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/HayoVanLoon/go-trie"
)

func TestAdd(t *testing.T) {
	type args struct {
		t trie.Trie
		s string
	}
	tests := []struct {
		name string
		args args
		want trie.Trie
	}{
		{
			"empty",
			args{trie.New(), "foo"},
			step1(),
		},
		{
			"add partial match",
			args{step1(), "fork"},
			step2(),
		},
		{
			"and now for something completely different",
			args{step2(), "bar"},
			step3(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := trie.Add(tt.args.t, tt.args.s)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		t trie.Trie
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"empty",
			args{trie.New(), "foo"},
			nil,
		},
		{
			"empty with empty input",
			args{trie.New(), ""},
			nil,
		},
		{
			"full",
			args{step3(), "fork"},
			[]string{"fork"},
		},
		{
			"partial",
			args{step3(), "fo"},
			[]string{"foo", "fork"},
		},
		{
			"past end",
			args{step3(), "forks"},
			nil,
		},
		{
			"empty input",
			args{step3(), ""},
			[]string{"foo", "fork", "bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := trie.Get(tt.args.t, tt.args.s)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestHas(t *testing.T) {
	type args struct {
		t trie.Trie
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"empty",
			args{trie.New(), "foo"},
			false,
		},
		{
			"empty with empty input",
			args{trie.New(), ""},
			false,
		},
		{
			"full",
			args{step3(), "fork"},
			true,
		},
		{
			"partial",
			args{step3(), "fo"},
			true,
		},
		{
			"past end",
			args{step3(), "forks"},
			false,
		},
		{
			"empty input",
			args{step3(), ""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := trie.Has(tt.args.t, tt.args.s)
			require.Equal(t, tt.want, actual)
		})
	}
}

func step1() trie.Trie {
	return trie.New2(
		map[rune]trie.Trie{
			'f': trie.New2(
				map[rune]trie.Trie{
					'o': trie.New2(
						map[rune]trie.Trie{
							'o': trie.New2(
								map[rune]trie.Trie{},
								[]string{"foo"},
							),
						},
						[]string{"foo"},
					),
				},
				[]string{"foo"},
			),
		},
		[]string{"foo"},
	)
}

func step2() trie.Trie {
	return trie.New2(
		map[rune]trie.Trie{
			'f': trie.New2(
				map[rune]trie.Trie{
					'o': trie.New2(
						map[rune]trie.Trie{
							'o': trie.New2(
								map[rune]trie.Trie{},
								[]string{"foo"},
							),
							'r': trie.New2(
								map[rune]trie.Trie{
									'k': trie.New2(
										map[rune]trie.Trie{},
										[]string{"fork"},
									),
								},
								[]string{"fork"},
							),
						},
						[]string{"foo", "fork"},
					),
				},
				[]string{"foo", "fork"},
			),
		},
		[]string{"foo", "fork"},
	)
}

func step3() trie.Trie {
	return trie.New2(
		map[rune]trie.Trie{
			'f': trie.New2(
				map[rune]trie.Trie{
					'o': trie.New2(
						map[rune]trie.Trie{
							'o': trie.New2(
								map[rune]trie.Trie{},
								[]string{"foo"},
							),
							'r': trie.New2(
								map[rune]trie.Trie{
									'k': trie.New2(
										map[rune]trie.Trie{},
										[]string{"fork"},
									),
								},
								[]string{"fork"},
							),
						},
						[]string{"foo", "fork"},
					),
				},
				[]string{"foo", "fork"},
			),
			'b': trie.New2(
				map[rune]trie.Trie{
					'a': trie.New2(
						map[rune]trie.Trie{
							'r': trie.New2(
								map[rune]trie.Trie{},
								[]string{"bar"},
							),
						},
						[]string{"bar"},
					),
				},
				[]string{"bar"},
			),
		},
		[]string{"foo", "fork", "bar"},
	)
}
