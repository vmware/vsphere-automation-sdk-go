package metadata

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

import "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"

func getError(msg string) error {
	args := map[string]string{
		"msg": msg,
	}
	return l10n.NewRuntimeError("vapi.metadata.parser.failure", args)
}
