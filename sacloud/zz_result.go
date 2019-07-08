// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-result'; DO NOT EDIT

package sacloud

// ArchiveFindResult represents the Result of API
type ArchiveFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Archives []*Archive `json:",omitempty" mapconv:"[]Archives,omitempty,recursive"`
}

// ArchiveCreateResult represents the Result of API
type ArchiveCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Archive *Archive `json:",omitempty" mapconv:"Archive,omitempty,recursive"`
}

// ArchiveCreateBlankResult represents the Result of API
type ArchiveCreateBlankResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Archive   *Archive   `json:",omitempty" mapconv:"Archive,omitempty,recursive"`
	FTPServer *FTPServer `json:",omitempty" mapconv:"FTPServer,omitempty,recursive"`
}

// ArchiveReadResult represents the Result of API
type ArchiveReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Archive *Archive `json:",omitempty" mapconv:"Archive,omitempty,recursive"`
}

// ArchiveUpdateResult represents the Result of API
type ArchiveUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Archive *Archive `json:",omitempty" mapconv:"Archive,omitempty,recursive"`
}

// ArchiveOpenFTPResult represents the Result of API
type ArchiveOpenFTPResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	FTPServer *FTPServer `json:",omitempty" mapconv:"FTPServer,omitempty,recursive"`
}

// AuthStatusReadResult represents the Result of API
type AuthStatusReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	AuthStatus *AuthStatus `json:",omitempty" mapconv:"AuthStatus,omitempty,recursive"`
}

// BridgeFindResult represents the Result of API
type BridgeFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Bridges []*Bridge `json:",omitempty" mapconv:"[]Bridges,omitempty,recursive"`
}

// BridgeCreateResult represents the Result of API
type BridgeCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Bridge *Bridge `json:",omitempty" mapconv:"Bridge,omitempty,recursive"`
}

// BridgeReadResult represents the Result of API
type BridgeReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Bridge *Bridge `json:",omitempty" mapconv:"Bridge,omitempty,recursive"`
}

// BridgeUpdateResult represents the Result of API
type BridgeUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Bridge *Bridge `json:",omitempty" mapconv:"Bridge,omitempty,recursive"`
}

// CDROMFindResult represents the Result of API
type CDROMFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	CDROMs []*CDROM `json:",omitempty" mapconv:"[]CDROMs,omitempty,recursive"`
}

// CDROMCreateResult represents the Result of API
type CDROMCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	CDROM     *CDROM     `json:",omitempty" mapconv:"CDROM,omitempty,recursive"`
	FTPServer *FTPServer `json:",omitempty" mapconv:"FTPServer,omitempty,recursive"`
}

// CDROMReadResult represents the Result of API
type CDROMReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	CDROM *CDROM `json:",omitempty" mapconv:"CDROM,omitempty,recursive"`
}

// CDROMUpdateResult represents the Result of API
type CDROMUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	CDROM *CDROM `json:",omitempty" mapconv:"CDROM,omitempty,recursive"`
}

// CDROMOpenFTPResult represents the Result of API
type CDROMOpenFTPResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	FTPServer *FTPServer `json:",omitempty" mapconv:"FTPServer,omitempty,recursive"`
}

// DiskFindResult represents the Result of API
type DiskFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Disks []*Disk `json:",omitempty" mapconv:"[]Disks,omitempty,recursive"`
}

// DiskCreateResult represents the Result of API
type DiskCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// DiskCreateDistantlyResult represents the Result of API
type DiskCreateDistantlyResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// DiskCreateWithConfigResult represents the Result of API
type DiskCreateWithConfigResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// DiskCreateWithConfigDistantlyResult represents the Result of API
type DiskCreateWithConfigDistantlyResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// DiskInstallDistantFromResult represents the Result of API
type DiskInstallDistantFromResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// DiskInstallResult represents the Result of API
type DiskInstallResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// DiskReadResult represents the Result of API
type DiskReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// DiskUpdateResult represents the Result of API
type DiskUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Disk *Disk `json:",omitempty" mapconv:"Disk,omitempty,recursive"`
}

// DiskMonitorResult represents the Result of API
type DiskMonitorResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	DiskActivity *DiskActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// GSLBFindResult represents the Result of API
type GSLBFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	GSLBs []*GSLB `json:",omitempty" mapconv:"[]CommonServiceItems,omitempty,recursive"`
}

