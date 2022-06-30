package string_generator

import (
	"bytes"
	"math/rand"
	"time"
)

// RandomStringGenerator implement GeneratorStr
type RandomStringGenerator struct {
	length   int
	alphabet []rune
}

func (r *RandomStringGenerator) GenerateString() string {
	rand.Seed(time.Now().Unix())
	var buf bytes.Buffer
	for i := 0; i < r.length; i++ {
		buf.WriteRune(r.alphabet[rand.Intn(len(r.alphabet))])
	}

	return buf.String()
}

func NewRandomStringGenerator(length int, alphabet []rune) *RandomStringGenerator {
	return &RandomStringGenerator{
		length:   length,
		alphabet: alphabet,
	}
}
