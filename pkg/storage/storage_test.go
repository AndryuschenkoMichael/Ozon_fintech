package storage

import (
	strgen "Ozon_fintech/pkg/string_generator"
	"testing"
)

func TestStorage_Get(t *testing.T) {
	tests := []struct {
		name     string
		f        Func
		keysPost []string
	}{
		{
			name:     "Test_1",
			f:        strgen.NewRandomStringGenerator(10, []rune{'a', 'b', 'c'}).GenerateString,
			keysPost: []string{"ozon", "fintech", "selection"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.f)
			var values []struct {
				key   string
				value string
			}

			for _, key := range tt.keysPost {
				got, _ := s.Post(key)
				values = append(values, struct {
					key   string
					value string
				}{key: key, value: got})
			}

			for _, value := range values {
				full, err := s.Get(value.value)
				if err != nil {
					t.Errorf("Get() = %v, %v; want err = nil", full, err)
				}

				if full != value.key {
					t.Errorf("Get() = %v, %v; want got = %v", full, err, value.key)
				}
			}

			full, err := s.Get("aaa")
			if err == nil || err.Error() != KeyError {
				t.Errorf("Get() = %v, %v; want err.Error() = %v", full, err, KeyError)
			}
		})
	}
}

func TestStorage_Post(t *testing.T) {
	tests := []struct {
		name string
		f    Func
		keys []string
	}{
		{
			name: "Test_1",
			f:    strgen.NewRandomStringGenerator(10, []rune{'a', 'b', 'c'}).GenerateString,
			keys: []string{"ozon", "fintech", "fintech", "selection"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.f)
			usedKey := make(map[string]struct{})
			usedValue := make(map[string]struct{})

			for _, key := range tt.keys {
				got, err := s.Post(key)
				if _, ok := usedKey[key]; ok {
					if err == nil || err.Error() != KeyExistError {
						t.Errorf("Post() = %v, %v; want err.Error() = %v", got, err, KeyExistError)
					}
				} else {
					if err != nil {
						t.Errorf("Post() = %v, %v; want err = nil", got, err)
					}

					if _, ok := usedValue[got]; ok {
						t.Errorf("Post() = %v, %v; want value must be unique", got, err)
					}
				}

				usedKey[key] = struct{}{}
				usedValue[got] = struct{}{}
			}
		})
	}
}
