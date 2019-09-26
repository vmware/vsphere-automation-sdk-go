/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Ntp
 * Used by client-side stubs.
 */

package ntp

import (
)

// ``Ntp`` interface provides methods Gets NTP configuration status and tests connection to ntp servers.
type NtpClient interface {


    // Test the connection to a list of ntp servers.
    //
    // @param serversParam List of host names or IP addresses of NTP servers.
    // @return List of test run statuses.
    // @throws Error Generic error
    Test(serversParam []string) ([]TestRunStatus, error) 


    // Set NTP servers. This method updates old NTP servers from configuration and sets the input NTP servers in the configuration. If NTP based time synchronization is used internally, the NTP daemon will be restarted to reload given NTP configuration. In case NTP based time synchronization is not used, this method only replaces servers in the NTP configuration.
    //
    // @param serversParam List of host names or ip addresses of ntp servers.
    // @throws Error Generic error
    Set(serversParam []string) error 


    // Get the NTP configuration status. If you run the 'timesync.get' command, you can retrieve the current time synchronization method (NTP- or VMware Tools-based). The 'ntp' command always returns the NTP server information, even when the time synchronization mode is not set to NTP. If the time synchronization mode is not NTP-based, the NTP server status is displayed as down.
    // @return List of NTP servers.
    // @throws Error Generic error
    Get() ([]string, error) 

}
