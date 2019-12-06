/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package args

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type properties struct {
	OptionsMap map[string]*propValue
	StringArgs []string
}

func (prop properties) GetPropertyValue(key string) interface{} {
	if prop.OptionsMap[key] != nil {
		if prop.OptionsMap[key].Value == nil {
			return prop.OptionsMap[key].DefaultValue
		}
		return prop.OptionsMap[key].Value
	}
	return nil
}

func (prop properties) GetStringPropertyValue(key string) (string, error) {
	if len(key) == 0 {
		return "", fmt.Errorf("Key cannot be empty")
	}
	value, validKey := prop.OptionsMap[key]
	if !validKey {
		return "", fmt.Errorf("Key %v is not present", key)
	} else if value == nil || value.Value == nil {
		return "", fmt.Errorf("Nil Value for the Key %v", key)
	} else if value.Kind != reflect.String {
		return "", fmt.Errorf("Key %v do not have String value, has %v value", key, value.Kind.String())
	}
	return value.Value.(string), nil
}

func (prop properties) GetIntPropertyValue(key string) (int64, error) {
	if len(key) == 0 {
		return 0, fmt.Errorf("Key cannot be empty")
	}
	value, validKey := prop.OptionsMap[key]
	if !validKey {
		return 0, fmt.Errorf("Key %v is not present", key)
	} else if value == nil || value.Value == nil {
		return 0, fmt.Errorf("Nil Value for the Key %v", key)
	} else if value.Kind != reflect.Int64 {
		return 0, fmt.Errorf("Key %v do not have Integer value, has %v value", key, value.Kind.String())
	}
	return value.Value.(int64), nil
}

func (prop properties) GetUIntPropertyValue(key string) (uint64, error) {
	if len(key) == 0 {
		return 0, fmt.Errorf("Key cannot be empty")
	}
	value, validKey := prop.OptionsMap[key]
	if !validKey {
		return 0, fmt.Errorf("Key %v is not present", key)
	} else if value == nil || value.Value == nil {
		return 0, fmt.Errorf("Nil Value for the Key %v", key)
	} else if value.Kind != reflect.Uint64 {
		return 0, fmt.Errorf("Key %v do not have Unsigned Integer value, has %v value", key, value.Kind.String())
	}
	return value.Value.(uint64), nil
}

func (prop properties) GetFloatPropertyValue(key string) (float64, error) {
	if len(key) == 0 {
		return 0, fmt.Errorf("Key cannot be empty")
	}
	value, validKey := prop.OptionsMap[key]
	if !validKey {
		return 0, fmt.Errorf("Key %v is not present", key)
	} else if value == nil || value.Value == nil {
		return 0, fmt.Errorf("Nil Value for the Key %v", key)
	} else if value.Kind != reflect.Float64 {
		return 0, fmt.Errorf("Key %v do not have Float value, has %v value", key, value.Kind.String())
	}
	return value.Value.(float64), nil
}

func (prop properties) GetBoolPropertyValue(key string) (bool, error) {
	if len(key) == 0 {
		return false, fmt.Errorf("Key cannot be empty")
	}
	value, validKey := prop.OptionsMap[key]
	if !validKey {
		return false, fmt.Errorf("Key %v is not present", key)
	} else if value == nil || value.Value == nil {
		return false, fmt.Errorf("Nil Value for the Key %v", key)
	} else if value.Kind != reflect.Bool {
		return false, fmt.Errorf("Key %v do not have Boolean value, has %v value", key, value.Kind.String())
	}
	return value.Value.(bool), nil
}

func (prop properties) Args() []string {
	return prop.StringArgs
}

func (prop properties) options() map[string]*propValue {
	return prop.OptionsMap
}

func (prop properties) Options() map[string]interface{} {
	var ops map[string]interface{} = make(map[string]interface{})
	for k := range prop.OptionsMap {
		ops[k] = prop.GetPropertyValue(k)
	}
	return ops
}

func (prop properties) Arg(index uint) string {
	return prop.StringArgs[index]
}

