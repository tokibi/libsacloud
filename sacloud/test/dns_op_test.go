package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestDNSOp_CRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testDNSCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createDNSExpected,
				IgnoreFields: ignoreDNSFields,
			}),
		},

		Read: &CRUDTestFunc{
			Func: testDNSRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createDNSExpected,
				IgnoreFields: ignoreDNSFields,
			}),
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testDNSUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateDNSExpected,
					IgnoreFields: ignoreDNSFields,
				}),
			},
			{
				Func: testDNSUpdateToMin,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateDNSToMinExpected,
					IgnoreFields: ignoreDNSFields,
				}),
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testDNSDelete,
		},
	})
}

var (
	ignoreDNSFields = []string{
		"ID",
		"Class",
		"SettingsHash",
		"FQDN",
		"CreatedAt",
		"ModifiedAt",
		"DNSNameServers",
	}
	createDNSParam = &sacloud.DNSCreateRequest{
		Name:        "libsacloud-test.com",
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		Records: []*sacloud.DNSRecord{
			{
				Name:  "host1",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.1",
			},
			{
				Name:  "host2",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.2",
			},
		},
	}
	createDNSExpected = &sacloud.DNS{
		Name:         createDNSParam.Name,
		Description:  createDNSParam.Description,
		Tags:         createDNSParam.Tags,
		Availability: types.Availabilities.Available,
		DNSZone:      createDNSParam.Name,
		Records:      createDNSParam.Records,
	}
	updateDNSParam = &sacloud.DNSUpdateRequest{
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		IconID:      testIconID,
		Records: []*sacloud.DNSRecord{
			{
				Name:  "host1",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.11",
			},
			{
				Name:  "host2",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.12",
			},
			{
				Name:  "host3",
				Type:  types.DNSRecordTypes.A,
				RData: "192.0.2.13",
			},
		},
	}
	updateDNSExpected = &sacloud.DNS{
		Name:         createDNSParam.Name,
		Description:  updateDNSParam.Description,
		IconID:       testIconID,
		Tags:         updateDNSParam.Tags,
		Availability: types.Availabilities.Available,
		DNSZone:      createDNSParam.Name,
		Records:      updateDNSParam.Records,
	}
	updateDNSToMinParam    = &sacloud.DNSUpdateRequest{}
	updateDNSToMinExpected = &sacloud.DNS{
		Name:         createDNSParam.Name,
		Availability: types.Availabilities.Available,
		DNSZone:      createDNSParam.Name,
	}
)

func testDNSCreate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Create(ctx, createDNSParam)
}

func testDNSRead(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testDNSUpdate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Update(ctx, ctx.ID, updateDNSParam)
}

func testDNSUpdateToMin(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDNSOp(caller)
	return client.Update(ctx, ctx.ID, updateDNSToMinParam)
}

func testDNSDelete(ctx *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewDNSOp(caller)
	return client.Delete(ctx, ctx.ID)
}
