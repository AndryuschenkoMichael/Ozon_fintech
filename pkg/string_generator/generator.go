// Package string_generator generate random string
package string_generator

// GeneratorStr is type of interface witch generate random string
type GeneratorStr interface {
	GenerateString() string
}

// StringGenerator needs for SOLID
type StringGenerator struct {
	GeneratorStr
}

// NewStringGeneratorRandom implement StringGenerator
func NewStringGeneratorRandom(length int) *StringGenerator {
	return &StringGenerator{
		GeneratorStr: NewRandomStringGenerator(length, generateAlphabet()),
	}
}

// NewStringGeneratorPermutation implement StringGenerator
func NewStringGeneratorPermutation(length int) *StringGenerator {
	return &StringGenerator{
		GeneratorStr: NewPermutationStringGenerator(length, generateAlphabet()),
	}
}

func generateAlphabet() []rune {
	alphabet := []rune{'_'}

	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, i)
	}

	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, i)
	}

	for i := '0'; i <= '9'; i++ {
		alphabet = append(alphabet, i)
	}

	return alphabet
}