func (prop *properties) DefineString(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []string) {
	var values []interface{} = nil
	if allowedValues != nil && len(allowedValues) > 0 {
		for _, value := range allowedValues {
			values = append(values, value)
		}
	}
	propVal := newPropValue(nil, nil, values, reflect.String, description, required, argOptionalCheck)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) DefineStringWithDefaultValue(key string, description string, defaultValue string, argOptionalCheck ArgOptionalIf, allowedValues []string) {
	var values []interface{} = nil
	if allowedValues != nil && len(allowedValues) > 0 {
		for _, value := range allowedValues {
			values = append(values, value)
		}
	}
	propVal := newPropValue(nil, defaultValue, values, reflect.String, description, false, argOptionalCheck)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) DefineInt(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []int64) {
	var values []interface{} = nil
	if allowedValues != nil && len(allowedValues) > 0 {
		for _, value := range allowedValues {
			values = append(values, value)
		}
	}
	propVal := newPropValue(nil, nil, values, reflect.Int64, description, required, argOptionalCheck)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) DefineIntWithDefaultValue(key string, description string, defaultValue int64, argOptionalCheck ArgOptionalIf, allowedValues []int64) {
	var values []interface{} = nil
	if allowedValues != nil && len(allowedValues) > 0 {
		for _, value := range allowedValues {
			values = append(values, value)
		}
	}
	propVal := newPropValue(nil, defaultValue, values, reflect.Int64, description, false, argOptionalCheck)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) DefineUInt(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []uint64) {
	var values []interface{} = nil
	if allowedValues != nil && len(allowedValues) > 0 {
		for _, value := range allowedValues {
			values = append(values, value)
		}
	}
	propVal := newPropValue(nil, nil, values, reflect.Uint64, description, required, argOptionalCheck)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) DefineUIntWithDefaultValue(key string, description string, defaultValue uint64, argOptionalCheck ArgOptionalIf, allowedValues []uint64) {
	var values []interface{} = nil
	if allowedValues != nil && len(allowedValues) > 0 {
		for _, value := range allowedValues {
			values = append(values, value)
		}
	}
	propVal := newPropValue(nil, defaultValue, values, reflect.Uint64, description, false, argOptionalCheck)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) DefineFloat(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []float64) {
	var values []interface{} = nil
	if allowedValues != nil && len(allowedValues) > 0 {
		for _, value := range allowedValues {
			values = append(values, value)
		}
	}
	propVal := newPropValue(nil, nil, values, reflect.Float64, description, required, argOptionalCheck)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) DefineFloatWithDefaultValue(key string, description string, defaultValue float64, argOptionalCheck ArgOptionalIf, allowedValues []float64) {
	var values []interface{} = nil
	if allowedValues != nil && len(allowedValues) > 0 {
		for _, value := range allowedValues {
			values = append(values, value)
		}
	}
	propVal := newPropValue(nil, defaultValue, values, reflect.Float64, description, false, argOptionalCheck)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) DefineBool(key string, description string) {
	propVal := newPropValue(nil, nil, nil, reflect.Bool, description, false, nil)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) DefineBoolWithDefaultValue(key string, description string, defaultValue bool) {
	propVal := newPropValue(nil, defaultValue, nil, reflect.Bool, description, false, nil)
	prop.OptionsMap[key] = propVal
}

func (prop *properties) SetProperty(key string, value interface{}) error {
	if len(key) == 0 && value == nil {
		return fmt.Errorf("Both Key and Value of Property cannot be empty")
	} else if len(key) == 0 && value != nil {
		prop.AddArg(value.(string))
	} else if prop.OptionsMap[key] == nil {
		propVal := newPropValue(value, nil, nil, getValueKind(value), "No Description Available.", false, nil)
		prop.OptionsMap[key] = propVal
	} else {
		prop.OptionsMap[key].setValue(value)
	}
	return nil
}

func (prop *properties) AddOption(key string, value interface{}) error {
	if len(key) == 0 {
		return fmt.Errorf("Key cannot be empty in an Option")
	}
	return prop.SetProperty(key, value)
}

func (prop *properties) AddFlag(key string) error {
	if len(key) == 0 {
		return fmt.Errorf("Key cannot be empty in a Flag")
	}
	prop.DefineBool(key, "No Description Available.")
	return prop.SetProperty(key, true)
}

