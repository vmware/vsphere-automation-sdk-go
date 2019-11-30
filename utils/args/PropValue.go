package args

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ArgOptionalIf func(properties Properties) bool

type propValue struct {
	Value             interface{}
	DefaultValue      interface{}
	Kind              reflect.Kind
	Description       string
	Required          bool
	OptionalCondition ArgOptionalIf
	AllowedValues     []interface{}
}

func newPropValue(value interface{}, defaultValue interface{},
	enumValues []interface{}, kind reflect.Kind, description string,
	required bool, optionalCondtion ArgOptionalIf) *propValue {
	return &propValue{
		Value:             value,
		DefaultValue:      defaultValue,
		AllowedValues:     enumValues,
		Kind:              kind,
		Description:       description,
		Required:          required,
		OptionalCondition: optionalCondtion,
	}
}

func (pv *propValue) setValue(val interface{}) {
	if val == nil {
		pv.Value = nil
	} else {
		pv.Value = parseValue(reflect.ValueOf(val).String())
	}
}

func parseValue(value string) interface{} {
	valueUInt, error := strconv.ParseUint(value, 10, 64)
	if error == nil {
		return valueUInt
	}

	valueBool, error := strconv.ParseBool(value)
	if error == nil {
		return valueBool
	}

	valueInt, error := strconv.ParseInt(value, 10, 64)
	if error == nil {
		return valueInt
	}

	valueFloat, error := strconv.ParseFloat(value, 64)
	if error == nil {
		return valueFloat
	}

	return value
}

func (pv *propValue) isValueValid(properties Properties) (isValid bool, err error) {
	if pv.Required && pv.Value == nil {
		if pv.OptionalCondition(properties) {
			return true, fmt.Errorf("Value is nil for the Argument")
		}
		return false, fmt.Errorf("Missing value for Required Argument")
	}

	if pv.Value != nil && pv.Kind != reflect.ValueOf(pv.Value).Kind() {
		return false, fmt.Errorf("Type mismatch for the value of the Argument")
	}
	if pv.AllowedValues != nil && len(pv.AllowedValues) > 0 {
		isValueAllowed := false
		for _, allowedValue := range pv.AllowedValues {
			if allowedValue == pv.Value {
				isValueAllowed = true
			}
			if isValueAllowed {
				break
			}
		}
		if !isValueAllowed {
			return false, fmt.Errorf("Invalid value for the Argument")
		}
	}
	return true, nil
}

func (pv *propValue) UsageDescription() string {
	var usageDescription strings.Builder
	usageDescription.WriteString(pv.Kind.String())
	if pv.Required {
		usageDescription.WriteString("\tMANDATORY")
	} else {
		usageDescription.WriteString("\tOPTIONAL")
	}
	if len(pv.Description) > 0 {
		usageDescription.WriteString("\n\tDescription: ")
		usageDescription.WriteString(pv.Description)
	} else {
		usageDescription.WriteString("\n\tDescription: No Description Available.")
	}
	if pv.DefaultValue != nil {
		usageDescription.WriteString("\n\tDefault-Value: ")
		usageDescription.WriteString(reflect.ValueOf(pv.DefaultValue).String())
	}
	if pv.AllowedValues != nil && len(pv.AllowedValues) > 0 {
		usageDescription.WriteString("\n\tAllowed Values: [")
		var allowedValuesSB strings.Builder
		for _, allowedValue := range pv.AllowedValues {
			if allowedValue != nil {
				if allowedValuesSB.Len() > 0 {
					allowedValuesSB.WriteString(", ")
				}
				allowedValuesSB.WriteString(fmt.Sprint(allowedValue))
			}
		}
		if allowedValuesSB.Len() > 0 {
			allowedValuesSB.WriteString("]")
			usageDescription.WriteString(allowedValuesSB.String())
		}
	}
	return usageDescription.String()
}
