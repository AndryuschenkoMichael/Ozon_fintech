package string_generator

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestPermutationStringGenerator_GenerateString(t *testing.T) {
	type fields struct {
		length   int
		alphabet []rune
		requests chan request
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "Test 1", fields: fields{
			length:   5,
			alphabet: generateAlphabet(),
		}},

		{name: "Test 2", fields: fields{
			length:   10,
			alphabet: generateAlphabet(),
		}},

		{name: "Test 3", fields: fields{
			length:   10,
			alphabet: []rune{'a', 'b', 'c'},
		}},

		{name: "Test 4", fields: fields{
			length:   10,
			alphabet: []rune{'a'},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPermutationStringGenerator(tt.fields.length, tt.fields.alphabet)

			got := p.GenerateString()
			if utf8.RuneCountInString(got) != p.length {
				t.Errorf("GenerateString() = %v, want length(%v) = %v", got, got, p.length)
			}

			mp := make(map[rune]struct{})

			for _, value := range p.alphabet {
				mp[value] = struct{}{}
			}

			for _, value := range got {
				if _, ok := mp[value]; !ok {
					t.Errorf("GenerateString() = %v, want symbols in %v must be in alphabet", got, got)
				}
			}
		})
	}
}

func TestRandomStringGenerator_GenerateString(t *testing.T) {
	type fields struct {
		length   int
		alphabet []rune
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "Test 1", fields: fields{
			length:   5,
			alphabet: generateAlphabet(),
		}},

		{name: "Test 2", fields: fields{
			length:   10,
			alphabet: generateAlphabet(),
		}},

		{name: "Test 3", fields: fields{
			length:   10,
			alphabet: []rune{'a', 'b', 'c'},
		}},

		{name: "Test 4", fields: fields{
			length:   10,
			alphabet: []rune{'a'},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RandomStringGenerator{
				length:   tt.fields.length,
				alphabet: tt.fields.alphabet,
			}
			got := r.GenerateString()
			if utf8.RuneCountInString(got) != r.length {
				t.Errorf("GenerateString() = %v, want length(%v) = %v", got, got, r.length)
			}

			mp := make(map[rune]struct{})

			for _, value := range r.alphabet {
				mp[value] = struct{}{}
			}

			for _, value := range got {
				if _, ok := mp[value]; !ok {
					t.Errorf("GenerateString() = %v, want symbols in %v must be in alphabet", got, got)
				}
			}
		})
	}
}

func Test_generateAlphabet(t *testing.T) {
	tests := []struct {
		name string
		want []rune
	}{
		{name: "Test 1", want: []rune{'_',
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
			'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
			'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
			'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateAlphabet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}
