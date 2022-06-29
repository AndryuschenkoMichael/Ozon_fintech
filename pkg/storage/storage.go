package storage

import (
	"errors"
	"time"
)

const (
	KeyExistError = "key is already exist"
	KeyError      = "key is not exist"
	timeoutError  = "timeout on generating unique value"
)

const (
	timeout = 10 * time.Second
)

type Query int

const (
	GET Query = iota
	POST
)

type Func func() string

type result struct {
	value string
	err   error
}

type request struct {
	queryType Query
	key       string
	response  chan<- result
}

type Storage struct {
	requests chan request
}

func New(f Func) *Storage {
	storage := &Storage{requests: make(chan request)}
	go storage.serve(f)
	return storage
}

func (s *Storage) Get(key string) (string, error) {
	response := make(chan result)
	s.requests <- request{key: key, response: response, queryType: GET}
	res := <-response
	return res.value, res.err
}

func (s *Storage) Post(key string) (string, error) {
	response := make(chan result)
	s.requests <- request{key: key, response: response, queryType: POST}
	res := <-response
	return res.value, res.err
}

func (s *Storage) Close() {
	close(s.requests)
}

func (s *Storage) serve(f Func) {
	storage := make(map[string]string)
	used := make(map[string]string)

	for req := range s.requests {
		if req.queryType == GET {
			value, ok := used[req.key]
			if !ok {
				req.response <- result{value: "", err: errors.New(KeyError)}
			} else {
				req.response <- result{value: value, err: nil}
			}
			continue
		}

		value, ok := storage[req.key]
		if !ok {
			for {
				value := f()
				if _, ok := used[value]; !ok {
					used[value] = req.key
					storage[req.key] = value
					req.response <- result{value: value, err: nil}
					break
				}
			}
		} else {
			req.response <- result{value: value, err: errors.New(KeyExistError)}
		}
	}
}
