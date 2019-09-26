/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/
package protocol

/**
 * ProtocolHanders are the classes that provide the endpoint for a given
 * protocol.
 */
type ProtocolHandler interface {
	Start()
	Stop()
}
