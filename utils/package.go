package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// NArg is the number of arguments remaining after flags and options have been processed.
func NArg() int {
	return len(Arguments.Args())
}

func PrintDefautls() {
	for k, v := range Arguments.Options() {
		fmt.Print("-" + k + "\t" + v.Kind.String())
		if v.Required {
			fmt.Println("\tMANDATORY")
		} else {
			fmt.Println("\tOPTIONAL")
		}
		if v.DefaultValue != nil {
			fmt.Println("\tDefault: " + reflect.ValueOf(v.DefaultValue).String())
		}
		fmt.Println("\tUsage: " + v.Description)
	}
	fmt.Println("*Note:")
	fmt.Println("- All Arguments should come after all Options and Flags.")
	fmt.Println("- Either last Flag should be followed by an Option or last Flag should be passed with Boolean value(true or false).")
	fmt.Println("- Command Line Arguments will be given higher precedence over config file.")

}

func GetArgsInstance() *Args {
	return &Arguments
}

//ParseArguments is function that parses all the command line properties or properties mentioned in the config file.
func ParseArguments() {

	if len(os.Args) < 2 {
		PrintDefautls()
		os.Exit(1)
	}

	parseConfigFileParameter()

	for i := 1; i < len(os.Args); i++ {
		key := os.Args[i]
		argsStarted := false

		if len(key) > 0 {
			isFlag, _ := regexp.MatchString(`-\d*\D+`, key)
			isNotOption := true
			value := ""
			if isFlag {
				if argsStarted {
					PrintDefautls()
					os.Exit(1)
				}
				key = key[1:]
				checkHelpFlag := strings.ToLower(strings.TrimLeft(key, "-"))
				if checkHelpFlag == "h" || checkHelpFlag == "help" {
					PrintDefautls()
					os.Exit(0)
				}
				if strings.Contains(key, "=") {
					value = key[strings.Index(key, "=")+1:]
					key = key[0:strings.Index(key, "=")]
					isNotOption = false
				} else if i+1 < len(os.Args) {
					value = os.Args[i+1]
					isNotOption, _ = regexp.MatchString(`-\d*\D+`, value)
					if !isNotOption {
						i++
					}
				}
				if isNotOption {
					Arguments.setValue(key, true)
				} else {
					Arguments.setValue(key, value)
				}
			} else {
				Arguments.addArg(key)
				argsStarted = true
			}
		}
	}
	validateArguments()
}

func GetSessionManager() SessionManager {
	var sessionManager SessionManager = &sessionManagerImpl{}
	return sessionManager
}

func CreateNewTask(taskName string, function TaskFunc) *Task {
	var task Task = &taskImpl{task: function, name: taskName}
	return &task
}

func parseConfigFileParameter() {
	for i := 1; i < len(os.Args)-1; i++ {
		if os.Args[i] == "-config-file" {
			configFilePath := os.Args[i+1]
			configFile, error := os.Open(configFilePath)
			if error != nil {
				fmt.Println("Invalid Config File (path): " + configFilePath)
				return
			}
			defer configFile.Close()
			byteValue, _ := ioutil.ReadAll(configFile)
			var configFileMap map[string]interface{}
			json.Unmarshal(byteValue, &configFileMap)
			for key, value := range configFileMap {
				Arguments.setValue(key, value)
			}
			return
		}
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

func validateArguments() {
	missingMandatoryKeys := ""
	typeMisMatchKeys := ""

	for k, v := range Arguments.Options() {
		if v.Required && v.Value == nil {
			if v.OptionalCondition() {
				continue
			}
			missingMandatoryKeys = missingMandatoryKeys + k + ", "
		}
		if v.Value != nil && v.Kind != reflect.TypeOf(v.Value).Kind() {
			typeMisMatchKeys = typeMisMatchKeys + k + ", "
		}
	}
	if len(missingMandatoryKeys) > 0 {
		fmt.Println("Following mandatory arguments are missing:\n", missingMandatoryKeys[:len(missingMandatoryKeys)-2])
	}
	if len(typeMisMatchKeys) > 0 {
		fmt.Println("Value type mismatch for the following arguments:\n", typeMisMatchKeys[:len(typeMisMatchKeys)-2])
	}
	if len(missingMandatoryKeys) > 0 || len(typeMisMatchKeys) > 0 {
		fmt.Println("-----------------------------------\nPlease check the usage directions:")
		PrintDefautls()
		os.Exit(1)
	}
}

func skipServerVerfification() bool {
	ssv := Arguments.Get("skip-server-verification")
	if ssv != nil {
		return ssv.(bool)
	}
	return false
}

func init() {
	Arguments = &properties{make(map[string]*propValue), []string{}}
	Arguments.DefineString("server", "Hostname of vCenter Server.", true)
	Arguments.DefineString("username", "Username to login to the vCenter Server.", true)
	Arguments.DefineString("password", "Password to login to the vCenter Server.", true)
	Arguments.DefineRequiredString("certfile", "Absolute path to the file containing the trusted server certificate. This option can be skipped if the parameter skip-server-verification is specified.", skipServerVerfification)
	Arguments.DefineRequiredString("certkeyfile", "Absolute path to the key file of the server certificate. This option can be skipped if the parameter skip-server-verification is specified.", skipServerVerfification)
	Arguments.DefineBool("skip-server-verification", "Specify this option if you do not want to perform SSL certificate verification.")
	Arguments.DefineString("config-file", "Absolute path to  the configuration file containing the sample options.", false)
}

// func main() {
// 	Arguments.DefineString("checkTest", "Check Test Description.", true)
// 	ParseArguments()
// 	fmt.Println(Arguments.Options())
// 	fmt.Println(Arguments.Args())
// }
