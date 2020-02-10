/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Ntp
 * Used by client-side stubs.
 */

package appliance


// ``Ntp`` interface provides methods Gets NTP configuration status and tests connection to ntp servers. This interface was added in vSphere API 6.7.
type NtpClient interface {

    // Test the connection to a list of ntp servers. This method was added in vSphere API 6.7.
    //
    // @param serversParam List of host names or IP addresses of NTP servers.
    // @return List of test run statuses.
    // @throws Error Generic error
	Test(serversParam []string) ([]NtpTestRunStatus, error)

    // Set NTP servers. This method updates old NTP servers from configuration and sets the input NTP servers in the configuration. If NTP based time synchronization is used internally, the NTP daemon will be restarted to reload given NTP configuration. In case NTP based time synchronization is not used, this method only replaces servers in the NTP configuration. This method was added in vSphere API 6.7.
    //
    // @param serversParam List of host names or ip addresses of ntp servers.
    // @throws Error Generic error
	Set(serversParam []string) error

    // Get the NTP configuration status. If you run the 'timesync.get' command, you can retrieve the current time synchronization method (NTP- or VMware Tools-based). The 'ntp' command always returns the NTP server information, even when the time synchronization mode is not set to NTP. If the time synchronization mode is not NTP-based, the NTP server status is displayed as down. This method was added in vSphere API 6.7.
    // @return List of NTP servers.
    // @throws Error Generic error
	Get() ([]string, error)
}
