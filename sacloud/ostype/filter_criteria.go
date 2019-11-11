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

package ostype

import (
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/search/keys"
)

// ArchiveCriteria OSTypeごとのアーカイブ検索条件
var ArchiveCriteria = map[ArchiveOSType]search.Filter{
	CentOS: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "distro-centos"),
	},
	CentOS6: {
		search.Key(keys.Tags): search.TagsAndEqual("distro-centos", "distro-ver-6.10"),
	},
	Ubuntu: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "distro-ubuntu"),
	},
	Debian: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "distro-debian"),
	},
	CoreOS: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "distro-coreos"),
	},
	RancherOS: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "distro-rancheros"),
	},
	K3OS: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "distro-k3os"),
	},
	Kusanagi: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "pkg-kusanagi"),
	},
	SophosUTM: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "pkg-sophosutm"),
	},
	FreeBSD: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "distro-freebsd"),
	},
	Netwiser: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "pkg-netwiserve"),
	},
	OPNsense: {
		search.Key(keys.Tags): search.TagsAndEqual("current-stable", "distro-opnsense"),
	},
	Windows2016: {
		search.Key(keys.Tags): search.TagsAndEqual("os-windows", "distro-ver-2016"),
		search.Key(keys.Name): search.OrEqual("Windows Server 2016 Datacenter Edition"),
	},
	Windows2016RDS: {
		search.Key(keys.Tags): search.TagsAndEqual("os-windows", "distro-ver-2016", "windows-rds"),
		search.Key(keys.Name): search.OrEqual("Windows Server 2016 for RDS"),
	},
	Windows2016RDSOffice: {
		search.Key(keys.Tags): search.TagsAndEqual("os-windows", "distro-ver-2016", "windows-rds", "with-office"),
		search.Key(keys.Name): search.OrEqual("Windows Server 2016 for RDS(MS Office付)"),
	},
	Windows2016SQLServerWeb: {
		search.Key(keys.Tags): search.TagsAndEqual("os-windows", "distro-ver-2016", "windows-sqlserver", "sqlserver-2016", "edition-web"),
		search.Key(keys.Name): search.OrEqual("Windows Server 2016 for MS SQL 2016(Web)"),
	},
	Windows2016SQLServerStandard: {
		search.Key(keys.Tags): search.TagsAndEqual("os-windows", "distro-ver-2016", "windows-sqlserver", "sqlserver-2016", "edition-standard"),
		search.Key(keys.Name): search.OrEqual("Windows Server 2016 for MS SQL 2016(Standard)"),
	},
	Windows2016SQLServer2017Standard: {
		search.Key(keys.Tags): search.TagsAndEqual("os-windows", "distro-ver-2016", "windows-sqlserver", "sqlserver-2017", "edition-standard"),
		search.Key(keys.Name): search.OrEqual("Windows Server 2016 for MS SQL 2017(Standard)"),
	},
	Windows2016SQLServerStandardAll: {
		search.Key(keys.Tags): search.TagsAndEqual("os-windows", "distro-ver-2016", "windows-sqlserver", "sqlserver-2016", "edition-standard", "windows-rds", "with-office"),
		search.Key(keys.Name): search.OrEqual("Windows Server 2016 for MS SQL 2016(Std) with RDS / MS Office"),
	},
	Windows2016SQLServer2017StandardAll: {
		search.Key(keys.Tags): search.TagsAndEqual("os-windows", "distro-ver-2016", "windows-sqlserver", "sqlserver-2017", "edition-standard", "windows-rds", "with-office"),
		search.Key(keys.Name): search.OrEqual("Windows Server 2016 for MS SQL 2017(Std) with RDS / MS Office"),
	},
	Windows2019: {
		search.Key(keys.Tags): search.TagsAndEqual("os-windows", "distro-ver-2019"),
		search.Key(keys.Name): search.OrEqual("Windows Server 2019 Datacenter Edition"),
	},
}
