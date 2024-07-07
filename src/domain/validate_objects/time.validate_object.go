package validate_object

import "time"

type TimeValidateObject struct {
	value    any
	input    string
	optional bool
	format   string
}

func NewTimeValueObject(value any, input string) *TimeValidateObject {
	return &TimeValidateObject{
		value:    value,
		optional: false,
		input:    input,
	}
}

func (t *TimeValidateObject) IsOptional() *TimeValidateObject {
	t.optional = true
	return t
}

func (t *TimeValidateObject) Format(format string) *TimeValidateObject {
	t.format = format
	return t
}

func (t *TimeValidateObject) Value() time.Time {
	if t.value == nil {
		return time.Time{}
	}
	return *(t.value.(*time.Time))
}

func (t *TimeValidateObject) ValueString() string {
	if t.value == nil || t.value.(*time.Time) == nil {
		return ""
	}
	if t.format != "" {
		return t.value.(*time.Time).Format(t.format)
	}
	return t.value.(*time.Time).String()
}
