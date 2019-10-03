package utils

import (
	"reflect"
)

func newPropDefaultValue(value interface{}, defaultValue interface{}, kind reflect.Kind, description string) *propValue {
	return &propValue{Value: value, DefaultValue: defaultValue, Kind: kind, Description: description, Required: false}
}

func newPropValue(value interface{}, kind reflect.Kind, description string, required bool, optionalCondtion ArgOptionalIf) *propValue {
	return &propValue{Value: value, DefaultValue: nil, Kind: kind, Description: description, Required: required, OptionalCondition: optionalCondtion}
}

func (pv *propValue) setValue(val interface{}) {
	if val == nil {
		pv.Value = nil
	} else {
		pv.Value = parseValue(reflect.ValueOf(val).String())
	}
}

func (prop properties) Get(key string) interface{} {
	if prop.options[key] != nil {
		if prop.options[key].Value == nil {
			return prop.options[key].DefaultValue
		}
		return prop.options[key].Value
	}
	return nil
}

func (prop properties) Args() []string {
	return prop.args
}

func (prop properties) Options() map[string]*propValue {
	return prop.options
}

func (prop properties) Arg(index uint) string {
	return prop.args[index]
}

func (prop *properties) DefineString(key string, description string, required bool) {
	propVal := newPropValue(nil, reflect.String, description, required, nil)
	prop.options[key] = propVal
}

func (prop *properties) DefineRequiredString(key string, description string, argOptionalCheck ArgOptionalIf) {
	propVal := newPropValue(nil, reflect.String, description, true, argOptionalCheck)
	prop.options[key] = propVal
}

func (prop *properties) DefineStringWithDefaultValue(key string, defaultValue string, description string) {
	propVal := newPropDefaultValue(nil, defaultValue, reflect.String, description)
	prop.options[key] = propVal
}

func (prop *properties) DefineInt(key string, description string, required bool) {
	propVal := newPropValue(nil, reflect.Int64, description, required, nil)
	prop.options[key] = propVal
}

func (prop *properties) DefineRequiredInt(key string, description string, argOptionalCheck ArgOptionalIf) {
	propVal := newPropValue(nil, reflect.Int64, description, true, argOptionalCheck)
	prop.options[key] = propVal
}

func (prop *properties) DefineIntWithDefaultValue(key string, defaultValue int64, description string) {
	propVal := newPropDefaultValue(nil, defaultValue, reflect.Int64, description)
	prop.options[key] = propVal
}

func (prop *properties) DefineUInt(key string, description string, required bool) {
	propVal := newPropValue(nil, reflect.Uint64, description, required, nil)
	prop.options[key] = propVal
}

func (prop *properties) DefineRequiredUInt(key string, description string, argOptionalCheck ArgOptionalIf) {
	propVal := newPropValue(nil, reflect.Uint64, description, true, argOptionalCheck)
	prop.options[key] = propVal
}

func (prop *properties) DefineUIntWithDefaultValue(key string, defaultValue uint64, description string) {
	propVal := newPropDefaultValue(nil, defaultValue, reflect.Uint64, description)
	prop.options[key] = propVal
}

func (prop *properties) DefineFloat(key string, description string, required bool) {
	propVal := newPropValue(nil, reflect.Float64, description, required, nil)
	prop.options[key] = propVal
}

func (prop *properties) DefineRequiredFloat(key string, description string, argOptionalCheck ArgOptionalIf) {
	propVal := newPropValue(nil, reflect.Float64, description, true, argOptionalCheck)
	prop.options[key] = propVal
}

func (prop *properties) DefineFloatWithDefaultValue(key string, defaultValue float64, description string) {
	propVal := newPropDefaultValue(nil, defaultValue, reflect.Float64, description)
	prop.options[key] = propVal
}

func (prop *properties) DefineBool(key string, description string) {
	propVal := newPropValue(nil, reflect.Bool, description, false, nil)
	prop.options[key] = propVal
}

func (prop *properties) DefineBoolWithDefaultValue(key string, defaultValue bool, description string) {
	propVal := newPropDefaultValue(nil, defaultValue, reflect.Bool, description)
	prop.options[key] = propVal
}

func (prop *properties) setValue(key string, value interface{}) {
	if prop.options[key] == nil {
		propVal := newPropValue(value, getValueKind(value), "No Description Available", false, nil)
		prop.options[key] = propVal
	} else {
		prop.options[key].setValue(value)
	}
}

func (prop *properties) addArg(value string) {
	prop.args = append(prop.args, value)
}

// func main() {
// 	Arguments.DefineString("checkTest", "Check Test Description.", true)
// 	ParseArguments()
// 	fmt.Println(Arguments.Options())
// 	fmt.Println(Arguments.Args())
// }