// GSLBCreateResult represents the Result of API
type GSLBCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	GSLB *GSLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// GSLBReadResult represents the Result of API
type GSLBReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	GSLB *GSLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// GSLBUpdateResult represents the Result of API
type GSLBUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	GSLB *GSLB `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// InterfaceFindResult represents the Result of API
type InterfaceFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Interfaces []*Interface `json:",omitempty" mapconv:"[]Interfaces,omitempty,recursive"`
}

// InterfaceCreateResult represents the Result of API
type InterfaceCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Interface *Interface `json:",omitempty" mapconv:"Interface,omitempty,recursive"`
}

// InterfaceReadResult represents the Result of API
type InterfaceReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Interface *Interface `json:",omitempty" mapconv:"Interface,omitempty,recursive"`
}

// InterfaceUpdateResult represents the Result of API
type InterfaceUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Interface *Interface `json:",omitempty" mapconv:"Interface,omitempty,recursive"`
}

// InterfaceMonitorResult represents the Result of API
type InterfaceMonitorResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	InterfaceActivity *InterfaceActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// InternetFindResult represents the Result of API
type InternetFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Internets []*Internet `json:",omitempty" mapconv:"[]Internets,omitempty,recursive"`
}

// InternetCreateResult represents the Result of API
type InternetCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Internet *Internet `json:",omitempty" mapconv:"Internet,omitempty,recursive"`
}

// InternetReadResult represents the Result of API
type InternetReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Internet *Internet `json:",omitempty" mapconv:"Internet,omitempty,recursive"`
}

// InternetUpdateResult represents the Result of API
type InternetUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Internet *Internet `json:",omitempty" mapconv:"Internet,omitempty,recursive"`
}

// InternetUpdateBandWidthResult represents the Result of API
type InternetUpdateBandWidthResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Internet *Internet `json:",omitempty" mapconv:"Internet,omitempty,recursive"`
}

// InternetAddSubnetResult represents the Result of API
type InternetAddSubnetResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Subnet *InternetSubnetOperationResult `json:",omitempty" mapconv:"Subnet,omitempty,recursive"`
}

// InternetUpdateSubnetResult represents the Result of API
type InternetUpdateSubnetResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Subnet *InternetSubnetOperationResult `json:",omitempty" mapconv:"Subnet,omitempty,recursive"`
}

// InternetMonitorResult represents the Result of API
type InternetMonitorResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	RouterActivity *RouterActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// LoadBalancerFindResult represents the Result of API
type LoadBalancerFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	LoadBalancers []*LoadBalancer `json:",omitempty" mapconv:"[]Appliances,omitempty,recursive"`
}

// LoadBalancerCreateResult represents the Result of API
type LoadBalancerCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	LoadBalancer *LoadBalancer `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// LoadBalancerReadResult represents the Result of API
type LoadBalancerReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	LoadBalancer *LoadBalancer `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// LoadBalancerUpdateResult represents the Result of API
type LoadBalancerUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	LoadBalancer *LoadBalancer `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// LoadBalancerMonitorInterfaceResult represents the Result of API
type LoadBalancerMonitorInterfaceResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	InterfaceActivity *InterfaceActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// LoadBalancerStatusResult represents the Result of API
type LoadBalancerStatusResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Status []*LoadBalancerStatus `json:",omitempty" mapconv:"[]LoadBalancer,omitempty,recursive"`
}

// NFSFindResult represents the Result of API
type NFSFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	NFS []*NFS `json:",omitempty" mapconv:"[]Appliances,omitempty,recursive"`
}

