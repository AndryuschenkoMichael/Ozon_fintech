package string_generator

type GeneratorStr interface {
	GenerateString() string
}

type StringGenerator struct {
	GeneratorStr
}

func NewStringGeneratorRandom(length int) *StringGenerator {
	return &StringGenerator{
		GeneratorStr: NewRandomStringGenerator(length, generateAlphabet()),
	}
}

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
