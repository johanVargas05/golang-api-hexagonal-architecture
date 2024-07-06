package validate_object

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type StringValidateObject struct {
	value     any
	input     string
	optional  bool
	isEmail   bool
	isID			bool
	maxLength int
	minLength int
}

func NewStringValueObject(value any, input string) *StringValidateObject {
	return &StringValidateObject{
		value:    value,
		optional: false,
		isEmail:  false,
		input:    input,
	}
}

func (s *StringValidateObject) IsOptional() *StringValidateObject {
	s.optional = true
	return s
}

func (s *StringValidateObject) IsEmail() *StringValidateObject {
	s.isEmail = true
	return s
}

func (s *StringValidateObject) IsID() *StringValidateObject {
	s.isID = true
	return s
}

func (s *StringValidateObject) MaxLength(maxLength int) *StringValidateObject {
	s.maxLength = maxLength
	return s
}

func (s *StringValidateObject) MinLength(minLength int) *StringValidateObject {
	s.minLength = minLength
	return s
}

func (s *StringValidateObject) TransformUpperCase() *StringValidateObject {
	if !isString(s.value) {
		return s
	}
	s.value = strings.ToUpper(s.Value())
	return s
}

func (s *StringValidateObject) TransformLowerCase() *StringValidateObject {
	if !isString(s.value) {
		return s
	}
	s.value = strings.ToLower(s.Value())
	return s
}

func (s *StringValidateObject) TransformSnakeCase() *StringValidateObject {
	if !isString(s.value) {
		return s
	}
	value := strings.ToLower(s.Value())
	s.value = strings.ReplaceAll(value, " ", "_")
	return s
}

func (s *StringValidateObject) TransformCamelCase() *StringValidateObject {
	if !isString(s.value) {
		return s
	}
	s.value = toCamelCase(s.Value())
	return s
}

func (s *StringValidateObject) Value() string {
	return s.value.(string)
}

func (s *StringValidateObject) Validate() error {
	ok := isString(s.value)
	if !ok {
		return fmt.Errorf("field %s is not a valid string", s.input)
	}

	value := s.value.(string)
	value = trimString(value)

	if s.optional && value == "" {
		return nil
	}

	if value == "" {
		return fmt.Errorf("field %s is required", s.input)
	}

	if s.isID && !isID(value) {
		return fmt.Errorf("field %s is not a valid ID", s.input)
	}

	if s.isEmail && !isEmail(value) {
		return fmt.Errorf("field %s is not a valid email", s.input)
	}

	if s.maxLength != 0 && len(value) > s.maxLength {
		return fmt.Errorf("field %s exceeds the maximum length of %d characters", s.input, s.maxLength)
	}

	if s.minLength != 0 && len(value) < s.minLength {
		return fmt.Errorf("field %s must have a minimum length of %d characters", s.input, s.minLength)
	}

	return nil
}

func isID(id string) bool {
	pattern := `^(?:[0-9a-fA-F]{24}|[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12})$`

	idRegex := regexp.MustCompile(pattern)
	return idRegex.MatchString(id)
}

func isEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	emailRegex := regexp.MustCompile(pattern)
	return emailRegex.MatchString(email)

}

func toCamelCase(value string) string {
	titleCase := cases.Title(language.Spanish)
	return titleCase.String(strings.ToLower(value))
}

func trimString(value string) string {
	return strings.TrimSpace(value)
}

func isString(value any) bool {
	_, ok := value.(string)
	return ok
}
