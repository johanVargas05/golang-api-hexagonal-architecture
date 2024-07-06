package validate_object

import (
	"fmt"
	"math"
)

type FloatValidateObject struct {
	value         any
	input         string
	optional      bool
	positive      bool
	differentZero bool
	decimals      int
}

func NewFloatValueObject(value any, input string) *FloatValidateObject {
	return &FloatValidateObject{
		value:         value,
		optional:      false,
		positive:      false,
		differentZero: false,
		input:         input,
		decimals:      2,
	}
}

func (n *FloatValidateObject) IsOptional() *FloatValidateObject {
	n.optional = true
	return n
}

func (n *FloatValidateObject) IsPositive() *FloatValidateObject {
	n.positive = true
	return n
}

func (n *FloatValidateObject) IsDifferentZero() *FloatValidateObject {
	n.differentZero = true
	return n
}

func (n *FloatValidateObject) Decimals(decimals int) *FloatValidateObject {
	n.decimals = decimals
	return n
}

func (n *FloatValidateObject) Value() float64 {
	value, ok := n.value.(float64)
	if !ok || math.IsNaN(value) {
		return 0
	}

	decimals := float64(n.decimals)
	factor := math.Pow(10, decimals)
	value = math.Round(value*factor) / factor

	return value
}

func (n *FloatValidateObject) Validate() error {

	value, ok := n.value.(float64)

	if !ok {
		return fmt.Errorf("%s must be a float64", n.input)
	}

	if n.optional && value == 0 {
		return nil
	}

	if n.positive && value < 0 {
		return fmt.Errorf("%s must be positive", n.input)
	}

	if n.differentZero && value == 0 {
		return fmt.Errorf("%s must be different from zero", n.input)
	}

	return nil
}
