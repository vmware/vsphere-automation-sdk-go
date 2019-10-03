package utils

type Args interface {
	Get(key string) interface{}

	Args() []string

	Options() map[string]*propValue

	Arg(index uint) string

	DefineString(key string, description string, required bool)

	DefineRequiredString(key string, description string, argOptionalCheck ArgOptionalIf)

	DefineStringWithDefaultValue(key string, defaultValue string, description string)

	DefineInt(key string, description string, required bool)

	DefineRequiredInt(key string, description string, argOptionalCheck ArgOptionalIf)

	DefineIntWithDefaultValue(key string, defaultValue int64, description string)

	DefineUInt(key string, description string, required bool)

	DefineRequiredUInt(key string, description string, argOptionalCheck ArgOptionalIf)

	DefineUIntWithDefaultValue(key string, defaultValue uint64, description string)

	DefineFloat(key string, description string, required bool)

	DefineRequiredFloat(key string, description string, argOptionalCheck ArgOptionalIf)

	DefineFloatWithDefaultValue(key string, defaultValue float64, description string)

	DefineBool(key string, description string)

	DefineBoolWithDefaultValue(key string, defaultValue bool, description string)

	setValue(key string, value interface{})

	addArg(value string)
}
