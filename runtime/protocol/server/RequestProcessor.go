package server
/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/
type RequestPreProcessor interface {
	Process(requestBody *map[string]interface{}) error
}
