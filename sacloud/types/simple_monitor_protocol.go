// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// ESimpleMonitorProtocol シンプル監視 プロトコル
type ESimpleMonitorProtocol string

// SimpleMonitorProtocols シンプル監視 プロトコル
var SimpleMonitorProtocols = struct {
	// Unknown 不明
	Unknown ESimpleMonitorProtocol
	// HTTP http
	HTTP ESimpleMonitorProtocol
	// HTTPS https
	HTTPS ESimpleMonitorProtocol
	// Ping ping
	Ping ESimpleMonitorProtocol
	// TCP tcp
	TCP ESimpleMonitorProtocol
	// DNS dns
	DNS ESimpleMonitorProtocol
	// SSH ssh
	SSH ESimpleMonitorProtocol
	// SMTP smtp
	SMTP ESimpleMonitorProtocol
	// POP3 pop3
	POP3 ESimpleMonitorProtocol
	// SNMP snmp
	SNMP ESimpleMonitorProtocol
	// SSLCertificate sslcertificate
	SSLCertificate ESimpleMonitorProtocol
}{
	Unknown:        ESimpleMonitorProtocol(""),
	HTTP:           ESimpleMonitorProtocol("http"),
	HTTPS:          ESimpleMonitorProtocol("https"),
	Ping:           ESimpleMonitorProtocol("ping"),
	TCP:            ESimpleMonitorProtocol("tcp"),
	DNS:            ESimpleMonitorProtocol("dns"),
	SSH:            ESimpleMonitorProtocol("ssh"),
	SMTP:           ESimpleMonitorProtocol("smtp"),
	POP3:           ESimpleMonitorProtocol("pop3"),
	SNMP:           ESimpleMonitorProtocol("snmp"),
	SSLCertificate: ESimpleMonitorProtocol("sslcertificate"),
}
