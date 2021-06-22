package validator

import (
	"reflect"
	"strings"
)

type ValidError struct {
	errorMessage []string
}

func (v *ValidError) Error() string {
	return strings.Join(v.errorMessage, ", ")
}

func (v *ValidError) AddError(err error) {
	v.errorMessage = append(v.errorMessage, err.Error())
}

func (v *ValidError) AddErrorMessage(msg string) {
	v.errorMessage = append(v.errorMessage, msg)
}

type Validator struct {
	ValidError
}

type Holder struct {
	ValidError
}

func (h *Holder) Check() error {
	if len(h.errorMessage) == 0 {
		return nil
	}
	return &h.ValidError
}

type ValueHolder struct {
	Holder
	value interface{}
}

type StringHolder struct {
	Holder
	value string
}

func StringOf(value string) *StringHolder {
	return &StringHolder{value: value}
}

func (vh *StringHolder) CheckLength(left, right int, msg string) *StringHolder {
	if len(vh.value) < left || len(vh.value) >= right {
		vh.AddErrorMessage(msg)
	}
	return vh
}

func (vh *StringHolder) MustEmpty(msg string) *StringHolder {
	if len(vh.value) > 0 {
		vh.AddErrorMessage(msg)
	}
	return vh
}

func (vh *StringHolder) MustNotEmpty(msg string) *StringHolder {
	if len(vh.value) == 0 {
		vh.AddErrorMessage(msg)
	}
	return vh
}

func (vh *StringHolder) MustHasSuffix(suffix string, msg string) *StringHolder {
	if !strings.HasSuffix(vh.value, suffix) {
		vh.AddErrorMessage(msg)
	}
	return vh
}

func (vh *StringHolder) MustHasPrefix(suffix string, msg string) *StringHolder {
	if !strings.HasPrefix(vh.value, suffix) {
		vh.AddErrorMessage(msg)
	}
	return vh
}

func Of(value interface{}) *ValueHolder {
	return &ValueHolder{value: value}
}

func (vh *ValueHolder) MustString(msg string) *StringHolder {
	result := &StringHolder{}
	if reflect.TypeOf(vh.value).Kind() == reflect.String {
		result.value = reflect.ValueOf(vh.value).String()
		return result
	}
	result.AddErrorMessage(msg)
	return result
}
