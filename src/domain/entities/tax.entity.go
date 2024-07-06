package entities

import (
	"math"
	"reflect"

	validate_object "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/validate_objects"
)

type Tax struct {
	id *validate_object.StringValidateObject
	typeTax *validate_object.StringValidateObject
	rate *validate_object.NumberValidateObject
}

func InitTax(id,typeTax string, rate float64) *Tax {
	return &Tax{
		id: validate_object.NewStringValueObject(id, "id tax"),
		typeTax: validate_object.NewStringValueObject(typeTax, "type tax"),
		rate: validate_object.NewNumberValueObject(rate, "rate tax").IsDifferentZero().IsPositive(),
	}
}

func (t *Tax) ID() string {
	return t.id.Value()
}

func (t *Tax) TypeTax() string {
	return t.typeTax.Value()
}

func (t *Tax) Rate() float64 {
	return math.Round(float64(t.rate.Value())*100) / 100
}

func (t *Tax) Validate() error {
	v := reflect.ValueOf(t).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if validator, ok := field.Interface().(interface{ Validate() error }); ok {
			if err := validator.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}