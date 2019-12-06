/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package keys defines json file keys or options for command line parameters
// to define host properties.
package keys

const (
	// AuthSchemeBasicUsernameKey defines key for User Name
	// for Basic Authenticaiton.
	AuthSchemeBasicUsernameKey = "username"
	// AuthSchemeBasicPasswordKey defines key for Password for
	// Basic Authenticaiton.
	AuthSchemeBasicPasswordKey = "password"
	// AuthSchemeOAuthCSPURLKey defines key for CSP URL for
	// OAuth Authenticaiton.
	AuthSchemeOAuthCSPURLKey = "csp-url"
	// AuthSchemeOAuthRefreshTokenKey defines key for Refresh Token for OAuth
	// Authenticaiton by Refresh Token.
	AuthSchemeOAuthRefreshTokenKey = "refresh-token"
	// AuthSchemeSAMLBearerTokenKey defines key for Bearer Token for
	// SAML Bearer Authenticaiton.
	AuthSchemeSAMLBearerTokenKey = "saml-bearer-token"
	// SkipServerVerifiationKey defines key for boolean property to skip server
	// verification. This shouldn't be set true for production instance.
	SkipServerVerifiationKey = "skip-server-verfication"
	// CertFileKey defines key for full path of Server Certificate file.
	CertFileKey = "certfile"
	// CertKeyFileKey defines key for full path of key file for
	// Server Certificate.
	CertKeyFileKey = "certkeyfile"
	// HostNameKey defines key for Host Name.
	HostNameKey = "name"
	// HostAddressKey defines key for Host Address.
	HostAddressKey = "server"
	// HostTypeKey defines key for Host Type.
	HostTypeKey = "server-type"
	// AuthSchemeKey defines key for Authentication Scheme.
	AuthSchemeKey = "auth-scheme"
	// ConfigFileKey defines key for full path of Config File.
	ConfigFileKey = "config-file"
)
