package service

const (
	validationError = "validation error"
)

type Valid struct{}

func NewValid() *Valid {
	return &Valid{}
}

func (v *Valid) ValidateLink(link string) error {
	return nil
}
