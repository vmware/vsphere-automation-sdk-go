/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package utils is utilities package that helps managing session,
//  authentication, execution as well as Host properties for fast
//  and easy development of Tasks/samples to be executed on the Host server.
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/host"
)

// Host defines the default host parsed from the Command Line parameters or the config file.
var Host host.Info

// Hosts defines the map of Host Name to Host details Info parsed from the multiple hosts defined in the config file.
var Hosts map[string]host.Info

// ParseArgs parses the command line parameters as well as config file, passed as command line parameter.
func ParseArgs() error {
	err := Host.ParseArgs(os.Args)
	if err == nil {
		configFile, configFileErr := Host.GetStringPropertyValue(keys.ConfigFileKey)
		if len(os.Args) < 3 && (configFileErr != nil || len(configFile) <= 0) {
			printDefaults(Host)
			return fmt.Errorf("Too few number of arguments")
		}
		if Host.Name() != "default" {
			Hosts[Host.Name()] = Host
		}
		if configFileErr != nil || len(configFile) == 0 {
			return nil
		}
		configFileParseErr := parseConfigFile(configFile)
		if configFileParseErr != nil {
			return configFileParseErr
		}
		setDefaultHost()
		allHostsValid := true
		for _, host := range Hosts {
			err := host.Validate()
			if err != nil {
				allHostsValid = false
			}
		}
		if !allHostsValid {
			err = fmt.Errorf("Invalid Arguments")
		}
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
	err = json.Unmarshal(byteValue, &configFileMap)
	if err != nil {
		return err
	}
	for key, value := range configFileMap {
		if key == "Hosts" && value != nil {
			fmt.Printf("Parsing Hosts from config file.")
			err := parseConfigFileHosts(value.([]interface{}))
			if err != nil {
				return err
			}
		} else {
			Host.SetProperty(key, value)
		}
	}
	return nil
}

func parseConfigFileHosts(hosts []interface{}) error {
	if hosts == nil || len(hosts) == 0 {
		return nil
	}
	for _, ht := range hosts {
		if ht == nil {
			continue
		}
		newHost, err := host.New(ht.(map[string]interface{}))
		if err != nil {
			return err
		}
		if newHost != nil {
			Hosts[newHost.Name()] = newHost
		}
	}
	return nil
}

func setDefaultHost() {
	defaultHostName := "default"
	if !Hosts[defaultHostName].IsSet() {
		for _, host := range Hosts {
			if host.Name() != defaultHostName {
				Hosts[defaultHostName] = host
				Host = Hosts[defaultHostName]
				return
			}
		}
	}
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
