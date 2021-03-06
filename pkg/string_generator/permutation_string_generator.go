package string_generator

import "bytes"

// result is container for string result
type result struct {
	value string
}

// request need to get result from server
type request struct {
	response chan<- result
}

// PermutationStringGenerator implement interface GeneratorStr
type PermutationStringGenerator struct {
	length   int
	alphabet []rune
	requests chan request
}

func (p *PermutationStringGenerator) GenerateString() string {
	response := make(chan result)
	p.requests <- request{response: response}
	res := <-response
	return res.value
}

func NewPermutationStringGenerator(length int, alphabet []rune) *PermutationStringGenerator {
	gen := &PermutationStringGenerator{
		length:   length,
		alphabet: alphabet,
		requests: make(chan request),
	}

	go gen.serve()

	return gen
}

// serve for serving requests
func (p *PermutationStringGenerator) serve() {
	perm := make([]int, p.length)
	for req := range p.requests {
		p.deliver(req.response, perm)
	}
}

// deliver get next permutation
func (p *PermutationStringGenerator) deliver(response chan<- result, perm []int) {
	i := p.length - 1

	for ; 0 <= i && perm[i] == len(p.alphabet)-1; i-- {
	}

	if i < 0 {
		for j := range perm {
			perm[j] = 0
		}
	} else {
		perm[i]++
		for j := i + 1; j < p.length; j++ {
			perm[j] = 0
		}
	}

	var buf bytes.Buffer
	for _, val := range perm {
		buf.WriteRune(p.alphabet[val])
	}

	response <- result{value: buf.String()}
}

func (p *PermutationStringGenerator) Close() {
	close(p.requests)
}
