/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: SessionManager
 * Used by client-side stubs.
 */

package sessionManager

import (
)

// Session management interface for the publish API
type SessionManagerClient interface {


    // Logs a user and returns a session id. Authentication information should be provided in the SecurityContext.
    Login() (string, error) 


    // Invalidates a session. Session Id should be provided in the SecurityContext.
    Logout() error 


    // Keeps a session id alive. Session Id should be provided in the SecurityContext.
    KeepAlive() error 

}
