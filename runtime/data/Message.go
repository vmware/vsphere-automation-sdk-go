/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/
package data

// TODO: Remove legacy data.Message when recommended change is past https://reviewboard.eng.vmware.com/r/1492104
type Message struct {
	id             string
	defaultMessage string
	args           []string
}

func NewMessage(id string, defaultMessage string, args []string) Message {
	return Message{id: id, defaultMessage: defaultMessage, args: args}
}

func (message *Message) GetID() string {
	return message.id
}

func (message *Message) GetDefaultMessage() string {
	return message.defaultMessage
}

func (message *Message) GetArgs() []string {
	return message.args
}
