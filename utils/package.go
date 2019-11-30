package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/host"
)

var Host host.Info
var Hosts map[string]host.Info

func ParseArgs() error {
	if len(os.Args) < 3 {
		printDefaults(Host)
		return fmt.Errorf("Too few number of arguments")
	}
	err := Host.ParseArgs(os.Args)
	if err == nil {
		if Host.Name() != "default" {
			Hosts[Host.Name()] = Host
		}
		configFile, configFileErr := Host.GetStringPropertyValue(keys.ConfigFileKey)
		if configFileErr != nil || len(configFile) == 0 {
			return nil
		}
		configFileParseErr := parseConfigFile(configFile)
		return configFileParseErr
	}
	return err
}

func parseConfigFile(configFilePath string) error {
	if len(configFilePath) == 0 {
		return fmt.Errorf("Config File Path cannot be empty")
	}
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return err
	}
	defer configFile.Close()
	byteValue, _ := ioutil.ReadAll(configFile)
	var configFileMap map[string]interface{}
	json.Unmarshal(byteValue, &configFileMap)

	for key, value := range configFileMap {
		if key == "Hosts" && value != nil {
			err := parseConfigFileHosts(value.([]map[string]interface{}))
			if err != nil {
				return err
			}
		} else {
			Host.SetProperty(key, value)
		}
	}
	return nil
}

func parseConfigFileHosts(hosts []map[string]interface{}) error {
	if hosts == nil || len(hosts) == 0 {
		return nil
	}
	for _, ht := range hosts {
		if ht == nil {
			continue
		}
		newHost, err := host.New(ht)
		if err != nil {
			return err
		}
		if newHost != nil {
			Hosts[newHost.Name()] = newHost
		}
	}
	return nil
}

func init() {
	Hosts = make(map[string]host.Info)
	Hosts["default"] = host.DefineDefault()
	Host = Hosts["default"]
}

func printDefaults(ht host.Info) {
	ht.PrintDefaults()
	if ht.Name() == "default" {
		fmt.Println("*Note:")
		fmt.Println("- All Arguments should come after all Options and Flags.")
		fmt.Println("- Either last Flag should be followed by an Option or last Flag should be passed with Boolean value(true or false).")
		fmt.Println("- Command Line Arguments will be given higher precedence over config file.")
		fmt.Println("- All Flags and Keys of Options are case insensitive.")
	}
}
