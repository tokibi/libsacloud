// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-fake-store'; DO NOT EDIT

package fake

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func (s *store) getArchive(zone string) []*sacloud.Archive {
	values := s.get(ResourceArchive, zone)
	var ret []*sacloud.Archive
	for _, v := range values {
		if v, ok := v.(*sacloud.Archive); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getArchiveByID(zone string, id types.ID) *sacloud.Archive {
	v := s.getByID(ResourceArchive, zone, id)
	if v, ok := v.(*sacloud.Archive); ok {
		return v
	}
	return nil
}

func (s *store) setArchive(zone string, value *sacloud.Archive) {
	s.set(ResourceArchive, zone, value)
}

func (s *store) getAuthStatus(zone string) []*sacloud.AuthStatus {
	values := s.get(ResourceAuthStatus, zone)
	var ret []*sacloud.AuthStatus
	for _, v := range values {
		if v, ok := v.(*sacloud.AuthStatus); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getAuthStatusByID(zone string, id types.ID) *sacloud.AuthStatus {
	v := s.getByID(ResourceAuthStatus, zone, id)
	if v, ok := v.(*sacloud.AuthStatus); ok {
		return v
	}
	return nil
}

func (s *store) setAuthStatus(zone string, value *sacloud.AuthStatus) {
	s.set(ResourceAuthStatus, zone, value)
}

func (s *store) getAutoBackup(zone string) []*sacloud.AutoBackup {
	values := s.get(ResourceAutoBackup, zone)
	var ret []*sacloud.AutoBackup
	for _, v := range values {
		if v, ok := v.(*sacloud.AutoBackup); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getAutoBackupByID(zone string, id types.ID) *sacloud.AutoBackup {
	v := s.getByID(ResourceAutoBackup, zone, id)
	if v, ok := v.(*sacloud.AutoBackup); ok {
		return v
	}
	return nil
}

func (s *store) setAutoBackup(zone string, value *sacloud.AutoBackup) {
	s.set(ResourceAutoBackup, zone, value)
}

func (s *store) getBill(zone string) []*sacloud.Bill {
	values := s.get(ResourceBill, zone)
	var ret []*sacloud.Bill
	for _, v := range values {
		if v, ok := v.(*sacloud.Bill); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getBillByID(zone string, id types.ID) *sacloud.Bill {
	v := s.getByID(ResourceBill, zone, id)
	if v, ok := v.(*sacloud.Bill); ok {
		return v
	}
	return nil
}

func (s *store) setBill(zone string, value *sacloud.Bill) {
	s.set(ResourceBill, zone, value)
}

func (s *store) getBridge(zone string) []*sacloud.Bridge {
	values := s.get(ResourceBridge, zone)
	var ret []*sacloud.Bridge
	for _, v := range values {
		if v, ok := v.(*sacloud.Bridge); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getBridgeByID(zone string, id types.ID) *sacloud.Bridge {
	v := s.getByID(ResourceBridge, zone, id)
	if v, ok := v.(*sacloud.Bridge); ok {
		return v
	}
	return nil
}

func (s *store) setBridge(zone string, value *sacloud.Bridge) {
	s.set(ResourceBridge, zone, value)
}

func (s *store) getCDROM(zone string) []*sacloud.CDROM {
	values := s.get(ResourceCDROM, zone)
	var ret []*sacloud.CDROM
	for _, v := range values {
		if v, ok := v.(*sacloud.CDROM); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getCDROMByID(zone string, id types.ID) *sacloud.CDROM {
	v := s.getByID(ResourceCDROM, zone, id)
	if v, ok := v.(*sacloud.CDROM); ok {
		return v
	}
	return nil
}

func (s *store) setCDROM(zone string, value *sacloud.CDROM) {
	s.set(ResourceCDROM, zone, value)
}

func (s *store) getCoupon(zone string) []*sacloud.Coupon {
	values := s.get(ResourceCoupon, zone)
	var ret []*sacloud.Coupon
	for _, v := range values {
		if v, ok := v.(*sacloud.Coupon); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getCouponByID(zone string, id types.ID) *sacloud.Coupon {
	v := s.getByID(ResourceCoupon, zone, id)
	if v, ok := v.(*sacloud.Coupon); ok {
		return v
	}
	return nil
}

func (s *store) setCoupon(zone string, value *sacloud.Coupon) {
	s.set(ResourceCoupon, zone, value)
}

func (s *store) getDatabase(zone string) []*sacloud.Database {
	values := s.get(ResourceDatabase, zone)
	var ret []*sacloud.Database
	for _, v := range values {
		if v, ok := v.(*sacloud.Database); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getDatabaseByID(zone string, id types.ID) *sacloud.Database {
	v := s.getByID(ResourceDatabase, zone, id)
	if v, ok := v.(*sacloud.Database); ok {
		return v
	}
	return nil
}

func (s *store) setDatabase(zone string, value *sacloud.Database) {
	s.set(ResourceDatabase, zone, value)
}

func (s *store) getDisk(zone string) []*sacloud.Disk {
	values := s.get(ResourceDisk, zone)
	var ret []*sacloud.Disk
	for _, v := range values {
		if v, ok := v.(*sacloud.Disk); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getDiskByID(zone string, id types.ID) *sacloud.Disk {
	v := s.getByID(ResourceDisk, zone, id)
	if v, ok := v.(*sacloud.Disk); ok {
		return v
	}
	return nil
}

func (s *store) setDisk(zone string, value *sacloud.Disk) {
	s.set(ResourceDisk, zone, value)
}

func (s *store) getDiskPlan(zone string) []*sacloud.DiskPlan {
	values := s.get(ResourceDiskPlan, zone)
	var ret []*sacloud.DiskPlan
	for _, v := range values {
		if v, ok := v.(*sacloud.DiskPlan); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getDiskPlanByID(zone string, id types.ID) *sacloud.DiskPlan {
	v := s.getByID(ResourceDiskPlan, zone, id)
	if v, ok := v.(*sacloud.DiskPlan); ok {
		return v
	}
	return nil
}

func (s *store) setDiskPlan(zone string, value *sacloud.DiskPlan) {
	s.set(ResourceDiskPlan, zone, value)
}

func (s *store) getDNS(zone string) []*sacloud.DNS {
	values := s.get(ResourceDNS, zone)
	var ret []*sacloud.DNS
	for _, v := range values {
		if v, ok := v.(*sacloud.DNS); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getDNSByID(zone string, id types.ID) *sacloud.DNS {
	v := s.getByID(ResourceDNS, zone, id)
	if v, ok := v.(*sacloud.DNS); ok {
		return v
	}
	return nil
}

func (s *store) setDNS(zone string, value *sacloud.DNS) {
	s.set(ResourceDNS, zone, value)
}

func (s *store) getGSLB(zone string) []*sacloud.GSLB {
	values := s.get(ResourceGSLB, zone)
	var ret []*sacloud.GSLB
	for _, v := range values {
		if v, ok := v.(*sacloud.GSLB); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getGSLBByID(zone string, id types.ID) *sacloud.GSLB {
	v := s.getByID(ResourceGSLB, zone, id)
	if v, ok := v.(*sacloud.GSLB); ok {
		return v
	}
	return nil
}

func (s *store) setGSLB(zone string, value *sacloud.GSLB) {
	s.set(ResourceGSLB, zone, value)
}

func (s *store) getIcon(zone string) []*sacloud.Icon {
	values := s.get(ResourceIcon, zone)
	var ret []*sacloud.Icon
	for _, v := range values {
		if v, ok := v.(*sacloud.Icon); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getIconByID(zone string, id types.ID) *sacloud.Icon {
	v := s.getByID(ResourceIcon, zone, id)
	if v, ok := v.(*sacloud.Icon); ok {
		return v
	}
	return nil
}

func (s *store) setIcon(zone string, value *sacloud.Icon) {
	s.set(ResourceIcon, zone, value)
}

func (s *store) getInterface(zone string) []*sacloud.Interface {
	values := s.get(ResourceInterface, zone)
	var ret []*sacloud.Interface
	for _, v := range values {
		if v, ok := v.(*sacloud.Interface); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getInterfaceByID(zone string, id types.ID) *sacloud.Interface {
	v := s.getByID(ResourceInterface, zone, id)
	if v, ok := v.(*sacloud.Interface); ok {
		return v
	}
	return nil
}

func (s *store) setInterface(zone string, value *sacloud.Interface) {
	s.set(ResourceInterface, zone, value)
}

func (s *store) getInternet(zone string) []*sacloud.Internet {
	values := s.get(ResourceInternet, zone)
	var ret []*sacloud.Internet
	for _, v := range values {
		if v, ok := v.(*sacloud.Internet); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getInternetByID(zone string, id types.ID) *sacloud.Internet {
	v := s.getByID(ResourceInternet, zone, id)
	if v, ok := v.(*sacloud.Internet); ok {
		return v
	}
	return nil
}

func (s *store) setInternet(zone string, value *sacloud.Internet) {
	s.set(ResourceInternet, zone, value)
}

func (s *store) getInternetPlan(zone string) []*sacloud.InternetPlan {
	values := s.get(ResourceInternetPlan, zone)
	var ret []*sacloud.InternetPlan
	for _, v := range values {
		if v, ok := v.(*sacloud.InternetPlan); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getInternetPlanByID(zone string, id types.ID) *sacloud.InternetPlan {
	v := s.getByID(ResourceInternetPlan, zone, id)
	if v, ok := v.(*sacloud.InternetPlan); ok {
		return v
	}
	return nil
}

func (s *store) setInternetPlan(zone string, value *sacloud.InternetPlan) {
	s.set(ResourceInternetPlan, zone, value)
}

func (s *store) getLicense(zone string) []*sacloud.License {
	values := s.get(ResourceLicense, zone)
	var ret []*sacloud.License
	for _, v := range values {
		if v, ok := v.(*sacloud.License); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getLicenseByID(zone string, id types.ID) *sacloud.License {
	v := s.getByID(ResourceLicense, zone, id)
	if v, ok := v.(*sacloud.License); ok {
		return v
	}
	return nil
}

func (s *store) setLicense(zone string, value *sacloud.License) {
	s.set(ResourceLicense, zone, value)
}

func (s *store) getLicenseInfo(zone string) []*sacloud.LicenseInfo {
	values := s.get(ResourceLicenseInfo, zone)
	var ret []*sacloud.LicenseInfo
	for _, v := range values {
		if v, ok := v.(*sacloud.LicenseInfo); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getLicenseInfoByID(zone string, id types.ID) *sacloud.LicenseInfo {
	v := s.getByID(ResourceLicenseInfo, zone, id)
	if v, ok := v.(*sacloud.LicenseInfo); ok {
		return v
	}
	return nil
}

func (s *store) setLicenseInfo(zone string, value *sacloud.LicenseInfo) {
	s.set(ResourceLicenseInfo, zone, value)
}

func (s *store) getLoadBalancer(zone string) []*sacloud.LoadBalancer {
	values := s.get(ResourceLoadBalancer, zone)
	var ret []*sacloud.LoadBalancer
	for _, v := range values {
		if v, ok := v.(*sacloud.LoadBalancer); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getLoadBalancerByID(zone string, id types.ID) *sacloud.LoadBalancer {
	v := s.getByID(ResourceLoadBalancer, zone, id)
	if v, ok := v.(*sacloud.LoadBalancer); ok {
		return v
	}
	return nil
}

func (s *store) setLoadBalancer(zone string, value *sacloud.LoadBalancer) {
	s.set(ResourceLoadBalancer, zone, value)
}

func (s *store) getNFS(zone string) []*sacloud.NFS {
	values := s.get(ResourceNFS, zone)
	var ret []*sacloud.NFS
	for _, v := range values {
		if v, ok := v.(*sacloud.NFS); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getNFSByID(zone string, id types.ID) *sacloud.NFS {
	v := s.getByID(ResourceNFS, zone, id)
	if v, ok := v.(*sacloud.NFS); ok {
		return v
	}
	return nil
}

func (s *store) setNFS(zone string, value *sacloud.NFS) {
	s.set(ResourceNFS, zone, value)
}

func (s *store) getNote(zone string) []*sacloud.Note {
	values := s.get(ResourceNote, zone)
	var ret []*sacloud.Note
	for _, v := range values {
		if v, ok := v.(*sacloud.Note); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getNoteByID(zone string, id types.ID) *sacloud.Note {
	v := s.getByID(ResourceNote, zone, id)
	if v, ok := v.(*sacloud.Note); ok {
		return v
	}
	return nil
}

func (s *store) setNote(zone string, value *sacloud.Note) {
	s.set(ResourceNote, zone, value)
}

func (s *store) getPacketFilter(zone string) []*sacloud.PacketFilter {
	values := s.get(ResourcePacketFilter, zone)
	var ret []*sacloud.PacketFilter
	for _, v := range values {
		if v, ok := v.(*sacloud.PacketFilter); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getPacketFilterByID(zone string, id types.ID) *sacloud.PacketFilter {
	v := s.getByID(ResourcePacketFilter, zone, id)
	if v, ok := v.(*sacloud.PacketFilter); ok {
		return v
	}
	return nil
}

func (s *store) setPacketFilter(zone string, value *sacloud.PacketFilter) {
	s.set(ResourcePacketFilter, zone, value)
}

func (s *store) getPrivateHost(zone string) []*sacloud.PrivateHost {
	values := s.get(ResourcePrivateHost, zone)
	var ret []*sacloud.PrivateHost
	for _, v := range values {
		if v, ok := v.(*sacloud.PrivateHost); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getPrivateHostByID(zone string, id types.ID) *sacloud.PrivateHost {
	v := s.getByID(ResourcePrivateHost, zone, id)
	if v, ok := v.(*sacloud.PrivateHost); ok {
		return v
	}
	return nil
}

func (s *store) setPrivateHost(zone string, value *sacloud.PrivateHost) {
	s.set(ResourcePrivateHost, zone, value)
}

func (s *store) getPrivateHostPlan(zone string) []*sacloud.PrivateHostPlan {
	values := s.get(ResourcePrivateHostPlan, zone)
	var ret []*sacloud.PrivateHostPlan
	for _, v := range values {
		if v, ok := v.(*sacloud.PrivateHostPlan); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getPrivateHostPlanByID(zone string, id types.ID) *sacloud.PrivateHostPlan {
	v := s.getByID(ResourcePrivateHostPlan, zone, id)
	if v, ok := v.(*sacloud.PrivateHostPlan); ok {
		return v
	}
	return nil
}

func (s *store) setPrivateHostPlan(zone string, value *sacloud.PrivateHostPlan) {
	s.set(ResourcePrivateHostPlan, zone, value)
}

func (s *store) getProxyLB(zone string) []*sacloud.ProxyLB {
	values := s.get(ResourceProxyLB, zone)
	var ret []*sacloud.ProxyLB
	for _, v := range values {
		if v, ok := v.(*sacloud.ProxyLB); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getProxyLBByID(zone string, id types.ID) *sacloud.ProxyLB {
	v := s.getByID(ResourceProxyLB, zone, id)
	if v, ok := v.(*sacloud.ProxyLB); ok {
		return v
	}
	return nil
}

func (s *store) setProxyLB(zone string, value *sacloud.ProxyLB) {
	s.set(ResourceProxyLB, zone, value)
}

func (s *store) getRegion(zone string) []*sacloud.Region {
	values := s.get(ResourceRegion, zone)
	var ret []*sacloud.Region
	for _, v := range values {
		if v, ok := v.(*sacloud.Region); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getRegionByID(zone string, id types.ID) *sacloud.Region {
	v := s.getByID(ResourceRegion, zone, id)
	if v, ok := v.(*sacloud.Region); ok {
		return v
	}
	return nil
}

func (s *store) setRegion(zone string, value *sacloud.Region) {
	s.set(ResourceRegion, zone, value)
}

func (s *store) getServer(zone string) []*sacloud.Server {
	values := s.get(ResourceServer, zone)
	var ret []*sacloud.Server
	for _, v := range values {
		if v, ok := v.(*sacloud.Server); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getServerByID(zone string, id types.ID) *sacloud.Server {
	v := s.getByID(ResourceServer, zone, id)
	if v, ok := v.(*sacloud.Server); ok {
		return v
	}
	return nil
}

func (s *store) setServer(zone string, value *sacloud.Server) {
	s.set(ResourceServer, zone, value)
}

func (s *store) getServerPlan(zone string) []*sacloud.ServerPlan {
	values := s.get(ResourceServerPlan, zone)
	var ret []*sacloud.ServerPlan
	for _, v := range values {
		if v, ok := v.(*sacloud.ServerPlan); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getServerPlanByID(zone string, id types.ID) *sacloud.ServerPlan {
	v := s.getByID(ResourceServerPlan, zone, id)
	if v, ok := v.(*sacloud.ServerPlan); ok {
		return v
	}
	return nil
}

func (s *store) setServerPlan(zone string, value *sacloud.ServerPlan) {
	s.set(ResourceServerPlan, zone, value)
}

func (s *store) getServiceClass(zone string) []*sacloud.ServiceClass {
	values := s.get(ResourceServiceClass, zone)
	var ret []*sacloud.ServiceClass
	for _, v := range values {
		if v, ok := v.(*sacloud.ServiceClass); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getServiceClassByID(zone string, id types.ID) *sacloud.ServiceClass {
	v := s.getByID(ResourceServiceClass, zone, id)
	if v, ok := v.(*sacloud.ServiceClass); ok {
		return v
	}
	return nil
}

func (s *store) setServiceClass(zone string, value *sacloud.ServiceClass) {
	s.set(ResourceServiceClass, zone, value)
}

func (s *store) getSIM(zone string) []*sacloud.SIM {
	values := s.get(ResourceSIM, zone)
	var ret []*sacloud.SIM
	for _, v := range values {
		if v, ok := v.(*sacloud.SIM); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getSIMByID(zone string, id types.ID) *sacloud.SIM {
	v := s.getByID(ResourceSIM, zone, id)
	if v, ok := v.(*sacloud.SIM); ok {
		return v
	}
	return nil
}

func (s *store) setSIM(zone string, value *sacloud.SIM) {
	s.set(ResourceSIM, zone, value)
}

func (s *store) getSimpleMonitor(zone string) []*sacloud.SimpleMonitor {
	values := s.get(ResourceSimpleMonitor, zone)
	var ret []*sacloud.SimpleMonitor
	for _, v := range values {
		if v, ok := v.(*sacloud.SimpleMonitor); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getSimpleMonitorByID(zone string, id types.ID) *sacloud.SimpleMonitor {
	v := s.getByID(ResourceSimpleMonitor, zone, id)
	if v, ok := v.(*sacloud.SimpleMonitor); ok {
		return v
	}
	return nil
}

func (s *store) setSimpleMonitor(zone string, value *sacloud.SimpleMonitor) {
	s.set(ResourceSimpleMonitor, zone, value)
}

func (s *store) getSSHKey(zone string) []*sacloud.SSHKey {
	values := s.get(ResourceSSHKey, zone)
	var ret []*sacloud.SSHKey
	for _, v := range values {
		if v, ok := v.(*sacloud.SSHKey); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getSSHKeyByID(zone string, id types.ID) *sacloud.SSHKey {
	v := s.getByID(ResourceSSHKey, zone, id)
	if v, ok := v.(*sacloud.SSHKey); ok {
		return v
	}
	return nil
}

func (s *store) setSSHKey(zone string, value *sacloud.SSHKey) {
	s.set(ResourceSSHKey, zone, value)
}

func (s *store) getSwitch(zone string) []*sacloud.Switch {
	values := s.get(ResourceSwitch, zone)
	var ret []*sacloud.Switch
	for _, v := range values {
		if v, ok := v.(*sacloud.Switch); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getSwitchByID(zone string, id types.ID) *sacloud.Switch {
	v := s.getByID(ResourceSwitch, zone, id)
	if v, ok := v.(*sacloud.Switch); ok {
		return v
	}
	return nil
}

func (s *store) setSwitch(zone string, value *sacloud.Switch) {
	s.set(ResourceSwitch, zone, value)
}

func (s *store) getVPCRouter(zone string) []*sacloud.VPCRouter {
	values := s.get(ResourceVPCRouter, zone)
	var ret []*sacloud.VPCRouter
	for _, v := range values {
		if v, ok := v.(*sacloud.VPCRouter); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getVPCRouterByID(zone string, id types.ID) *sacloud.VPCRouter {
	v := s.getByID(ResourceVPCRouter, zone, id)
	if v, ok := v.(*sacloud.VPCRouter); ok {
		return v
	}
	return nil
}

func (s *store) setVPCRouter(zone string, value *sacloud.VPCRouter) {
	s.set(ResourceVPCRouter, zone, value)
}

func (s *store) getZone(zone string) []*sacloud.Zone {
	values := s.get(ResourceZone, zone)
	var ret []*sacloud.Zone
	for _, v := range values {
		if v, ok := v.(*sacloud.Zone); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getZoneByID(zone string, id types.ID) *sacloud.Zone {
	v := s.getByID(ResourceZone, zone, id)
	if v, ok := v.(*sacloud.Zone); ok {
		return v
	}
	return nil
}

func (s *store) setZone(zone string, value *sacloud.Zone) {
	s.set(ResourceZone, zone, value)
}
