package validate_object

import "fmt"

type NumberValidateObject struct {
	value         any
	input         string
	optional      bool
	positive      bool
	differentZero bool
}

func NewNumberValueObject(value any, input string) *NumberValidateObject {
	return &NumberValidateObject{
		value:         value,
		optional:      false,
		positive:      false,
		differentZero: false,
		input:         input,
	}
}

func (n *NumberValidateObject) IsOptional() *NumberValidateObject {
	n.optional = true
	return n
}

func (n *NumberValidateObject) IsPositive() *NumberValidateObject {
	n.positive = true
	return n
}

func (n *NumberValidateObject) IsDifferentZero() *NumberValidateObject {
	n.differentZero = true
	return n
}

func (n *NumberValidateObject) Value() int {
	return n.value.(int)
}

func (n *NumberValidateObject) Validate() error {
	if n.optional && n.value == nil {
		return nil
	}

	if n.value == "NaN" {
		n.value = 0
	}

	if _, ok := n.value.(int); !ok {
		return fmt.Errorf("%s must be an integer", n.input)
	}

	if n.optional && n.value.(int) == 0 {
		return nil
	}

	if n.value.(int) == 0 {
		return fmt.Errorf("%s must be greater than zero", n.input)
	}

	if n.positive && n.value.(int) < 0 {
		return fmt.Errorf("%s must be positive", n.input)
	}

	if n.differentZero && n.value.(int) == 0 {
		return fmt.Errorf("%s must be different from zero", n.input)
	}

	return nil
}
