package validator

type EvaluableEntity interface {
	Validate() error
}

type Validator interface {
	Validate(entity EvaluableEntity) error
}

type CustomValidator struct {
}

var _ Validator = (*CustomValidator)(nil)

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{}
}

func (v *CustomValidator) Validate(entity EvaluableEntity) error {
	return entity.Validate()
}