func (prop *properties) AddArg(value string) {
	prop.StringArgs = append(prop.StringArgs, value)
}

func (prop *properties) Flag(flag string) bool {
	ssv := prop.GetPropertyValue(flag)
	if ssv != nil {
		return ssv.(bool)
	}
	return false
}

func (prop *properties) NArg() int {
	return len(prop.Args())
}

func (prop *properties) PrintDefaults() {
	fmt.Println(prop.UsageDescription())
}

func (prop *properties) UsageDescription() string {
	var usageSB strings.Builder
	for k, v := range prop.options() {
		usageSB.WriteString(fmt.Sprintf("\n-%v\t", k))
		usageSB.WriteString(v.UsageDescription())
	}
	return usageSB.String()
}

func (prop *properties) ParseArgs(argList []string) error {
	allPropertiesValid := true
	for i := 1; i < len(argList); i++ {
		key := argList[i]
		argsStarted := false

		if len(key) > 0 {
			isKey, _ := regexp.MatchString(`-\d*\D+`, key)
			isFlag := true
			value := ""
			if isKey {
				if argsStarted {
					prop.PrintDefaults()
					os.Exit(1)
				}
				key = key[1:]
				checkHelpFlag := strings.ToLower(strings.TrimLeft(key, "-"))
				if checkHelpFlag == "h" || checkHelpFlag == "help" {
					prop.PrintDefaults()
					os.Exit(0)
				}
				if strings.Contains(key, "=") {
					value = key[strings.Index(key, "=")+1:]
					key = key[0:strings.Index(key, "=")]
					isFlag = false
				} else if i+1 < len(argList) {
					value = argList[i+1]
					isFlag, _ = regexp.MatchString(`^-\d*\D+`, value)
					if !isFlag {
						i++
					}
				}
				var err error = nil
				if isFlag {
					err = prop.SetProperty(key, "true")
				} else {
					err = prop.SetProperty(key, value)
				}
				if err != nil {
					fmt.Printf("\n%v %v", err.Error(), key)
					allPropertiesValid = false
				}
			} else {
				prop.AddArg(key)
				argsStarted = true
			}
		}
	}
	if !allPropertiesValid {
		fmt.Print("\n-----------------------------------\nPlease check the usage directions:")
		prop.PrintDefaults()
		return fmt.Errorf("Invalid Properties")
	}
	return nil
}

func (prop *properties) AddProperties(jsonMap map[string]interface{}) error {
	var originalMap map[string]interface{} = prop.Options()
	for key, value := range jsonMap {
		prop.SetProperty(key, value)
	}
	err := prop.Validate()
	if err != nil {
		prop.restoreValues(originalMap)
	}
	return err
}

func (prop *properties) Validate() error {
	allPropertiesValid := true
	for k, v := range prop.options() {
		isValid, err := v.Validate(Properties(prop))
		allPropertiesValid = allPropertiesValid && isValid
		if !isValid {
			fmt.Printf("\n%v %v", err.Error(), k)
		} else if err != nil {
			delete(prop.OptionsMap, k)
		}
	}
	if !allPropertiesValid {
		fmt.Print("\n-----------------------------------\nPlease check the usage directions:")
		prop.PrintDefaults()
		return fmt.Errorf("Invalid Properties")
	}
	return nil
}

func (prop *properties) restoreValues(originalValues map[string]interface{}) {
	var allOrigKeySB strings.Builder
	for key, origValue := range originalValues {
		allOrigKeySB.WriteString(key)
		allOrigKeySB.WriteString(", ")
		prop.SetProperty(key, origValue)
	}
	allOrigKeyString := allOrigKeySB.String()
	for key := range prop.Options() {
		if !strings.Contains(allOrigKeyString, key) {
			delete(prop.OptionsMap, key)
		}
	}
}

func getValueKind(value interface{}) reflect.Kind {
	switch value.(type) {
	case bool:
		return reflect.Bool
	case int64, int32, int16, int8, int:
		return reflect.Int64
	case uint64, uint32, uint16, uint8, uint:
		return reflect.Uint64
	case float64, float32:
		return reflect.Float64
	default:
		return reflect.String
	}
}
