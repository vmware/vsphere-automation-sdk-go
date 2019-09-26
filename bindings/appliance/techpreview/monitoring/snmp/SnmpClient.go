/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Snmp
 * Used by client-side stubs.
 */

package snmp

import (
)

// ``Snmp`` interface provides methods SNMP agent operations.
type SnmpClient interface {


    // Restore settings to factory defaults.
    // @throws Error Generic error
    Reset() error 


    // Start a disabled SNMP agent.
    // @throws Error Generic error
    Enable() error 


    // Generate localized keys for secure SNMPv3 communications.
    //
    // @param configParam SNMP hash configuration.
    // @return SNMP hash result
    // @throws Error Generic error
    Hash(configParam SNMPHashConfig) (SNMPHashResults, error) 


    // Get SNMP limits information.
    // @return SNMP limits structure
    // @throws Error Generic error
    Limits() (SNMPLimits, error) 


    // Return an SNMP agent configuration.
    // @return SNMP config structure
    // @throws Error Generic error
    Get() (SNMPConfigReadOnly, error) 


    // Stop an enabled SNMP agent.
    // @throws Error Generic error
    Disable() error 


    // Set SNMP configuration.
    //
    // @param configParam SNMP configuration.
    // @throws Error Generic error
    Set(configParam SNMPConfig) error 


    // Send a warmStart notification to all configured traps and inform destinations (see RFC 3418).
    // @return SNMP test result
    // @throws Error Generic error
    Test() (SNMPTestResults, error) 


    // Generate diagnostics report for snmp agent.
    // @return SNMP stats
    // @throws Error Generic error
    Stats() (SNMPStats, error) 

}
