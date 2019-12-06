/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package args

// Properties represents Options, Flags and Arguments from
// Command Line or Config file in json format.
type Properties interface {
	// GetPropertyValue fetches value for the Option/Flag defined by the key.
	GetPropertyValue(key string) interface{}

	// GetStringPropertyValue fetches string value for the Option/Flag defined by the key.
	GetStringPropertyValue(key string) (string, error)

	// GetIntPropertyValue fetches value of type int64 for the Option/Flag defined by the key.
	GetIntPropertyValue(key string) (int64, error)

	// GetUIntPropertyValue fetches value of type uint64 for the Option/Flag defined by the key.
	GetUIntPropertyValue(key string) (uint64, error)

	// GetFloatPropertyValue fetches value of type float64 for the Option/Flag defined by the key.
	GetFloatPropertyValue(key string) (float64, error)

	// GetBoolPropertyValue fetches bool value for the Option/Flag defined by the key.
	GetBoolPropertyValue(key string) (bool, error)

	// Args fetches arguments as array of string.
	Args() []string

	// Options fetches all Options and Flags as Map of string to interface
	Options() map[string]interface{}

	// options fetches all Options and Flags as Map of string to type propValue
	options() map[string]*propValue

	// Arg fetches argument defined at the index.
	Arg(index uint) string

	// NArg gives the number of arguments.
	NArg() int

	// Define new property of type string.
	DefineString(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []string)

	// Define new property of type string with default value.
	DefineStringWithDefaultValue(key string, description string, defaultValue string, argOptionalCheck ArgOptionalIf, allowedValues []string)

	// Define new property of type int64.
	DefineInt(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []int64)

	// Define new property of type int64 with default value.
	DefineIntWithDefaultValue(key string, description string, defaultValue int64, argOptionalCheck ArgOptionalIf, allowedValues []int64)

	// Define new property of type uint64.
	DefineUInt(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []uint64)

	// Define new property of type uint64 with default value.
	DefineUIntWithDefaultValue(key string, description string, defaultValue uint64, argOptionalCheck ArgOptionalIf, allowedValues []uint64)

	// Define new property of type float64.
	DefineFloat(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []float64)

	// Define new property of type float64 with default value.
	DefineFloatWithDefaultValue(key string, description string, defaultValue float64, argOptionalCheck ArgOptionalIf, allowedValues []float64)

	// Define new property of type bool.
	DefineBool(key string, description string)

	// Define new property of type bool with default value.
	DefineBoolWithDefaultValue(key string, description string, defaultValue bool)

	// PrintDefaults prints the usage and descriptions of all the Defined Properties.
	PrintDefaults()

	//ParseArgs is function that parses all the command line properties or properties mentioned in the config file.
	ParseArgs(argList []string) error

	// AddProperties adds all the properties passed as map of string to interface.
	AddProperties(jsonMap map[string]interface{}) error

	// Validate validates all the options and flags defined.
	Validate() error

	// Flag checks if the flag is defined and if defined return bool value.
	Flag(flag string) bool

	// SetProperty sets the value for the key.
	SetProperty(key string, value interface{}) error

	// AddOption adds an option value for the key.
	AddOption(key string, value interface{}) error

	// AddFlag adds flag with the key.
	AddFlag(key string) error

	// AddArg adds argument to the list of arguments.
	AddArg(value string)
}
