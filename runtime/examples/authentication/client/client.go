/* Copyright © 2021-2022 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package main

import (
	"fmt"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/examples/authentication/client/bindings/src/generated/vmodl/examples/greeter"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
)

const correctUsername = "username"
const correctPassword = "password"

const wrongUsername = "345username"
const wrongPassword = "123password"

func callProvider(greeterClient greeter.GreeterClient) {
	var formalGreeting, err = greeterClient.FormalyGreet()
	if err != nil {
		fmt.Printf("Error: %s\n", err.(errors.Unauthenticated).Messages[0].DefaultMessage)
	} else {
		fmt.Printf("Response: %s\n", formalGreeting)
	}

	informalGreeting, err := greeterClient.InformalyGreet()
	if err != nil {
		fmt.Printf("Error: %s\n", err.(errors.Unauthenticated).Messages[0].DefaultMessage)
	} else {
		fmt.Printf("Response: %s\n", informalGreeting)
	}
}

func main() {
	var sctxCorrectCredentials = security.NewUserPasswordSecurityContext(correctUsername, correctPassword)
	var connector = client.NewConnector(
		"http://localhost:8088",
		client.WithSecurityContext(sctxCorrectCredentials))
	var greeterClient = greeter.NewGreeterClient(connector)

	fmt.Printf("Username: %s\nPassword: %s\n", correctUsername, correctPassword)
	callProvider(greeterClient)

	var sctxWrongCredentials = security.NewUserPasswordSecurityContext(wrongUsername, wrongPassword)
	connector = client.NewConnector(
		"http://localhost:8088",
		client.WithSecurityContext(sctxWrongCredentials))
	greeterClient = greeter.NewGreeterClient(connector)
	fmt.Printf("Username: %s\nPassword: %s\n", wrongUsername, wrongPassword)
	callProvider(greeterClient)
}