// NFSCreateResult represents the Result of API
type NFSCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	NFS *NFS `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// NFSReadResult represents the Result of API
type NFSReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	NFS *NFS `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// NFSUpdateResult represents the Result of API
type NFSUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	NFS *NFS `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// NFSMonitorFreeDiskSizeResult represents the Result of API
type NFSMonitorFreeDiskSizeResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	FreeDiskSizeActivity *FreeDiskSizeActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// NFSMonitorInterfaceResult represents the Result of API
type NFSMonitorInterfaceResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	InterfaceActivity *InterfaceActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// NoteFindResult represents the Result of API
type NoteFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Notes []*Note `json:",omitempty" mapconv:"[]Notes,omitempty,recursive"`
}

// NoteCreateResult represents the Result of API
type NoteCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Note *Note `json:",omitempty" mapconv:"Note,omitempty,recursive"`
}

// NoteReadResult represents the Result of API
type NoteReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Note *Note `json:",omitempty" mapconv:"Note,omitempty,recursive"`
}

// NoteUpdateResult represents the Result of API
type NoteUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Note *Note `json:",omitempty" mapconv:"Note,omitempty,recursive"`
}

// PacketFilterFindResult represents the Result of API
type PacketFilterFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	PacketFilters []*PacketFilter `json:",omitempty" mapconv:"[]PacketFilters,omitempty,recursive"`
}

// PacketFilterCreateResult represents the Result of API
type PacketFilterCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	PacketFilter *PacketFilter `json:",omitempty" mapconv:"PacketFilter,omitempty,recursive"`
}

// PacketFilterReadResult represents the Result of API
type PacketFilterReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	PacketFilter *PacketFilter `json:",omitempty" mapconv:"PacketFilter,omitempty,recursive"`
}

// PacketFilterUpdateResult represents the Result of API
type PacketFilterUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	PacketFilter *PacketFilter `json:",omitempty" mapconv:"PacketFilter,omitempty,recursive"`
}

// ServerFindResult represents the Result of API
type ServerFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Servers []*Server `json:",omitempty" mapconv:"[]Servers,omitempty,recursive"`
}

// ServerCreateResult represents the Result of API
type ServerCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Server *Server `json:",omitempty" mapconv:"Server,omitempty,recursive"`
}

// ServerReadResult represents the Result of API
type ServerReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Server *Server `json:",omitempty" mapconv:"Server,omitempty,recursive"`
}

// ServerUpdateResult represents the Result of API
type ServerUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Server *Server `json:",omitempty" mapconv:"Server,omitempty,recursive"`
}

// ServerChangePlanResult represents the Result of API
type ServerChangePlanResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Server *Server `json:",omitempty" mapconv:"Server,omitempty,recursive"`
}

// ServerMonitorResult represents the Result of API
type ServerMonitorResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	CPUTimeActivity *CPUTimeActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// SIMFindResult represents the Result of API
type SIMFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	SIMs []*SIM `json:",omitempty" mapconv:"[]CommonServiceItems,omitempty,recursive"`
}

// SIMCreateResult represents the Result of API
type SIMCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SIM *SIM `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// SIMReadResult represents the Result of API
type SIMReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SIM *SIM `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// SIMUpdateResult represents the Result of API
type SIMUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	SIM *SIM `json:",omitempty" mapconv:"CommonServiceItem,omitempty,recursive"`
}

// SIMLogsResult represents the Result of API
type SIMLogsResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Logs []*SIMLog `json:",omitempty" mapconv:"[]Logs,omitempty,recursive"`
}

// SIMGetNetworkOperatorResult represents the Result of API
type SIMGetNetworkOperatorResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Configs []*SIMNetworkOperatorConfig `json:",omitempty" mapconv:"[]NetworkOperationConfigs,omitempty,recursive"`
}

// SIMMonitorSIMResult represents the Result of API
type SIMMonitorSIMResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	LinkActivity *LinkActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// SwitchFindResult represents the Result of API
type SwitchFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Switches []*Switch `json:",omitempty" mapconv:"[]Switches,omitempty,recursive"`
}

// SwitchCreateResult represents the Result of API
type SwitchCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Switch *Switch `json:",omitempty" mapconv:"Switch,omitempty,recursive"`
}

// SwitchReadResult represents the Result of API
type SwitchReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Switch *Switch `json:",omitempty" mapconv:"Switch,omitempty,recursive"`
}

// SwitchUpdateResult represents the Result of API
type SwitchUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Switch *Switch `json:",omitempty" mapconv:"Switch,omitempty,recursive"`
}

// VPCRouterFindResult represents the Result of API
type VPCRouterFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	VPCRouters []*VPCRouter `json:",omitempty" mapconv:"[]Appliances,omitempty,recursive"`
}

// VPCRouterCreateResult represents the Result of API
type VPCRouterCreateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	VPCRouter *VPCRouter `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// VPCRouterReadResult represents the Result of API
type VPCRouterReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	VPCRouter *VPCRouter `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// VPCRouterUpdateResult represents the Result of API
type VPCRouterUpdateResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	VPCRouter *VPCRouter `json:",omitempty" mapconv:"Appliance,omitempty,recursive"`
}

// VPCRouterMonitorInterfaceResult represents the Result of API
type VPCRouterMonitorInterfaceResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	InterfaceActivity *InterfaceActivity `json:",omitempty" mapconv:"Data,omitempty,recursive"`
}

// ZoneFindResult represents the Result of API
type ZoneFindResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Zones []*Zone `json:",omitempty" mapconv:"[]Zones,omitempty,recursive"`
}

// ZoneReadResult represents the Result of API
type ZoneReadResult struct {
	IsOk bool `json:",omitempty"` // is_ok

	Zone *Zone `json:",omitempty" mapconv:"Zone,omitempty,recursive"`
}
