package args

type Properties interface {
	GetPropertyValue(key string) interface{}

	GetStringPropertyValue(key string) (string, error)

	GetIntPropertyValue(key string) (int64, error)

	GetUIntPropertyValue(key string) (uint64, error)

	GetFloatPropertyValue(key string) (float64, error)

	GetBoolPropertyValue(key string) (bool, error)

	Args() []string

	Options() map[string]interface{}

	options() map[string]*propValue

	Arg(index uint) string

	// NArg is the number of arguments remaining after flags and options have been processed.
	NArg() int

	DefineString(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []string)

	DefineStringWithDefaultValue(key string, description string, defaultValue string, argOptionalCheck ArgOptionalIf, allowedValues []string)

	DefineInt(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []int64)

	DefineIntWithDefaultValue(key string, description string, defaultValue int64, argOptionalCheck ArgOptionalIf, allowedValues []int64)

	DefineUInt(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []uint64)

	DefineUIntWithDefaultValue(key string, description string, defaultValue uint64, argOptionalCheck ArgOptionalIf, allowedValues []uint64)

	DefineFloat(key string, description string, required bool, argOptionalCheck ArgOptionalIf, allowedValues []float64)

	DefineFloatWithDefaultValue(key string, description string, defaultValue float64, argOptionalCheck ArgOptionalIf, allowedValues []float64)

	DefineBool(key string, description string)

	DefineBoolWithDefaultValue(key string, description string, defaultValue bool)

	PrintDefaults()

	//ParseArgs is function that parses all the command line properties or properties mentioned in the config file.
	ParseArgs(argList []string) error

	AddProperties(jsonMap map[string]interface{}) error

	Validate() error

	Flag(flag string) bool

	SetProperty(key string, value interface{}) error

	AddOption(key string, value interface{}) error

	AddFlag(key string) error

	AddArg(value string)
}
